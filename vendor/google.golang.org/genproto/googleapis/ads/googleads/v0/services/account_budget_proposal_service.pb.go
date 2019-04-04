// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/account_budget_proposal_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

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

// Request message for
// [AccountBudgetProposalService.GetAccountBudgetProposal][google.ads.googleads.v0.services.AccountBudgetProposalService.GetAccountBudgetProposal].
type GetAccountBudgetProposalRequest struct {
	// The resource name of the account-level budget proposal to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountBudgetProposalRequest) Reset()         { *m = GetAccountBudgetProposalRequest{} }
func (m *GetAccountBudgetProposalRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountBudgetProposalRequest) ProtoMessage()    {}
func (*GetAccountBudgetProposalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_service_e914761b5f55b9ea, []int{0}
}
func (m *GetAccountBudgetProposalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountBudgetProposalRequest.Unmarshal(m, b)
}
func (m *GetAccountBudgetProposalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountBudgetProposalRequest.Marshal(b, m, deterministic)
}
func (dst *GetAccountBudgetProposalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountBudgetProposalRequest.Merge(dst, src)
}
func (m *GetAccountBudgetProposalRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountBudgetProposalRequest.Size(m)
}
func (m *GetAccountBudgetProposalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountBudgetProposalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountBudgetProposalRequest proto.InternalMessageInfo

func (m *GetAccountBudgetProposalRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [AccountBudgetProposalService.MutateAccountBudgetProposal][google.ads.googleads.v0.services.AccountBudgetProposalService.MutateAccountBudgetProposal].
type MutateAccountBudgetProposalRequest struct {
	// The ID of the customer.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform on an individual account-level budget proposal.
	Operation            *AccountBudgetProposalOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MutateAccountBudgetProposalRequest) Reset()         { *m = MutateAccountBudgetProposalRequest{} }
func (m *MutateAccountBudgetProposalRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAccountBudgetProposalRequest) ProtoMessage()    {}
func (*MutateAccountBudgetProposalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_service_e914761b5f55b9ea, []int{1}
}
func (m *MutateAccountBudgetProposalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAccountBudgetProposalRequest.Unmarshal(m, b)
}
func (m *MutateAccountBudgetProposalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAccountBudgetProposalRequest.Marshal(b, m, deterministic)
}
func (dst *MutateAccountBudgetProposalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAccountBudgetProposalRequest.Merge(dst, src)
}
func (m *MutateAccountBudgetProposalRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAccountBudgetProposalRequest.Size(m)
}
func (m *MutateAccountBudgetProposalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAccountBudgetProposalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAccountBudgetProposalRequest proto.InternalMessageInfo

func (m *MutateAccountBudgetProposalRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAccountBudgetProposalRequest) GetOperation() *AccountBudgetProposalOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// A single operation to propose the creation of a new account-level budget or
// edit/end/remove an existing one.
type AccountBudgetProposalOperation struct {
	// FieldMask that determines which budget fields are modified.  While budgets
	// may be modified, proposals that propose such modifications are final.
	// Therefore, update operations are not supported for proposals.
	//
	// Proposals that modify budgets have the 'update' proposal type.  Specifying
	// a mask for any other proposal type is considered an error.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AccountBudgetProposalOperation_Create
	//	*AccountBudgetProposalOperation_Remove
	Operation            isAccountBudgetProposalOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *AccountBudgetProposalOperation) Reset()         { *m = AccountBudgetProposalOperation{} }
func (m *AccountBudgetProposalOperation) String() string { return proto.CompactTextString(m) }
func (*AccountBudgetProposalOperation) ProtoMessage()    {}
func (*AccountBudgetProposalOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_service_e914761b5f55b9ea, []int{2}
}
func (m *AccountBudgetProposalOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudgetProposalOperation.Unmarshal(m, b)
}
func (m *AccountBudgetProposalOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudgetProposalOperation.Marshal(b, m, deterministic)
}
func (dst *AccountBudgetProposalOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudgetProposalOperation.Merge(dst, src)
}
func (m *AccountBudgetProposalOperation) XXX_Size() int {
	return xxx_messageInfo_AccountBudgetProposalOperation.Size(m)
}
func (m *AccountBudgetProposalOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudgetProposalOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudgetProposalOperation proto.InternalMessageInfo

func (m *AccountBudgetProposalOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAccountBudgetProposalOperation_Operation interface {
	isAccountBudgetProposalOperation_Operation()
}

type AccountBudgetProposalOperation_Create struct {
	Create *resources.AccountBudgetProposal `protobuf:"bytes,2,opt,name=create,proto3,oneof"`
}

type AccountBudgetProposalOperation_Remove struct {
	Remove string `protobuf:"bytes,1,opt,name=remove,proto3,oneof"`
}

func (*AccountBudgetProposalOperation_Create) isAccountBudgetProposalOperation_Operation() {}

func (*AccountBudgetProposalOperation_Remove) isAccountBudgetProposalOperation_Operation() {}

func (m *AccountBudgetProposalOperation) GetOperation() isAccountBudgetProposalOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AccountBudgetProposalOperation) GetCreate() *resources.AccountBudgetProposal {
	if x, ok := m.GetOperation().(*AccountBudgetProposalOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AccountBudgetProposalOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AccountBudgetProposalOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AccountBudgetProposalOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AccountBudgetProposalOperation_OneofMarshaler, _AccountBudgetProposalOperation_OneofUnmarshaler, _AccountBudgetProposalOperation_OneofSizer, []interface{}{
		(*AccountBudgetProposalOperation_Create)(nil),
		(*AccountBudgetProposalOperation_Remove)(nil),
	}
}

func _AccountBudgetProposalOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AccountBudgetProposalOperation)
	// operation
	switch x := m.Operation.(type) {
	case *AccountBudgetProposalOperation_Create:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *AccountBudgetProposalOperation_Remove:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("AccountBudgetProposalOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _AccountBudgetProposalOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AccountBudgetProposalOperation)
	switch tag {
	case 2: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.AccountBudgetProposal)
		err := b.DecodeMessage(msg)
		m.Operation = &AccountBudgetProposalOperation_Create{msg}
		return true, err
	case 1: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &AccountBudgetProposalOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _AccountBudgetProposalOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AccountBudgetProposalOperation)
	// operation
	switch x := m.Operation.(type) {
	case *AccountBudgetProposalOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountBudgetProposalOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for account-level budget mutate operations.
type MutateAccountBudgetProposalResponse struct {
	// The result of the mutate.
	Result               *MutateAccountBudgetProposalResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *MutateAccountBudgetProposalResponse) Reset()         { *m = MutateAccountBudgetProposalResponse{} }
func (m *MutateAccountBudgetProposalResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAccountBudgetProposalResponse) ProtoMessage()    {}
func (*MutateAccountBudgetProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_service_e914761b5f55b9ea, []int{3}
}
func (m *MutateAccountBudgetProposalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAccountBudgetProposalResponse.Unmarshal(m, b)
}
func (m *MutateAccountBudgetProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAccountBudgetProposalResponse.Marshal(b, m, deterministic)
}
func (dst *MutateAccountBudgetProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAccountBudgetProposalResponse.Merge(dst, src)
}
func (m *MutateAccountBudgetProposalResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAccountBudgetProposalResponse.Size(m)
}
func (m *MutateAccountBudgetProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAccountBudgetProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAccountBudgetProposalResponse proto.InternalMessageInfo

func (m *MutateAccountBudgetProposalResponse) GetResult() *MutateAccountBudgetProposalResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// The result for the account budget proposal mutate.
type MutateAccountBudgetProposalResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAccountBudgetProposalResult) Reset()         { *m = MutateAccountBudgetProposalResult{} }
func (m *MutateAccountBudgetProposalResult) String() string { return proto.CompactTextString(m) }
func (*MutateAccountBudgetProposalResult) ProtoMessage()    {}
func (*MutateAccountBudgetProposalResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_service_e914761b5f55b9ea, []int{4}
}
func (m *MutateAccountBudgetProposalResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAccountBudgetProposalResult.Unmarshal(m, b)
}
func (m *MutateAccountBudgetProposalResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAccountBudgetProposalResult.Marshal(b, m, deterministic)
}
func (dst *MutateAccountBudgetProposalResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAccountBudgetProposalResult.Merge(dst, src)
}
func (m *MutateAccountBudgetProposalResult) XXX_Size() int {
	return xxx_messageInfo_MutateAccountBudgetProposalResult.Size(m)
}
func (m *MutateAccountBudgetProposalResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAccountBudgetProposalResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAccountBudgetProposalResult proto.InternalMessageInfo

func (m *MutateAccountBudgetProposalResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAccountBudgetProposalRequest)(nil), "google.ads.googleads.v0.services.GetAccountBudgetProposalRequest")
	proto.RegisterType((*MutateAccountBudgetProposalRequest)(nil), "google.ads.googleads.v0.services.MutateAccountBudgetProposalRequest")
	proto.RegisterType((*AccountBudgetProposalOperation)(nil), "google.ads.googleads.v0.services.AccountBudgetProposalOperation")
	proto.RegisterType((*MutateAccountBudgetProposalResponse)(nil), "google.ads.googleads.v0.services.MutateAccountBudgetProposalResponse")
	proto.RegisterType((*MutateAccountBudgetProposalResult)(nil), "google.ads.googleads.v0.services.MutateAccountBudgetProposalResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountBudgetProposalServiceClient is the client API for AccountBudgetProposalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountBudgetProposalServiceClient interface {
	// Returns an account-level budget proposal in full detail.
	GetAccountBudgetProposal(ctx context.Context, in *GetAccountBudgetProposalRequest, opts ...grpc.CallOption) (*resources.AccountBudgetProposal, error)
	// Creates, updates, or removes account budget proposals.  Operation statuses
	// are returned.
	MutateAccountBudgetProposal(ctx context.Context, in *MutateAccountBudgetProposalRequest, opts ...grpc.CallOption) (*MutateAccountBudgetProposalResponse, error)
}

type accountBudgetProposalServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountBudgetProposalServiceClient(cc *grpc.ClientConn) AccountBudgetProposalServiceClient {
	return &accountBudgetProposalServiceClient{cc}
}

func (c *accountBudgetProposalServiceClient) GetAccountBudgetProposal(ctx context.Context, in *GetAccountBudgetProposalRequest, opts ...grpc.CallOption) (*resources.AccountBudgetProposal, error) {
	out := new(resources.AccountBudgetProposal)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AccountBudgetProposalService/GetAccountBudgetProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountBudgetProposalServiceClient) MutateAccountBudgetProposal(ctx context.Context, in *MutateAccountBudgetProposalRequest, opts ...grpc.CallOption) (*MutateAccountBudgetProposalResponse, error) {
	out := new(MutateAccountBudgetProposalResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AccountBudgetProposalService/MutateAccountBudgetProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountBudgetProposalServiceServer is the server API for AccountBudgetProposalService service.
type AccountBudgetProposalServiceServer interface {
	// Returns an account-level budget proposal in full detail.
	GetAccountBudgetProposal(context.Context, *GetAccountBudgetProposalRequest) (*resources.AccountBudgetProposal, error)
	// Creates, updates, or removes account budget proposals.  Operation statuses
	// are returned.
	MutateAccountBudgetProposal(context.Context, *MutateAccountBudgetProposalRequest) (*MutateAccountBudgetProposalResponse, error)
}

func RegisterAccountBudgetProposalServiceServer(s *grpc.Server, srv AccountBudgetProposalServiceServer) {
	s.RegisterService(&_AccountBudgetProposalService_serviceDesc, srv)
}

func _AccountBudgetProposalService_GetAccountBudgetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBudgetProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountBudgetProposalServiceServer).GetAccountBudgetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AccountBudgetProposalService/GetAccountBudgetProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountBudgetProposalServiceServer).GetAccountBudgetProposal(ctx, req.(*GetAccountBudgetProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountBudgetProposalService_MutateAccountBudgetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAccountBudgetProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountBudgetProposalServiceServer).MutateAccountBudgetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AccountBudgetProposalService/MutateAccountBudgetProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountBudgetProposalServiceServer).MutateAccountBudgetProposal(ctx, req.(*MutateAccountBudgetProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountBudgetProposalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.AccountBudgetProposalService",
	HandlerType: (*AccountBudgetProposalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountBudgetProposal",
			Handler:    _AccountBudgetProposalService_GetAccountBudgetProposal_Handler,
		},
		{
			MethodName: "MutateAccountBudgetProposal",
			Handler:    _AccountBudgetProposalService_MutateAccountBudgetProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/account_budget_proposal_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/account_budget_proposal_service.proto", fileDescriptor_account_budget_proposal_service_e914761b5f55b9ea)
}

var fileDescriptor_account_budget_proposal_service_e914761b5f55b9ea = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x3f, 0x6f, 0xd3, 0x40,
	0x1c, 0xc5, 0xae, 0x14, 0xa9, 0x17, 0x58, 0x6e, 0x8a, 0x42, 0x45, 0x53, 0xb7, 0x43, 0x95, 0xc1,
	0x8e, 0xca, 0x52, 0x5d, 0x54, 0x11, 0x07, 0x48, 0xc2, 0x50, 0x88, 0x8c, 0x94, 0x01, 0x22, 0xac,
	0x8b, 0x7d, 0xb5, 0xac, 0xda, 0x3e, 0xe3, 0x3b, 0x67, 0xa9, 0xba, 0x74, 0x62, 0xe7, 0x1b, 0x20,
	0xb1, 0xf0, 0x51, 0x58, 0x59, 0x60, 0x67, 0x62, 0x41, 0xe2, 0x13, 0x20, 0xfb, 0xee, 0xd2, 0x56,
	0x4a, 0x6c, 0x44, 0xb7, 0x9f, 0xcf, 0xef, 0xde, 0x7b, 0xbf, 0x7f, 0x07, 0x46, 0x01, 0xa5, 0x41,
	0x44, 0x2c, 0xec, 0x33, 0x4b, 0x84, 0x45, 0xb4, 0xec, 0x59, 0x8c, 0x64, 0xcb, 0xd0, 0x23, 0xcc,
	0xc2, 0x9e, 0x47, 0xf3, 0x84, 0xbb, 0x8b, 0xdc, 0x0f, 0x08, 0x77, 0xd3, 0x8c, 0xa6, 0x94, 0xe1,
	0xc8, 0x95, 0x00, 0x33, 0xcd, 0x28, 0xa7, 0xb0, 0x23, 0x2e, 0x9b, 0xd8, 0x67, 0xe6, 0x8a, 0xc7,
	0x5c, 0xf6, 0x4c, 0xc5, 0xd3, 0x7e, 0xb2, 0x49, 0x29, 0x23, 0x8c, 0xe6, 0x59, 0x85, 0x94, 0x90,
	0x68, 0xef, 0x28, 0x82, 0x34, 0xb4, 0x70, 0x92, 0x50, 0x8e, 0x79, 0x48, 0x13, 0x26, 0xff, 0x4a,
	0x03, 0x56, 0xf9, 0xb5, 0xc8, 0xcf, 0xac, 0xb3, 0x90, 0x44, 0xbe, 0x1b, 0x63, 0x76, 0x2e, 0x10,
	0xc6, 0x08, 0xec, 0x8e, 0x09, 0xb7, 0x85, 0xc6, 0xb0, 0x94, 0x98, 0x4a, 0x05, 0x87, 0xbc, 0xcf,
	0x09, 0xe3, 0x70, 0x1f, 0x3c, 0x50, 0x6e, 0xdc, 0x04, 0xc7, 0xa4, 0xa5, 0x75, 0xb4, 0xc3, 0x6d,
	0xe7, 0xbe, 0x3a, 0x7c, 0x89, 0x63, 0x62, 0x7c, 0xd6, 0x80, 0x71, 0x9a, 0x73, 0xcc, 0x49, 0x25,
	0xd7, 0x2e, 0x68, 0x7a, 0x39, 0xe3, 0x34, 0x26, 0x99, 0x1b, 0xfa, 0x92, 0x09, 0xa8, 0xa3, 0x17,
	0x3e, 0x7c, 0x07, 0xb6, 0x69, 0x4a, 0xb2, 0x32, 0x8b, 0x96, 0xde, 0xd1, 0x0e, 0x9b, 0x47, 0x03,
	0xb3, 0xae, 0x8c, 0xe6, 0x5a, 0xcd, 0x57, 0x8a, 0xc7, 0xb9, 0xa6, 0x34, 0xbe, 0x6b, 0xe0, 0x51,
	0x35, 0x1a, 0xf6, 0x41, 0x33, 0x4f, 0x7d, 0xcc, 0x49, 0x59, 0xa7, 0xd6, 0x56, 0x69, 0xa2, 0xad,
	0x4c, 0xa8, 0x52, 0x9a, 0xa3, 0xa2, 0x94, 0xa7, 0x98, 0x9d, 0x3b, 0x40, 0xc0, 0x8b, 0x18, 0x3a,
	0xa0, 0xe1, 0x65, 0x04, 0x73, 0x22, 0xcd, 0x1f, 0x6f, 0x34, 0xbf, 0xea, 0xf0, 0x7a, 0xf7, 0x93,
	0x7b, 0x8e, 0x64, 0x82, 0x2d, 0xd0, 0xc8, 0x48, 0x4c, 0x97, 0xb2, 0xf2, 0xc5, 0x1f, 0xf1, 0x3d,
	0x6c, 0xde, 0xa8, 0x96, 0x71, 0xa5, 0x81, 0xfd, 0xca, 0x16, 0xb0, 0x94, 0x26, 0x8c, 0xc0, 0xb7,
	0x05, 0x1d, 0xcb, 0x23, 0x2e, 0x2d, 0x3e, 0xad, 0xaf, 0x6f, 0x35, 0x6d, 0x1e, 0x71, 0x47, 0x52,
	0x1a, 0x13, 0xb0, 0x57, 0x0b, 0xfe, 0xa7, 0x89, 0x3a, 0xfa, 0xbd, 0x05, 0x76, 0xd6, 0x92, 0xbc,
	0x16, 0xae, 0xe0, 0x0f, 0x0d, 0xb4, 0x36, 0xcd, 0x2e, 0xb4, 0xeb, 0x93, 0xaa, 0x99, 0xfb, 0xf6,
	0x7f, 0xb7, 0xce, 0x18, 0x5c, 0x7d, 0xfb, 0xf9, 0x51, 0x47, 0xf0, 0xb8, 0xd8, 0xe4, 0x8b, 0x5b,
	0xa9, 0x9e, 0xa8, 0x59, 0x67, 0x56, 0x57, 0xad, 0xf6, 0xed, 0xdb, 0xcc, 0xea, 0x5e, 0xc2, 0x3f,
	0x1a, 0x78, 0x58, 0x51, 0x47, 0xf8, 0xec, 0x8e, 0x3d, 0x13, 0x19, 0x3e, 0xbf, 0x6b, 0xe7, 0xcb,
	0x81, 0x32, 0x46, 0x65, 0xba, 0x03, 0xa3, 0x5f, 0xa4, 0x7b, 0x9d, 0xdf, 0xc5, 0x8d, 0x4d, 0x3f,
	0xe9, 0x5e, 0x6e, 0xc8, 0x16, 0xc5, 0xa5, 0x02, 0xd2, 0xba, 0xc3, 0x0f, 0x3a, 0x38, 0xf0, 0x68,
	0x5c, 0x6b, 0x6a, 0xb8, 0x57, 0x35, 0x17, 0xd3, 0x62, 0x41, 0xa7, 0xda, 0x9b, 0x89, 0xa4, 0x09,
	0x68, 0x84, 0x93, 0xc0, 0xa4, 0x59, 0x60, 0x05, 0x24, 0x29, 0xd7, 0x57, 0x3d, 0xb5, 0x69, 0xc8,
	0x36, 0xbf, 0xf1, 0x7d, 0x15, 0x7c, 0xd2, 0xb7, 0xc6, 0xb6, 0xfd, 0x45, 0xef, 0x8c, 0x05, 0xa1,
	0xed, 0x33, 0x53, 0x84, 0x45, 0x34, 0xeb, 0x99, 0x52, 0x98, 0x7d, 0x55, 0x90, 0xb9, 0xed, 0xb3,
	0xf9, 0x0a, 0x32, 0x9f, 0xf5, 0xe6, 0x0a, 0xf2, 0x4b, 0x3f, 0x10, 0xe7, 0x08, 0xd9, 0x3e, 0x43,
	0x68, 0x05, 0x42, 0x68, 0xd6, 0x43, 0x48, 0xc1, 0x16, 0x8d, 0xd2, 0xe7, 0xe3, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd2, 0xc8, 0x94, 0x6f, 0x8a, 0x06, 0x00, 0x00,
}
