// Code generated by protoc-gen-go. DO NOT EDIT.
// source: call_aggregate.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

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

type CallAggregateRequest struct {
	// 聚合id
	AggregateId string `protobuf:"bytes,1,opt,name=aggregateId,proto3" json:"aggregateId,omitempty"`
	// 聚合类型
	AggregateType string `protobuf:"bytes,2,opt,name=aggregateType,proto3" json:"aggregateType,omitempty"`
	// 聚合
	Aggregate []byte `protobuf:"bytes,3,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
	// 需要应用的命令
	Command              *Command `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallAggregateRequest) Reset()         { *m = CallAggregateRequest{} }
func (m *CallAggregateRequest) String() string { return proto.CompactTextString(m) }
func (*CallAggregateRequest) ProtoMessage()    {}
func (*CallAggregateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_aggregate_a92e186390522670, []int{0}
}
func (m *CallAggregateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallAggregateRequest.Unmarshal(m, b)
}
func (m *CallAggregateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallAggregateRequest.Marshal(b, m, deterministic)
}
func (dst *CallAggregateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallAggregateRequest.Merge(dst, src)
}
func (m *CallAggregateRequest) XXX_Size() int {
	return xxx_messageInfo_CallAggregateRequest.Size(m)
}
func (m *CallAggregateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallAggregateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallAggregateRequest proto.InternalMessageInfo

func (m *CallAggregateRequest) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *CallAggregateRequest) GetAggregateType() string {
	if m != nil {
		return m.AggregateType
	}
	return ""
}

func (m *CallAggregateRequest) GetAggregate() []byte {
	if m != nil {
		return m.Aggregate
	}
	return nil
}

func (m *CallAggregateRequest) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

// /返回
type CallAggregateResponse struct {
	// 事件
	Events               []*CallAggregateResponseEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CallAggregateResponse) Reset()         { *m = CallAggregateResponse{} }
func (m *CallAggregateResponse) String() string { return proto.CompactTextString(m) }
func (*CallAggregateResponse) ProtoMessage()    {}
func (*CallAggregateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_aggregate_a92e186390522670, []int{1}
}
func (m *CallAggregateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallAggregateResponse.Unmarshal(m, b)
}
func (m *CallAggregateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallAggregateResponse.Marshal(b, m, deterministic)
}
func (dst *CallAggregateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallAggregateResponse.Merge(dst, src)
}
func (m *CallAggregateResponse) XXX_Size() int {
	return xxx_messageInfo_CallAggregateResponse.Size(m)
}
func (m *CallAggregateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallAggregateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallAggregateResponse proto.InternalMessageInfo

func (m *CallAggregateResponse) GetEvents() []*CallAggregateResponseEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

// 事件的详细信息
type CallAggregateResponseEvent struct {
	// 事件id
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 事件内容
	Data                 *any.Any `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallAggregateResponseEvent) Reset()         { *m = CallAggregateResponseEvent{} }
func (m *CallAggregateResponseEvent) String() string { return proto.CompactTextString(m) }
func (*CallAggregateResponseEvent) ProtoMessage()    {}
func (*CallAggregateResponseEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_aggregate_a92e186390522670, []int{2}
}
func (m *CallAggregateResponseEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallAggregateResponseEvent.Unmarshal(m, b)
}
func (m *CallAggregateResponseEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallAggregateResponseEvent.Marshal(b, m, deterministic)
}
func (dst *CallAggregateResponseEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallAggregateResponseEvent.Merge(dst, src)
}
func (m *CallAggregateResponseEvent) XXX_Size() int {
	return xxx_messageInfo_CallAggregateResponseEvent.Size(m)
}
func (m *CallAggregateResponseEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CallAggregateResponseEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CallAggregateResponseEvent proto.InternalMessageInfo

func (m *CallAggregateResponseEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CallAggregateResponseEvent) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

// 命令
type Command struct {
	// 命令类型
	CmdType string `protobuf:"bytes,2,opt,name=cmdType,proto3" json:"cmdType,omitempty"`
	// 事件内容
	Data                 []byte   `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_aggregate_a92e186390522670, []int{3}
}
func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (dst *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(dst, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetCmdType() string {
	if m != nil {
		return m.CmdType
	}
	return ""
}

func (m *Command) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CallAggregateRequest)(nil), "rpc.CallAggregateRequest")
	proto.RegisterType((*CallAggregateResponse)(nil), "rpc.CallAggregateResponse")
	proto.RegisterType((*CallAggregateResponseEvent)(nil), "rpc.CallAggregateResponseEvent")
	proto.RegisterType((*Command)(nil), "rpc.Command")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CallAggregateClient is the client API for CallAggregate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CallAggregateClient interface {
	// 调用聚合处理事件
	CallAggregate(ctx context.Context, in *CallAggregateRequest, opts ...grpc.CallOption) (*CallAggregateResponse, error)
}

type callAggregateClient struct {
	cc *grpc.ClientConn
}

func NewCallAggregateClient(cc *grpc.ClientConn) CallAggregateClient {
	return &callAggregateClient{cc}
}

func (c *callAggregateClient) CallAggregate(ctx context.Context, in *CallAggregateRequest, opts ...grpc.CallOption) (*CallAggregateResponse, error) {
	out := new(CallAggregateResponse)
	err := c.cc.Invoke(ctx, "/rpc.CallAggregate/CallAggregate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallAggregateServer is the server API for CallAggregate service.
type CallAggregateServer interface {
	// 调用聚合处理事件
	CallAggregate(context.Context, *CallAggregateRequest) (*CallAggregateResponse, error)
}

func RegisterCallAggregateServer(s *grpc.Server, srv CallAggregateServer) {
	s.RegisterService(&_CallAggregate_serviceDesc, srv)
}

func _CallAggregate_CallAggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallAggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallAggregateServer).CallAggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CallAggregate/CallAggregate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallAggregateServer).CallAggregate(ctx, req.(*CallAggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CallAggregate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.CallAggregate",
	HandlerType: (*CallAggregateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallAggregate",
			Handler:    _CallAggregate_CallAggregate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "call_aggregate.proto",
}

func init() {
	proto.RegisterFile("call_aggregate.proto", fileDescriptor_call_aggregate_a92e186390522670)
}

var fileDescriptor_call_aggregate_a92e186390522670 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6a, 0x83, 0x40,
	0x10, 0x87, 0xbb, 0x1a, 0x22, 0x19, 0x4d, 0x0f, 0x83, 0x85, 0x8d, 0x14, 0x2a, 0x52, 0x8a, 0xa7,
	0x0d, 0xd8, 0x43, 0xce, 0x21, 0x2d, 0x34, 0xb7, 0xb2, 0x94, 0x42, 0x4f, 0x65, 0xa3, 0x5b, 0x2f,
	0xea, 0x5a, 0x35, 0x05, 0xdf, 0xa7, 0x0f, 0x5a, 0xea, 0xbf, 0x98, 0x90, 0x1c, 0xf7, 0x9b, 0x6f,
	0x86, 0xdf, 0xcc, 0x82, 0x1d, 0x8a, 0x24, 0xf9, 0x14, 0x71, 0x5c, 0xc8, 0x58, 0x54, 0x92, 0xe5,
	0x85, 0xaa, 0x14, 0xea, 0x45, 0x1e, 0x3a, 0x8b, 0x58, 0xa9, 0x38, 0x91, 0xcb, 0x06, 0xed, 0xf6,
	0x5f, 0x4b, 0x91, 0xd5, 0x6d, 0xdd, 0xfb, 0x25, 0x60, 0x6f, 0x44, 0x92, 0xac, 0xfb, 0x3e, 0x2e,
	0xbf, 0xf7, 0xb2, 0xac, 0xd0, 0x05, 0x73, 0x98, 0xb5, 0x8d, 0x28, 0x71, 0x89, 0x3f, 0xe3, 0x63,
	0x84, 0xf7, 0x30, 0x1f, 0x9e, 0x6f, 0x75, 0x2e, 0xa9, 0xd6, 0x38, 0xc7, 0x10, 0x6f, 0x61, 0x36,
	0x00, 0xaa, 0xbb, 0xc4, 0xb7, 0xf8, 0x01, 0xe0, 0x03, 0x18, 0xa1, 0x4a, 0x53, 0x91, 0x45, 0x74,
	0xe2, 0x12, 0xdf, 0x0c, 0x2c, 0x56, 0xe4, 0x21, 0xdb, 0xb4, 0x8c, 0xf7, 0x45, 0xef, 0x15, 0x6e,
	0x4e, 0x52, 0x96, 0xb9, 0xca, 0x4a, 0x89, 0x2b, 0x98, 0xca, 0x1f, 0x99, 0x55, 0x25, 0x25, 0xae,
	0xee, 0x9b, 0xc1, 0x5d, 0xdb, 0x7f, 0xce, 0x7d, 0xfe, 0xf7, 0x78, 0xa7, 0x7b, 0xef, 0xe0, 0x5c,
	0xb6, 0xf0, 0x1a, 0xb4, 0x61, 0x69, 0x6d, 0x1b, 0xa1, 0x0f, 0x93, 0x27, 0x51, 0x89, 0x2e, 0xa4,
	0xcd, 0xda, 0x83, 0xb2, 0xfe, 0xa0, 0x6c, 0x9d, 0xd5, 0xbc, 0x31, 0xbc, 0x15, 0x18, 0x5d, 0x7a,
	0xa4, 0x60, 0x84, 0x69, 0x34, 0x3a, 0x4d, 0xff, 0x44, 0x1c, 0x8d, 0xb3, 0xda, 0xc6, 0xe0, 0x03,
	0xe6, 0x47, 0x81, 0xf0, 0xe5, 0x14, 0x2c, 0xce, 0xed, 0xd6, 0xfc, 0x96, 0xe3, 0x5c, 0x5e, 0xdb,
	0xbb, 0xda, 0x4d, 0x9b, 0x9c, 0x8f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x88, 0x6c, 0x41,
	0x23, 0x02, 0x00, 0x00,
}
