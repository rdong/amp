// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/tests/integration/data/storage/etcd/store_test.proto
// DO NOT EDIT!

/*
Package etcd is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/tests/integration/data/storage/etcd/store_test.proto

It has these top-level messages:
	TestMessage
*/
package etcd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestMessage struct {
	Id    string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Names []string `protobuf:"bytes,3,rep,name=names" json:"names,omitempty"`
}

func (m *TestMessage) Reset()                    { *m = TestMessage{} }
func (m *TestMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()               {}
func (*TestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TestMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestMessage) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*TestMessage)(nil), "etcd.TestMessage")
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/tests/integration/data/storage/etcd/store_test.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0x8c, 0xbb, 0xaa, 0xc3, 0x30,
	0x10, 0x44, 0xf1, 0xe3, 0x5e, 0xb0, 0x02, 0x29, 0x44, 0x0a, 0x97, 0x26, 0x95, 0x2b, 0x6f, 0x91,
	0x8f, 0x48, 0x15, 0x08, 0x26, 0x7d, 0x58, 0x5b, 0x8b, 0x22, 0x88, 0xbd, 0x42, 0xbb, 0xf9, 0xff,
	0x20, 0xa5, 0x9a, 0x39, 0x87, 0x61, 0xcc, 0xdd, 0x07, 0x7d, 0x7d, 0x96, 0x69, 0xe5, 0x0d, 0x30,
	0xc6, 0x95, 0xde, 0x94, 0x50, 0x39, 0x01, 0x6e, 0x11, 0x94, 0x44, 0x05, 0xc2, 0xae, 0xe4, 0x13,
	0x6a, 0xe0, 0x1d, 0x1c, 0x2a, 0x82, 0x28, 0x27, 0xf4, 0x04, 0xa4, 0xab, 0x2b, 0x40, 0xcf, 0xbc,
	0x9c, 0x62, 0x62, 0x65, 0xdb, 0x66, 0x7d, 0xbe, 0x9a, 0xc3, 0x83, 0x44, 0x6f, 0x24, 0x82, 0x9e,
	0xec, 0xd1, 0xd4, 0xc1, 0xf5, 0xd5, 0x50, 0x8d, 0xdd, 0x5c, 0x07, 0x67, 0xad, 0x69, 0x77, 0xdc,
	0xa8, 0xaf, 0x8b, 0x29, 0xdd, 0x9e, 0xcc, 0x5f, 0x4e, 0xe9, 0x9b, 0xa1, 0x19, 0xbb, 0xf9, 0x07,
	0xcb, 0x7f, 0x79, 0xbd, 0x7c, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x91, 0x14, 0x26, 0xa9, 0x00,
	0x00, 0x00,
}
