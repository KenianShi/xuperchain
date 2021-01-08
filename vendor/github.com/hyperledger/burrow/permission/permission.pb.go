// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: permission.proto

package permission

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	github_com_hyperledger_burrow_crypto "github.com/hyperledger/burrow/crypto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AccountPermissions struct {
	Base             BasePermissions `protobuf:"bytes,1,opt,name=Base" json:"Base"`
	Roles            []string        `protobuf:"bytes,2,rep,name=Roles" json:"Roles,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *AccountPermissions) Reset()         { *m = AccountPermissions{} }
func (m *AccountPermissions) String() string { return proto.CompactTextString(m) }
func (*AccountPermissions) ProtoMessage()    {}
func (*AccountPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{0}
}
func (m *AccountPermissions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountPermissions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountPermissions.Merge(m, src)
}
func (m *AccountPermissions) XXX_Size() int {
	return m.Size()
}
func (m *AccountPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_AccountPermissions proto.InternalMessageInfo

func (m *AccountPermissions) GetBase() BasePermissions {
	if m != nil {
		return m.Base
	}
	return BasePermissions{}
}

func (m *AccountPermissions) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (*AccountPermissions) XXX_MessageName() string {
	return "permission.AccountPermissions"
}

type BasePermissions struct {
	Perms            PermFlag `protobuf:"varint,1,opt,name=Perms,casttype=PermFlag" json:"Perms"`
	SetBit           PermFlag `protobuf:"varint,2,opt,name=SetBit,casttype=PermFlag" json:"SetBit"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BasePermissions) Reset()      { *m = BasePermissions{} }
func (*BasePermissions) ProtoMessage() {}
func (*BasePermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{1}
}
func (m *BasePermissions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasePermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasePermissions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasePermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasePermissions.Merge(m, src)
}
func (m *BasePermissions) XXX_Size() int {
	return m.Size()
}
func (m *BasePermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_BasePermissions.DiscardUnknown(m)
}

var xxx_messageInfo_BasePermissions proto.InternalMessageInfo

func (m *BasePermissions) GetPerms() PermFlag {
	if m != nil {
		return m.Perms
	}
	return 0
}

func (m *BasePermissions) GetSetBit() PermFlag {
	if m != nil {
		return m.SetBit
	}
	return 0
}

func (*BasePermissions) XXX_MessageName() string {
	return "permission.BasePermissions"
}

type PermArgs struct {
	// The permission function
	Action PermFlag `protobuf:"varint,1,opt,name=Action,casttype=PermFlag" json:"Action"`
	// The target of the action
	Target *github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,2,opt,name=Target,customtype=github.com/hyperledger/burrow/crypto.Address" json:"Target,omitempty"`
	// Possible arguments
	Permission           *PermFlag `protobuf:"varint,3,opt,name=Permission,casttype=PermFlag" json:"Permission,omitempty"`
	Role                 *string   `protobuf:"bytes,4,opt,name=Role" json:"Role,omitempty"`
	Value                *bool     `protobuf:"varint,5,opt,name=Value" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PermArgs) Reset()      { *m = PermArgs{} }
func (*PermArgs) ProtoMessage() {}
func (*PermArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{2}
}
func (m *PermArgs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PermArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PermArgs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PermArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermArgs.Merge(m, src)
}
func (m *PermArgs) XXX_Size() int {
	return m.Size()
}
func (m *PermArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_PermArgs.DiscardUnknown(m)
}

var xxx_messageInfo_PermArgs proto.InternalMessageInfo

func (m *PermArgs) GetAction() PermFlag {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *PermArgs) GetPermission() PermFlag {
	if m != nil && m.Permission != nil {
		return *m.Permission
	}
	return 0
}

func (m *PermArgs) GetRole() string {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return ""
}

func (m *PermArgs) GetValue() bool {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return false
}

func (*PermArgs) XXX_MessageName() string {
	return "permission.PermArgs"
}
func init() {
	proto.RegisterType((*AccountPermissions)(nil), "permission.AccountPermissions")
	golang_proto.RegisterType((*AccountPermissions)(nil), "permission.AccountPermissions")
	proto.RegisterType((*BasePermissions)(nil), "permission.BasePermissions")
	golang_proto.RegisterType((*BasePermissions)(nil), "permission.BasePermissions")
	proto.RegisterType((*PermArgs)(nil), "permission.PermArgs")
	golang_proto.RegisterType((*PermArgs)(nil), "permission.PermArgs")
}

func init() { proto.RegisterFile("permission.proto", fileDescriptor_c837ef01cbda0ad8) }
func init() { golang_proto.RegisterFile("permission.proto", fileDescriptor_c837ef01cbda0ad8) }

var fileDescriptor_c837ef01cbda0ad8 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x4f, 0xc2, 0x40,
	0x14, 0xc7, 0x7b, 0x50, 0x48, 0x39, 0x49, 0x24, 0x17, 0x87, 0x46, 0x93, 0xb6, 0x61, 0x30, 0x1d,
	0xb0, 0x35, 0x26, 0x2e, 0x0c, 0x26, 0xed, 0x60, 0x1c, 0x4d, 0x35, 0x0e, 0x6e, 0xa5, 0x9c, 0xa5,
	0xb1, 0x70, 0xf5, 0xee, 0x1a, 0xc3, 0xb7, 0x70, 0x64, 0x54, 0x3e, 0x89, 0x23, 0x23, 0xb3, 0x03,
	0x21, 0xf0, 0x2d, 0x9c, 0xcc, 0x5d, 0x09, 0x54, 0x13, 0xdc, 0xee, 0xff, 0xfe, 0xff, 0xf7, 0x7e,
	0x7d, 0xaf, 0xb0, 0x95, 0x61, 0x3a, 0x4c, 0x18, 0x4b, 0xc8, 0xc8, 0xc9, 0x28, 0xe1, 0x04, 0xc1,
	0x5d, 0xe5, 0xf8, 0x2c, 0x4e, 0xf8, 0x20, 0xef, 0x39, 0x11, 0x19, 0xba, 0x31, 0x89, 0x89, 0x2b,
	0x23, 0xbd, 0xfc, 0x49, 0x2a, 0x29, 0xe4, 0xab, 0x68, 0x6d, 0x3f, 0x43, 0xe4, 0x45, 0x11, 0xc9,
	0x47, 0xfc, 0x76, 0x3b, 0x83, 0xa1, 0x4b, 0xa8, 0xfa, 0x21, 0xc3, 0x3a, 0xb0, 0x80, 0x7d, 0x70,
	0x71, 0xe2, 0x94, 0x88, 0xa2, 0x5e, 0x8a, 0xfa, 0xea, 0x6c, 0x61, 0x2a, 0x81, 0x8c, 0xa3, 0x23,
	0x58, 0x0b, 0x48, 0x8a, 0x99, 0x5e, 0xb1, 0xaa, 0x76, 0x23, 0x28, 0x44, 0x57, 0x7b, 0x9b, 0x9a,
	0xca, 0x64, 0x6a, 0x2a, 0xed, 0x17, 0x78, 0xf8, 0xa7, 0x1d, 0x9d, 0xc2, 0x9a, 0x90, 0x4c, 0xa2,
	0x54, 0xbf, 0x25, 0xa6, 0x7d, 0x2f, 0x4c, 0x4d, 0x14, 0xaf, 0xd3, 0x30, 0x0e, 0x0a, 0x1b, 0xd9,
	0xb0, 0x7e, 0x87, 0xb9, 0x9f, 0x70, 0xbd, 0xb2, 0x27, 0xb8, 0xf1, 0xbb, 0xcd, 0xc9, 0xbb, 0xa9,
	0x6c, 0x91, 0x4b, 0x00, 0x65, 0xc4, 0xa3, 0xb1, 0x1c, 0xe2, 0x45, 0x3c, 0x21, 0xa3, 0xbd, 0xb4,
	0x8d, 0x8f, 0x6e, 0x60, 0xfd, 0x3e, 0xa4, 0x31, 0x2e, 0x70, 0x4d, 0xff, 0xfc, 0x6b, 0x61, 0x76,
	0x4a, 0x97, 0x1d, 0x8c, 0x33, 0x4c, 0x53, 0xdc, 0x8f, 0x31, 0x75, 0x7b, 0x39, 0xa5, 0xe4, 0xd5,
	0x8d, 0xe8, 0x38, 0xe3, 0xc4, 0xf1, 0xfa, 0x7d, 0x8a, 0x19, 0x0b, 0x36, 0xfd, 0xa8, 0x03, 0xe1,
	0x6e, 0x5f, 0xbd, 0x2a, 0xb9, 0xcd, 0x5f, 0xcc, 0x92, 0x8f, 0x10, 0x54, 0xc5, 0xd1, 0x74, 0xd5,
	0x02, 0x76, 0x23, 0x90, 0x6f, 0x71, 0xd5, 0x87, 0x30, 0xcd, 0xb1, 0x5e, 0xb3, 0x80, 0xad, 0x05,
	0x85, 0xe8, 0x6a, 0x62, 0xcd, 0xf9, 0x87, 0xa9, 0xf8, 0x57, 0xb3, 0x95, 0x01, 0xe6, 0x2b, 0x03,
	0x2c, 0x57, 0x06, 0xf8, 0x5c, 0x1b, 0x60, 0xb6, 0x36, 0xc0, 0xa3, 0xfd, 0xff, 0xd7, 0xee, 0xfe,
	0xe8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xa6, 0x3c, 0xf5, 0x50, 0x02, 0x00, 0x00,
}

func (m *AccountPermissions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountPermissions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountPermissions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Roles[iNdEx])
			copy(dAtA[i:], m.Roles[iNdEx])
			i = encodeVarintPermission(dAtA, i, uint64(len(m.Roles[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Base.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPermission(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BasePermissions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasePermissions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasePermissions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	i = encodeVarintPermission(dAtA, i, uint64(m.SetBit))
	i--
	dAtA[i] = 0x10
	i = encodeVarintPermission(dAtA, i, uint64(m.Perms))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *PermArgs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PermArgs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PermArgs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		i--
		if *m.Value {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Role != nil {
		i -= len(*m.Role)
		copy(dAtA[i:], *m.Role)
		i = encodeVarintPermission(dAtA, i, uint64(len(*m.Role)))
		i--
		dAtA[i] = 0x22
	}
	if m.Permission != nil {
		i = encodeVarintPermission(dAtA, i, uint64(*m.Permission))
		i--
		dAtA[i] = 0x18
	}
	if m.Target != nil {
		{
			size := m.Target.Size()
			i -= size
			if _, err := m.Target.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintPermission(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	i = encodeVarintPermission(dAtA, i, uint64(m.Action))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintPermission(dAtA []byte, offset int, v uint64) int {
	offset -= sovPermission(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountPermissions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Base.Size()
	n += 1 + l + sovPermission(uint64(l))
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovPermission(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BasePermissions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovPermission(uint64(m.Perms))
	n += 1 + sovPermission(uint64(m.SetBit))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PermArgs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovPermission(uint64(m.Action))
	if m.Target != nil {
		l = m.Target.Size()
		n += 1 + l + sovPermission(uint64(l))
	}
	if m.Permission != nil {
		n += 1 + sovPermission(uint64(*m.Permission))
	}
	if m.Role != nil {
		l = len(*m.Role)
		n += 1 + l + sovPermission(uint64(l))
	}
	if m.Value != nil {
		n += 2
	}
	return n
}

func sovPermission(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPermission(x uint64) (n int) {
	return sovPermission(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountPermissions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
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
			return fmt.Errorf("proto: AccountPermissions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountPermissions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
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
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Base.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
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
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BasePermissions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
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
			return fmt.Errorf("proto: BasePermissions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasePermissions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Perms", wireType)
			}
			m.Perms = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Perms |= PermFlag(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetBit", wireType)
			}
			m.SetBit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SetBit |= PermFlag(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PermArgs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
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
			return fmt.Errorf("proto: PermArgs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PermArgs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= PermFlag(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_hyperledger_burrow_crypto.Address
			m.Target = &v
			if err := m.Target.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var v PermFlag
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= PermFlag(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Permission = &v
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
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
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Role = &s
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
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
			b := bool(v != 0)
			m.Value = &b
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermission
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
func skipPermission(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPermission
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
					return 0, ErrIntOverflowPermission
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
					return 0, ErrIntOverflowPermission
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
				return 0, ErrInvalidLengthPermission
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPermission
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPermission
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPermission        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPermission          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPermission = fmt.Errorf("proto: unexpected end of group")
)