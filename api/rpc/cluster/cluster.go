package cluster

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"

	"github.com/appcelerator/amp/data/accounts"
	"github.com/appcelerator/amp/data/clusters"
	"github.com/appcelerator/amp/pkg/docker"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/docker/docker/api/types"

	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
)

// Server is used to implement cluster.ClusterServer
type Server struct {
	Accounts accounts.Interface
	Docker   *docker.Docker
	Clusters clusters.Interface
}

func convertError(err error) error {
	switch err {
	case clusters.InvalidName:
		return status.Errorf(codes.InvalidArgument, err.Error())
	case clusters.InvalidProvider:
		return status.Errorf(codes.InvalidArgument, err.Error())
	case clusters.ClusterAlreadyExists:
		return status.Errorf(codes.AlreadyExists, err.Error())
	case clusters.ClusterNotFound:
		return status.Errorf(codes.NotFound, err.Error())
	case accounts.NotAuthorized:
		return status.Errorf(codes.PermissionDenied, err.Error())
	}
	return status.Errorf(codes.Internal, err.Error())
}

func (s *Server) createOnAws(ctx context.Context, in *CreateRequest) (string, error) {
	stackName := fmt.Sprintf("amp-%s", in.OrganizationName)
	// TODO: get the credentials from the organization
	// TODO: get the default keypair from the organization
	sess := session.Must(session.NewSession())

	svc := cloudformation.New(sess)

	params := &cloudformation.CreateStackInput{
		StackName: aws.String(stackName),
		Capabilities: []*string{
			aws.String("CAPABILITY_IAM"),
		},
		OnFailure: aws.String("DELETE"),
		Parameters: []*cloudformation.Parameter{
			{
				ParameterKey:     aws.String("Organization"),
				ParameterValue:   aws.String(in.OrganizationName),
				UsePreviousValue: aws.Bool(true),
			},
		},
		Tags: []*cloudformation.Tag{
			{
				Key:   aws.String("Organization"),
				Value: aws.String(in.OrganizationName),
			},
		},
		TemplateBody:     aws.String("TemplateBody"),
		TemplateURL:      aws.String("TemplateURL"),
		TimeoutInMinutes: aws.Int64(10),
	}
	resp, err := svc.CreateStack(params)

	if err != nil {
		return "", err
	}
	stackId := awsutil.StringValue(resp)
	log.Printf("Stack ID = %s\n", stackId)

	// Pretty-print the response data.
	log.Println(resp)
	return stackId, nil
}

// Validation of the inputs of a create request
func (s *Server) validateCreateInputs(in *CreateRequest) error {
	return nil
}

// Create implements cluster.Server
func (s *Server) Create(ctx context.Context, in *CreateRequest) (*CreateReply, error) {
	log.Println("[cluster] Create", in.String())

	if err := s.validateCreateInputs(in); err != nil {
		return nil, convertError(err)
	}
	if c, _ := s.Clusters.GetClusterByName(ctx, in.Name, in.OrganizationName); c != nil {
		return nil, convertError(clusters.ClusterAlreadyExists)
	}
	// TODO: check authorization to create clusters

	log.Println("cluster creation request: %s on %s for %s", in.Name, in.Provider, in.OrganizationName)
	var id string
	var err error
	switch in.Provider {
	// hardcoded, to avoid circular imports
	case clusters.CloudProvider_AWS:
		id, err = s.createOnAws(ctx, in)
		if err != nil {
			return nil, convertError(err)
		}
	default:
		return nil, convertError(clusters.InvalidProvider)
	}

	log.Println("[cluster] Success: created cluster")
	return &CreateReply{Id: id}, nil
}

// List implements cluster.Server
func (s *Server) List(ctx context.Context, in *ListRequest) (*ListReply, error) {
	log.Println("[cluster] List", in.String())

	log.Println("[cluster] Success: list")
	return &ListReply{}, nil
}

// Status implements cluster.Server
func (s *Server) Status(ctx context.Context, in *StatusRequest) (*StatusReply, error) {
	log.Println("[cluster] Status", in.String())

	log.Println("[cluster] Success: list")
	return &StatusReply{}, nil
}

// Update implements cluster.Server
func (s *Server) Update(ctx context.Context, in *UpdateRequest) (*UpdateReply, error) {
	log.Println("[cluster] Update", in.String())

	log.Println("[cluster] Success: list")
	return &UpdateReply{}, nil
}

// Remove implements cluster.Server
func (s *Server) Remove(ctx context.Context, in *RemoveRequest) (*RemoveReply, error) {
	log.Println("[cluster] Remove", in.String())

	log.Println("[cluster] Success: removed", in.Id)
	return &RemoveReply{}, nil
}

// NodeList get cluster node list
func (s *Server) NodeList(ctx context.Context, in *NodeListRequest) (*NodeListReply, error) {
	log.Println("[cluster] NodeList", in.String())

	list, err := s.Docker.GetClient().NodeList(ctx, types.NodeListOptions{})
	if err != nil {
		return nil, convertError(err)
	}
	ret := &NodeListReply{}
	for _, node := range list {
		leader := false
		if node.ManagerStatus != nil {
			leader = node.ManagerStatus.Leader
		}
		ret.Nodes = append(ret.Nodes, &NodeReply{
			Id:            node.ID,
			Hostname:      node.Description.Hostname,
			Status:        string(node.Status.State),
			Availability:  string(node.Spec.Availability),
			Role:          string(node.Spec.Role),
			ManagerLeader: leader,
		})
	}
	return ret, nil
}
