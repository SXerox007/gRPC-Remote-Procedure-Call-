// Code generated by protoc-gen-go. DO NOT EDIT.
// source: expert/deadline/proto/deadline.proto

package errorpb

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

type Color int32

const (
	Color_UNKNOWN Color = 0
	Color_RED     Color = 1
	Color_BLUE    Color = 2
	Color_YELLOW  Color = 3
)

var Color_name = map[int32]string{
	0: "UNKNOWN",
	1: "RED",
	2: "BLUE",
	3: "YELLOW",
}

var Color_value = map[string]int32{
	"UNKNOWN": 0,
	"RED":     1,
	"BLUE":    2,
	"YELLOW":  3,
}

func (x Color) String() string {
	return proto.EnumName(Color_name, int32(x))
}

func (Color) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dee6e907266ade1a, []int{0}
}

type EvenNumberRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvenNumberRequest) Reset()         { *m = EvenNumberRequest{} }
func (m *EvenNumberRequest) String() string { return proto.CompactTextString(m) }
func (*EvenNumberRequest) ProtoMessage()    {}
func (*EvenNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dee6e907266ade1a, []int{0}
}

func (m *EvenNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvenNumberRequest.Unmarshal(m, b)
}
func (m *EvenNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvenNumberRequest.Marshal(b, m, deterministic)
}
func (m *EvenNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvenNumberRequest.Merge(m, src)
}
func (m *EvenNumberRequest) XXX_Size() int {
	return xxx_messageInfo_EvenNumberRequest.Size(m)
}
func (m *EvenNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvenNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvenNumberRequest proto.InternalMessageInfo

func (m *EvenNumberRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type Name struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName           string   `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_dee6e907266ade1a, []int{1}
}

func (m *Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Name.Unmarshal(m, b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Name.Marshal(b, m, deterministic)
}
func (m *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(m, src)
}
func (m *Name) XXX_Size() int {
	return xxx_messageInfo_Name.Size(m)
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Name) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *Name) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type Data struct {
	Name                 *Name    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Color                Color    `protobuf:"varint,2,opt,name=color,proto3,enum=Color" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_dee6e907266ade1a, []int{2}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Data) GetColor() Color {
	if m != nil {
		return m.Color
	}
	return Color_UNKNOWN
}

type SuccessResponse struct {
	//repeated as data is multiple
	SuccessInfo          []*SuccessResponse_SuccessInfo `protobuf:"bytes,1,rep,name=success_info,json=successInfo,proto3" json:"success_info,omitempty"`
	IsEven               bool                           `protobuf:"varint,2,opt,name=isEven,proto3" json:"isEven,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *SuccessResponse) Reset()         { *m = SuccessResponse{} }
func (m *SuccessResponse) String() string { return proto.CompactTextString(m) }
func (*SuccessResponse) ProtoMessage()    {}
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dee6e907266ade1a, []int{3}
}

func (m *SuccessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessResponse.Unmarshal(m, b)
}
func (m *SuccessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessResponse.Marshal(b, m, deterministic)
}
func (m *SuccessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessResponse.Merge(m, src)
}
func (m *SuccessResponse) XXX_Size() int {
	return xxx_messageInfo_SuccessResponse.Size(m)
}
func (m *SuccessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessResponse proto.InternalMessageInfo

func (m *SuccessResponse) GetSuccessInfo() []*SuccessResponse_SuccessInfo {
	if m != nil {
		return m.SuccessInfo
	}
	return nil
}

func (m *SuccessResponse) GetIsEven() bool {
	if m != nil {
		return m.IsEven
	}
	return false
}

type SuccessResponse_SuccessInfo struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Data                 *Data    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessResponse_SuccessInfo) Reset()         { *m = SuccessResponse_SuccessInfo{} }
func (m *SuccessResponse_SuccessInfo) String() string { return proto.CompactTextString(m) }
func (*SuccessResponse_SuccessInfo) ProtoMessage()    {}
func (*SuccessResponse_SuccessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dee6e907266ade1a, []int{3, 0}
}

func (m *SuccessResponse_SuccessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessResponse_SuccessInfo.Unmarshal(m, b)
}
func (m *SuccessResponse_SuccessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessResponse_SuccessInfo.Marshal(b, m, deterministic)
}
func (m *SuccessResponse_SuccessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessResponse_SuccessInfo.Merge(m, src)
}
func (m *SuccessResponse_SuccessInfo) XXX_Size() int {
	return xxx_messageInfo_SuccessResponse_SuccessInfo.Size(m)
}
func (m *SuccessResponse_SuccessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessResponse_SuccessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessResponse_SuccessInfo proto.InternalMessageInfo

func (m *SuccessResponse_SuccessInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SuccessResponse_SuccessInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SuccessResponse_SuccessInfo) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("Color", Color_name, Color_value)
	proto.RegisterType((*EvenNumberRequest)(nil), "EvenNumberRequest")
	proto.RegisterType((*Name)(nil), "Name")
	proto.RegisterType((*Data)(nil), "Data")
	proto.RegisterType((*SuccessResponse)(nil), "SuccessResponse")
	proto.RegisterType((*SuccessResponse_SuccessInfo)(nil), "SuccessResponse.SuccessInfo")
}

func init() {
	proto.RegisterFile("expert/deadline/proto/deadline.proto", fileDescriptor_dee6e907266ade1a)
}

var fileDescriptor_dee6e907266ade1a = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x51, 0x6b, 0xd4, 0x40,
	0x14, 0x85, 0x4d, 0x77, 0xb3, 0xbb, 0xb9, 0x11, 0x8d, 0xf3, 0xa0, 0x6b, 0xad, 0x58, 0x82, 0x0f,
	0x45, 0x21, 0x85, 0xf4, 0xc5, 0xb7, 0x42, 0xdd, 0x80, 0xe2, 0x92, 0xc2, 0x2c, 0xa5, 0xd8, 0x97,
	0x32, 0x9b, 0xdc, 0x2d, 0x81, 0x64, 0x26, 0xce, 0x4c, 0x8a, 0xbf, 0xd0, 0xdf, 0x25, 0x73, 0x93,
	0xac, 0xb2, 0x7d, 0x0a, 0xdf, 0x39, 0x77, 0x6e, 0x38, 0x67, 0x06, 0x3e, 0xe2, 0xef, 0x16, 0xb5,
	0x3d, 0x2f, 0x51, 0x94, 0x75, 0x25, 0xf1, 0xbc, 0xd5, 0xca, 0xaa, 0x3d, 0x26, 0x84, 0xf1, 0x67,
	0x78, 0x95, 0x3d, 0xa2, 0xcc, 0xbb, 0x66, 0x8b, 0x9a, 0xe3, 0xaf, 0x0e, 0x8d, 0x65, 0xaf, 0x61,
	0x26, 0x49, 0x58, 0x7a, 0xa7, 0xde, 0x99, 0xcf, 0x07, 0x8a, 0x0b, 0x98, 0xe6, 0xa2, 0x41, 0xf6,
	0x1e, 0x60, 0x57, 0x69, 0x63, 0xef, 0xa5, 0x68, 0x90, 0x66, 0x02, 0x1e, 0x90, 0x42, 0xf6, 0x07,
	0x08, 0x9b, 0xaa, 0x2c, 0x6b, 0xec, 0xfd, 0x23, 0xf2, 0xa1, 0x97, 0x68, 0xe0, 0x1d, 0x04, 0xb5,
	0x18, 0x8f, 0x4f, 0xc8, 0x5e, 0x38, 0xc1, 0x99, 0xf1, 0x25, 0x4c, 0x57, 0xc2, 0x0a, 0xf6, 0x16,
	0xa6, 0xfb, 0xf5, 0x61, 0xea, 0x27, 0xce, 0xe4, 0x24, 0xb1, 0x13, 0xf0, 0x0b, 0x55, 0x2b, 0x4d,
	0xab, 0x5f, 0xa4, 0xb3, 0xe4, 0xab, 0x23, 0xde, 0x8b, 0xf1, 0x1f, 0x0f, 0x5e, 0x6e, 0xba, 0xa2,
	0x40, 0x63, 0x38, 0x9a, 0x56, 0x49, 0x83, 0xec, 0x12, 0x9e, 0x9b, 0x5e, 0xba, 0xaf, 0xe4, 0x4e,
	0x2d, 0xbd, 0xd3, 0xc9, 0x59, 0x98, 0x9e, 0x24, 0x07, 0x73, 0x23, 0x7f, 0x97, 0x3b, 0xc5, 0x43,
	0xf3, 0x0f, 0x5c, 0x25, 0x95, 0x71, 0x4d, 0xd1, 0x3f, 0x17, 0x7c, 0xa0, 0xe3, 0x3b, 0x08, 0xff,
	0x3b, 0xc3, 0x96, 0x30, 0x6f, 0xd0, 0x18, 0xf1, 0x30, 0xd6, 0x32, 0xa2, 0x5b, 0x60, 0xac, 0xb0,
	0x9d, 0xa1, 0x05, 0x3e, 0x1f, 0xc8, 0xc5, 0x2c, 0x85, 0x15, 0x54, 0x83, 0x8b, 0xe9, 0xb2, 0x73,
	0x92, 0x3e, 0x5d, 0x80, 0x4f, 0xc1, 0x58, 0x08, 0xf3, 0x9b, 0xfc, 0x47, 0x7e, 0x7d, 0x9b, 0x47,
	0xcf, 0xd8, 0x1c, 0x26, 0x3c, 0x5b, 0x45, 0x1e, 0x5b, 0xc0, 0xf4, 0x6a, 0x7d, 0x93, 0x45, 0x47,
	0x0c, 0x60, 0xf6, 0x33, 0x5b, 0xaf, 0xaf, 0x6f, 0xa3, 0x49, 0xba, 0x81, 0x37, 0xab, 0xe1, 0x8a,
	0xbf, 0x09, 0xe9, 0xbe, 0x0f, 0x1b, 0xd4, 0x8f, 0x55, 0x81, 0xec, 0x0b, 0x44, 0x87, 0x16, 0x63,
	0xc9, 0x93, 0xeb, 0x3f, 0x8e, 0x0e, 0x6b, 0xb9, 0x0a, 0xee, 0xe6, 0xa8, 0xb5, 0xd2, 0xed, 0x76,
	0x3b, 0xa3, 0x77, 0x73, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x88, 0xa2, 0x4f, 0xbd, 0x5f, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeadlineHandlingServiceClient is the client API for DeadlineHandlingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeadlineHandlingServiceClient interface {
	DeadlineHandling(ctx context.Context, in *EvenNumberRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type deadlineHandlingServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeadlineHandlingServiceClient(cc *grpc.ClientConn) DeadlineHandlingServiceClient {
	return &deadlineHandlingServiceClient{cc}
}

func (c *deadlineHandlingServiceClient) DeadlineHandling(ctx context.Context, in *EvenNumberRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/DeadlineHandlingService/DeadlineHandling", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeadlineHandlingServiceServer is the server API for DeadlineHandlingService service.
type DeadlineHandlingServiceServer interface {
	DeadlineHandling(context.Context, *EvenNumberRequest) (*SuccessResponse, error)
}

func RegisterDeadlineHandlingServiceServer(s *grpc.Server, srv DeadlineHandlingServiceServer) {
	s.RegisterService(&_DeadlineHandlingService_serviceDesc, srv)
}

func _DeadlineHandlingService_DeadlineHandling_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvenNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadlineHandlingServiceServer).DeadlineHandling(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeadlineHandlingService/DeadlineHandling",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadlineHandlingServiceServer).DeadlineHandling(ctx, req.(*EvenNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeadlineHandlingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DeadlineHandlingService",
	HandlerType: (*DeadlineHandlingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeadlineHandling",
			Handler:    _DeadlineHandlingService_DeadlineHandling_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "expert/deadline/proto/deadline.proto",
}
