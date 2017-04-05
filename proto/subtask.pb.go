// Code generated by protoc-gen-go.
// source: subtask.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	subtask.proto

It has these top-level messages:
	SaveRequest
	SaveResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type SaveRequest struct {
	Operation string `protobuf:"bytes,1,opt,name=operation" json:"operation,omitempty"`
}

func (m *SaveRequest) Reset()                    { *m = SaveRequest{} }
func (m *SaveRequest) String() string            { return proto1.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()               {}
func (*SaveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SaveRequest) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

type SaveResponse struct {
	ExitCode string `protobuf:"bytes,1,opt,name=exitCode" json:"exitCode,omitempty"`
}

func (m *SaveResponse) Reset()                    { *m = SaveResponse{} }
func (m *SaveResponse) String() string            { return proto1.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()               {}
func (*SaveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SaveResponse) GetExitCode() string {
	if m != nil {
		return m.ExitCode
	}
	return ""
}

func init() {
	proto1.RegisterType((*SaveRequest)(nil), "proto.SaveRequest")
	proto1.RegisterType((*SaveResponse)(nil), "proto.SaveResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SubTask service

type SubTaskClient interface {
	Save(ctx context.Context, in *SaveRequest, opts ...client.CallOption) (*SaveResponse, error)
}

type subTaskClient struct {
	c           client.Client
	serviceName string
}

func NewSubTaskClient(serviceName string, c client.Client) SubTaskClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "proto"
	}
	return &subTaskClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *subTaskClient) Save(ctx context.Context, in *SaveRequest, opts ...client.CallOption) (*SaveResponse, error) {
	req := c.c.NewRequest(c.serviceName, "SubTask.Save", in)
	out := new(SaveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SubTask service

type SubTaskHandler interface {
	Save(context.Context, *SaveRequest, *SaveResponse) error
}

func RegisterSubTaskHandler(s server.Server, hdlr SubTaskHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&SubTask{hdlr}, opts...))
}

type SubTask struct {
	SubTaskHandler
}

func (h *SubTask) Save(ctx context.Context, in *SaveRequest, out *SaveResponse) error {
	return h.SubTaskHandler.Save(ctx, in, out)
}

func init() { proto1.RegisterFile("subtask.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0x4d, 0x2a,
	0x49, 0x2c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xda, 0x5c,
	0xdc, 0xc1, 0x89, 0x65, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x32, 0x5c, 0x9c,
	0xf9, 0x05, 0xa9, 0x45, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x08, 0x01, 0x25, 0x2d, 0x2e, 0x1e, 0x88, 0xe2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x29,
	0x2e, 0x8e, 0xd4, 0x8a, 0xcc, 0x12, 0xe7, 0xfc, 0x94, 0x54, 0xa8, 0x62, 0x38, 0xdf, 0xc8, 0x86,
	0x8b, 0x3d, 0xb8, 0x34, 0x29, 0x24, 0xb1, 0x38, 0x5b, 0xc8, 0x90, 0x8b, 0x05, 0xa4, 0x4d, 0x48,
	0x08, 0x62, 0xb5, 0x1e, 0x92, 0x85, 0x52, 0xc2, 0x28, 0x62, 0x10, 0x73, 0x95, 0x18, 0x92, 0xd8,
	0xc0, 0xa2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0x7b, 0xe1, 0x7b, 0xb5, 0x00, 0x00,
	0x00,
}