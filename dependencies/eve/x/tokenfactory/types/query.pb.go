// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eve/tokenfactory/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryGetDenomRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryGetDenomRequest) Reset()         { *m = QueryGetDenomRequest{} }
func (m *QueryGetDenomRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetDenomRequest) ProtoMessage()    {}
func (*QueryGetDenomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64cf0549e5ace280, []int{0}
}
func (m *QueryGetDenomRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDenomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDenomRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDenomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDenomRequest.Merge(m, src)
}
func (m *QueryGetDenomRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDenomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDenomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDenomRequest proto.InternalMessageInfo

func (m *QueryGetDenomRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type QueryGetDenomResponse struct {
	Denom Denom `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom"`
}

func (m *QueryGetDenomResponse) Reset()         { *m = QueryGetDenomResponse{} }
func (m *QueryGetDenomResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetDenomResponse) ProtoMessage()    {}
func (*QueryGetDenomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64cf0549e5ace280, []int{1}
}
func (m *QueryGetDenomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDenomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDenomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDenomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDenomResponse.Merge(m, src)
}
func (m *QueryGetDenomResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDenomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDenomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDenomResponse proto.InternalMessageInfo

func (m *QueryGetDenomResponse) GetDenom() Denom {
	if m != nil {
		return m.Denom
	}
	return Denom{}
}

type QueryAllDenomRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllDenomRequest) Reset()         { *m = QueryAllDenomRequest{} }
func (m *QueryAllDenomRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllDenomRequest) ProtoMessage()    {}
func (*QueryAllDenomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64cf0549e5ace280, []int{2}
}
func (m *QueryAllDenomRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllDenomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDenomRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllDenomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDenomRequest.Merge(m, src)
}
func (m *QueryAllDenomRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllDenomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllDenomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllDenomRequest proto.InternalMessageInfo

func (m *QueryAllDenomRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllDenomResponse struct {
	Denom      []Denom             `protobuf:"bytes,1,rep,name=denom,proto3" json:"denom"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllDenomResponse) Reset()         { *m = QueryAllDenomResponse{} }
func (m *QueryAllDenomResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllDenomResponse) ProtoMessage()    {}
func (*QueryAllDenomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64cf0549e5ace280, []int{3}
}
func (m *QueryAllDenomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllDenomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDenomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllDenomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDenomResponse.Merge(m, src)
}
func (m *QueryAllDenomResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllDenomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllDenomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllDenomResponse proto.InternalMessageInfo

func (m *QueryAllDenomResponse) GetDenom() []Denom {
	if m != nil {
		return m.Denom
	}
	return nil
}

func (m *QueryAllDenomResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetDenomRequest)(nil), "eve.tokenfactory.v1beta1.QueryGetDenomRequest")
	proto.RegisterType((*QueryGetDenomResponse)(nil), "eve.tokenfactory.v1beta1.QueryGetDenomResponse")
	proto.RegisterType((*QueryAllDenomRequest)(nil), "eve.tokenfactory.v1beta1.QueryAllDenomRequest")
	proto.RegisterType((*QueryAllDenomResponse)(nil), "eve.tokenfactory.v1beta1.QueryAllDenomResponse")
}

func init() {
	proto.RegisterFile("eve/tokenfactory/v1beta1/query.proto", fileDescriptor_64cf0549e5ace280)
}

var fileDescriptor_64cf0549e5ace280 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6b, 0x1a, 0x41,
	0x14, 0xc7, 0x77, 0x6c, 0x2d, 0xed, 0xf4, 0x36, 0x6c, 0x41, 0xa4, 0xac, 0x65, 0x69, 0x6b, 0x29,
	0x75, 0x06, 0x2d, 0x3d, 0xf5, 0xa4, 0x94, 0x7a, 0x6d, 0xa5, 0xa7, 0x1e, 0x0a, 0xb3, 0xf6, 0x65,
	0x91, 0xac, 0x33, 0xab, 0x33, 0x4a, 0x24, 0xe4, 0x92, 0x4f, 0x10, 0xc8, 0x2d, 0x10, 0xf2, 0x6d,
	0x82, 0x47, 0x21, 0x97, 0x9c, 0x42, 0xd0, 0x7c, 0x90, 0xb0, 0x33, 0x6b, 0x5c, 0x35, 0xa2, 0x39,
	0xed, 0xec, 0xf2, 0xff, 0xff, 0xdf, 0xef, 0xbd, 0xb7, 0x83, 0xdf, 0xc3, 0x10, 0x98, 0x96, 0xfb,
	0x20, 0xf6, 0x78, 0x5b, 0xcb, 0xfe, 0x88, 0x0d, 0xab, 0x01, 0x68, 0x5e, 0x65, 0xbd, 0x01, 0xf4,
	0x47, 0x34, 0xee, 0x4b, 0x2d, 0x49, 0x01, 0x86, 0x40, 0xb3, 0x2a, 0x9a, 0xaa, 0x8a, 0x6f, 0x43,
	0x29, 0xc3, 0x08, 0x18, 0x8f, 0x3b, 0x8c, 0x0b, 0x21, 0x35, 0xd7, 0x1d, 0x29, 0x94, 0xf5, 0x15,
	0x3f, 0xb7, 0xa5, 0xea, 0x4a, 0xc5, 0x02, 0xae, 0xc0, 0x06, 0x3e, 0xc4, 0xc7, 0x3c, 0xec, 0x08,
	0x23, 0x4e, 0xb5, 0x9b, 0x49, 0xfe, 0x83, 0x90, 0xdd, 0x54, 0xe5, 0x86, 0x32, 0x94, 0xe6, 0xc8,
	0x92, 0x93, 0xfd, 0xea, 0x7f, 0xc1, 0xee, 0xef, 0x24, 0xbd, 0x09, 0xfa, 0x47, 0x22, 0x6e, 0x41,
	0x6f, 0x00, 0x4a, 0x13, 0x17, 0xe7, 0x8d, 0xb9, 0x80, 0xde, 0xa1, 0x4f, 0xaf, 0x5a, 0xf6, 0xc5,
	0xff, 0x83, 0xdf, 0xac, 0xa8, 0x55, 0x2c, 0x85, 0x02, 0xf2, 0x3d, 0x2b, 0x7f, 0x5d, 0x2b, 0xd1,
	0x4d, 0x6d, 0x53, 0xe3, 0x6b, 0x3c, 0x1f, 0xdf, 0x94, 0x9c, 0x79, 0xea, 0xbf, 0x94, 0xa1, 0x1e,
	0x45, 0x4b, 0x0c, 0x3f, 0x31, 0x5e, 0xf4, 0x9a, 0x26, 0x7f, 0xa4, 0x76, 0x30, 0x34, 0x19, 0x0c,
	0xb5, 0x93, 0x9e, 0x47, 0xff, 0xe2, 0x21, 0xa4, 0xde, 0x56, 0xc6, 0xe9, 0x9f, 0xa3, 0x14, 0x7b,
	0x51, 0x60, 0x1d, 0xfb, 0xd9, 0x53, 0xb1, 0x49, 0x73, 0x09, 0x2f, 0x67, 0xf0, 0xca, 0x5b, 0xf1,
	0x6c, 0xe5, 0x2c, 0x5f, 0xed, 0x32, 0x87, 0xf3, 0x86, 0x8f, 0x5c, 0x20, 0x9c, 0x37, 0x95, 0x08,
	0xdd, 0x8c, 0xf2, 0xd8, 0xbe, 0x8a, 0x6c, 0x67, 0xbd, 0x05, 0xf0, 0xbf, 0x1d, 0x5f, 0xdd, 0x9d,
	0xe6, 0x18, 0xa9, 0x30, 0x21, 0x13, 0x0a, 0x1e, 0x55, 0x22, 0x1e, 0x28, 0xb6, 0xf6, 0x2f, 0x99,
	0x66, 0xd9, 0xa1, 0x79, 0x1c, 0x91, 0x33, 0x84, 0x5f, 0x9a, 0xa0, 0x7a, 0x14, 0x6d, 0x85, 0x5c,
	0x59, 0xe8, 0x56, 0xc8, 0xd5, 0xfd, 0xf8, 0x15, 0x03, 0x59, 0x26, 0x1f, 0x76, 0x82, 0x6c, 0xd0,
	0xf1, 0xd4, 0x43, 0x93, 0xa9, 0x87, 0x6e, 0xa7, 0x1e, 0x3a, 0x99, 0x79, 0xce, 0x64, 0xe6, 0x39,
	0xd7, 0x33, 0xcf, 0xf9, 0xeb, 0x1e, 0x2c, 0xeb, 0xf5, 0x28, 0x06, 0x15, 0xbc, 0x30, 0x77, 0xe0,
	0xeb, 0x7d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0xca, 0x46, 0xa4, 0xcb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a denom by index.
	Denom(ctx context.Context, in *QueryGetDenomRequest, opts ...grpc.CallOption) (*QueryGetDenomResponse, error)
	// Queries a list of denom items.
	DenomAll(ctx context.Context, in *QueryAllDenomRequest, opts ...grpc.CallOption) (*QueryAllDenomResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Denom(ctx context.Context, in *QueryGetDenomRequest, opts ...grpc.CallOption) (*QueryGetDenomResponse, error) {
	out := new(QueryGetDenomResponse)
	err := c.cc.Invoke(ctx, "/eve.tokenfactory.v1beta1.Query/Denom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomAll(ctx context.Context, in *QueryAllDenomRequest, opts ...grpc.CallOption) (*QueryAllDenomResponse, error) {
	out := new(QueryAllDenomResponse)
	err := c.cc.Invoke(ctx, "/eve.tokenfactory.v1beta1.Query/DenomAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a denom by index.
	Denom(context.Context, *QueryGetDenomRequest) (*QueryGetDenomResponse, error)
	// Queries a list of denom items.
	DenomAll(context.Context, *QueryAllDenomRequest) (*QueryAllDenomResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Denom(ctx context.Context, req *QueryGetDenomRequest) (*QueryGetDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Denom not implemented")
}
func (*UnimplementedQueryServer) DenomAll(ctx context.Context, req *QueryAllDenomRequest) (*QueryAllDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Denom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Denom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eve.tokenfactory.v1beta1.Query/Denom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Denom(ctx, req.(*QueryGetDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eve.tokenfactory.v1beta1.Query/DenomAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomAll(ctx, req.(*QueryAllDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eve.tokenfactory.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Denom",
			Handler:    _Query_Denom_Handler,
		},
		{
			MethodName: "DenomAll",
			Handler:    _Query_DenomAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eve/tokenfactory/v1beta1/query.proto",
}

func (m *QueryGetDenomRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDenomRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDenomRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetDenomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDenomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDenomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Denom.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllDenomRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDenomRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDenomRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllDenomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDenomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDenomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		for iNdEx := len(m.Denom) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Denom[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetDenomRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetDenomResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Denom.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllDenomRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllDenomResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Denom) > 0 {
		for _, e := range m.Denom {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetDenomRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetDenomRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDenomRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetDenomResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetDenomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDenomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Denom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllDenomRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllDenomRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDenomRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllDenomResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllDenomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDenomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = append(m.Denom, Denom{})
			if err := m.Denom[len(m.Denom)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
