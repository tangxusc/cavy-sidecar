// Code generated by protoc-gen-go. DO NOT EDIT.
// source: call_event_handler.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type CallEventHandlerRequest struct {
	Events               []*CallEventHandlerRequestEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *CallEventHandlerRequest) Reset()         { *m = CallEventHandlerRequest{} }
func (m *CallEventHandlerRequest) String() string { return proto.CompactTextString(m) }
func (*CallEventHandlerRequest) ProtoMessage()    {}
func (*CallEventHandlerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_event_handler_5400891be3a99dc9, []int{0}
}
func (m *CallEventHandlerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallEventHandlerRequest.Unmarshal(m, b)
}
func (m *CallEventHandlerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallEventHandlerRequest.Marshal(b, m, deterministic)
}
func (dst *CallEventHandlerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallEventHandlerRequest.Merge(dst, src)
}
func (m *CallEventHandlerRequest) XXX_Size() int {
	return xxx_messageInfo_CallEventHandlerRequest.Size(m)
}
func (m *CallEventHandlerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallEventHandlerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallEventHandlerRequest proto.InternalMessageInfo

func (m *CallEventHandlerRequest) GetEvents() []*CallEventHandlerRequestEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type CallEventHandlerResponse struct {
	Data                 []*CallEventHandlerResponseMark `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *CallEventHandlerResponse) Reset()         { *m = CallEventHandlerResponse{} }
func (m *CallEventHandlerResponse) String() string { return proto.CompactTextString(m) }
func (*CallEventHandlerResponse) ProtoMessage()    {}
func (*CallEventHandlerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_event_handler_5400891be3a99dc9, []int{1}
}
func (m *CallEventHandlerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallEventHandlerResponse.Unmarshal(m, b)
}
func (m *CallEventHandlerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallEventHandlerResponse.Marshal(b, m, deterministic)
}
func (dst *CallEventHandlerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallEventHandlerResponse.Merge(dst, src)
}
func (m *CallEventHandlerResponse) XXX_Size() int {
	return xxx_messageInfo_CallEventHandlerResponse.Size(m)
}
func (m *CallEventHandlerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallEventHandlerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallEventHandlerResponse proto.InternalMessageInfo

func (m *CallEventHandlerResponse) GetData() []*CallEventHandlerResponseMark {
	if m != nil {
		return m.Data
	}
	return nil
}

// 事件的详细信息
type CallEventHandlerRequestEvent struct {
	// 事件id
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventType            string               `protobuf:"bytes,2,opt,name=eventType,proto3" json:"eventType,omitempty"`
	AggId                string               `protobuf:"bytes,3,opt,name=aggId,proto3" json:"aggId,omitempty"`
	AggType              string               `protobuf:"bytes,4,opt,name=aggType,proto3" json:"aggType,omitempty"`
	Create               *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create,proto3" json:"create,omitempty"`
	Data                 []byte               `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CallEventHandlerRequestEvent) Reset()         { *m = CallEventHandlerRequestEvent{} }
func (m *CallEventHandlerRequestEvent) String() string { return proto.CompactTextString(m) }
func (*CallEventHandlerRequestEvent) ProtoMessage()    {}
func (*CallEventHandlerRequestEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_event_handler_5400891be3a99dc9, []int{2}
}
func (m *CallEventHandlerRequestEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallEventHandlerRequestEvent.Unmarshal(m, b)
}
func (m *CallEventHandlerRequestEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallEventHandlerRequestEvent.Marshal(b, m, deterministic)
}
func (dst *CallEventHandlerRequestEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallEventHandlerRequestEvent.Merge(dst, src)
}
func (m *CallEventHandlerRequestEvent) XXX_Size() int {
	return xxx_messageInfo_CallEventHandlerRequestEvent.Size(m)
}
func (m *CallEventHandlerRequestEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CallEventHandlerRequestEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CallEventHandlerRequestEvent proto.InternalMessageInfo

func (m *CallEventHandlerRequestEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CallEventHandlerRequestEvent) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *CallEventHandlerRequestEvent) GetAggId() string {
	if m != nil {
		return m.AggId
	}
	return ""
}

func (m *CallEventHandlerRequestEvent) GetAggType() string {
	if m != nil {
		return m.AggType
	}
	return ""
}

func (m *CallEventHandlerRequestEvent) GetCreate() *timestamp.Timestamp {
	if m != nil {
		return m.Create
	}
	return nil
}

func (m *CallEventHandlerRequestEvent) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CallEventHandlerResponseMark struct {
	// 事件id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallEventHandlerResponseMark) Reset()         { *m = CallEventHandlerResponseMark{} }
func (m *CallEventHandlerResponseMark) String() string { return proto.CompactTextString(m) }
func (*CallEventHandlerResponseMark) ProtoMessage()    {}
func (*CallEventHandlerResponseMark) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_event_handler_5400891be3a99dc9, []int{3}
}
func (m *CallEventHandlerResponseMark) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallEventHandlerResponseMark.Unmarshal(m, b)
}
func (m *CallEventHandlerResponseMark) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallEventHandlerResponseMark.Marshal(b, m, deterministic)
}
func (dst *CallEventHandlerResponseMark) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallEventHandlerResponseMark.Merge(dst, src)
}
func (m *CallEventHandlerResponseMark) XXX_Size() int {
	return xxx_messageInfo_CallEventHandlerResponseMark.Size(m)
}
func (m *CallEventHandlerResponseMark) XXX_DiscardUnknown() {
	xxx_messageInfo_CallEventHandlerResponseMark.DiscardUnknown(m)
}

var xxx_messageInfo_CallEventHandlerResponseMark proto.InternalMessageInfo

func (m *CallEventHandlerResponseMark) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CallEventHandlerResponseMark) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*CallEventHandlerRequest)(nil), "rpc.CallEventHandlerRequest")
	proto.RegisterType((*CallEventHandlerResponse)(nil), "rpc.CallEventHandlerResponse")
	proto.RegisterType((*CallEventHandlerRequestEvent)(nil), "rpc.CallEventHandlerRequestEvent")
	proto.RegisterType((*CallEventHandlerResponseMark)(nil), "rpc.CallEventHandlerResponseMark")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CallEventHandlerClient is the client API for CallEventHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CallEventHandlerClient interface {
	// 调用聚合处理事件
	CallEventHandler(ctx context.Context, in *CallEventHandlerRequest, opts ...grpc.CallOption) (*CallEventHandlerResponse, error)
}

type callEventHandlerClient struct {
	cc *grpc.ClientConn
}

func NewCallEventHandlerClient(cc *grpc.ClientConn) CallEventHandlerClient {
	return &callEventHandlerClient{cc}
}

func (c *callEventHandlerClient) CallEventHandler(ctx context.Context, in *CallEventHandlerRequest, opts ...grpc.CallOption) (*CallEventHandlerResponse, error) {
	out := new(CallEventHandlerResponse)
	err := c.cc.Invoke(ctx, "/rpc.CallEventHandler/CallEventHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallEventHandlerServer is the server API for CallEventHandler service.
type CallEventHandlerServer interface {
	// 调用聚合处理事件
	CallEventHandler(context.Context, *CallEventHandlerRequest) (*CallEventHandlerResponse, error)
}

func RegisterCallEventHandlerServer(s *grpc.Server, srv CallEventHandlerServer) {
	s.RegisterService(&_CallEventHandler_serviceDesc, srv)
}

func _CallEventHandler_CallEventHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallEventHandlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallEventHandlerServer).CallEventHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CallEventHandler/CallEventHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallEventHandlerServer).CallEventHandler(ctx, req.(*CallEventHandlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CallEventHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.CallEventHandler",
	HandlerType: (*CallEventHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallEventHandler",
			Handler:    _CallEventHandler_CallEventHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "call_event_handler.proto",
}

func init() {
	proto.RegisterFile("call_event_handler.proto", fileDescriptor_call_event_handler_5400891be3a99dc9)
}

var fileDescriptor_call_event_handler_5400891be3a99dc9 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0xbf, 0xe9, 0x9f, 0xe9, 0xd7, 0x5b, 0x11, 0x09, 0x82, 0xa1, 0x54, 0x1c, 0x67, 0x35,
	0xab, 0x14, 0x46, 0x5c, 0xb8, 0x16, 0xa1, 0x2e, 0x5c, 0x34, 0x74, 0x5f, 0xd2, 0xcc, 0x35, 0x16,
	0xd3, 0x4e, 0x4c, 0x52, 0xc1, 0xf7, 0xf3, 0xc1, 0xc4, 0xcc, 0x0c, 0x05, 0x5b, 0xc7, 0xe5, 0x3d,
	0x39, 0xf7, 0x24, 0xbf, 0x13, 0xa0, 0x52, 0x68, 0xbd, 0xc4, 0x77, 0xdc, 0xfa, 0xe5, 0x8b, 0xd8,
	0x16, 0x1a, 0x2d, 0x33, 0xb6, 0xf4, 0x25, 0xe9, 0x5a, 0x23, 0xc7, 0x57, 0xaa, 0x2c, 0x95, 0xc6,
	0x69, 0x90, 0x56, 0xbb, 0xe7, 0xa9, 0x5f, 0x6f, 0xd0, 0x79, 0xb1, 0x31, 0x95, 0x2b, 0x5d, 0xc0,
	0xc5, 0xbd, 0xd0, 0xfa, 0xe1, 0x3b, 0x60, 0x56, 0xed, 0x73, 0x7c, 0xdb, 0xa1, 0xf3, 0xe4, 0x0e,
	0xe2, 0x90, 0xeb, 0x68, 0x94, 0x74, 0xb3, 0x51, 0x7e, 0xcd, 0xac, 0x91, 0xec, 0x17, 0x77, 0x90,
	0x78, 0xbd, 0x90, 0xce, 0x81, 0x1e, 0xfa, 0x9c, 0x29, 0xb7, 0x0e, 0xc9, 0x2d, 0xf4, 0x0a, 0xe1,
	0xc5, 0x1f, 0xa1, 0x95, 0xf9, 0x49, 0xd8, 0x57, 0x1e, 0xec, 0xe9, 0x67, 0x04, 0x93, 0xb6, 0xbb,
	0xc9, 0x29, 0x74, 0xd6, 0x05, 0x8d, 0x92, 0x28, 0x1b, 0xf2, 0xce, 0xba, 0x20, 0x13, 0x18, 0x86,
	0xd7, 0x2c, 0x3e, 0x0c, 0xd2, 0x4e, 0x90, 0xf7, 0x02, 0x39, 0x87, 0xbe, 0x50, 0xea, 0xb1, 0xa0,
	0xdd, 0x70, 0x52, 0x0d, 0x84, 0xc2, 0x40, 0x28, 0x15, 0x36, 0x7a, 0x41, 0x6f, 0x46, 0x92, 0x43,
	0x2c, 0x2d, 0x0a, 0x8f, 0xb4, 0x9f, 0x44, 0xd9, 0x28, 0x1f, 0xb3, 0xaa, 0x59, 0xd6, 0x34, 0xcb,
	0x16, 0x4d, 0xb3, 0xbc, 0x76, 0x12, 0x52, 0x93, 0xc6, 0x49, 0x94, 0x9d, 0xd4, 0x18, 0xb3, 0x63,
	0x14, 0x7b, 0xd8, 0x03, 0x0a, 0x0a, 0x03, 0xb7, 0x93, 0x12, 0x9d, 0x0b, 0x0c, 0xff, 0x79, 0x33,
	0xe6, 0x08, 0x67, 0x3f, 0x93, 0xc8, 0xfc, 0x88, 0x36, 0x69, 0xfb, 0xb6, 0xf1, 0x65, 0x6b, 0xff,
	0xe9, 0xbf, 0x55, 0x1c, 0x00, 0x6f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x54, 0x50, 0x85, 0xc6,
	0x69, 0x02, 0x00, 0x00,
}
