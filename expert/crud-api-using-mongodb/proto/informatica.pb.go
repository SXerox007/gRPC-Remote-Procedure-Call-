// Code generated by protoc-gen-go. DO NOT EDIT.
// source: expert/crud-api-using-mongodb/proto/informatica.proto

package informaticapb

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

type Informatica struct {
	Sequence             int32    `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Info                 string   `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	HostName             string   `protobuf:"bytes,4,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Informatica) Reset()         { *m = Informatica{} }
func (m *Informatica) String() string { return proto.CompactTextString(m) }
func (*Informatica) ProtoMessage()    {}
func (*Informatica) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3c05494cc022a7, []int{0}
}

func (m *Informatica) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Informatica.Unmarshal(m, b)
}
func (m *Informatica) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Informatica.Marshal(b, m, deterministic)
}
func (m *Informatica) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Informatica.Merge(m, src)
}
func (m *Informatica) XXX_Size() int {
	return xxx_messageInfo_Informatica.Size(m)
}
func (m *Informatica) XXX_DiscardUnknown() {
	xxx_messageInfo_Informatica.DiscardUnknown(m)
}

var xxx_messageInfo_Informatica proto.InternalMessageInfo

func (m *Informatica) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Informatica) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Informatica) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Informatica) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

type CommonResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3c05494cc022a7, []int{1}
}

func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CommonResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CommonRequest struct {
	//as of now it will be bool
	// but access_token should be string given when login
	// it just for sample purpose
	AccessToken          bool     `protobuf:"varint,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonRequest) Reset()         { *m = CommonRequest{} }
func (m *CommonRequest) String() string { return proto.CompactTextString(m) }
func (*CommonRequest) ProtoMessage()    {}
func (*CommonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3c05494cc022a7, []int{2}
}

func (m *CommonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonRequest.Unmarshal(m, b)
}
func (m *CommonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonRequest.Marshal(b, m, deterministic)
}
func (m *CommonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonRequest.Merge(m, src)
}
func (m *CommonRequest) XXX_Size() int {
	return xxx_messageInfo_CommonRequest.Size(m)
}
func (m *CommonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommonRequest proto.InternalMessageInfo

func (m *CommonRequest) GetAccessToken() bool {
	if m != nil {
		return m.AccessToken
	}
	return false
}

func (m *CommonRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type InformaticaRequest struct {
	Informatica          *Informatica `protobuf:"bytes,1,opt,name=informatica,proto3" json:"informatica,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *InformaticaRequest) Reset()         { *m = InformaticaRequest{} }
func (m *InformaticaRequest) String() string { return proto.CompactTextString(m) }
func (*InformaticaRequest) ProtoMessage()    {}
func (*InformaticaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3c05494cc022a7, []int{3}
}

func (m *InformaticaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InformaticaRequest.Unmarshal(m, b)
}
func (m *InformaticaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InformaticaRequest.Marshal(b, m, deterministic)
}
func (m *InformaticaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InformaticaRequest.Merge(m, src)
}
func (m *InformaticaRequest) XXX_Size() int {
	return xxx_messageInfo_InformaticaRequest.Size(m)
}
func (m *InformaticaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InformaticaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InformaticaRequest proto.InternalMessageInfo

func (m *InformaticaRequest) GetInformatica() *Informatica {
	if m != nil {
		return m.Informatica
	}
	return nil
}

type InformaticaResponse struct {
	CommonResponse       *CommonResponse `protobuf:"bytes,1,opt,name=common_response,json=commonResponse,proto3" json:"common_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *InformaticaResponse) Reset()         { *m = InformaticaResponse{} }
func (m *InformaticaResponse) String() string { return proto.CompactTextString(m) }
func (*InformaticaResponse) ProtoMessage()    {}
func (*InformaticaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3c05494cc022a7, []int{4}
}

func (m *InformaticaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InformaticaResponse.Unmarshal(m, b)
}
func (m *InformaticaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InformaticaResponse.Marshal(b, m, deterministic)
}
func (m *InformaticaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InformaticaResponse.Merge(m, src)
}
func (m *InformaticaResponse) XXX_Size() int {
	return xxx_messageInfo_InformaticaResponse.Size(m)
}
func (m *InformaticaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InformaticaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InformaticaResponse proto.InternalMessageInfo

func (m *InformaticaResponse) GetCommonResponse() *CommonResponse {
	if m != nil {
		return m.CommonResponse
	}
	return nil
}

type GetInformaticaResponse struct {
	CommonResponse       *CommonResponse `protobuf:"bytes,1,opt,name=common_response,json=commonResponse,proto3" json:"common_response,omitempty"`
	Data                 []*Informatica  `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetInformaticaResponse) Reset()         { *m = GetInformaticaResponse{} }
func (m *GetInformaticaResponse) String() string { return proto.CompactTextString(m) }
func (*GetInformaticaResponse) ProtoMessage()    {}
func (*GetInformaticaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3c05494cc022a7, []int{5}
}

func (m *GetInformaticaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInformaticaResponse.Unmarshal(m, b)
}
func (m *GetInformaticaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInformaticaResponse.Marshal(b, m, deterministic)
}
func (m *GetInformaticaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInformaticaResponse.Merge(m, src)
}
func (m *GetInformaticaResponse) XXX_Size() int {
	return xxx_messageInfo_GetInformaticaResponse.Size(m)
}
func (m *GetInformaticaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInformaticaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInformaticaResponse proto.InternalMessageInfo

func (m *GetInformaticaResponse) GetCommonResponse() *CommonResponse {
	if m != nil {
		return m.CommonResponse
	}
	return nil
}

func (m *GetInformaticaResponse) GetData() []*Informatica {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdateInformaticaRequest struct {
	Informatica          *Informatica `protobuf:"bytes,1,opt,name=informatica,proto3" json:"informatica,omitempty"`
	UpdateSequence       int32        `protobuf:"varint,2,opt,name=update_sequence,json=updateSequence,proto3" json:"update_sequence,omitempty"`
	HostName             string       `protobuf:"bytes,3,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateInformaticaRequest) Reset()         { *m = UpdateInformaticaRequest{} }
func (m *UpdateInformaticaRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInformaticaRequest) ProtoMessage()    {}
func (*UpdateInformaticaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3c05494cc022a7, []int{6}
}

func (m *UpdateInformaticaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInformaticaRequest.Unmarshal(m, b)
}
func (m *UpdateInformaticaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInformaticaRequest.Marshal(b, m, deterministic)
}
func (m *UpdateInformaticaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInformaticaRequest.Merge(m, src)
}
func (m *UpdateInformaticaRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateInformaticaRequest.Size(m)
}
func (m *UpdateInformaticaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInformaticaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInformaticaRequest proto.InternalMessageInfo

func (m *UpdateInformaticaRequest) GetInformatica() *Informatica {
	if m != nil {
		return m.Informatica
	}
	return nil
}

func (m *UpdateInformaticaRequest) GetUpdateSequence() int32 {
	if m != nil {
		return m.UpdateSequence
	}
	return 0
}

func (m *UpdateInformaticaRequest) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func init() {
	proto.RegisterType((*Informatica)(nil), "informatica.Informatica")
	proto.RegisterType((*CommonResponse)(nil), "informatica.CommonResponse")
	proto.RegisterType((*CommonRequest)(nil), "informatica.CommonRequest")
	proto.RegisterType((*InformaticaRequest)(nil), "informatica.InformaticaRequest")
	proto.RegisterType((*InformaticaResponse)(nil), "informatica.InformaticaResponse")
	proto.RegisterType((*GetInformaticaResponse)(nil), "informatica.GetInformaticaResponse")
	proto.RegisterType((*UpdateInformaticaRequest)(nil), "informatica.UpdateInformaticaRequest")
}

func init() {
	proto.RegisterFile("expert/crud-api-using-mongodb/proto/informatica.proto", fileDescriptor_be3c05494cc022a7)
}

var fileDescriptor_be3c05494cc022a7 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x55, 0x76, 0xb7, 0x65, 0x3b, 0x4b, 0xb3, 0xea, 0x80, 0x2a, 0x6b, 0x7b, 0x20, 0x04, 0x21,
	0x7a, 0x60, 0xbb, 0x52, 0x11, 0x17, 0x8e, 0x2d, 0x12, 0x70, 0x41, 0x90, 0x02, 0x07, 0x40, 0x8a,
	0xbc, 0xce, 0xb0, 0x44, 0xd4, 0x76, 0x1a, 0x3b, 0x88, 0x8f, 0xe0, 0x13, 0xe0, 0x5f, 0x51, 0x9c,
	0x64, 0x89, 0x05, 0x81, 0x0b, 0xdc, 0xf2, 0x9e, 0xed, 0x37, 0x6f, 0xde, 0x8c, 0x02, 0x0f, 0xe9,
	0x4b, 0x41, 0xa5, 0x5d, 0x89, 0xb2, 0xca, 0x96, 0xbc, 0xc8, 0x97, 0x95, 0xc9, 0xd5, 0x66, 0x29,
	0xb5, 0xda, 0xe8, 0x6c, 0xbd, 0x2a, 0x4a, 0x6d, 0xf5, 0x2a, 0x57, 0x1f, 0x74, 0x29, 0xb9, 0xcd,
	0x05, 0x3f, 0x71, 0x0c, 0xce, 0x7a, 0x54, 0x5c, 0xc0, 0xec, 0xd9, 0x4f, 0x88, 0x0b, 0x98, 0x1a,
	0xba, 0xaa, 0x48, 0x09, 0x62, 0x41, 0x14, 0x1c, 0xef, 0x24, 0x5b, 0x8c, 0x37, 0x61, 0xc7, 0xe6,
	0xf6, 0x92, 0xd8, 0x28, 0x0a, 0x8e, 0xf7, 0x92, 0x06, 0x20, 0xc2, 0xa4, 0xd6, 0x63, 0x63, 0x47,
	0xba, 0x6f, 0x3c, 0x82, 0xbd, 0x8f, 0xda, 0xd8, 0x54, 0x71, 0x49, 0x6c, 0xe2, 0x0e, 0xa6, 0x35,
	0xf1, 0x9c, 0x4b, 0x8a, 0xcf, 0x20, 0x3c, 0xd7, 0x52, 0x6a, 0x95, 0x90, 0x29, 0xb4, 0x32, 0x84,
	0x87, 0xb0, 0x6b, 0x2c, 0xb7, 0x95, 0x69, 0x4b, 0xb6, 0x08, 0x19, 0x5c, 0x93, 0x64, 0x0c, 0xdf,
	0x74, 0x25, 0x3b, 0x18, 0x3f, 0x85, 0xfd, 0x4e, 0xe3, 0xaa, 0x22, 0x63, 0xf1, 0x36, 0x5c, 0xe7,
	0x42, 0x90, 0x31, 0xa9, 0xd5, 0x9f, 0x48, 0x39, 0xa1, 0x69, 0x32, 0x6b, 0xb8, 0x57, 0x35, 0x55,
	0xdb, 0x27, 0xc9, 0xf3, 0xcb, 0xce, 0xbe, 0x03, 0xf1, 0x0b, 0xc0, 0x5e, 0xff, 0x9d, 0xdc, 0x23,
	0xe8, 0x87, 0xe4, 0xd4, 0x66, 0xa7, 0xec, 0xa4, 0x9f, 0x65, 0xff, 0x95, 0x97, 0xe8, 0x3b, 0xb8,
	0xe1, 0x29, 0xb6, 0x4d, 0x3e, 0x86, 0xb9, 0x70, 0x96, 0xd3, 0xb2, 0xa5, 0x5a, 0xd9, 0x23, 0x4f,
	0xd6, 0x8f, 0x26, 0x09, 0x85, 0x87, 0xe3, 0xaf, 0x01, 0x1c, 0x3e, 0x21, 0xfb, 0xdf, 0x0a, 0xe0,
	0x7d, 0x98, 0x64, 0xdc, 0x72, 0x36, 0x8a, 0xc6, 0x7f, 0x6c, 0xd9, 0xdd, 0x8a, 0xbf, 0x05, 0xc0,
	0x5e, 0x17, 0x19, 0xb7, 0xf4, 0x6f, 0x43, 0xc4, 0x7b, 0x30, 0xaf, 0x9c, 0x6e, 0xba, 0x5d, 0xc7,
	0x91, 0xdb, 0x8d, 0xb0, 0xa1, 0x2f, 0xba, 0xa5, 0xf4, 0x56, 0x6d, 0xec, 0xaf, 0xda, 0xe9, 0xf7,
	0x91, 0x37, 0xdd, 0x0b, 0x2a, 0x3f, 0xe7, 0x82, 0xf0, 0x0d, 0x1c, 0x9c, 0x97, 0xe4, 0x9b, 0xc6,
	0x5b, 0x83, 0xc6, 0x9a, 0x76, 0x16, 0xd1, 0xf0, 0x85, 0x36, 0xbb, 0x97, 0x10, 0xfa, 0xb3, 0xc1,
	0xc5, 0x6f, 0xa3, 0x6f, 0xf4, 0xee, 0x78, 0x67, 0x03, 0x43, 0x7d, 0x0f, 0x07, 0xbf, 0xe4, 0x8b,
	0x77, 0xbd, 0x97, 0x43, 0xf9, 0xff, 0xdd, 0xf0, 0xd9, 0xfc, 0xed, 0x7e, 0xef, 0x4a, 0xb1, 0x5e,
	0xef, 0xba, 0x3f, 0xc4, 0x83, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xf2, 0x75, 0x6d, 0x5a,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InformaticaServiceClient is the client API for InformaticaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InformaticaServiceClient interface {
	CreateInformatica(ctx context.Context, in *InformaticaRequest, opts ...grpc.CallOption) (*InformaticaResponse, error)
	GetInformatica(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetInformaticaResponse, error)
	UpdateInformatica(ctx context.Context, in *UpdateInformaticaRequest, opts ...grpc.CallOption) (*InformaticaResponse, error)
}

type informaticaServiceClient struct {
	cc *grpc.ClientConn
}

func NewInformaticaServiceClient(cc *grpc.ClientConn) InformaticaServiceClient {
	return &informaticaServiceClient{cc}
}

func (c *informaticaServiceClient) CreateInformatica(ctx context.Context, in *InformaticaRequest, opts ...grpc.CallOption) (*InformaticaResponse, error) {
	out := new(InformaticaResponse)
	err := c.cc.Invoke(ctx, "/informatica.InformaticaService/CreateInformatica", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informaticaServiceClient) GetInformatica(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetInformaticaResponse, error) {
	out := new(GetInformaticaResponse)
	err := c.cc.Invoke(ctx, "/informatica.InformaticaService/GetInformatica", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informaticaServiceClient) UpdateInformatica(ctx context.Context, in *UpdateInformaticaRequest, opts ...grpc.CallOption) (*InformaticaResponse, error) {
	out := new(InformaticaResponse)
	err := c.cc.Invoke(ctx, "/informatica.InformaticaService/UpdateInformatica", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformaticaServiceServer is the server API for InformaticaService service.
type InformaticaServiceServer interface {
	CreateInformatica(context.Context, *InformaticaRequest) (*InformaticaResponse, error)
	GetInformatica(context.Context, *CommonRequest) (*GetInformaticaResponse, error)
	UpdateInformatica(context.Context, *UpdateInformaticaRequest) (*InformaticaResponse, error)
}

func RegisterInformaticaServiceServer(s *grpc.Server, srv InformaticaServiceServer) {
	s.RegisterService(&_InformaticaService_serviceDesc, srv)
}

func _InformaticaService_CreateInformatica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InformaticaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformaticaServiceServer).CreateInformatica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/informatica.InformaticaService/CreateInformatica",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformaticaServiceServer).CreateInformatica(ctx, req.(*InformaticaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InformaticaService_GetInformatica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformaticaServiceServer).GetInformatica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/informatica.InformaticaService/GetInformatica",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformaticaServiceServer).GetInformatica(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InformaticaService_UpdateInformatica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInformaticaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformaticaServiceServer).UpdateInformatica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/informatica.InformaticaService/UpdateInformatica",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformaticaServiceServer).UpdateInformatica(ctx, req.(*UpdateInformaticaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InformaticaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "informatica.InformaticaService",
	HandlerType: (*InformaticaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInformatica",
			Handler:    _InformaticaService_CreateInformatica_Handler,
		},
		{
			MethodName: "GetInformatica",
			Handler:    _InformaticaService_GetInformatica_Handler,
		},
		{
			MethodName: "UpdateInformatica",
			Handler:    _InformaticaService_UpdateInformatica_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "expert/crud-api-using-mongodb/proto/informatica.proto",
}