// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: misestm/rest_tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

type RestCreateDidRequest struct {
	Did  string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Pkey string `protobuf:"bytes,3,opt,name=pkey,proto3" json:"pkey,omitempty"`
}

func (m *RestCreateDidRequest) Reset()         { *m = RestCreateDidRequest{} }
func (m *RestCreateDidRequest) String() string { return proto.CompactTextString(m) }
func (*RestCreateDidRequest) ProtoMessage()    {}
func (*RestCreateDidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a4b9f04197513a, []int{0}
}
func (m *RestCreateDidRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RestCreateDidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RestCreateDidRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RestCreateDidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestCreateDidRequest.Merge(m, src)
}
func (m *RestCreateDidRequest) XXX_Size() int {
	return m.Size()
}
func (m *RestCreateDidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RestCreateDidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RestCreateDidRequest proto.InternalMessageInfo

func (m *RestCreateDidRequest) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *RestCreateDidRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RestCreateDidRequest) GetPkey() string {
	if m != nil {
		return m.Pkey
	}
	return ""
}

type RestCreateDidResponse struct {
	Ret        uint64            `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	TxResponse *types.TxResponse `protobuf:"bytes,2,opt,name=tx_response,json=txResponse,proto3" json:"tx_response,omitempty"`
}

func (m *RestCreateDidResponse) Reset()         { *m = RestCreateDidResponse{} }
func (m *RestCreateDidResponse) String() string { return proto.CompactTextString(m) }
func (*RestCreateDidResponse) ProtoMessage()    {}
func (*RestCreateDidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a4b9f04197513a, []int{1}
}
func (m *RestCreateDidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RestCreateDidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RestCreateDidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RestCreateDidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestCreateDidResponse.Merge(m, src)
}
func (m *RestCreateDidResponse) XXX_Size() int {
	return m.Size()
}
func (m *RestCreateDidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RestCreateDidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RestCreateDidResponse proto.InternalMessageInfo

func (m *RestCreateDidResponse) GetRet() uint64 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *RestCreateDidResponse) GetTxResponse() *types.TxResponse {
	if m != nil {
		return m.TxResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*RestCreateDidRequest)(nil), "misesid.misestm.misestm.RestCreateDidRequest")
	proto.RegisterType((*RestCreateDidResponse)(nil), "misesid.misestm.misestm.RestCreateDidResponse")
}

func init() { proto.RegisterFile("misestm/rest_tx.proto", fileDescriptor_30a4b9f04197513a) }

var fileDescriptor_30a4b9f04197513a = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xbf, 0x6b, 0xe3, 0x30,
	0x14, 0x8e, 0x92, 0x10, 0x88, 0x32, 0xdc, 0x61, 0x2e, 0x5c, 0x2e, 0x1c, 0xe6, 0xf0, 0xdd, 0x70,
	0x84, 0x8b, 0x44, 0x72, 0x5b, 0xc7, 0xfe, 0x9a, 0x8b, 0xc9, 0xd4, 0x25, 0xc8, 0xf1, 0xc3, 0x15,
	0xad, 0x2d, 0xc7, 0x7a, 0x29, 0xf6, 0xd8, 0xce, 0x1d, 0x0a, 0xfd, 0xa7, 0x3a, 0x06, 0xba, 0x74,
	0x2c, 0x49, 0xff, 0x90, 0x22, 0xd9, 0x0e, 0xa5, 0xb4, 0xd0, 0x49, 0x9f, 0xbe, 0xf7, 0xbd, 0xef,
	0xe9, 0xd3, 0xa3, 0xfd, 0x58, 0x6a, 0xd0, 0x18, 0xf3, 0x0c, 0x34, 0xce, 0x31, 0x67, 0x69, 0xa6,
	0x50, 0x39, 0xdf, 0x2d, 0x2d, 0x43, 0x56, 0x95, 0xeb, 0x73, 0xf8, 0x33, 0x52, 0x2a, 0xba, 0x00,
	0x2e, 0x52, 0xc9, 0x45, 0x92, 0x28, 0x14, 0x28, 0x55, 0xa2, 0xcb, 0xb6, 0xe1, 0xef, 0x85, 0xd2,
	0xb1, 0xd2, 0x3c, 0x10, 0x1a, 0xb8, 0x08, 0x16, 0x92, 0x5f, 0x4e, 0x02, 0x40, 0x31, 0xb1, 0x97,
	0x4a, 0x34, 0x7a, 0x2d, 0x5a, 0xae, 0x20, 0x2b, 0x76, 0xaa, 0x54, 0x44, 0x32, 0xb1, 0x8e, 0x95,
	0xf6, 0x47, 0xfd, 0xbc, 0x43, 0x19, 0xfa, 0x10, 0x49, 0x8d, 0x59, 0x51, 0x96, 0xbc, 0x13, 0xfa,
	0xcd, 0x07, 0x8d, 0x07, 0x19, 0x08, 0x04, 0x5b, 0x5e, 0xae, 0x40, 0xa3, 0xf3, 0x95, 0xb6, 0x42,
	0x19, 0x0e, 0xc8, 0x2f, 0xf2, 0xb7, 0xeb, 0x1b, 0xe8, 0x38, 0xb4, 0x8d, 0x45, 0x0a, 0x83, 0xa6,
	0xa5, 0x2c, 0x36, 0x5c, 0x7a, 0x0e, 0xc5, 0xa0, 0x55, 0x72, 0x06, 0x7b, 0x29, 0xed, 0xbf, 0x71,
	0xd4, 0xa9, 0x4a, 0x34, 0x18, 0xcb, 0x0c, 0xd0, 0x5a, 0xb6, 0x7d, 0x03, 0x9d, 0x23, 0xda, 0xc3,
	0x7c, 0x9e, 0x55, 0x02, 0xeb, 0xdc, 0x9b, 0xfe, 0x61, 0x65, 0x32, 0x66, 0x92, 0x31, 0x9b, 0xb8,
	0x0a, 0xc6, 0x66, 0x79, 0x6d, 0xe6, 0x53, 0xdc, 0xe1, 0xe9, 0x0d, 0xa1, 0x1d, 0x33, 0x72, 0x96,
	0x3b, 0x57, 0x84, 0x76, 0x77, 0x93, 0x9d, 0x31, 0xfb, 0x60, 0x01, 0xec, 0xbd, 0xcc, 0x43, 0xf6,
	0x59, 0x79, 0x39, 0xd7, 0xeb, 0x5f, 0x3f, 0x3c, 0xdf, 0x35, 0xbf, 0x78, 0x94, 0x5b, 0x1d, 0x0f,
	0x65, 0xb8, 0x47, 0x46, 0xfb, 0xc7, 0xf7, 0x1b, 0x97, 0xac, 0x37, 0x2e, 0x79, 0xda, 0xb8, 0xe4,
	0x76, 0xeb, 0x36, 0xd6, 0x5b, 0xb7, 0xf1, 0xb8, 0x75, 0x1b, 0xa7, 0xff, 0x22, 0x89, 0x67, 0xab,
	0x80, 0x2d, 0x54, 0x5c, 0xb6, 0x8c, 0x65, 0x58, 0x01, 0x8c, 0x79, 0xce, 0xeb, 0x35, 0x99, 0xbf,
	0xd5, 0x41, 0xc7, 0x6e, 0xe8, 0xff, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x6c, 0xa5, 0xb9,
	0x5d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RestTxClient is the client API for RestTx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestTxClient interface {
	// create a did
	CreateDid(ctx context.Context, in *RestCreateDidRequest, opts ...grpc.CallOption) (*RestCreateDidResponse, error)
}

type restTxClient struct {
	cc grpc1.ClientConn
}

func NewRestTxClient(cc grpc1.ClientConn) RestTxClient {
	return &restTxClient{cc}
}

func (c *restTxClient) CreateDid(ctx context.Context, in *RestCreateDidRequest, opts ...grpc.CallOption) (*RestCreateDidResponse, error) {
	out := new(RestCreateDidResponse)
	err := c.cc.Invoke(ctx, "/misesid.misestm.misestm.RestTx/CreateDid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestTxServer is the server API for RestTx service.
type RestTxServer interface {
	// create a did
	CreateDid(context.Context, *RestCreateDidRequest) (*RestCreateDidResponse, error)
}

// UnimplementedRestTxServer can be embedded to have forward compatible implementations.
type UnimplementedRestTxServer struct {
}

func (*UnimplementedRestTxServer) CreateDid(ctx context.Context, req *RestCreateDidRequest) (*RestCreateDidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDid not implemented")
}

func RegisterRestTxServer(s grpc1.Server, srv RestTxServer) {
	s.RegisterService(&_RestTx_serviceDesc, srv)
}

func _RestTx_CreateDid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestCreateDidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestTxServer).CreateDid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/misesid.misestm.misestm.RestTx/CreateDid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestTxServer).CreateDid(ctx, req.(*RestCreateDidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestTx_serviceDesc = grpc.ServiceDesc{
	ServiceName: "misesid.misestm.misestm.RestTx",
	HandlerType: (*RestTxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDid",
			Handler:    _RestTx_CreateDid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "misestm/rest_tx.proto",
}

func (m *RestCreateDidRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RestCreateDidRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RestCreateDidRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pkey) > 0 {
		i -= len(m.Pkey)
		copy(dAtA[i:], m.Pkey)
		i = encodeVarintRestTx(dAtA, i, uint64(len(m.Pkey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintRestTx(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintRestTx(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RestCreateDidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RestCreateDidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RestCreateDidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TxResponse != nil {
		{
			size, err := m.TxResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRestTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Ret != 0 {
		i = encodeVarintRestTx(dAtA, i, uint64(m.Ret))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRestTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovRestTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RestCreateDidRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovRestTx(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovRestTx(uint64(l))
	}
	l = len(m.Pkey)
	if l > 0 {
		n += 1 + l + sovRestTx(uint64(l))
	}
	return n
}

func (m *RestCreateDidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ret != 0 {
		n += 1 + sovRestTx(uint64(m.Ret))
	}
	if m.TxResponse != nil {
		l = m.TxResponse.Size()
		n += 1 + l + sovRestTx(uint64(l))
	}
	return n
}

func sovRestTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRestTx(x uint64) (n int) {
	return sovRestTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RestCreateDidRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRestTx
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
			return fmt.Errorf("proto: RestCreateDidRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestCreateDidRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestTx
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
				return ErrInvalidLengthRestTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestTx
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
				return ErrInvalidLengthRestTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestTx
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
				return ErrInvalidLengthRestTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRestTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRestTx
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
func (m *RestCreateDidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRestTx
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
			return fmt.Errorf("proto: RestCreateDidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestCreateDidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			m.Ret = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ret |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestTx
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
				return ErrInvalidLengthRestTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRestTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TxResponse == nil {
				m.TxResponse = &types.TxResponse{}
			}
			if err := m.TxResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRestTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRestTx
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
func skipRestTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRestTx
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
					return 0, ErrIntOverflowRestTx
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
					return 0, ErrIntOverflowRestTx
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
				return 0, ErrInvalidLengthRestTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRestTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRestTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRestTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRestTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRestTx = fmt.Errorf("proto: unexpected end of group")
)
