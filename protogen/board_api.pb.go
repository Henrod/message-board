// Code generated by protoc-gen-go. DO NOT EDIT.
// source: board_api.proto

package protov1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateMessageRequest struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessageRequest) Reset()         { *m = CreateMessageRequest{} }
func (m *CreateMessageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMessageRequest) ProtoMessage()    {}
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ae269f57042a290, []int{0}
}

func (m *CreateMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageRequest.Unmarshal(m, b)
}
func (m *CreateMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageRequest.Marshal(b, m, deterministic)
}
func (m *CreateMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageRequest.Merge(m, src)
}
func (m *CreateMessageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMessageRequest.Size(m)
}
func (m *CreateMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageRequest proto.InternalMessageInfo

func (m *CreateMessageRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *CreateMessageRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type CreateMessageResponse struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessageResponse) Reset()         { *m = CreateMessageResponse{} }
func (m *CreateMessageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMessageResponse) ProtoMessage()    {}
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ae269f57042a290, []int{1}
}

func (m *CreateMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageResponse.Unmarshal(m, b)
}
func (m *CreateMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageResponse.Marshal(b, m, deterministic)
}
func (m *CreateMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageResponse.Merge(m, src)
}
func (m *CreateMessageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMessageResponse.Size(m)
}
func (m *CreateMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageResponse proto.InternalMessageInfo

func (m *CreateMessageResponse) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type ListMessagesRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMessagesRequest) Reset()         { *m = ListMessagesRequest{} }
func (m *ListMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*ListMessagesRequest) ProtoMessage()    {}
func (*ListMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ae269f57042a290, []int{2}
}

func (m *ListMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMessagesRequest.Unmarshal(m, b)
}
func (m *ListMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMessagesRequest.Marshal(b, m, deterministic)
}
func (m *ListMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMessagesRequest.Merge(m, src)
}
func (m *ListMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_ListMessagesRequest.Size(m)
}
func (m *ListMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMessagesRequest proto.InternalMessageInfo

func (m *ListMessagesRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ListMessagesResponse struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListMessagesResponse) Reset()         { *m = ListMessagesResponse{} }
func (m *ListMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMessagesResponse) ProtoMessage()    {}
func (*ListMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ae269f57042a290, []int{3}
}

func (m *ListMessagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMessagesResponse.Unmarshal(m, b)
}
func (m *ListMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMessagesResponse.Marshal(b, m, deterministic)
}
func (m *ListMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMessagesResponse.Merge(m, src)
}
func (m *ListMessagesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMessagesResponse.Size(m)
}
func (m *ListMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMessagesResponse proto.InternalMessageInfo

func (m *ListMessagesResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateMessageRequest)(nil), "proto.v1.CreateMessageRequest")
	proto.RegisterType((*CreateMessageResponse)(nil), "proto.v1.CreateMessageResponse")
	proto.RegisterType((*ListMessagesRequest)(nil), "proto.v1.ListMessagesRequest")
	proto.RegisterType((*ListMessagesResponse)(nil), "proto.v1.ListMessagesResponse")
}

func init() { proto.RegisterFile("board_api.proto", fileDescriptor_7ae269f57042a290) }

var fileDescriptor_7ae269f57042a290 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xca, 0x4f, 0x2c,
	0x4a, 0x89, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x7a,
	0x65, 0x86, 0x52, 0xbc, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x10, 0x09, 0x29, 0x99, 0xf4,
	0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4,
	0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0xac, 0x52, 0x0c, 0x97, 0x88, 0x73, 0x51, 0x6a, 0x62, 0x49,
	0xaa, 0x2f, 0x44, 0x53, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x36, 0x17, 0x3b, 0xd4,
	0x18, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x41, 0x3d, 0x98, 0x05, 0x7a, 0x30, 0xa5, 0x30,
	0x15, 0x42, 0xe2, 0x5c, 0xec, 0xa5, 0xc5, 0xa9, 0x45, 0xf1, 0x99, 0x29, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x6c, 0x20, 0xae, 0x67, 0x8a, 0x92, 0x0b, 0x97, 0x28, 0x9a, 0xe9, 0xc5, 0x05,
	0xf9, 0x79, 0xc5, 0xa9, 0x24, 0x19, 0xaf, 0xa4, 0xc7, 0x25, 0xec, 0x93, 0x59, 0x5c, 0x02, 0x15,
	0x2f, 0x86, 0x39, 0x11, 0xc9, 0x56, 0x46, 0x14, 0x5b, 0x5d, 0xb9, 0x44, 0x50, 0xd5, 0x43, 0x2d,
	0xd5, 0xe5, 0xe2, 0x80, 0x1a, 0x59, 0x2c, 0xc1, 0xa8, 0xc0, 0x8c, 0xdd, 0x56, 0xb8, 0x12, 0xa3,
	0x5b, 0x8c, 0x5c, 0x1c, 0x4e, 0xa0, 0x50, 0x76, 0x0c, 0xf0, 0x14, 0xaa, 0xe7, 0xe2, 0x45, 0xf1,
	0x89, 0x90, 0x1c, 0x42, 0x2b, 0xb6, 0x00, 0x94, 0x92, 0xc7, 0x29, 0x0f, 0x71, 0x8d, 0x92, 0x6e,
	0xd3, 0xe5, 0x27, 0x93, 0x99, 0xd4, 0x95, 0x64, 0xf4, 0xcb, 0x0c, 0xf5, 0x41, 0x2e, 0x2f, 0xd6,
	0xaf, 0x86, 0xfa, 0xa7, 0x56, 0x1f, 0xe6, 0x08, 0x2b, 0x78, 0x18, 0xfb, 0x72, 0xf1, 0x20, 0x7b,
	0x4a, 0x48, 0x16, 0x61, 0x3e, 0x96, 0xc0, 0x91, 0x92, 0xc3, 0x25, 0x0d, 0xb1, 0xdd, 0xc9, 0x85,
	0x8b, 0x27, 0x39, 0x3f, 0x17, 0xae, 0xc8, 0x89, 0x17, 0xe2, 0xd3, 0x82, 0xcc, 0x00, 0x90, 0x48,
	0x00, 0x63, 0x14, 0x3b, 0x58, 0xaa, 0xcc, 0x70, 0x11, 0x13, 0x73, 0x40, 0x44, 0xc4, 0x2a, 0x26,
	0x0e, 0xb0, 0x84, 0x5e, 0x98, 0xe1, 0x29, 0x28, 0x33, 0x26, 0xcc, 0x30, 0x89, 0x0d, 0xac, 0xc8,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x6e, 0x63, 0x7c, 0x8e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BoardAPIClient is the client API for BoardAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BoardAPIClient interface {
	// CreateMessage posts users message on theirs boards.
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
	ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error)
}

type boardAPIClient struct {
	cc *grpc.ClientConn
}

func NewBoardAPIClient(cc *grpc.ClientConn) BoardAPIClient {
	return &boardAPIClient{cc}
}

func (c *boardAPIClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := c.cc.Invoke(ctx, "/proto.v1.BoardAPI/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardAPIClient) ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error) {
	out := new(ListMessagesResponse)
	err := c.cc.Invoke(ctx, "/proto.v1.BoardAPI/ListMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardAPIServer is the server API for BoardAPI service.
type BoardAPIServer interface {
	// CreateMessage posts users message on theirs boards.
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
	ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error)
}

// UnimplementedBoardAPIServer can be embedded to have forward compatible implementations.
type UnimplementedBoardAPIServer struct {
}

func (*UnimplementedBoardAPIServer) CreateMessage(ctx context.Context, req *CreateMessageRequest) (*CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (*UnimplementedBoardAPIServer) ListMessages(ctx context.Context, req *ListMessagesRequest) (*ListMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessages not implemented")
}

func RegisterBoardAPIServer(s *grpc.Server, srv BoardAPIServer) {
	s.RegisterService(&_BoardAPI_serviceDesc, srv)
}

func _BoardAPI_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardAPIServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.v1.BoardAPI/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardAPIServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardAPI_ListMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardAPIServer).ListMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.v1.BoardAPI/ListMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardAPIServer).ListMessages(ctx, req.(*ListMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BoardAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.BoardAPI",
	HandlerType: (*BoardAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _BoardAPI_CreateMessage_Handler,
		},
		{
			MethodName: "ListMessages",
			Handler:    _BoardAPI_ListMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "board_api.proto",
}
