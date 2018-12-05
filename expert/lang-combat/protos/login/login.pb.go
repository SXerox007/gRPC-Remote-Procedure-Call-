// Code generated by protoc-gen-go. DO NOT EDIT.
// source: expert/lang-combat/protos/login/login.proto

package login

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

type LoginRequestData struct {
	EmailOrPhone         string   `protobuf:"bytes,1,opt,name=EmailOrPhone,proto3" json:"EmailOrPhone,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequestData) Reset()         { *m = LoginRequestData{} }
func (m *LoginRequestData) String() string { return proto.CompactTextString(m) }
func (*LoginRequestData) ProtoMessage()    {}
func (*LoginRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d2a1f8a79770da, []int{0}
}

func (m *LoginRequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequestData.Unmarshal(m, b)
}
func (m *LoginRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequestData.Marshal(b, m, deterministic)
}
func (m *LoginRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequestData.Merge(m, src)
}
func (m *LoginRequestData) XXX_Size() int {
	return xxx_messageInfo_LoginRequestData.Size(m)
}
func (m *LoginRequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequestData.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequestData proto.InternalMessageInfo

func (m *LoginRequestData) GetEmailOrPhone() string {
	if m != nil {
		return m.EmailOrPhone
	}
	return ""
}

func (m *LoginRequestData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponseData struct {
	Success              string   `protobuf:"bytes,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponseData) Reset()         { *m = LoginResponseData{} }
func (m *LoginResponseData) String() string { return proto.CompactTextString(m) }
func (*LoginResponseData) ProtoMessage()    {}
func (*LoginResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d2a1f8a79770da, []int{1}
}

func (m *LoginResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponseData.Unmarshal(m, b)
}
func (m *LoginResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponseData.Marshal(b, m, deterministic)
}
func (m *LoginResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponseData.Merge(m, src)
}
func (m *LoginResponseData) XXX_Size() int {
	return xxx_messageInfo_LoginResponseData.Size(m)
}
func (m *LoginResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponseData proto.InternalMessageInfo

func (m *LoginResponseData) GetSuccess() string {
	if m != nil {
		return m.Success
	}
	return ""
}

func (m *LoginResponseData) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type LoginRequest struct {
	LoginRequestData     *LoginRequestData `protobuf:"bytes,1,opt,name=loginRequestData,proto3" json:"loginRequestData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d2a1f8a79770da, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetLoginRequestData() *LoginRequestData {
	if m != nil {
		return m.LoginRequestData
	}
	return nil
}

type LoginResponse struct {
	LoginResponseData    *LoginResponseData `protobuf:"bytes,1,opt,name=loginResponseData,proto3" json:"loginResponseData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d2a1f8a79770da, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetLoginResponseData() *LoginResponseData {
	if m != nil {
		return m.LoginResponseData
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginRequestData)(nil), "LoginRequestData")
	proto.RegisterType((*LoginResponseData)(nil), "LoginResponseData")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
}

func init() {
	proto.RegisterFile("expert/lang-combat/protos/login/login.proto", fileDescriptor_56d2a1f8a79770da)
}

var fileDescriptor_56d2a1f8a79770da = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0x9d, 0xb0, 0xa9, 0xcf, 0x4d, 0xd6, 0x77, 0x90, 0xb2, 0x93, 0xe4, 0x34, 0x10, 0x5b,
	0x98, 0x17, 0x2f, 0x82, 0x07, 0x77, 0x53, 0x9c, 0xe9, 0x2f, 0xc8, 0xea, 0x63, 0x16, 0xb2, 0xa4,
	0xe6, 0xa5, 0xea, 0xcf, 0x17, 0x63, 0x14, 0xd3, 0x5e, 0x0a, 0xdf, 0xeb, 0xe3, 0xe3, 0x7b, 0x81,
	0x4b, 0xfa, 0x6c, 0xc9, 0xf9, 0x52, 0x2b, 0xb3, 0xbb, 0xaa, 0xed, 0x7e, 0xab, 0x7c, 0xd9, 0x3a,
	0xeb, 0x2d, 0x97, 0xda, 0xee, 0x1a, 0xf3, 0xf3, 0x2d, 0xc2, 0x48, 0x48, 0x98, 0x3f, 0x7c, 0xa3,
	0xa4, 0xb7, 0x8e, 0xd8, 0xdf, 0x2b, 0xaf, 0x50, 0xc0, 0x74, 0xbd, 0x57, 0x8d, 0x7e, 0x72, 0x9b,
	0x57, 0x6b, 0x28, 0x1f, 0x5d, 0x8c, 0x96, 0x27, 0x32, 0x99, 0xe1, 0x02, 0x8e, 0x37, 0x8a, 0xf9,
	0xc3, 0xba, 0x97, 0xfc, 0x30, 0xfc, 0xff, 0x63, 0xb1, 0x86, 0x2c, 0x3a, 0xb9, 0xb5, 0x86, 0x29,
	0x48, 0x73, 0x38, 0xaa, 0xba, 0xba, 0x26, 0xe6, 0xe8, 0xfb, 0x45, 0x3c, 0x87, 0x49, 0xe5, 0x95,
	0xef, 0x38, 0x88, 0xc6, 0x32, 0x92, 0x78, 0x84, 0xe9, 0xff, 0x34, 0xbc, 0x85, 0xb9, 0xee, 0xa5,
	0x06, 0xd5, 0xe9, 0x2a, 0x2b, 0xfa, 0x37, 0xc8, 0xc1, 0xaa, 0x78, 0x86, 0x59, 0x52, 0x85, 0x77,
	0x90, 0xe9, 0x7e, 0x66, 0x14, 0x62, 0x31, 0x38, 0x40, 0x0e, 0x97, 0x57, 0x37, 0xb1, 0xb0, 0x22,
	0xf7, 0xde, 0xd4, 0x84, 0x4b, 0x18, 0x07, 0xc6, 0x59, 0x12, 0xb4, 0x38, 0x4b, 0x75, 0xe2, 0x60,
	0x3b, 0x09, 0xaf, 0x7f, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x94, 0x26, 0xb8, 0x0e, 0xac, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type loginServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoginServiceClient(cc *grpc.ClientConn) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/LoginService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
type LoginServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "expert/lang-combat/protos/login/login.proto",
}
