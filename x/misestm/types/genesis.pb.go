// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: misestm/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the misestm module's genesis state.
type GenesisState struct {
	MisesAccountList []*MisesAccount `protobuf:"bytes,9,rep,name=MisesAccountList,proto3" json:"MisesAccountList,omitempty"`
	// this line is used by starport scaffolding # genesis/proto/state
	UserInfoList      []*UserInfo     `protobuf:"bytes,7,rep,name=UserInfoList,proto3" json:"UserInfoList,omitempty"`
	UserInfoCount     uint64          `protobuf:"varint,8,opt,name=UserInfoCount,proto3" json:"UserInfoCount,omitempty"`
	UserRelationList  []*UserRelation `protobuf:"bytes,5,rep,name=UserRelationList,proto3" json:"UserRelationList,omitempty"`
	UserRelationCount uint64          `protobuf:"varint,6,opt,name=UserRelationCount,proto3" json:"UserRelationCount,omitempty"`
	AppInfoList       []*AppInfo      `protobuf:"bytes,3,rep,name=AppInfoList,proto3" json:"AppInfoList,omitempty"`
	AppInfoCount      uint64          `protobuf:"varint,4,opt,name=AppInfoCount,proto3" json:"AppInfoCount,omitempty"`
	DidRegistryList   []*DidRegistry  `protobuf:"bytes,1,rep,name=DidRegistryList,proto3" json:"DidRegistryList,omitempty"`
	DidRegistryCount  uint64          `protobuf:"varint,2,opt,name=DidRegistryCount,proto3" json:"DidRegistryCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_26f6a90bdd027bdd, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetMisesAccountList() []*MisesAccount {
	if m != nil {
		return m.MisesAccountList
	}
	return nil
}

func (m *GenesisState) GetUserInfoList() []*UserInfo {
	if m != nil {
		return m.UserInfoList
	}
	return nil
}

func (m *GenesisState) GetUserInfoCount() uint64 {
	if m != nil {
		return m.UserInfoCount
	}
	return 0
}

func (m *GenesisState) GetUserRelationList() []*UserRelation {
	if m != nil {
		return m.UserRelationList
	}
	return nil
}

func (m *GenesisState) GetUserRelationCount() uint64 {
	if m != nil {
		return m.UserRelationCount
	}
	return 0
}

func (m *GenesisState) GetAppInfoList() []*AppInfo {
	if m != nil {
		return m.AppInfoList
	}
	return nil
}

func (m *GenesisState) GetAppInfoCount() uint64 {
	if m != nil {
		return m.AppInfoCount
	}
	return 0
}

func (m *GenesisState) GetDidRegistryList() []*DidRegistry {
	if m != nil {
		return m.DidRegistryList
	}
	return nil
}

func (m *GenesisState) GetDidRegistryCount() uint64 {
	if m != nil {
		return m.DidRegistryCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "misesid.misestm.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("misestm/v1beta1/genesis.proto", fileDescriptor_26f6a90bdd027bdd) }

var fileDescriptor_26f6a90bdd027bdd = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcb, 0x4e, 0x3a, 0x31,
	0x14, 0xc6, 0x99, 0x3f, 0xfc, 0x51, 0x0b, 0x46, 0xec, 0x46, 0x42, 0x62, 0x03, 0x04, 0x13, 0x62,
	0x70, 0x26, 0xe8, 0x13, 0x80, 0xb7, 0x98, 0xa8, 0x89, 0x35, 0x6e, 0xdc, 0x71, 0xa9, 0xd8, 0xc4,
	0xb9, 0x84, 0x1e, 0x8c, 0xbc, 0x85, 0xef, 0xe2, 0x4b, 0xb8, 0x64, 0xe9, 0xd2, 0xc0, 0x8b, 0x18,
	0x4e, 0x3b, 0x49, 0x61, 0x9c, 0xb8, 0x9b, 0x73, 0xce, 0xf7, 0x7d, 0xbf, 0x9e, 0x69, 0xc9, 0xbe,
	0x2f, 0x95, 0x50, 0xe0, 0x7b, 0xaf, 0xed, 0xbe, 0x80, 0x5e, 0xdb, 0x1b, 0x89, 0x40, 0x28, 0xa9,
	0xdc, 0x68, 0x1c, 0x42, 0x48, 0xf7, 0x70, 0x2c, 0x87, 0xae, 0x91, 0xb9, 0x46, 0x56, 0x61, 0xeb,
	0xbe, 0x07, 0x25, 0xc6, 0x57, 0xc1, 0x53, 0xa8, 0x8d, 0x95, 0xfa, 0x6f, 0x73, 0x2e, 0x5e, 0x7a,
	0x20, 0xc3, 0xc0, 0x68, 0x12, 0xec, 0x4e, 0x14, 0x59, 0x11, 0xb5, 0xf5, 0xf1, 0x99, 0x1c, 0x72,
	0x31, 0x92, 0x0a, 0xc6, 0xd3, 0x34, 0xca, 0xcd, 0xb2, 0xee, 0x0c, 0x06, 0xe1, 0x24, 0x00, 0xad,
	0xa9, 0x7f, 0xe4, 0x48, 0xf1, 0x52, 0x2f, 0x75, 0x0f, 0x3d, 0x10, 0xf4, 0x8e, 0x94, 0x6c, 0xd9,
	0xb5, 0x54, 0x50, 0xde, 0xaa, 0x66, 0x9b, 0x85, 0xe3, 0x03, 0x37, 0x65, 0x5d, 0xd7, 0x36, 0xf0,
	0x84, 0x9d, 0x9e, 0x93, 0x62, 0xbc, 0x3f, 0xc6, 0x6d, 0x60, 0x5c, 0x2d, 0x35, 0x2e, 0x16, 0xf3,
	0x15, 0x1b, 0x6d, 0x90, 0xed, 0xb8, 0x3e, 0x5d, 0x66, 0x97, 0x37, 0xab, 0x4e, 0x33, 0xc7, 0x57,
	0x9b, 0xcb, 0xf3, 0xdb, 0x3f, 0x13, 0x81, 0xff, 0xff, 0x38, 0xbf, 0x6d, 0xe0, 0x09, 0x3b, 0x6d,
	0x91, 0x5d, 0xbb, 0xa7, 0xe1, 0x79, 0x84, 0x27, 0x07, 0xb4, 0x4b, 0x0a, 0xe6, 0xa6, 0x90, 0x9d,
	0x45, 0x76, 0x35, 0x95, 0x6d, 0xb4, 0xdc, 0x36, 0xd1, 0x3a, 0x29, 0x9a, 0x52, 0xc3, 0x72, 0x08,
	0x5b, 0xe9, 0xd1, 0x5b, 0xb2, 0x63, 0x5d, 0x39, 0xb2, 0x1c, 0x64, 0x35, 0x52, 0x59, 0x96, 0x9e,
	0xaf, 0x9b, 0xe9, 0x21, 0x29, 0x59, 0x2d, 0xcd, 0xfd, 0x87, 0xdc, 0x44, 0xbf, 0x7b, 0xf1, 0x39,
	0x67, 0xce, 0x6c, 0xce, 0x9c, 0xef, 0x39, 0x73, 0xde, 0x17, 0x2c, 0x33, 0x5b, 0xb0, 0xcc, 0xd7,
	0x82, 0x65, 0x1e, 0x5b, 0x23, 0x09, 0xcf, 0x93, 0xbe, 0x3b, 0x08, 0x7d, 0x0f, 0xf1, 0x47, 0x72,
	0x68, 0x3e, 0xc0, 0xf7, 0xde, 0xbc, 0xf8, 0x49, 0xc2, 0x34, 0x12, 0xaa, 0x9f, 0xc7, 0x47, 0x78,
	0xf2, 0x13, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x4a, 0x88, 0xac, 0x68, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MisesAccountList) > 0 {
		for iNdEx := len(m.MisesAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MisesAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.UserInfoCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.UserInfoCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.UserInfoList) > 0 {
		for iNdEx := len(m.UserInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.UserRelationCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.UserRelationCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.UserRelationList) > 0 {
		for iNdEx := len(m.UserRelationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserRelationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.AppInfoCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AppInfoCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AppInfoList) > 0 {
		for iNdEx := len(m.AppInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AppInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.DidRegistryCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.DidRegistryCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DidRegistryList) > 0 {
		for iNdEx := len(m.DidRegistryList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DidRegistryList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DidRegistryList) > 0 {
		for _, e := range m.DidRegistryList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.DidRegistryCount != 0 {
		n += 1 + sovGenesis(uint64(m.DidRegistryCount))
	}
	if len(m.AppInfoList) > 0 {
		for _, e := range m.AppInfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.AppInfoCount != 0 {
		n += 1 + sovGenesis(uint64(m.AppInfoCount))
	}
	if len(m.UserRelationList) > 0 {
		for _, e := range m.UserRelationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.UserRelationCount != 0 {
		n += 1 + sovGenesis(uint64(m.UserRelationCount))
	}
	if len(m.UserInfoList) > 0 {
		for _, e := range m.UserInfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.UserInfoCount != 0 {
		n += 1 + sovGenesis(uint64(m.UserInfoCount))
	}
	if len(m.MisesAccountList) > 0 {
		for _, e := range m.MisesAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidRegistryList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidRegistryList = append(m.DidRegistryList, &DidRegistry{})
			if err := m.DidRegistryList[len(m.DidRegistryList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidRegistryCount", wireType)
			}
			m.DidRegistryCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DidRegistryCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppInfoList = append(m.AppInfoList, &AppInfo{})
			if err := m.AppInfoList[len(m.AppInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppInfoCount", wireType)
			}
			m.AppInfoCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppInfoCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserRelationList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserRelationList = append(m.UserRelationList, &UserRelation{})
			if err := m.UserRelationList[len(m.UserRelationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserRelationCount", wireType)
			}
			m.UserRelationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserRelationCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserInfoList = append(m.UserInfoList, &UserInfo{})
			if err := m.UserInfoList[len(m.UserInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserInfoCount", wireType)
			}
			m.UserInfoCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserInfoCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MisesAccountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MisesAccountList = append(m.MisesAccountList, &MisesAccount{})
			if err := m.MisesAccountList[len(m.MisesAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
