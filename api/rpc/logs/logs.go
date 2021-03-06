package logs

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/appcelerator/amp/pkg/elasticsearch"
	"github.com/appcelerator/amp/pkg/labels"
	"github.com/appcelerator/amp/pkg/nats-streaming"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats-streaming"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/olivere/elastic.v5"
)

// Server is used to implement log.LogServer
type Server struct {
	ES *elasticsearch.Elasticsearch
	NS *ns.NatsStreaming
}

// Get implements logs.LogsServer
func (s *Server) Get(ctx context.Context, in *GetRequest) (*GetReply, error) {
	if err := s.ES.Connect(); err != nil {
		return nil, errors.New("unable to connect to elasticsearch service")
	}
	log.Println("rpc-logs: Get", in.String())

	// Prepares indices
	indices := []string{}
	date := time.Now().UTC()
	for i := 0; i < 2; i++ {
		indices = append(indices, "ampbeat-"+date.Format("2006.01.02"))
		date = date.AddDate(0, 0, -1)
	}

	// Prepare request to elasticsearch
	request := s.ES.GetClient().Search().Index(indices...).IgnoreUnavailable(true)
	request.Type("logs")
	request.Sort("@timestamp", false)
	if in.Size != 0 {
		request.Size(int(in.Size))
	} else {
		request.Size(100)
	}

	masterQuery := elastic.NewBoolQuery()
	if in.Container != "" {
		boolQuery := elastic.NewBoolQuery()
		masterQuery.Filter(
			boolQuery.Should(elastic.NewPrefixQuery("container_id", in.Container)),
			boolQuery.Should(elastic.NewPrefixQuery("container_name", in.Container)),
		)
	}
	if in.Service != "" {
		boolQuery := elastic.NewBoolQuery()
		masterQuery.Filter(
			boolQuery.Should(elastic.NewPrefixQuery("service_id", in.Service)),
			boolQuery.Should(elastic.NewPrefixQuery("service_name", in.Service)),
		)
	}
	if in.Stack != "" {
		masterQuery.Filter(elastic.NewPrefixQuery("stack_name", in.Stack))
	}
	if in.Task != "" {
		masterQuery.Filter(elastic.NewPrefixQuery("task_id", in.Task))
	}
	if in.Node != "" {
		masterQuery.Filter(elastic.NewPrefixQuery("node_id", in.Node))
	}
	if in.Message != "" {
		queryString := elastic.NewSimpleQueryStringQuery(in.Message)
		queryString.Field("msg")
		masterQuery.Filter(queryString)
	}
	if !in.Infra {
		masterQuery.MustNot(elastic.NewTermQuery(convertLabelNameToESName(labels.LabelsNameRole), labels.LabelsValuesRoleInfrastructure))
		//For now Tools role is manage as Infrastructure role
		//Later: ToolsRole should be manage with a user premission (admin)
		masterQuery.MustNot(elastic.NewTermQuery(convertLabelNameToESName(labels.LabelsNameRole), labels.LabelsValuesRoleTools))
	}

	// Perform request
	searchResult, err := request.Query(masterQuery).Do(ctx)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%v", err)
	}

	// Build reply (from elasticsearch response)
	reply := GetReply{}
	reply.Entries = make([]*LogEntry, len(searchResult.Hits.Hits))
	for i, hit := range searchResult.Hits.Hits {
		entry := &LogEntry{}
		if err := s.unmarshal(*hit.Source, entry); err != nil {
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
		reply.Entries[i] = entry
	}

	// Reverse entries
	for i, j := 0, len(reply.Entries)-1; i < j; i, j = i+1, j-1 {
		reply.Entries[i], reply.Entries[j] = reply.Entries[j], reply.Entries[i]
	}
	log.Printf("rpc-logs: Get successful, returned %d entries\n", len(reply.Entries))
	return &reply, nil
}

//custom unmarshal for @timestamp
func (s *Server) unmarshal(data []byte, entry *LogEntry) error {
	type Alias LogEntry
	aux := &struct {
		TimestampTmp string `json:"@timestamp"`
		*Alias
	}{
		Alias: (*Alias)(entry),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	entry.Timestamp = aux.TimestampTmp
	return nil
}

// GetStream implements log.LogServer
func (s *Server) GetStream(in *GetRequest, stream Logs_GetStreamServer) error {
	if err := s.NS.Connect(); err != nil {
		return errors.New("unable to connect to nats service")
	}
	log.Println("rpc-logs: GetStream", in.String())

	sub, err := s.NS.GetClient().Subscribe(ns.LogsSubject, func(msg *stan.Msg) {
		entry := &LogEntry{}
		if err := proto.Unmarshal(msg.Data, entry); err != nil {
			return
		}
		if filter(entry, in) {
			stream.Send(entry)
		}
	})
	if err != nil {
		sub.Unsubscribe()
		return status.Errorf(codes.Internal, "%v", err)
	}
	for {
		select {
		case <-stream.Context().Done():
			sub.Unsubscribe()
			return stream.Context().Err()
		}
	}
}

func filter(entry *LogEntry, in *GetRequest) bool {
	match := true
	if in.Container != "" {
		containerID := strings.ToLower(entry.ContainerId)
		containerName := strings.ToLower(entry.ContainerName)
		match = strings.HasPrefix(containerID, strings.ToLower(in.Container)) || strings.HasPrefix(containerName, strings.ToLower(in.Container))
	}
	if in.Service != "" {
		serviceID := strings.ToLower(entry.ServiceId)
		serviceName := strings.ToLower(entry.ServiceName)
		match = strings.HasPrefix(serviceID, strings.ToLower(in.Service)) || strings.HasPrefix(serviceName, strings.ToLower(in.Service))
	}
	if in.Stack != "" {
		match = strings.HasPrefix(strings.ToLower(entry.StackName), strings.ToLower(in.Stack))
	}
	if in.Task != "" {
		match = strings.HasPrefix(strings.ToLower(entry.TaskId), strings.ToLower(in.Task))
	}
	if in.Node != "" {
		match = strings.HasPrefix(strings.ToLower(entry.NodeId), strings.ToLower(in.Node))
	}
	if in.Message != "" {
		match = strings.Contains(strings.ToLower(entry.Msg), strings.ToLower(in.Message))
	}
	if !in.Infra {
		role := entry.Labels[labels.LabelsNameRole]
		match = role != labels.LabelsValuesRoleInfrastructure && role != labels.LabelsValuesRoleTools
	}
	return match
}

func convertLabelNameToESName(name string) string {
	return "labels." + strings.Replace(name, ".", "-", -1)
}
