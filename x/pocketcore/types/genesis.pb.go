// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/pocketcore/genesis.proto

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

// GenesisState defines the auth module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Claims Claims `protobuf:"bytes,2,rep,name=claims,proto3,castrepeated=Claims" json:"claims"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28f1f4b52eeeb96, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClaims() Claims {
	if m != nil {
		return m.Claims
	}
	return nil
}

type Params struct {
	SessionNodeCount           int64    `protobuf:"varint,1,opt,name=sessionNodeCount,proto3" json:"session_node_count"`
	ClaimSubmissionWindow      int64    `protobuf:"varint,2,opt,name=claimSubmissionWindow,proto3" json:"proof_waiting_period"`
	SupportedBlockchains       []string `protobuf:"bytes,3,rep,name=supportedBlockchains,proto3" json:"supported_blockchains"`
	ClaimExpiration            int64    `protobuf:"varint,4,opt,name=claimExpiration,proto3" json:"claim_expiration"`
	ReplayAttackBurnMultiplier int64    `protobuf:"varint,5,opt,name=replayAttackBurnMultiplier,proto3" json:"replay_attack_burn_multiplier"`
	MinimumNumberOfProofs      int64    `protobuf:"varint,6,opt,name=minimumNumberOfProofs,proto3" json:"minimum_number_of_proofs"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28f1f4b52eeeb96, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetSessionNodeCount() int64 {
	if m != nil {
		return m.SessionNodeCount
	}
	return 0
}

func (m *Params) GetClaimSubmissionWindow() int64 {
	if m != nil {
		return m.ClaimSubmissionWindow
	}
	return 0
}

func (m *Params) GetSupportedBlockchains() []string {
	if m != nil {
		return m.SupportedBlockchains
	}
	return nil
}

func (m *Params) GetClaimExpiration() int64 {
	if m != nil {
		return m.ClaimExpiration
	}
	return 0
}

func (m *Params) GetReplayAttackBurnMultiplier() int64 {
	if m != nil {
		return m.ReplayAttackBurnMultiplier
	}
	return 0
}

func (m *Params) GetMinimumNumberOfProofs() int64 {
	if m != nil {
		return m.MinimumNumberOfProofs
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "x.pocketcore.GenesisState")
	proto.RegisterType((*Params)(nil), "x.pocketcore.Params")
}

func init() { proto.RegisterFile("x/pocketcore/genesis.proto", fileDescriptor_d28f1f4b52eeeb96) }

var fileDescriptor_d28f1f4b52eeeb96 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x86, 0x9b, 0xed, 0x1a, 0x70, 0x76, 0xd1, 0x25, 0xb4, 0x4b, 0xb6, 0x68, 0x52, 0xf7, 0xd4,
	0xcb, 0xb6, 0xb0, 0x82, 0x07, 0x11, 0xc1, 0x2c, 0xa2, 0x97, 0xd6, 0x92, 0x3d, 0x08, 0x5e, 0x86,
	0x49, 0x3a, 0xcd, 0x0e, 0x4d, 0xe6, 0x1b, 0x66, 0x26, 0xb4, 0xfb, 0x2f, 0x3c, 0x89, 0x47, 0xcf,
	0xfe, 0x92, 0x3d, 0xee, 0xd1, 0x53, 0x94, 0xf6, 0x22, 0xf9, 0x15, 0xd2, 0x49, 0xec, 0x5a, 0x2d,
	0x9e, 0x32, 0xe4, 0x79, 0xde, 0xf7, 0x1b, 0x86, 0x0f, 0x75, 0x16, 0x03, 0x01, 0xf1, 0x8c, 0xea,
	0x18, 0x24, 0x1d, 0x24, 0x94, 0x53, 0xc5, 0x54, 0x5f, 0x48, 0xd0, 0xe0, 0x1c, 0x2e, 0xfa, 0x77,
	0xac, 0xd3, 0x4a, 0x20, 0x01, 0x03, 0x06, 0xeb, 0x53, 0xe5, 0x74, 0x4e, 0xb6, 0xf2, 0xd5, 0xb1,
	0x42, 0xa7, 0x9f, 0x2c, 0x74, 0xf8, 0xa6, 0x2a, 0xbc, 0xd4, 0x44, 0x53, 0xe7, 0x05, 0xb2, 0x05,
	0x91, 0x24, 0x53, 0xae, 0xd5, 0xb5, 0x7a, 0x07, 0xe7, 0xad, 0xfe, 0x9f, 0x03, 0xfa, 0x63, 0xc3,
	0x82, 0x07, 0x37, 0x85, 0xdf, 0x28, 0x0b, 0xbf, 0x76, 0xc3, 0xfa, 0xeb, 0xbc, 0x45, 0x76, 0x9c,
	0x12, 0x96, 0x29, 0x77, 0xaf, 0xdb, 0xec, 0x1d, 0x9c, 0x1f, 0x6f, 0xa7, 0x87, 0x2a, 0xb9, 0x58,
	0xe3, 0xc0, 0xfd, 0x9d, 0xaf, 0xec, 0xaf, 0xdf, 0x7d, 0xdb, 0x00, 0x15, 0xd6, 0x7f, 0x4e, 0x7f,
	0x36, 0x91, 0x5d, 0x0d, 0x73, 0x02, 0x74, 0xa4, 0xa8, 0x52, 0x0c, 0xf8, 0x08, 0x26, 0xf4, 0x02,
	0x72, 0xae, 0xcd, 0xe5, 0x9a, 0xc1, 0x71, 0x59, 0xf8, 0x4e, 0xcd, 0x30, 0x87, 0x09, 0xc5, 0xf1,
	0x9a, 0x86, 0xff, 0xf8, 0xce, 0x08, 0xb5, 0x4d, 0xf1, 0x65, 0x1e, 0x65, 0xcc, 0xb0, 0xf7, 0x8c,
	0x4f, 0x60, 0xee, 0xee, 0x99, 0x22, 0xb7, 0x2c, 0xfc, 0x96, 0x90, 0x00, 0x53, 0x3c, 0x27, 0x4c,
	0x33, 0x9e, 0x60, 0x41, 0x25, 0x83, 0x49, 0xb8, 0x3b, 0xe6, 0x0c, 0x51, 0x4b, 0xe5, 0x42, 0x80,
	0xd4, 0x74, 0x12, 0xa4, 0x10, 0xcf, 0xe2, 0x2b, 0xc2, 0xb8, 0x72, 0x9b, 0xdd, 0x66, 0xef, 0x7e,
	0x70, 0x52, 0x16, 0x7e, 0x7b, 0xc3, 0x71, 0x74, 0x27, 0x84, 0x3b, 0x63, 0xce, 0x4b, 0xf4, 0xd0,
	0xcc, 0x79, 0xbd, 0x10, 0x4c, 0x12, 0xcd, 0x80, 0xbb, 0xfb, 0xe6, 0x62, 0xad, 0xb2, 0xf0, 0x8f,
	0x0c, 0xc2, 0x74, 0xc3, 0xc2, 0xbf, 0x65, 0x87, 0xa0, 0x8e, 0xa4, 0x22, 0x25, 0xd7, 0xaf, 0xb4,
	0x26, 0xf1, 0x2c, 0xc8, 0x25, 0x1f, 0xe6, 0xa9, 0x66, 0x22, 0x65, 0x54, 0xba, 0xf7, 0x4c, 0xd5,
	0x93, 0xb2, 0xf0, 0x1f, 0x57, 0x16, 0x26, 0x46, 0xc3, 0x51, 0x2e, 0x39, 0xce, 0x36, 0x62, 0xf8,
	0x9f, 0x12, 0x27, 0x44, 0xed, 0x8c, 0x71, 0x96, 0xe5, 0xd9, 0x28, 0xcf, 0x22, 0x2a, 0xdf, 0x4d,
	0xc7, 0xeb, 0xf7, 0x52, 0xae, 0x6d, 0xda, 0x1f, 0x95, 0x85, 0xef, 0xd6, 0x02, 0xe6, 0xc6, 0xc0,
	0x30, 0xc5, 0xe6, 0x4d, 0x55, 0xb8, 0x3b, 0xfa, 0x7c, 0xff, 0xf3, 0x17, 0xbf, 0x11, 0x8c, 0x6f,
	0x96, 0x9e, 0x75, 0xbb, 0xf4, 0xac, 0x1f, 0x4b, 0xcf, 0xfa, 0xb8, 0xf2, 0x1a, 0xb7, 0x2b, 0xaf,
	0xf1, 0x6d, 0xe5, 0x35, 0x3e, 0x3c, 0x4b, 0x98, 0xbe, 0xca, 0xa3, 0x7e, 0x0c, 0xd9, 0x40, 0xc0,
	0x4c, 0x9f, 0x71, 0xaa, 0xe7, 0x20, 0x67, 0xf5, 0x0e, 0x9f, 0x99, 0x7d, 0xde, 0x5a, 0x6e, 0x7d,
	0x2d, 0xa8, 0x8a, 0x6c, 0xb3, 0xdc, 0x4f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xa3, 0xd6,
	0x15, 0x39, 0x03, 0x00, 0x00,
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
	if len(m.Claims) > 0 {
		for iNdEx := len(m.Claims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinimumNumberOfProofs != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MinimumNumberOfProofs))
		i--
		dAtA[i] = 0x30
	}
	if m.ReplayAttackBurnMultiplier != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ReplayAttackBurnMultiplier))
		i--
		dAtA[i] = 0x28
	}
	if m.ClaimExpiration != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ClaimExpiration))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SupportedBlockchains) > 0 {
		for iNdEx := len(m.SupportedBlockchains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SupportedBlockchains[iNdEx])
			copy(dAtA[i:], m.SupportedBlockchains[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.SupportedBlockchains[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ClaimSubmissionWindow != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ClaimSubmissionWindow))
		i--
		dAtA[i] = 0x10
	}
	if m.SessionNodeCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SessionNodeCount))
		i--
		dAtA[i] = 0x8
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Claims) > 0 {
		for _, e := range m.Claims {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SessionNodeCount != 0 {
		n += 1 + sovGenesis(uint64(m.SessionNodeCount))
	}
	if m.ClaimSubmissionWindow != 0 {
		n += 1 + sovGenesis(uint64(m.ClaimSubmissionWindow))
	}
	if len(m.SupportedBlockchains) > 0 {
		for _, s := range m.SupportedBlockchains {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ClaimExpiration != 0 {
		n += 1 + sovGenesis(uint64(m.ClaimExpiration))
	}
	if m.ReplayAttackBurnMultiplier != 0 {
		n += 1 + sovGenesis(uint64(m.ReplayAttackBurnMultiplier))
	}
	if m.MinimumNumberOfProofs != 0 {
		n += 1 + sovGenesis(uint64(m.MinimumNumberOfProofs))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claims", wireType)
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
			m.Claims = append(m.Claims, MsgClaim{})
			if err := m.Claims[len(m.Claims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionNodeCount", wireType)
			}
			m.SessionNodeCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionNodeCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimSubmissionWindow", wireType)
			}
			m.ClaimSubmissionWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimSubmissionWindow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportedBlockchains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportedBlockchains = append(m.SupportedBlockchains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimExpiration", wireType)
			}
			m.ClaimExpiration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimExpiration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplayAttackBurnMultiplier", wireType)
			}
			m.ReplayAttackBurnMultiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplayAttackBurnMultiplier |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumNumberOfProofs", wireType)
			}
			m.MinimumNumberOfProofs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinimumNumberOfProofs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
