// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/storage/storage.proto
// DO NOT EDIT!

/*
Package storage is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/storage/storage.proto

It has these top-level messages:
	StorageEntry
	GetRequest
	GetReply
	PutRequest
	PutReply
	ListRequest
	ListReply
	DeleteRequest
	DeleteReply
*/
package storage

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

type StorageEntry struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
}

func (m *StorageEntry) Reset()                    { *m = StorageEntry{} }
func (m *StorageEntry) String() string            { return proto.CompactTextString(m) }
func (*StorageEntry) ProtoMessage()               {}
func (*StorageEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StorageEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StorageEntry) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type GetRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetReply struct {
	Entry *StorageEntry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetReply) GetEntry() *StorageEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type PutRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PutRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutRequest) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type PutReply struct {
	Entry *StorageEntry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
}

func (m *PutReply) Reset()                    { *m = PutReply{} }
func (m *PutReply) String() string            { return proto.CompactTextString(m) }
func (*PutReply) ProtoMessage()               {}
func (*PutReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PutReply) GetEntry() *StorageEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListReply struct {
	Entries []*StorageEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListReply) GetEntries() []*StorageEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type DeleteRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteReply struct {
	Entry *StorageEntry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
}

func (m *DeleteReply) Reset()                    { *m = DeleteReply{} }
func (m *DeleteReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteReply) ProtoMessage()               {}
func (*DeleteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeleteReply) GetEntry() *StorageEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func init() {
	proto.RegisterType((*StorageEntry)(nil), "storage.StorageEntry")
	proto.RegisterType((*GetRequest)(nil), "storage.GetRequest")
	proto.RegisterType((*GetReply)(nil), "storage.GetReply")
	proto.RegisterType((*PutRequest)(nil), "storage.PutRequest")
	proto.RegisterType((*PutReply)(nil), "storage.PutReply")
	proto.RegisterType((*ListRequest)(nil), "storage.ListRequest")
	proto.RegisterType((*ListReply)(nil), "storage.ListReply")
	proto.RegisterType((*DeleteRequest)(nil), "storage.DeleteRequest")
	proto.RegisterType((*DeleteReply)(nil), "storage.DeleteReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Storage service

type StorageClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error)
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/storage.Storage/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error) {
	out := new(PutReply)
	err := grpc.Invoke(ctx, "/storage.Storage/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/storage.Storage/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error) {
	out := new(DeleteReply)
	err := grpc.Invoke(ctx, "/storage.Storage/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Storage service

type StorageServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	Put(context.Context, *PutRequest) (*PutReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Delete(context.Context, *DeleteRequest) (*DeleteReply, error)
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Storage_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Storage_Put_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Storage_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Storage_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/storage/storage.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/storage/storage.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6e, 0xda, 0x40,
	0x14, 0xc6, 0x65, 0xdc, 0xf2, 0xe7, 0xf1, 0x47, 0xed, 0x94, 0xb6, 0xc8, 0x6a, 0x2b, 0x3a, 0xab,
	0x88, 0x48, 0x4c, 0x42, 0x16, 0x28, 0x28, 0x8b, 0x48, 0x49, 0xc4, 0x86, 0x05, 0x22, 0x27, 0x18,
	0xd0, 0x93, 0x63, 0x31, 0x78, 0x26, 0xf6, 0xd8, 0x92, 0x85, 0xd8, 0xe4, 0x0a, 0xb9, 0x46, 0x6e,
	0x93, 0x2b, 0xe4, 0x20, 0x91, 0x6d, 0x0c, 0x26, 0x84, 0x05, 0x2b, 0x7b, 0xbe, 0xf9, 0xde, 0xef,
	0xbd, 0xf7, 0xd9, 0x70, 0x69, 0x3b, 0xfa, 0x21, 0x98, 0x76, 0x67, 0x72, 0xc1, 0xb8, 0x52, 0x33,
	0x14, 0xe8, 0x71, 0x2d, 0x3d, 0xc6, 0x17, 0x8a, 0x71, 0xe5, 0x30, 0x4f, 0xcd, 0x98, 0xaf, 0xa5,
	0xc7, 0x6d, 0xcc, 0x9e, 0x5d, 0xe5, 0x49, 0x2d, 0x49, 0x69, 0x7d, 0xb4, 0xfe, 0xd8, 0x52, 0xda,
	0x02, 0x13, 0x3b, 0x77, 0x5d, 0xa9, 0xb9, 0x76, 0xa4, 0xeb, 0xa7, 0x36, 0xda, 0x83, 0xda, 0x7d,
	0x6a, 0xbc, 0x73, 0xb5, 0x17, 0x91, 0x6f, 0x60, 0xce, 0x31, 0x6a, 0x19, 0x6d, 0xe3, 0xa4, 0x32,
	0x89, 0x5f, 0x63, 0x25, 0xe4, 0xa2, 0x55, 0x48, 0x95, 0x90, 0x0b, 0xfa, 0x0f, 0x60, 0x88, 0x7a,
	0x82, 0x8f, 0x01, 0xfa, 0x7a, 0xbf, 0x82, 0xf6, 0xa1, 0x9c, 0xdc, 0x2b, 0x11, 0x91, 0x53, 0xf8,
	0x8a, 0x31, 0x38, 0xb9, 0xaf, 0xf6, 0x7e, 0x76, 0xb3, 0x29, 0xf3, 0x5d, 0x27, 0xa9, 0x87, 0x9e,
	0x01, 0x8c, 0x83, 0xc3, 0xe0, 0x4f, 0x46, 0xe9, 0x43, 0x39, 0xa9, 0x38, 0xba, 0x55, 0x1d, 0xaa,
	0x23, 0xc7, 0xcf, 0x7a, 0xd1, 0x2b, 0xa8, 0xa4, 0xc7, 0x18, 0xc4, 0xa0, 0x14, 0x9b, 0x1c, 0xf4,
	0x5b, 0x46, 0xdb, 0x3c, 0x8c, 0xca, 0x5c, 0xf4, 0x3f, 0xd4, 0x6f, 0x51, 0xa0, 0xc6, 0xc3, 0x99,
	0x0c, 0xa0, 0x9a, 0x59, 0x8e, 0x9d, 0xb5, 0xf7, 0x52, 0x80, 0xd2, 0x5a, 0x27, 0x37, 0x60, 0x0e,
	0x51, 0x93, 0x1f, 0x9b, 0x82, 0xed, 0x97, 0xb0, 0xbe, 0xef, 0x8a, 0x4a, 0x44, 0xb4, 0xf9, 0xf4,
	0xfa, 0xf6, 0x5c, 0x68, 0x90, 0x1a, 0x0b, 0xcf, 0xd9, 0x3c, 0x64, 0xcb, 0x39, 0x46, 0x2b, 0x32,
	0x02, 0x73, 0x1c, 0xe4, 0x21, 0xdb, 0xd4, 0x73, 0x90, 0x2c, 0x58, 0xfa, 0x37, 0x81, 0xfc, 0xa6,
	0x24, 0x0f, 0x61, 0xcb, 0x90, 0x8b, 0xd5, 0xc0, 0xe8, 0x90, 0x6b, 0xf8, 0x12, 0x67, 0x47, 0x9a,
	0x9b, 0xca, 0x5c, 0xb2, 0x16, 0xf9, 0xa0, 0xc6, 0xc0, 0x46, 0x02, 0x2c, 0x93, 0x62, 0x0a, 0x24,
	0x23, 0x28, 0xa6, 0xe1, 0x90, 0x5f, 0x1b, 0xf7, 0x4e, 0xa0, 0x56, 0x73, 0x4f, 0xcf, 0x6d, 0xd7,
	0xd9, 0xd9, 0x6e, 0x5a, 0x4c, 0xfe, 0xec, 0x8b, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x60,
	0x3c, 0x2c, 0x3d, 0x03, 0x00, 0x00,
}
