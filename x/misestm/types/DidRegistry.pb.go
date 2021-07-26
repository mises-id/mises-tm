// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: misestm/DidRegistry.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type DidRegistry struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id            uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Did           string `protobuf:"bytes,3,opt,name=did,proto3" json:"did,omitempty"`
	PkeyDid       string `protobuf:"bytes,4,opt,name=pkeyDid,proto3" json:"pkeyDid,omitempty"`
	PkeyType      string `protobuf:"bytes,5,opt,name=pkeyType,proto3" json:"pkeyType,omitempty"`
	PkeyMultibase string `protobuf:"bytes,6,opt,name=pkeyMultibase,proto3" json:"pkeyMultibase,omitempty"`
	Version       uint64 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *DidRegistry) Reset()         { *m = DidRegistry{} }
func (m *DidRegistry) String() string { return proto.CompactTextString(m) }
func (*DidRegistry) ProtoMessage()    {}
func (*DidRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_38637afb3080373b, []int{0}
}
func (m *DidRegistry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DidRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DidRegistry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DidRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DidRegistry.Merge(m, src)
}
func (m *DidRegistry) XXX_Size() int {
	return m.Size()
}
func (m *DidRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_DidRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_DidRegistry proto.InternalMessageInfo

func (m *DidRegistry) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *DidRegistry) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DidRegistry) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *DidRegistry) GetPkeyDid() string {
	if m != nil {
		return m.PkeyDid
	}
	return ""
}

func (m *DidRegistry) GetPkeyType() string {
	if m != nil {
		return m.PkeyType
	}
	return ""
}

func (m *DidRegistry) GetPkeyMultibase() string {
	if m != nil {
		return m.PkeyMultibase
	}
	return ""
}

func (m *DidRegistry) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*DidRegistry)(nil), "misesid.misestm.misestm.DidRegistry")
}

func init() { proto.RegisterFile("misestm/DidRegistry.proto", fileDescriptor_38637afb3080373b) }

var fileDescriptor_38637afb3080373b = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0xb4, 0xb4, 0x60, 0x04, 0x42, 0x16, 0x12, 0xa6, 0x83, 0x55, 0x21, 0x86, 0x0e,
	0x10, 0x0f, 0xbc, 0x01, 0xaa, 0xd8, 0x58, 0x22, 0x26, 0xb6, 0xa6, 0xb6, 0xc2, 0x09, 0x82, 0x23,
	0xdb, 0x45, 0xe4, 0x2d, 0x78, 0x23, 0x56, 0xc6, 0x8e, 0x8c, 0x28, 0x79, 0x11, 0xe4, 0x4b, 0x8c,
	0xe8, 0x94, 0xef, 0xcf, 0x7d, 0x77, 0xb2, 0x7e, 0x7a, 0x5e, 0x81, 0xd3, 0xce, 0x57, 0x72, 0x09,
	0x2a, 0xd7, 0x25, 0x38, 0x6f, 0x9b, 0xac, 0xb6, 0xc6, 0x1b, 0x76, 0x86, 0x23, 0x50, 0xd9, 0xa0,
	0xc4, 0xef, 0xec, 0xb4, 0x34, 0xa5, 0x41, 0x47, 0x06, 0xea, 0xf5, 0x8b, 0x4f, 0x42, 0x0f, 0xff,
	0x1d, 0x61, 0x9c, 0x4e, 0xd7, 0x56, 0xaf, 0xbc, 0xb1, 0x9c, 0xcc, 0xc9, 0xe2, 0x20, 0x8f, 0x91,
	0x1d, 0xd3, 0x14, 0x14, 0x4f, 0xe7, 0x64, 0x31, 0xce, 0x53, 0x50, 0xec, 0x84, 0x8e, 0x14, 0x28,
	0x3e, 0x42, 0x2b, 0x60, 0xd8, 0xad, 0x9f, 0x75, 0xb3, 0x04, 0xc5, 0xc7, 0xfd, 0xee, 0x10, 0xd9,
	0x8c, 0xee, 0x07, 0x7c, 0x68, 0x6a, 0xcd, 0xf7, 0x70, 0xf4, 0x97, 0xd9, 0x25, 0x3d, 0x0a, 0x7c,
	0xbf, 0x79, 0xf1, 0x50, 0xac, 0x9c, 0xe6, 0x13, 0x14, 0x76, 0x7f, 0x86, 0xdb, 0x6f, 0xda, 0x3a,
	0x30, 0xaf, 0x7c, 0x8a, 0x4f, 0x88, 0xf1, 0xf6, 0xee, 0xab, 0x15, 0x64, 0xdb, 0x0a, 0xf2, 0xd3,
	0x0a, 0xf2, 0xd1, 0x89, 0x64, 0xdb, 0x89, 0xe4, 0xbb, 0x13, 0xc9, 0xe3, 0x55, 0x09, 0xfe, 0x69,
	0x53, 0x64, 0x6b, 0x53, 0x49, 0x6c, 0xe1, 0x1a, 0xd4, 0x00, 0xbe, 0x92, 0xef, 0x32, 0x96, 0xe8,
	0x9b, 0x5a, 0xbb, 0x62, 0x82, 0x85, 0xdc, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xc6, 0xc5,
	0x47, 0x5c, 0x01, 0x00, 0x00,
}

func (m *DidRegistry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DidRegistry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DidRegistry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintDidRegistry(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x38
	}
	if len(m.PkeyMultibase) > 0 {
		i -= len(m.PkeyMultibase)
		copy(dAtA[i:], m.PkeyMultibase)
		i = encodeVarintDidRegistry(dAtA, i, uint64(len(m.PkeyMultibase)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PkeyType) > 0 {
		i -= len(m.PkeyType)
		copy(dAtA[i:], m.PkeyType)
		i = encodeVarintDidRegistry(dAtA, i, uint64(len(m.PkeyType)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PkeyDid) > 0 {
		i -= len(m.PkeyDid)
		copy(dAtA[i:], m.PkeyDid)
		i = encodeVarintDidRegistry(dAtA, i, uint64(len(m.PkeyDid)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintDidRegistry(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintDidRegistry(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDidRegistry(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDidRegistry(dAtA []byte, offset int, v uint64) int {
	offset -= sovDidRegistry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DidRegistry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDidRegistry(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovDidRegistry(uint64(m.Id))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovDidRegistry(uint64(l))
	}
	l = len(m.PkeyDid)
	if l > 0 {
		n += 1 + l + sovDidRegistry(uint64(l))
	}
	l = len(m.PkeyType)
	if l > 0 {
		n += 1 + l + sovDidRegistry(uint64(l))
	}
	l = len(m.PkeyMultibase)
	if l > 0 {
		n += 1 + l + sovDidRegistry(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovDidRegistry(uint64(m.Version))
	}
	return n
}

func sovDidRegistry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDidRegistry(x uint64) (n int) {
	return sovDidRegistry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DidRegistry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDidRegistry
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
			return fmt.Errorf("proto: DidRegistry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DidRegistry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidRegistry
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
				return ErrInvalidLengthDidRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDidRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidRegistry
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
				return ErrInvalidLengthDidRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDidRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PkeyDid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidRegistry
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
				return ErrInvalidLengthDidRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDidRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PkeyDid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PkeyType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidRegistry
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
				return ErrInvalidLengthDidRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDidRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PkeyType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PkeyMultibase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidRegistry
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
				return ErrInvalidLengthDidRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDidRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PkeyMultibase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDidRegistry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDidRegistry
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
func skipDidRegistry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDidRegistry
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
					return 0, ErrIntOverflowDidRegistry
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
					return 0, ErrIntOverflowDidRegistry
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
				return 0, ErrInvalidLengthDidRegistry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDidRegistry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDidRegistry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDidRegistry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDidRegistry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDidRegistry = fmt.Errorf("proto: unexpected end of group")
)