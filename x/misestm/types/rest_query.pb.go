// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: misestm/rest_query.proto

package types

import (
	context "context"
	fmt "fmt"
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

type RestQueryDidRequest struct {
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
}

func (m *RestQueryDidRequest) Reset()         { *m = RestQueryDidRequest{} }
func (m *RestQueryDidRequest) String() string { return proto.CompactTextString(m) }
func (*RestQueryDidRequest) ProtoMessage()    {}
func (*RestQueryDidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d6cefbd419a531, []int{0}
}
func (m *RestQueryDidRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RestQueryDidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RestQueryDidRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RestQueryDidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestQueryDidRequest.Merge(m, src)
}
func (m *RestQueryDidRequest) XXX_Size() int {
	return m.Size()
}
func (m *RestQueryDidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RestQueryDidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RestQueryDidRequest proto.InternalMessageInfo

func (m *RestQueryDidRequest) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

type RestQueryDidResponse struct {
	DidRegistry *DidRegistry `protobuf:"bytes,1,opt,name=didRegistry,proto3" json:"didRegistry,omitempty"`
}

func (m *RestQueryDidResponse) Reset()         { *m = RestQueryDidResponse{} }
func (m *RestQueryDidResponse) String() string { return proto.CompactTextString(m) }
func (*RestQueryDidResponse) ProtoMessage()    {}
func (*RestQueryDidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d6cefbd419a531, []int{1}
}
func (m *RestQueryDidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RestQueryDidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RestQueryDidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RestQueryDidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestQueryDidResponse.Merge(m, src)
}
func (m *RestQueryDidResponse) XXX_Size() int {
	return m.Size()
}
func (m *RestQueryDidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RestQueryDidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RestQueryDidResponse proto.InternalMessageInfo

func (m *RestQueryDidResponse) GetDidRegistry() *DidRegistry {
	if m != nil {
		return m.DidRegistry
	}
	return nil
}

func init() {
	proto.RegisterType((*RestQueryDidRequest)(nil), "misesid.misestm.misestm.RestQueryDidRequest")
	proto.RegisterType((*RestQueryDidResponse)(nil), "misesid.misestm.misestm.RestQueryDidResponse")
}

func init() { proto.RegisterFile("misestm/rest_query.proto", fileDescriptor_56d6cefbd419a531) }

var fileDescriptor_56d6cefbd419a531 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0x3b, 0x31,
	0x10, 0xc7, 0x9b, 0xdf, 0x0f, 0xc4, 0xa6, 0x97, 0x12, 0x05, 0x6b, 0x91, 0x45, 0x8a, 0xa0, 0x48,
	0xbb, 0x43, 0xeb, 0x1b, 0x88, 0xf4, 0xee, 0x1e, 0x3d, 0x28, 0xd9, 0x26, 0xac, 0x03, 0xee, 0x66,
	0xdb, 0x99, 0x8a, 0x45, 0x3c, 0xe8, 0x0b, 0x28, 0xf8, 0x52, 0x1e, 0x0b, 0x5e, 0x3c, 0x4a, 0xeb,
	0x83, 0x48, 0xd3, 0xae, 0x56, 0xb0, 0xe0, 0x29, 0x03, 0xf3, 0xf9, 0xfe, 0xc9, 0xc8, 0x5a, 0x8a,
	0x64, 0x89, 0x53, 0x18, 0x58, 0xe2, 0x8b, 0xfe, 0xd0, 0x0e, 0x46, 0x61, 0x3e, 0x70, 0xec, 0xd4,
	0x96, 0xdf, 0xa0, 0x09, 0x17, 0x44, 0xf1, 0xd6, 0x77, 0x12, 0xe7, 0x92, 0x2b, 0x0b, 0x3a, 0x47,
	0xd0, 0x59, 0xe6, 0x58, 0x33, 0xba, 0x8c, 0xe6, 0xb2, 0xfa, 0x61, 0xcf, 0x51, 0xea, 0x08, 0x62,
	0x4d, 0x16, 0xbc, 0x1f, 0x5c, 0xb7, 0x63, 0xcb, 0xba, 0x0d, 0xb9, 0x4e, 0x30, 0xf3, 0xf0, 0x82,
	0xdd, 0x2e, 0xc2, 0x4f, 0xd0, 0x44, 0x36, 0x41, 0xe2, 0x22, 0xbd, 0xb1, 0x2f, 0x37, 0x22, 0x4b,
	0x7c, 0x3a, 0x33, 0xf0, 0xdb, 0xfe, 0xd0, 0x12, 0xab, 0xaa, 0xfc, 0x6f, 0xd0, 0xd4, 0xc4, 0xae,
	0x38, 0x28, 0x47, 0xb3, 0xb1, 0x71, 0x2e, 0x37, 0x7f, 0x82, 0x94, 0xbb, 0x8c, 0xac, 0xea, 0xca,
	0x8a, 0xf9, 0x76, 0xf5, 0x8a, 0x4a, 0x67, 0x2f, 0x5c, 0xf1, 0xa9, 0x70, 0xa9, 0x41, 0xb4, 0x2c,
	0xec, 0x3c, 0x0a, 0x59, 0xfe, 0x0a, 0x50, 0xf7, 0x42, 0xae, 0x17, 0x51, 0xaa, 0xb9, 0xd2, 0xed,
	0x97, 0xea, 0xf5, 0xd6, 0x1f, 0xe9, 0x79, 0xff, 0x46, 0xed, 0xe1, 0xf5, 0xe3, 0xf9, 0x9f, 0x52,
	0x55, 0xf0, 0x18, 0x18, 0x34, 0x70, 0x6b, 0xd0, 0xdc, 0x1d, 0x77, 0x5f, 0x26, 0x81, 0x18, 0x4f,
	0x02, 0xf1, 0x3e, 0x09, 0xc4, 0xd3, 0x34, 0x28, 0x8d, 0xa7, 0x41, 0xe9, 0x6d, 0x1a, 0x94, 0xce,
	0x9a, 0x09, 0xf2, 0xe5, 0x30, 0x0e, 0x7b, 0x2e, 0x9d, 0xab, 0x5a, 0x68, 0x16, 0x03, 0xa7, 0x70,
	0x03, 0xc5, 0xb9, 0x79, 0x94, 0x5b, 0x8a, 0xd7, 0xfc, 0xa5, 0x8f, 0x3e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0xb4, 0xe8, 0x94, 0x03, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RestQueryClient is the client API for RestQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestQueryClient interface {
	// query a did
	QueryDid(ctx context.Context, in *RestQueryDidRequest, opts ...grpc.CallOption) (*RestQueryDidResponse, error)
}

type restQueryClient struct {
	cc grpc1.ClientConn
}

func NewRestQueryClient(cc grpc1.ClientConn) RestQueryClient {
	return &restQueryClient{cc}
}

func (c *restQueryClient) QueryDid(ctx context.Context, in *RestQueryDidRequest, opts ...grpc.CallOption) (*RestQueryDidResponse, error) {
	out := new(RestQueryDidResponse)
	err := c.cc.Invoke(ctx, "/misesid.misestm.misestm.RestQuery/QueryDid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestQueryServer is the server API for RestQuery service.
type RestQueryServer interface {
	// query a did
	QueryDid(context.Context, *RestQueryDidRequest) (*RestQueryDidResponse, error)
}

// UnimplementedRestQueryServer can be embedded to have forward compatible implementations.
type UnimplementedRestQueryServer struct {
}

func (*UnimplementedRestQueryServer) QueryDid(ctx context.Context, req *RestQueryDidRequest) (*RestQueryDidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDid not implemented")
}

func RegisterRestQueryServer(s grpc1.Server, srv RestQueryServer) {
	s.RegisterService(&_RestQuery_serviceDesc, srv)
}

func _RestQuery_QueryDid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestQueryDidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestQueryServer).QueryDid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/misesid.misestm.misestm.RestQuery/QueryDid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestQueryServer).QueryDid(ctx, req.(*RestQueryDidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestQuery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "misesid.misestm.misestm.RestQuery",
	HandlerType: (*RestQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryDid",
			Handler:    _RestQuery_QueryDid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "misestm/rest_query.proto",
}

func (m *RestQueryDidRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RestQueryDidRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RestQueryDidRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintRestQuery(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RestQueryDidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RestQueryDidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RestQueryDidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DidRegistry != nil {
		{
			size, err := m.DidRegistry.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRestQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRestQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovRestQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RestQueryDidRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovRestQuery(uint64(l))
	}
	return n
}

func (m *RestQueryDidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DidRegistry != nil {
		l = m.DidRegistry.Size()
		n += 1 + l + sovRestQuery(uint64(l))
	}
	return n
}

func sovRestQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRestQuery(x uint64) (n int) {
	return sovRestQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RestQueryDidRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRestQuery
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
			return fmt.Errorf("proto: RestQueryDidRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestQueryDidRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestQuery
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
				return ErrInvalidLengthRestQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRestQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRestQuery
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
func (m *RestQueryDidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRestQuery
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
			return fmt.Errorf("proto: RestQueryDidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestQueryDidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidRegistry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestQuery
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
				return ErrInvalidLengthRestQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRestQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DidRegistry == nil {
				m.DidRegistry = &DidRegistry{}
			}
			if err := m.DidRegistry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRestQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRestQuery
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
func skipRestQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRestQuery
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
					return 0, ErrIntOverflowRestQuery
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
					return 0, ErrIntOverflowRestQuery
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
				return 0, ErrInvalidLengthRestQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRestQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRestQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRestQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRestQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRestQuery = fmt.Errorf("proto: unexpected end of group")
)
