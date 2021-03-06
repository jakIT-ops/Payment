// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment.proto

package pb

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

type TransactionRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	TransactionInfor     string   `protobuf:"bytes,2,opt,name=TransactionInfor,proto3" json:"TransactionInfor,omitempty"`
	DebitAmount          int64    `protobuf:"varint,3,opt,name=DebitAmount,proto3" json:"DebitAmount,omitempty"`
	CreditAmount         int64    `protobuf:"varint,4,opt,name=CreditAmount,proto3" json:"CreditAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{0}
}

func (m *TransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionRequest.Unmarshal(m, b)
}
func (m *TransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
}
func (m *TransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *TransactionRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionRequest.Size(m)
}
func (m *TransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func (m *TransactionRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *TransactionRequest) GetTransactionInfor() string {
	if m != nil {
		return m.TransactionInfor
	}
	return ""
}

func (m *TransactionRequest) GetDebitAmount() int64 {
	if m != nil {
		return m.DebitAmount
	}
	return 0
}

func (m *TransactionRequest) GetCreditAmount() int64 {
	if m != nil {
		return m.CreditAmount
	}
	return 0
}

type TransactionResponse struct {
	TransactionInfor     string   `protobuf:"bytes,1,opt,name=TransactionInfor,proto3" json:"TransactionInfor,omitempty"`
	TransactionId        int32    `protobuf:"varint,2,opt,name=TransactionId,proto3" json:"TransactionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{1}
}

func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetTransactionInfor() string {
	if m != nil {
		return m.TransactionInfor
	}
	return ""
}

func (m *TransactionResponse) GetTransactionId() int32 {
	if m != nil {
		return m.TransactionId
	}
	return 0
}

type Balance struct {
	Amount               int64    `protobuf:"varint,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{2}
}

func (m *Balance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Balance.Unmarshal(m, b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return xxx_messageInfo_Balance.Size(m)
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Account_Id struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account_Id) Reset()         { *m = Account_Id{} }
func (m *Account_Id) String() string { return proto.CompactTextString(m) }
func (*Account_Id) ProtoMessage()    {}
func (*Account_Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{3}
}

func (m *Account_Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account_Id.Unmarshal(m, b)
}
func (m *Account_Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account_Id.Marshal(b, m, deterministic)
}
func (m *Account_Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account_Id.Merge(m, src)
}
func (m *Account_Id) XXX_Size() int {
	return xxx_messageInfo_Account_Id.Size(m)
}
func (m *Account_Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Account_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Account_Id proto.InternalMessageInfo

func (m *Account_Id) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionRequest)(nil), "protos.TransactionRequest")
	proto.RegisterType((*TransactionResponse)(nil), "protos.TransactionResponse")
	proto.RegisterType((*Balance)(nil), "protos.Balance")
	proto.RegisterType((*Account_Id)(nil), "protos.Account_Id")
}

func init() { proto.RegisterFile("payment.proto", fileDescriptor_6362648dfa63d410) }

var fileDescriptor_6362648dfa63d410 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xac, 0xcc,
	0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x4b,
	0x18, 0xb9, 0x84, 0x42, 0x8a, 0x12, 0xf3, 0x8a, 0x13, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x64, 0xb8, 0x38, 0x1d, 0x93, 0x93, 0xf3, 0x4b, 0xf3, 0x4a,
	0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x10, 0x02, 0x42, 0x5a, 0x5c, 0x02, 0x48,
	0x7a, 0x3c, 0xf3, 0xd2, 0xf2, 0x8b, 0x24, 0x98, 0xc0, 0x8a, 0x30, 0xc4, 0x85, 0x14, 0xb8, 0xb8,
	0x5d, 0x52, 0x93, 0x32, 0x4b, 0x1c, 0x73, 0x41, 0x9a, 0x25, 0x98, 0x15, 0x18, 0x35, 0x98, 0x83,
	0x90, 0x85, 0x84, 0x94, 0xb8, 0x78, 0x9c, 0x8b, 0x52, 0x53, 0xe0, 0x4a, 0x58, 0xc0, 0x4a, 0x50,
	0xc4, 0x94, 0xd2, 0xb9, 0x84, 0x51, 0x5c, 0x59, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x8a, 0xd5, 0x21,
	0x8c, 0x38, 0x1c, 0xa2, 0xc2, 0xc5, 0x8b, 0x2c, 0x96, 0x02, 0x76, 0x31, 0x6b, 0x10, 0xaa, 0xa0,
	0x92, 0x22, 0x17, 0xbb, 0x53, 0x62, 0x4e, 0x62, 0x5e, 0x72, 0xaa, 0x90, 0x18, 0x17, 0x1b, 0xd4,
	0x45, 0x8c, 0x60, 0x17, 0x41, 0x79, 0x4a, 0x4a, 0x5c, 0x5c, 0xd0, 0xa0, 0x88, 0xf7, 0x4c, 0x11,
	0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x85, 0xda, 0x0b, 0xe1, 0x18, 0x5d, 0x64, 0xe4,
	0xe2, 0x46, 0x32, 0x58, 0xc8, 0x0d, 0x25, 0x14, 0x84, 0xa4, 0x20, 0xb1, 0x50, 0xac, 0x87, 0x19,
	0xf4, 0x52, 0xd2, 0x58, 0xe5, 0xa0, 0x1e, 0x76, 0x47, 0x0d, 0x2b, 0xf2, 0x0d, 0x32, 0xe4, 0xe2,
	0x72, 0x4f, 0x2d, 0x81, 0x79, 0x55, 0x08, 0xa6, 0x14, 0xe1, 0x31, 0x29, 0x7e, 0x98, 0x18, 0x54,
	0x91, 0x13, 0x4b, 0x14, 0x53, 0x41, 0x52, 0x12, 0x24, 0xe1, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x6c, 0x80, 0xa7, 0xb7, 0x50, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	DebitAmount(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	CreditAmount(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	GetBalance(ctx context.Context, in *Account_Id, opts ...grpc.CallOption) (*Balance, error)
}

type transactionClient struct {
	cc *grpc.ClientConn
}

func NewTransactionClient(cc *grpc.ClientConn) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) DebitAmount(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/protos.Transaction/DebitAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) CreditAmount(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/protos.Transaction/CreditAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) GetBalance(ctx context.Context, in *Account_Id, opts ...grpc.CallOption) (*Balance, error) {
	out := new(Balance)
	err := c.cc.Invoke(ctx, "/protos.Transaction/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	DebitAmount(context.Context, *TransactionRequest) (*TransactionResponse, error)
	CreditAmount(context.Context, *TransactionRequest) (*TransactionResponse, error)
	GetBalance(context.Context, *Account_Id) (*Balance, error)
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) DebitAmount(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DebitAmount not implemented")
}
func (*UnimplementedTransactionServer) CreditAmount(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreditAmount not implemented")
}
func (*UnimplementedTransactionServer) GetBalance(ctx context.Context, req *Account_Id) (*Balance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}

func RegisterTransactionServer(s *grpc.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_DebitAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).DebitAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Transaction/DebitAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).DebitAmount(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_CreditAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).CreditAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Transaction/CreditAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).CreditAmount(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account_Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Transaction/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).GetBalance(ctx, req.(*Account_Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DebitAmount",
			Handler:    _Transaction_DebitAmount_Handler,
		},
		{
			MethodName: "CreditAmount",
			Handler:    _Transaction_CreditAmount_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Transaction_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}