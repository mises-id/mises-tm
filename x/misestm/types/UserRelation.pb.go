// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: misestm/v1beta1/UserRelation.proto

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

type UserRelation struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id           uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	UidFrom      string `protobuf:"bytes,3,opt,name=uidFrom,proto3" json:"uidFrom,omitempty"`
	UidTo        string `protobuf:"bytes,4,opt,name=uidTo,proto3" json:"uidTo,omitempty"`
	IsFollowing  bool   `protobuf:"varint,5,opt,name=isFollowing,proto3" json:"isFollowing,omitempty"`
	IsBlocking   bool   `protobuf:"varint,6,opt,name=isBlocking,proto3" json:"isBlocking,omitempty"`
	IsReferredBy bool   `protobuf:"varint,7,opt,name=isReferredBy,proto3" json:"isReferredBy,omitempty"`
	Version      uint64 `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *UserRelation) Reset()         { *m = UserRelation{} }
func (m *UserRelation) String() string { return proto.CompactTextString(m) }
func (*UserRelation) ProtoMessage()    {}
func (*UserRelation) Descriptor() ([]byte, []int) {
	return fileDescriptor_8abe9dd47d555cd5, []int{0}
}
func (m *UserRelation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserRelation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserRelation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserRelation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRelation.Merge(m, src)
}
func (m *UserRelation) XXX_Size() int {
	return m.Size()
}
func (m *UserRelation) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRelation.DiscardUnknown(m)
}

var xxx_messageInfo_UserRelation proto.InternalMessageInfo

func (m *UserRelation) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *UserRelation) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRelation) GetUidFrom() string {
	if m != nil {
		return m.UidFrom
	}
	return ""
}

func (m *UserRelation) GetUidTo() string {
	if m != nil {
		return m.UidTo
	}
	return ""
}

func (m *UserRelation) GetIsFollowing() bool {
	if m != nil {
		return m.IsFollowing
	}
	return false
}

func (m *UserRelation) GetIsBlocking() bool {
	if m != nil {
		return m.IsBlocking
	}
	return false
}

func (m *UserRelation) GetIsReferredBy() bool {
	if m != nil {
		return m.IsReferredBy
	}
	return false
}

func (m *UserRelation) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRelation)(nil), "misesid.misestm.v1beta1.UserRelation")
}

func init() {
	proto.RegisterFile("misestm/v1beta1/UserRelation.proto", fileDescriptor_8abe9dd47d555cd5)
}

var fileDescriptor_8abe9dd47d555cd5 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xd0, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0x06, 0xf0, 0xba, 0xf4, 0x1f, 0xa6, 0x62, 0xb0, 0x2a, 0x61, 0x31, 0x58, 0x51, 0xa7, 0x0e,
	0x10, 0xab, 0xe2, 0x06, 0x1d, 0x7a, 0x80, 0x08, 0x16, 0xb6, 0x36, 0x36, 0xe1, 0x89, 0x24, 0xaf,
	0xb2, 0x9d, 0x42, 0x6f, 0xc1, 0xb1, 0x18, 0x3b, 0x32, 0xa2, 0x64, 0xe1, 0x18, 0x28, 0x4e, 0x22,
	0x85, 0xcd, 0xdf, 0xa7, 0x9f, 0xac, 0xa7, 0x8f, 0x2e, 0x33, 0xb0, 0xda, 0xba, 0x4c, 0x1e, 0xd7,
	0x7b, 0xed, 0x76, 0x6b, 0xf9, 0x64, 0xb5, 0x89, 0x74, 0xba, 0x73, 0x80, 0x79, 0x78, 0x30, 0xe8,
	0x90, 0xdd, 0x78, 0x03, 0x2a, 0x6c, 0x6d, 0xd8, 0xda, 0xdb, 0x45, 0x82, 0x09, 0x7a, 0x23, 0xeb,
	0x57, 0xc3, 0x97, 0xbf, 0x84, 0xce, 0xfb, 0xbf, 0x30, 0x4e, 0xa7, 0xb1, 0xd1, 0x3b, 0x87, 0x86,
	0x93, 0x80, 0xac, 0x2e, 0xa3, 0x2e, 0xb2, 0x6b, 0x3a, 0x04, 0xc5, 0x87, 0x01, 0x59, 0x8d, 0xa2,
	0x21, 0xa8, 0x5a, 0x16, 0xa0, 0xb6, 0x06, 0x33, 0x7e, 0xd1, 0xc8, 0x36, 0xb2, 0x05, 0x1d, 0x17,
	0xa0, 0x1e, 0x91, 0x8f, 0x7c, 0xdf, 0x04, 0x16, 0xd0, 0x2b, 0xb0, 0x5b, 0x4c, 0x53, 0x7c, 0x87,
	0x3c, 0xe1, 0xe3, 0x80, 0xac, 0x66, 0x51, 0xbf, 0x62, 0x82, 0x52, 0xb0, 0x9b, 0x14, 0xe3, 0xb7,
	0x1a, 0x4c, 0x3c, 0xe8, 0x35, 0x6c, 0x49, 0xe7, 0x60, 0x23, 0xfd, 0xa2, 0x8d, 0xd1, 0x6a, 0x73,
	0xe2, 0x53, 0x2f, 0xfe, 0x75, 0xf5, 0x55, 0x47, 0x6d, 0x2c, 0x60, 0xce, 0x67, 0xfe, 0xd4, 0x2e,
	0x6e, 0xb6, 0x5f, 0xa5, 0x20, 0xe7, 0x52, 0x90, 0x9f, 0x52, 0x90, 0xcf, 0x4a, 0x0c, 0xce, 0x95,
	0x18, 0x7c, 0x57, 0x62, 0xf0, 0x7c, 0x97, 0x80, 0x7b, 0x2d, 0xf6, 0x61, 0x8c, 0x99, 0xf4, 0xb3,
	0xdd, 0x83, 0x6a, 0x1f, 0x2e, 0x93, 0x1f, 0xb2, 0x9b, 0xdd, 0x9d, 0x0e, 0xda, 0xee, 0x27, 0x7e,
	0xb9, 0x87, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0x6f, 0x4e, 0x3b, 0x8e, 0x01, 0x00, 0x00,
}

func (m *UserRelation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserRelation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserRelation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintUserRelation(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x40
	}
	if m.IsReferredBy {
		i--
		if m.IsReferredBy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.IsBlocking {
		i--
		if m.IsBlocking {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.IsFollowing {
		i--
		if m.IsFollowing {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.UidTo) > 0 {
		i -= len(m.UidTo)
		copy(dAtA[i:], m.UidTo)
		i = encodeVarintUserRelation(dAtA, i, uint64(len(m.UidTo)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.UidFrom) > 0 {
		i -= len(m.UidFrom)
		copy(dAtA[i:], m.UidFrom)
		i = encodeVarintUserRelation(dAtA, i, uint64(len(m.UidFrom)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintUserRelation(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintUserRelation(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUserRelation(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserRelation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserRelation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovUserRelation(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovUserRelation(uint64(m.Id))
	}
	l = len(m.UidFrom)
	if l > 0 {
		n += 1 + l + sovUserRelation(uint64(l))
	}
	l = len(m.UidTo)
	if l > 0 {
		n += 1 + l + sovUserRelation(uint64(l))
	}
	if m.IsFollowing {
		n += 2
	}
	if m.IsBlocking {
		n += 2
	}
	if m.IsReferredBy {
		n += 2
	}
	if m.Version != 0 {
		n += 1 + sovUserRelation(uint64(m.Version))
	}
	return n
}

func sovUserRelation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserRelation(x uint64) (n int) {
	return sovUserRelation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserRelation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserRelation
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
			return fmt.Errorf("proto: UserRelation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserRelation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRelation
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
				return ErrInvalidLengthUserRelation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserRelation
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
					return ErrIntOverflowUserRelation
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
				return fmt.Errorf("proto: wrong wireType = %d for field UidFrom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRelation
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
				return ErrInvalidLengthUserRelation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UidFrom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UidTo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRelation
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
				return ErrInvalidLengthUserRelation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UidTo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsFollowing", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsFollowing = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsBlocking", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsBlocking = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsReferredBy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsReferredBy = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRelation
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
			skippy, err := skipUserRelation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserRelation
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
func skipUserRelation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserRelation
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
					return 0, ErrIntOverflowUserRelation
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
					return 0, ErrIntOverflowUserRelation
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
				return 0, ErrInvalidLengthUserRelation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserRelation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserRelation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserRelation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserRelation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserRelation = fmt.Errorf("proto: unexpected end of group")
)
