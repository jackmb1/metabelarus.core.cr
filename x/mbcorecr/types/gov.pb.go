// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mbcorecr/gov.proto

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

type MsgCreateSuperIdentity struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	WalletPath string `protobuf:"bytes,4,opt,name=walletPath,proto3" json:"walletPath,omitempty"`
}

func (m *MsgCreateSuperIdentity) Reset()         { *m = MsgCreateSuperIdentity{} }
func (m *MsgCreateSuperIdentity) String() string { return proto.CompactTextString(m) }
func (*MsgCreateSuperIdentity) ProtoMessage()    {}
func (*MsgCreateSuperIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef5434c92c47b57d, []int{0}
}
func (m *MsgCreateSuperIdentity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateSuperIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateSuperIdentity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateSuperIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateSuperIdentity.Merge(m, src)
}
func (m *MsgCreateSuperIdentity) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateSuperIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateSuperIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateSuperIdentity proto.InternalMessageInfo

func (m *MsgCreateSuperIdentity) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateSuperIdentity) GetWalletPath() string {
	if m != nil {
		return m.WalletPath
	}
	return ""
}

type IdentityAccount struct {
	Uid        string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PubKey     string `protobuf:"bytes,3,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	PrivKey    string `protobuf:"bytes,4,opt,name=privKey,proto3" json:"privKey,omitempty"`
	Mnemonic   string `protobuf:"bytes,5,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	IdentityID string `protobuf:"bytes,6,opt,name=identityID,proto3" json:"identityID,omitempty"`
}

func (m *IdentityAccount) Reset()         { *m = IdentityAccount{} }
func (m *IdentityAccount) String() string { return proto.CompactTextString(m) }
func (*IdentityAccount) ProtoMessage()    {}
func (*IdentityAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef5434c92c47b57d, []int{1}
}
func (m *IdentityAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentityAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentityAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentityAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityAccount.Merge(m, src)
}
func (m *IdentityAccount) XXX_Size() int {
	return m.Size()
}
func (m *IdentityAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityAccount.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityAccount proto.InternalMessageInfo

func (m *IdentityAccount) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *IdentityAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IdentityAccount) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *IdentityAccount) GetPrivKey() string {
	if m != nil {
		return m.PrivKey
	}
	return ""
}

func (m *IdentityAccount) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

func (m *IdentityAccount) GetIdentityID() string {
	if m != nil {
		return m.IdentityID
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgCreateSuperIdentity)(nil), "metabelarus.mbcorecr.mbcorecr.MsgCreateSuperIdentity")
	proto.RegisterType((*IdentityAccount)(nil), "metabelarus.mbcorecr.mbcorecr.IdentityAccount")
}

func init() { proto.RegisterFile("mbcorecr/gov.proto", fileDescriptor_ef5434c92c47b57d) }

var fileDescriptor_ef5434c92c47b57d = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0x4d, 0x4a, 0xc4, 0x30,
	0x18, 0x9d, 0x38, 0x5a, 0x35, 0x1b, 0x25, 0xc8, 0x10, 0x06, 0x0c, 0x32, 0x2b, 0x57, 0xad, 0xe0,
	0x09, 0xfc, 0xd9, 0x8c, 0x22, 0xc8, 0xb8, 0x73, 0x97, 0xa6, 0x1f, 0x9d, 0x40, 0xdb, 0x94, 0x34,
	0x19, 0xed, 0x2d, 0xbc, 0x87, 0x17, 0x71, 0x39, 0x4b, 0x97, 0xd2, 0x5e, 0x44, 0x92, 0xfe, 0xcc,
	0xec, 0xde, 0x4f, 0xbe, 0x97, 0xc7, 0xc3, 0x24, 0x8f, 0x85, 0xd2, 0x20, 0x74, 0x94, 0xaa, 0x4d,
	0x58, 0x6a, 0x65, 0x14, 0xb9, 0xcc, 0xc1, 0xf0, 0x18, 0x32, 0xae, 0x6d, 0x15, 0x0e, 0xfe, 0x08,
	0xe6, 0x17, 0xa9, 0x4a, 0x95, 0x7f, 0x19, 0x39, 0xd4, 0x1d, 0x2d, 0x56, 0x78, 0xf6, 0x52, 0xa5,
	0x0f, 0x1a, 0xb8, 0x81, 0x37, 0x5b, 0x82, 0x5e, 0x26, 0x50, 0x18, 0x69, 0x6a, 0x42, 0xf1, 0xb1,
	0x70, 0xb2, 0xd2, 0x14, 0x5d, 0xa1, 0xeb, 0xd3, 0xd5, 0x40, 0x09, 0xc3, 0xf8, 0x83, 0x67, 0x19,
	0x98, 0x57, 0x6e, 0xd6, 0xf4, 0xd0, 0x9b, 0x7b, 0xca, 0xe2, 0x1b, 0xe1, 0xb3, 0x21, 0xe6, 0x4e,
	0x08, 0x65, 0x0b, 0x43, 0xce, 0xf1, 0xd4, 0xca, 0xa4, 0x4f, 0x72, 0xd0, 0xe5, 0xf3, 0x24, 0xd1,
	0x50, 0x55, 0xf4, 0xa0, 0xcb, 0xef, 0x29, 0x99, 0xe1, 0xa0, 0xb4, 0xf1, 0x33, 0xd4, 0x74, 0xea,
	0x8d, 0x9e, 0xb9, 0x8b, 0x52, 0xcb, 0x8d, 0x33, 0xba, 0x4f, 0x07, 0x4a, 0xe6, 0xf8, 0x24, 0x2f,
	0x20, 0x57, 0x85, 0x14, 0xf4, 0xc8, 0x5b, 0x23, 0x77, 0x6d, 0x65, 0x5f, 0x66, 0xf9, 0x48, 0x83,
	0xae, 0xed, 0x4e, 0xb9, 0x7f, 0xfa, 0x69, 0x18, 0xda, 0x36, 0x0c, 0xfd, 0x35, 0x0c, 0x7d, 0xb5,
	0x6c, 0xb2, 0x6d, 0xd9, 0xe4, 0xb7, 0x65, 0x93, 0xf7, 0x9b, 0x54, 0x9a, 0xb5, 0x8d, 0x43, 0xa1,
	0xf2, 0x68, 0x6f, 0xdb, 0x68, 0xdc, 0xfe, 0x73, 0x07, 0x4d, 0x5d, 0x42, 0x15, 0x07, 0x7e, 0xd4,
	0xdb, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x44, 0x82, 0xc4, 0x9f, 0x01, 0x00, 0x00,
}

func (m *MsgCreateSuperIdentity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateSuperIdentity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateSuperIdentity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WalletPath) > 0 {
		i -= len(m.WalletPath)
		copy(dAtA[i:], m.WalletPath)
		i = encodeVarintGov(dAtA, i, uint64(len(m.WalletPath)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IdentityAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentityAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentityAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IdentityID) > 0 {
		i -= len(m.IdentityID)
		copy(dAtA[i:], m.IdentityID)
		i = encodeVarintGov(dAtA, i, uint64(len(m.IdentityID)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Mnemonic) > 0 {
		i -= len(m.Mnemonic)
		copy(dAtA[i:], m.Mnemonic)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Mnemonic)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PrivKey) > 0 {
		i -= len(m.PrivKey)
		copy(dAtA[i:], m.PrivKey)
		i = encodeVarintGov(dAtA, i, uint64(len(m.PrivKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintGov(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateSuperIdentity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.WalletPath)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func (m *IdentityAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.PrivKey)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Mnemonic)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.IdentityID)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateSuperIdentity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: MsgCreateSuperIdentity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateSuperIdentity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WalletPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *IdentityAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: IdentityAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentityAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrivKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mnemonic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mnemonic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentityID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)