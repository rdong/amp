// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/logs/logs.proto
// DO NOT EDIT!

/*
Package logs is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/logs/logs.proto

It has these top-level messages:
	LogEntry
	GetRequest
	GetReply
*/
package logs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogEntry struct {
	Timestamp          string            `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	ContainerId        string            `protobuf:"bytes,2,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	ContainerName      string            `protobuf:"bytes,3,opt,name=container_name,json=containerName" json:"container_name,omitempty"`
	ContainerShortName string            `protobuf:"bytes,4,opt,name=container_short_name,json=containerShortName" json:"container_short_name,omitempty"`
	ContainerState     string            `protobuf:"bytes,5,opt,name=container_state,json=containerState" json:"container_state,omitempty"`
	ServiceName        string            `protobuf:"bytes,6,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	ServiceId          string            `protobuf:"bytes,7,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
	TaskId             string            `protobuf:"bytes,8,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	StackName          string            `protobuf:"bytes,9,opt,name=stack_name,json=stackName" json:"stack_name,omitempty"`
	NodeId             string            `protobuf:"bytes,10,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Labels             map[string]string `protobuf:"bytes,11,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Msg                string            `protobuf:"bytes,12,opt,name=msg" json:"msg,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogEntry) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *LogEntry) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *LogEntry) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *LogEntry) GetContainerShortName() string {
	if m != nil {
		return m.ContainerShortName
	}
	return ""
}

func (m *LogEntry) GetContainerState() string {
	if m != nil {
		return m.ContainerState
	}
	return ""
}

func (m *LogEntry) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *LogEntry) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *LogEntry) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *LogEntry) GetStackName() string {
	if m != nil {
		return m.StackName
	}
	return ""
}

func (m *LogEntry) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *LogEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LogEntry) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GetRequest struct {
	Container string `protobuf:"bytes,1,opt,name=container" json:"container,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Node      string `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
	Size      int64  `protobuf:"zigzag64,4,opt,name=size" json:"size,omitempty"`
	Service   string `protobuf:"bytes,5,opt,name=service" json:"service,omitempty"`
	Stack     string `protobuf:"bytes,6,opt,name=stack" json:"stack,omitempty"`
	Task      string `protobuf:"bytes,7,opt,name=task" json:"task,omitempty"`
	Infra     bool   `protobuf:"varint,8,opt,name=infra" json:"infra,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetRequest) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *GetRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *GetRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *GetRequest) GetStack() string {
	if m != nil {
		return m.Stack
	}
	return ""
}

func (m *GetRequest) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *GetRequest) GetInfra() bool {
	if m != nil {
		return m.Infra
	}
	return false
}

type GetReply struct {
	Entries []*LogEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetReply) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "logs.LogEntry")
	proto.RegisterType((*GetRequest)(nil), "logs.GetRequest")
	proto.RegisterType((*GetReply)(nil), "logs.GetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Logs service

type LogsClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	GetStream(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (Logs_GetStreamClient, error)
}

type logsClient struct {
	cc *grpc.ClientConn
}

func NewLogsClient(cc *grpc.ClientConn) LogsClient {
	return &logsClient{cc}
}

func (c *logsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/logs.Logs/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsClient) GetStream(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (Logs_GetStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Logs_serviceDesc.Streams[0], c.cc, "/logs.Logs/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &logsGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Logs_GetStreamClient interface {
	Recv() (*LogEntry, error)
	grpc.ClientStream
}

type logsGetStreamClient struct {
	grpc.ClientStream
}

func (x *logsGetStreamClient) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Logs service

type LogsServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	GetStream(*GetRequest, Logs_GetStreamServer) error
}

func RegisterLogsServer(s *grpc.Server, srv LogsServer) {
	s.RegisterService(&_Logs_serviceDesc, srv)
}

func _Logs_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logs.Logs/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logs_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogsServer).GetStream(m, &logsGetStreamServer{stream})
}

type Logs_GetStreamServer interface {
	Send(*LogEntry) error
	grpc.ServerStream
}

type logsGetStreamServer struct {
	grpc.ServerStream
}

func (x *logsGetStreamServer) Send(m *LogEntry) error {
	return x.ServerStream.SendMsg(m)
}

var _Logs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logs.Logs",
	HandlerType: (*LogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Logs_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Logs_GetStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/appcelerator/amp/api/rpc/logs/logs.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/logs/logs.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0xda, 0x6e, 0x7f, 0xa6, 0x4b, 0xb7, 0x58, 0x2b, 0x35, 0xaa, 0x16, 0x69, 0xa9, 0x84,
	0xe8, 0xa9, 0x59, 0xba, 0x1c, 0x58, 0xee, 0xa8, 0x54, 0xaa, 0x38, 0xa4, 0x0f, 0x80, 0xdc, 0x64,
	0xc8, 0x5a, 0x4d, 0xe2, 0x60, 0xbb, 0x95, 0xca, 0x91, 0x2b, 0x47, 0xde, 0x89, 0x17, 0xe0, 0x09,
	0x90, 0x78, 0x10, 0xe4, 0xb1, 0xd3, 0x96, 0xc3, 0x5e, 0xa2, 0x99, 0x6f, 0xbe, 0xf9, 0x62, 0x7f,
	0x9f, 0x0c, 0xf7, 0x99, 0x30, 0x8f, 0xbb, 0xcd, 0x2c, 0x91, 0x45, 0xc4, 0xab, 0x2a, 0xc1, 0x1c,
	0x15, 0x37, 0x52, 0x45, 0xbc, 0xa8, 0x22, 0x5e, 0x89, 0x48, 0x55, 0x49, 0x94, 0xcb, 0x4c, 0xd3,
	0x67, 0x56, 0x29, 0x69, 0x24, 0x6b, 0xd9, 0x7a, 0x7c, 0x93, 0x49, 0x99, 0xe5, 0x48, 0x2c, 0x5e,
	0x96, 0xd2, 0x70, 0x23, 0x64, 0xe9, 0x39, 0x93, 0x3f, 0x4d, 0xe8, 0xae, 0x64, 0xf6, 0xa1, 0x34,
	0xea, 0xc0, 0x6e, 0xa0, 0x67, 0x44, 0x81, 0xda, 0xf0, 0xa2, 0x0a, 0x83, 0xdb, 0x60, 0xda, 0x8b,
	0x4f, 0x00, 0x7b, 0x09, 0x97, 0x89, 0x2c, 0x0d, 0x17, 0x25, 0xaa, 0xcf, 0x22, 0x0d, 0x1b, 0x44,
	0xe8, 0x1f, 0xb1, 0x65, 0xca, 0x5e, 0xc1, 0xe0, 0x44, 0x29, 0x79, 0x81, 0x61, 0x93, 0x48, 0xcf,
	0x8e, 0xe8, 0x27, 0x5e, 0x20, 0xbb, 0x83, 0xeb, 0x13, 0x4d, 0x3f, 0x4a, 0x65, 0x1c, 0xb9, 0x45,
	0x64, 0x76, 0x9c, 0xad, 0xed, 0x88, 0x36, 0x5e, 0xc3, 0xd5, 0xd9, 0x86, 0xe1, 0x06, 0xc3, 0x0b,
	0x22, 0x9f, 0xfe, 0xb7, 0xb6, 0xa8, 0x3d, 0xa4, 0x46, 0xb5, 0x17, 0x09, 0x3a, 0xc9, 0xb6, 0x3b,
	0xa4, 0xc7, 0x48, 0xeb, 0x05, 0x40, 0x4d, 0x11, 0x69, 0xd8, 0x71, 0xd7, 0xf4, 0xc8, 0x32, 0x65,
	0x23, 0xe8, 0x18, 0xae, 0xb7, 0x76, 0xd6, 0xa5, 0x59, 0xdb, 0xb6, 0xcb, 0x94, 0xf6, 0x0c, 0x4f,
	0xb6, 0x4e, 0xb8, 0xe7, 0xf7, 0x2c, 0x42, 0xb2, 0x23, 0xe8, 0x94, 0x32, 0x25, 0x4d, 0x70, 0x7b,
	0xb6, 0x5d, 0xa6, 0x6c, 0x0e, 0xed, 0x9c, 0x6f, 0x30, 0xd7, 0x61, 0xff, 0xb6, 0x39, 0xed, 0xcf,
	0xc7, 0x33, 0xca, 0xa8, 0x76, 0x7d, 0xb6, 0xa2, 0x21, 0xd5, 0xb1, 0x67, 0xb2, 0x21, 0x34, 0x0b,
	0x9d, 0x85, 0x97, 0x24, 0x64, 0xcb, 0xf1, 0x03, 0xf4, 0xcf, 0x88, 0x96, 0xb0, 0xc5, 0x83, 0x0f,
	0xc9, 0x96, 0xec, 0x1a, 0x2e, 0xf6, 0x3c, 0xdf, 0xa1, 0xcf, 0xc5, 0x35, 0xef, 0x1b, 0xef, 0x82,
	0xc9, 0xaf, 0x00, 0x60, 0x81, 0x26, 0xc6, 0xaf, 0x3b, 0xd4, 0xc6, 0xa6, 0x7c, 0x34, 0xad, 0x4e,
	0xf9, 0x08, 0xb0, 0x10, 0x3a, 0x05, 0x6a, 0xcd, 0xb3, 0x5a, 0xa8, 0x6e, 0x19, 0x83, 0x96, 0xbd,
	0x91, 0x8f, 0x94, 0x6a, 0x8b, 0x69, 0xf1, 0xcd, 0x25, 0xc7, 0x62, 0xaa, 0xad, 0x82, 0x77, 0xd3,
	0x67, 0x54, 0xb7, 0xf6, 0x88, 0xe4, 0x97, 0x4f, 0xc5, 0x35, 0x56, 0xc3, 0x3a, 0xec, 0x93, 0xa0,
	0xda, 0x32, 0x45, 0xf9, 0x45, 0x71, 0x8a, 0xa0, 0x1b, 0xbb, 0x66, 0xf2, 0x16, 0xba, 0x74, 0x8f,
	0x2a, 0x3f, 0xb0, 0x29, 0x74, 0xb0, 0x34, 0x4a, 0xa0, 0x0e, 0x03, 0xb2, 0x75, 0xf0, 0xbf, 0xad,
	0x71, 0x3d, 0x9e, 0xff, 0x08, 0xa0, 0xb5, 0x92, 0x99, 0x66, 0x0f, 0xd0, 0x5c, 0xa0, 0x61, 0x43,
	0x47, 0x3c, 0x39, 0x32, 0x1e, 0x9c, 0x21, 0x55, 0x7e, 0x98, 0x0c, 0xbf, 0xff, 0xfe, 0xfb, 0xb3,
	0x01, 0xac, 0x1b, 0xed, 0xdf, 0xd0, 0x83, 0x62, 0x1f, 0xa1, 0xb7, 0x40, 0xb3, 0x36, 0x0a, 0x79,
	0xf1, 0xb4, 0x40, 0xfd, 0xef, 0xc9, 0x88, 0x04, 0x9e, 0xb3, 0xab, 0x5a, 0x20, 0xd2, 0xb4, 0x7a,
	0x17, 0x6c, 0xda, 0xf4, 0xee, 0xee, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x29, 0x2d, 0x2d, 0x37,
	0xd2, 0x03, 0x00, 0x00,
}
