// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package stream

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// The response message
type ResponseScale struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Subtype              string   `protobuf:"bytes,4,opt,name=subtype,proto3" json:"subtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseScale) Reset()         { *m = ResponseScale{} }
func (m *ResponseScale) String() string { return proto.CompactTextString(m) }
func (*ResponseScale) ProtoMessage()    {}
func (*ResponseScale) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{1}
}

func (m *ResponseScale) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseScale.Unmarshal(m, b)
}
func (m *ResponseScale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseScale.Marshal(b, m, deterministic)
}
func (m *ResponseScale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseScale.Merge(m, src)
}
func (m *ResponseScale) XXX_Size() int {
	return xxx_messageInfo_ResponseScale.Size(m)
}
func (m *ResponseScale) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseScale.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseScale proto.InternalMessageInfo

func (m *ResponseScale) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ResponseScale) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResponseScale) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ResponseScale) GetSubtype() string {
	if m != nil {
		return m.Subtype
	}
	return ""
}

type RequestTareValue struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestTareValue) Reset()         { *m = RequestTareValue{} }
func (m *RequestTareValue) String() string { return proto.CompactTextString(m) }
func (*RequestTareValue) ProtoMessage()    {}
func (*RequestTareValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{2}
}

func (m *RequestTareValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestTareValue.Unmarshal(m, b)
}
func (m *RequestTareValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestTareValue.Marshal(b, m, deterministic)
}
func (m *RequestTareValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestTareValue.Merge(m, src)
}
func (m *RequestTareValue) XXX_Size() int {
	return xxx_messageInfo_RequestTareValue.Size(m)
}
func (m *RequestTareValue) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestTareValue.DiscardUnknown(m)
}

var xxx_messageInfo_RequestTareValue proto.InternalMessageInfo

func (m *RequestTareValue) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RequestScale struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Subtype              string   `protobuf:"bytes,3,opt,name=subtype,proto3" json:"subtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestScale) Reset()         { *m = RequestScale{} }
func (m *RequestScale) String() string { return proto.CompactTextString(m) }
func (*RequestScale) ProtoMessage()    {}
func (*RequestScale) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{3}
}

func (m *RequestScale) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestScale.Unmarshal(m, b)
}
func (m *RequestScale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestScale.Marshal(b, m, deterministic)
}
func (m *RequestScale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestScale.Merge(m, src)
}
func (m *RequestScale) XXX_Size() int {
	return xxx_messageInfo_RequestScale.Size(m)
}
func (m *RequestScale) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestScale.DiscardUnknown(m)
}

var xxx_messageInfo_RequestScale proto.InternalMessageInfo

func (m *RequestScale) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RequestScale) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RequestScale) GetSubtype() string {
	if m != nil {
		return m.Subtype
	}
	return ""
}

type ResponseSetScale struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseSetScale) Reset()         { *m = ResponseSetScale{} }
func (m *ResponseSetScale) String() string { return proto.CompactTextString(m) }
func (*ResponseSetScale) ProtoMessage()    {}
func (*ResponseSetScale) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{4}
}

func (m *ResponseSetScale) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseSetScale.Unmarshal(m, b)
}
func (m *ResponseSetScale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseSetScale.Marshal(b, m, deterministic)
}
func (m *ResponseSetScale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseSetScale.Merge(m, src)
}
func (m *ResponseSetScale) XXX_Size() int {
	return xxx_messageInfo_ResponseSetScale.Size(m)
}
func (m *ResponseSetScale) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseSetScale.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseSetScale proto.InternalMessageInfo

func (m *ResponseSetScale) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ResponseInstantWeight struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseInstantWeight) Reset()         { *m = ResponseInstantWeight{} }
func (m *ResponseInstantWeight) String() string { return proto.CompactTextString(m) }
func (*ResponseInstantWeight) ProtoMessage()    {}
func (*ResponseInstantWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{5}
}

func (m *ResponseInstantWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseInstantWeight.Unmarshal(m, b)
}
func (m *ResponseInstantWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseInstantWeight.Marshal(b, m, deterministic)
}
func (m *ResponseInstantWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseInstantWeight.Merge(m, src)
}
func (m *ResponseInstantWeight) XXX_Size() int {
	return xxx_messageInfo_ResponseInstantWeight.Size(m)
}
func (m *ResponseInstantWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseInstantWeight.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseInstantWeight proto.InternalMessageInfo

func (m *ResponseInstantWeight) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ResponseInstantWeight) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "stream.Empty")
	proto.RegisterType((*ResponseScale)(nil), "stream.ResponseScale")
	proto.RegisterType((*RequestTareValue)(nil), "stream.RequestTareValue")
	proto.RegisterType((*RequestScale)(nil), "stream.RequestScale")
	proto.RegisterType((*ResponseSetScale)(nil), "stream.ResponseSetScale")
	proto.RegisterType((*ResponseInstantWeight)(nil), "stream.ResponseInstantWeight")
}

func init() {
	proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54)
}

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdd, 0x4a, 0xf3, 0x40,
	0x10, 0x86, 0x9b, 0xfe, 0x7e, 0xdf, 0xd0, 0x4a, 0x59, 0x5a, 0x0c, 0x05, 0x41, 0xf6, 0xa8, 0x07,
	0x52, 0xa4, 0x7a, 0x03, 0x5a, 0x25, 0x08, 0x8a, 0x90, 0x48, 0x05, 0xcf, 0xb6, 0x32, 0xb4, 0x85,
	0xfc, 0xb9, 0x3b, 0x39, 0xe8, 0x1d, 0x78, 0xd9, 0xd2, 0xdd, 0x4d, 0x4d, 0x52, 0x5b, 0xe8, 0xd9,
	0xce, 0xe4, 0x99, 0x99, 0xf7, 0x9d, 0x21, 0xd0, 0x55, 0x24, 0x51, 0x44, 0x93, 0x54, 0x26, 0x94,
	0xb0, 0xb6, 0x89, 0x78, 0x07, 0x5a, 0x8f, 0x51, 0x4a, 0x1b, 0x1e, 0x41, 0xcf, 0x47, 0x95, 0x26,
	0xb1, 0xc2, 0xe0, 0x53, 0x84, 0xc8, 0x06, 0xd0, 0x42, 0x29, 0x13, 0xe9, 0x3a, 0x97, 0xce, 0xf8,
	0xbf, 0x6f, 0x02, 0xe6, 0x42, 0x27, 0x42, 0xa5, 0xc4, 0x12, 0xdd, 0xba, 0xce, 0xe7, 0x21, 0x63,
	0xd0, 0xa4, 0x4d, 0x8a, 0x6e, 0x43, 0xa7, 0xf5, 0x7b, 0x4b, 0xab, 0x6c, 0xa1, 0xd3, 0x4d, 0x43,
	0xdb, 0x90, 0x5f, 0x41, 0xdf, 0xc7, 0xaf, 0x0c, 0x15, 0xbd, 0x09, 0x89, 0x73, 0x11, 0x66, 0x58,
	0xec, 0xed, 0x94, 0x7a, 0xf3, 0x39, 0x74, 0x2d, 0x6d, 0xb4, 0x1d, 0x24, 0x77, 0x2a, 0xea, 0x7f,
	0xab, 0x68, 0x94, 0x55, 0x8c, 0xb7, 0x2a, 0xac, 0x69, 0xa4, 0x23, 0xbe, 0xb9, 0x07, 0xc3, 0x9c,
	0x7c, 0x8a, 0x15, 0x89, 0x98, 0xde, 0x71, 0xbd, 0x5c, 0xd1, 0xa9, 0x6b, 0x9a, 0x7e, 0x37, 0xe0,
	0xec, 0x2e, 0x5d, 0xcf, 0x44, 0x18, 0xa2, 0x34, 0x13, 0x9f, 0xe1, 0x5c, 0x3f, 0xd4, 0x8b, 0x61,
	0x5e, 0x33, 0x9a, 0xad, 0x44, 0x1c, 0x63, 0xc8, 0x06, 0x13, 0x7b, 0xb5, 0xa2, 0xfd, 0xd1, 0xf0,
	0x37, 0x5b, 0xb8, 0x18, 0xaf, 0x8d, 0x9d, 0x6b, 0x87, 0xdd, 0x42, 0x27, 0x40, 0xbd, 0x55, 0xd6,
	0xcb, 0x39, 0x7d, 0xe2, 0x91, 0xbb, 0x57, 0x66, 0x3d, 0xf3, 0x1a, 0x7b, 0x80, 0xae, 0xad, 0xb2,
	0xb7, 0xa8, 0x0c, 0xde, 0x7d, 0x39, 0xda, 0xc5, 0xcc, 0xfe, 0x40, 0x99, 0x9c, 0x32, 0xfb, 0x1e,
	0xfa, 0x1e, 0x52, 0x79, 0xad, 0x95, 0xf2, 0x8b, 0x6a, 0x79, 0x89, 0xe6, 0x35, 0x36, 0x85, 0x7f,
	0x1e, 0x52, 0x40, 0x82, 0xf6, 0x6c, 0x1f, 0xda, 0xd6, 0xa2, 0xad, 0x7f, 0x85, 0x9b, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0x0a, 0x62, 0xf0, 0x1a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApiCallerScaleClient is the client API for ApiCallerScale service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiCallerScaleClient interface {
	ScalesMessageOutChannel(ctx context.Context, opts ...grpc.CallOption) (ApiCallerScale_ScalesMessageOutChannelClient, error)
	SetTare(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseSetScale, error)
	SetTareValue(ctx context.Context, in *RequestTareValue, opts ...grpc.CallOption) (*ResponseSetScale, error)
	SetZero(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseSetScale, error)
	// +
	GetInstantWeight(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseInstantWeight, error)
	// +
	GetState(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseScale, error)
}

type apiCallerScaleClient struct {
	cc grpc.ClientConnInterface
}

func NewApiCallerScaleClient(cc grpc.ClientConnInterface) ApiCallerScaleClient {
	return &apiCallerScaleClient{cc}
}

func (c *apiCallerScaleClient) ScalesMessageOutChannel(ctx context.Context, opts ...grpc.CallOption) (ApiCallerScale_ScalesMessageOutChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ApiCallerScale_serviceDesc.Streams[0], "/stream.ApiCallerScale/ScalesMessageOutChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiCallerScaleScalesMessageOutChannelClient{stream}
	return x, nil
}

type ApiCallerScale_ScalesMessageOutChannelClient interface {
	Send(*RequestScale) error
	Recv() (*ResponseScale, error)
	grpc.ClientStream
}

type apiCallerScaleScalesMessageOutChannelClient struct {
	grpc.ClientStream
}

func (x *apiCallerScaleScalesMessageOutChannelClient) Send(m *RequestScale) error {
	return x.ClientStream.SendMsg(m)
}

func (x *apiCallerScaleScalesMessageOutChannelClient) Recv() (*ResponseScale, error) {
	m := new(ResponseScale)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *apiCallerScaleClient) SetTare(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseSetScale, error) {
	out := new(ResponseSetScale)
	err := c.cc.Invoke(ctx, "/stream.ApiCallerScale/SetTare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiCallerScaleClient) SetTareValue(ctx context.Context, in *RequestTareValue, opts ...grpc.CallOption) (*ResponseSetScale, error) {
	out := new(ResponseSetScale)
	err := c.cc.Invoke(ctx, "/stream.ApiCallerScale/SetTareValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiCallerScaleClient) SetZero(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseSetScale, error) {
	out := new(ResponseSetScale)
	err := c.cc.Invoke(ctx, "/stream.ApiCallerScale/SetZero", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiCallerScaleClient) GetInstantWeight(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseInstantWeight, error) {
	out := new(ResponseInstantWeight)
	err := c.cc.Invoke(ctx, "/stream.ApiCallerScale/GetInstantWeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiCallerScaleClient) GetState(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseScale, error) {
	out := new(ResponseScale)
	err := c.cc.Invoke(ctx, "/stream.ApiCallerScale/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiCallerScaleServer is the server API for ApiCallerScale service.
type ApiCallerScaleServer interface {
	ScalesMessageOutChannel(ApiCallerScale_ScalesMessageOutChannelServer) error
	SetTare(context.Context, *Empty) (*ResponseSetScale, error)
	SetTareValue(context.Context, *RequestTareValue) (*ResponseSetScale, error)
	SetZero(context.Context, *Empty) (*ResponseSetScale, error)
	// +
	GetInstantWeight(context.Context, *Empty) (*ResponseInstantWeight, error)
	// +
	GetState(context.Context, *Empty) (*ResponseScale, error)
}

// UnimplementedApiCallerScaleServer can be embedded to have forward compatible implementations.
type UnimplementedApiCallerScaleServer struct {
}

func (*UnimplementedApiCallerScaleServer) ScalesMessageOutChannel(srv ApiCallerScale_ScalesMessageOutChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method ScalesMessageOutChannel not implemented")
}
func (*UnimplementedApiCallerScaleServer) SetTare(ctx context.Context, req *Empty) (*ResponseSetScale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTare not implemented")
}
func (*UnimplementedApiCallerScaleServer) SetTareValue(ctx context.Context, req *RequestTareValue) (*ResponseSetScale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTareValue not implemented")
}
func (*UnimplementedApiCallerScaleServer) SetZero(ctx context.Context, req *Empty) (*ResponseSetScale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetZero not implemented")
}
func (*UnimplementedApiCallerScaleServer) GetInstantWeight(ctx context.Context, req *Empty) (*ResponseInstantWeight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstantWeight not implemented")
}
func (*UnimplementedApiCallerScaleServer) GetState(ctx context.Context, req *Empty) (*ResponseScale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}

func RegisterApiCallerScaleServer(s *grpc.Server, srv ApiCallerScaleServer) {
	s.RegisterService(&_ApiCallerScale_serviceDesc, srv)
}

func _ApiCallerScale_ScalesMessageOutChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ApiCallerScaleServer).ScalesMessageOutChannel(&apiCallerScaleScalesMessageOutChannelServer{stream})
}

type ApiCallerScale_ScalesMessageOutChannelServer interface {
	Send(*ResponseScale) error
	Recv() (*RequestScale, error)
	grpc.ServerStream
}

type apiCallerScaleScalesMessageOutChannelServer struct {
	grpc.ServerStream
}

func (x *apiCallerScaleScalesMessageOutChannelServer) Send(m *ResponseScale) error {
	return x.ServerStream.SendMsg(m)
}

func (x *apiCallerScaleScalesMessageOutChannelServer) Recv() (*RequestScale, error) {
	m := new(RequestScale)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ApiCallerScale_SetTare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCallerScaleServer).SetTare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.ApiCallerScale/SetTare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCallerScaleServer).SetTare(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiCallerScale_SetTareValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTareValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCallerScaleServer).SetTareValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.ApiCallerScale/SetTareValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCallerScaleServer).SetTareValue(ctx, req.(*RequestTareValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiCallerScale_SetZero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCallerScaleServer).SetZero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.ApiCallerScale/SetZero",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCallerScaleServer).SetZero(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiCallerScale_GetInstantWeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCallerScaleServer).GetInstantWeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.ApiCallerScale/GetInstantWeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCallerScaleServer).GetInstantWeight(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiCallerScale_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCallerScaleServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.ApiCallerScale/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCallerScaleServer).GetState(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiCallerScale_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stream.ApiCallerScale",
	HandlerType: (*ApiCallerScaleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetTare",
			Handler:    _ApiCallerScale_SetTare_Handler,
		},
		{
			MethodName: "SetTareValue",
			Handler:    _ApiCallerScale_SetTareValue_Handler,
		},
		{
			MethodName: "SetZero",
			Handler:    _ApiCallerScale_SetZero_Handler,
		},
		{
			MethodName: "GetInstantWeight",
			Handler:    _ApiCallerScale_GetInstantWeight_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _ApiCallerScale_GetState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ScalesMessageOutChannel",
			Handler:       _ApiCallerScale_ScalesMessageOutChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
