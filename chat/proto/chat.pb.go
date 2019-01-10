// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat/proto/chat.proto

package goChat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_825b1469f80f958d, []int{0}
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

type ChatMessage struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver             string   `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_825b1469f80f958d, []int{1}
}

func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (m *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(m, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *ChatMessage) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *ChatMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ClientInfo struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_825b1469f80f958d, []int{2}
}

func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (m *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(m, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type GroupInfo struct {
	Client               string   `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	GroupName            string   `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupInfo) Reset()         { *m = GroupInfo{} }
func (m *GroupInfo) String() string { return proto.CompactTextString(m) }
func (*GroupInfo) ProtoMessage()    {}
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_825b1469f80f958d, []int{3}
}

func (m *GroupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupInfo.Unmarshal(m, b)
}
func (m *GroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupInfo.Marshal(b, m, deterministic)
}
func (m *GroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupInfo.Merge(m, src)
}
func (m *GroupInfo) XXX_Size() int {
	return xxx_messageInfo_GroupInfo.Size(m)
}
func (m *GroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GroupInfo proto.InternalMessageInfo

func (m *GroupInfo) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *GroupInfo) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type GroupList struct {
	Groups               []string `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupList) Reset()         { *m = GroupList{} }
func (m *GroupList) String() string { return proto.CompactTextString(m) }
func (*GroupList) ProtoMessage()    {}
func (*GroupList) Descriptor() ([]byte, []int) {
	return fileDescriptor_825b1469f80f958d, []int{4}
}

func (m *GroupList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupList.Unmarshal(m, b)
}
func (m *GroupList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupList.Marshal(b, m, deterministic)
}
func (m *GroupList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupList.Merge(m, src)
}
func (m *GroupList) XXX_Size() int {
	return xxx_messageInfo_GroupList.Size(m)
}
func (m *GroupList) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupList.DiscardUnknown(m)
}

var xxx_messageInfo_GroupList proto.InternalMessageInfo

func (m *GroupList) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

type ClientList struct {
	Clients              []string `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientList) Reset()         { *m = ClientList{} }
func (m *ClientList) String() string { return proto.CompactTextString(m) }
func (*ClientList) ProtoMessage()    {}
func (*ClientList) Descriptor() ([]byte, []int) {
	return fileDescriptor_825b1469f80f958d, []int{5}
}

func (m *ClientList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientList.Unmarshal(m, b)
}
func (m *ClientList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientList.Marshal(b, m, deterministic)
}
func (m *ClientList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientList.Merge(m, src)
}
func (m *ClientList) XXX_Size() int {
	return xxx_messageInfo_ClientList.Size(m)
}
func (m *ClientList) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientList.DiscardUnknown(m)
}

var xxx_messageInfo_ClientList proto.InternalMessageInfo

func (m *ClientList) GetClients() []string {
	if m != nil {
		return m.Clients
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "goChat.Empty")
	proto.RegisterType((*ChatMessage)(nil), "goChat.ChatMessage")
	proto.RegisterType((*ClientInfo)(nil), "goChat.ClientInfo")
	proto.RegisterType((*GroupInfo)(nil), "goChat.GroupInfo")
	proto.RegisterType((*GroupList)(nil), "goChat.GroupList")
	proto.RegisterType((*ClientList)(nil), "goChat.ClientList")
}

func init() { proto.RegisterFile("chat/proto/chat.proto", fileDescriptor_825b1469f80f958d) }

var fileDescriptor_825b1469f80f958d = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0xac, 0xb6, 0xcd, 0xd4, 0x1e, 0x1c, 0x51, 0x42, 0xf0, 0x50, 0x56, 0x91, 0x9e, 0x1a,
	0x5b, 0xbd, 0x89, 0x07, 0x29, 0x52, 0x94, 0xea, 0x21, 0xe0, 0xc9, 0x53, 0xac, 0x63, 0x1a, 0x30,
	0xd9, 0xb2, 0xd9, 0x16, 0xfc, 0xc1, 0xfe, 0x0f, 0xd9, 0xdd, 0x6c, 0xfa, 0x41, 0x84, 0x7a, 0xdb,
	0x37, 0xf3, 0xde, 0x9b, 0x99, 0x17, 0x02, 0x27, 0xd3, 0x59, 0x24, 0x83, 0xb9, 0xe0, 0x92, 0x07,
	0xea, 0xd9, 0xd7, 0x4f, 0x6c, 0xc4, 0x7c, 0x34, 0x8b, 0x24, 0x6b, 0xc2, 0xc1, 0x43, 0x3a, 0x97,
	0xdf, 0xec, 0x0d, 0xda, 0xaa, 0xf0, 0x4c, 0x79, 0x1e, 0xc5, 0x84, 0xa7, 0xd0, 0xc8, 0x29, 0xfb,
	0x20, 0xe1, 0x39, 0x5d, 0xa7, 0xe7, 0x86, 0x05, 0x42, 0x1f, 0x5a, 0x82, 0xa6, 0x94, 0x2c, 0x49,
	0x78, 0x7b, 0xba, 0x53, 0x62, 0xf4, 0xa0, 0x99, 0x1a, 0xb9, 0x57, 0xd7, 0x2d, 0x0b, 0xd9, 0x05,
	0xc0, 0xe8, 0x2b, 0xa1, 0x4c, 0x3e, 0x66, 0x9f, 0xfc, 0x2f, 0x6f, 0x76, 0x0f, 0xee, 0x58, 0xf0,
	0xc5, 0xdc, 0x92, 0xa6, 0x5a, 0x62, 0x49, 0x06, 0xe1, 0x19, 0xb8, 0xb1, 0x22, 0xbd, 0x44, 0x29,
	0x15, 0x1b, 0xac, 0x0a, 0xec, 0xbc, 0xb0, 0x98, 0x24, 0xb9, 0x54, 0x16, 0xba, 0x93, 0x7b, 0x4e,
	0xb7, 0xae, 0x2c, 0x0c, 0x62, 0x97, 0x76, 0x1b, 0xcd, 0xf2, 0xa0, 0x69, 0xac, 0x2d, 0xcd, 0xc2,
	0xe1, 0x4f, 0x1d, 0xf6, 0x55, 0x26, 0x78, 0x0b, 0x6e, 0xc8, 0x17, 0x92, 0x34, 0x38, 0xee, 0x9b,
	0xe8, 0xfa, 0x6b, 0x71, 0xf9, 0x55, 0x45, 0x56, 0xeb, 0x39, 0x57, 0x0e, 0x0e, 0x00, 0x5e, 0xb3,
	0x90, 0xe2, 0x24, 0x97, 0x24, 0x10, 0x4b, 0x62, 0x99, 0x87, 0xdf, 0xb1, 0x35, 0xf3, 0x25, 0x6a,
	0x18, 0x40, 0xeb, 0x7f, 0x82, 0x01, 0xb4, 0x47, 0x82, 0x22, 0x49, 0xfa, 0x78, 0x3c, 0xb2, 0xfd,
	0x32, 0xce, 0xaa, 0x19, 0xee, 0x13, 0x4f, 0xb2, 0xdd, 0x05, 0x43, 0x38, 0x1c, 0x93, 0x5c, 0xa5,
	0xbb, 0x49, 0xf0, 0x37, 0x2d, 0x14, 0x83, 0xd5, 0xf0, 0x0e, 0xd0, 0x6a, 0xd6, 0x12, 0xaf, 0x98,
	0xb6, 0x75, 0x65, 0x21, 0xbf, 0x81, 0xce, 0x98, 0xe4, 0x9a, 0x72, 0x6b, 0x66, 0xb5, 0x2a, 0x00,
	0x77, 0x42, 0xd1, 0x92, 0x42, 0xce, 0xd3, 0x5d, 0x2e, 0x7b, 0x6f, 0xe8, 0x5f, 0xe2, 0xfa, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x20, 0x55, 0x88, 0x6d, 0x2b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (Chat_RouteChatClient, error)
	UnRegister(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*Empty, error)
	Register(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*Empty, error)
	CreateGroup(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error)
	JoinGroup(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error)
	GetGroupList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GroupList, error)
	GetGroupClientList(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*ClientList, error)
	GetClientList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ClientList, error)
	LeaveRoom(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (Chat_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/goChat.Chat/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatRouteChatClient{stream}
	return x, nil
}

type Chat_RouteChatClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatRouteChatClient struct {
	grpc.ClientStream
}

func (x *chatRouteChatClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatRouteChatClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) UnRegister(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/goChat.Chat/UnRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Register(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/goChat.Chat/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) CreateGroup(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/goChat.Chat/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) JoinGroup(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/goChat.Chat/JoinGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetGroupList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GroupList, error) {
	out := new(GroupList)
	err := c.cc.Invoke(ctx, "/goChat.Chat/GetGroupList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetGroupClientList(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*ClientList, error) {
	out := new(ClientList)
	err := c.cc.Invoke(ctx, "/goChat.Chat/GetGroupClientList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetClientList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ClientList, error) {
	out := new(ClientList)
	err := c.cc.Invoke(ctx, "/goChat.Chat/GetClientList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) LeaveRoom(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/goChat.Chat/LeaveRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	RouteChat(Chat_RouteChatServer) error
	UnRegister(context.Context, *ClientInfo) (*Empty, error)
	Register(context.Context, *ClientInfo) (*Empty, error)
	CreateGroup(context.Context, *GroupInfo) (*Empty, error)
	JoinGroup(context.Context, *GroupInfo) (*Empty, error)
	GetGroupList(context.Context, *Empty) (*GroupList, error)
	GetGroupClientList(context.Context, *GroupInfo) (*ClientList, error)
	GetClientList(context.Context, *Empty) (*ClientList, error)
	LeaveRoom(context.Context, *GroupInfo) (*Empty, error)
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).RouteChat(&chatRouteChatServer{stream})
}

type Chat_RouteChatServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatRouteChatServer struct {
	grpc.ServerStream
}

func (x *chatRouteChatServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatRouteChatServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_UnRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).UnRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goChat.Chat/UnRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).UnRegister(ctx, req.(*ClientInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goChat.Chat/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Register(ctx, req.(*ClientInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goChat.Chat/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CreateGroup(ctx, req.(*GroupInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_JoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).JoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goChat.Chat/JoinGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).JoinGroup(ctx, req.(*GroupInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetGroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetGroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goChat.Chat/GetGroupList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetGroupList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetGroupClientList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetGroupClientList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goChat.Chat/GetGroupClientList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetGroupClientList(ctx, req.(*GroupInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetClientList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetClientList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goChat.Chat/GetClientList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetClientList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_LeaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).LeaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goChat.Chat/LeaveRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).LeaveRoom(ctx, req.(*GroupInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goChat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnRegister",
			Handler:    _Chat_UnRegister_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Chat_Register_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _Chat_CreateGroup_Handler,
		},
		{
			MethodName: "JoinGroup",
			Handler:    _Chat_JoinGroup_Handler,
		},
		{
			MethodName: "GetGroupList",
			Handler:    _Chat_GetGroupList_Handler,
		},
		{
			MethodName: "GetGroupClientList",
			Handler:    _Chat_GetGroupClientList_Handler,
		},
		{
			MethodName: "GetClientList",
			Handler:    _Chat_GetClientList_Handler,
		},
		{
			MethodName: "LeaveRoom",
			Handler:    _Chat_LeaveRoom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouteChat",
			Handler:       _Chat_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat/proto/chat.proto",
}
