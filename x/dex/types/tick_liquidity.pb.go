// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/tick_liquidity.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type TickLiquidity struct {
	PairId            *PairId                                 `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	TokenIn           string                                  `protobuf:"bytes,2,opt,name=tokenIn,proto3" json:"tokenIn,omitempty"`
	TickIndex         int64                                   `protobuf:"varint,3,opt,name=tickIndex,proto3" json:"tickIndex,omitempty"`
	LiquidityType     string                                  `protobuf:"bytes,4,opt,name=liquidityType,proto3" json:"liquidityType,omitempty"`
	LiquidityIndex    uint64                                  `protobuf:"varint,5,opt,name=liquidityIndex,proto3" json:"liquidityIndex,omitempty"`
	LPReserve         *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=LPReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"LPReserve" yaml:"LPReserve"`
	LimitOrderTranche *LimitOrderTranche                      `protobuf:"bytes,7,opt,name=limitOrderTranche,proto3" json:"limitOrderTranche,omitempty"`
}

func (m *TickLiquidity) Reset()         { *m = TickLiquidity{} }
func (m *TickLiquidity) String() string { return proto.CompactTextString(m) }
func (*TickLiquidity) ProtoMessage()    {}
func (*TickLiquidity) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bf4777d3c75e20c, []int{0}
}
func (m *TickLiquidity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickLiquidity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickLiquidity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickLiquidity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickLiquidity.Merge(m, src)
}
func (m *TickLiquidity) XXX_Size() int {
	return m.Size()
}
func (m *TickLiquidity) XXX_DiscardUnknown() {
	xxx_messageInfo_TickLiquidity.DiscardUnknown(m)
}

var xxx_messageInfo_TickLiquidity proto.InternalMessageInfo

func (m *TickLiquidity) GetPairId() *PairId {
	if m != nil {
		return m.PairId
	}
	return nil
}

func (m *TickLiquidity) GetTokenIn() string {
	if m != nil {
		return m.TokenIn
	}
	return ""
}

func (m *TickLiquidity) GetTickIndex() int64 {
	if m != nil {
		return m.TickIndex
	}
	return 0
}

func (m *TickLiquidity) GetLiquidityType() string {
	if m != nil {
		return m.LiquidityType
	}
	return ""
}

func (m *TickLiquidity) GetLiquidityIndex() uint64 {
	if m != nil {
		return m.LiquidityIndex
	}
	return 0
}

func (m *TickLiquidity) GetLimitOrderTranche() *LimitOrderTranche {
	if m != nil {
		return m.LimitOrderTranche
	}
	return nil
}

func init() {
	proto.RegisterType((*TickLiquidity)(nil), "nicholasdotsol.duality.dex.TickLiquidity")
}

func init() { proto.RegisterFile("duality/dex/tick_liquidity.proto", fileDescriptor_1bf4777d3c75e20c) }

var fileDescriptor_1bf4777d3c75e20c = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0xf5, 0xd4, 0xae, 0x8d, 0xa7, 0xb8, 0xb4, 0x43, 0x17, 0xaa, 0x29, 0x92, 0x10, 0xad, 0x11,
	0x14, 0x4b, 0xd0, 0xee, 0xbc, 0x34, 0x74, 0x21, 0x30, 0xd8, 0x08, 0xaf, 0xda, 0x85, 0x90, 0x35,
	0x83, 0x3d, 0xe8, 0x31, 0xaa, 0x66, 0x1c, 0xa4, 0xbf, 0xc8, 0x0f, 0xe4, 0x7f, 0xbc, 0xf4, 0x32,
	0x64, 0x21, 0x82, 0xbd, 0xcb, 0x32, 0x5f, 0x10, 0xf4, 0xf0, 0x23, 0x0e, 0xc9, 0x6a, 0xe6, 0x9e,
	0x7b, 0xce, 0xe1, 0xde, 0x33, 0x03, 0x55, 0xbc, 0x76, 0x03, 0x2a, 0x32, 0x13, 0x93, 0xd4, 0x14,
	0xd4, 0xf3, 0x9d, 0x80, 0xfe, 0x5f, 0x53, 0x4c, 0x45, 0x66, 0xc4, 0x09, 0x13, 0x0c, 0xf5, 0x23,
	0xea, 0xad, 0x58, 0xe0, 0x72, 0xcc, 0x04, 0x67, 0x81, 0x51, 0x0b, 0x0c, 0x4c, 0xd2, 0xfe, 0x97,
	0x25, 0x5b, 0xb2, 0x92, 0x66, 0x16, 0xb7, 0x4a, 0xd1, 0xff, 0x71, 0xee, 0x19, 0xd0, 0x90, 0x0a,
	0x87, 0x25, 0x98, 0x24, 0x8e, 0x48, 0xdc, 0xc8, 0x5b, 0x91, 0x9a, 0xf6, 0xf5, 0x9c, 0x16, 0xbb,
	0x34, 0x71, 0x28, 0xae, 0x5a, 0xda, 0x4d, 0x13, 0xf6, 0xe6, 0xd4, 0xf3, 0x27, 0x87, 0x59, 0xd0,
	0x08, 0xb6, 0x0b, 0x8a, 0x85, 0x25, 0xa0, 0x02, 0xfd, 0xc3, 0x2f, 0xcd, 0x78, 0x7d, 0x2c, 0x63,
	0x56, 0x32, 0xed, 0x5a, 0x81, 0x24, 0xd8, 0x11, 0xcc, 0x27, 0x91, 0x15, 0x49, 0xef, 0x54, 0xa0,
	0x77, 0xed, 0x43, 0x89, 0xbe, 0xc1, 0x6e, 0xb1, 0xb3, 0x15, 0x61, 0x92, 0x4a, 0x4d, 0x15, 0xe8,
	0x4d, 0xfb, 0x04, 0xa0, 0xef, 0xb0, 0x77, 0x0c, 0x63, 0x9e, 0xc5, 0x44, 0x6a, 0x95, 0xea, 0xe7,
	0x20, 0x1a, 0xc0, 0x8f, 0x47, 0xa0, 0x32, 0x7a, 0xaf, 0x02, 0xbd, 0x65, 0x5f, 0xa0, 0x28, 0x84,
	0xdd, 0xc9, 0xcc, 0x26, 0x9c, 0x24, 0x57, 0x44, 0x6a, 0x17, 0x4e, 0xe3, 0xe9, 0x26, 0x57, 0xc0,
	0x5d, 0xae, 0x0c, 0x96, 0x54, 0xac, 0xd6, 0x0b, 0xc3, 0x63, 0xa1, 0xe9, 0x31, 0x1e, 0x32, 0x5e,
	0x1f, 0x43, 0x8e, 0x7d, 0x53, 0x64, 0x31, 0xe1, 0x86, 0x15, 0x89, 0x87, 0x5c, 0x39, 0x59, 0x3c,
	0xe6, 0xca, 0xa7, 0xcc, 0x0d, 0x83, 0x91, 0x76, 0x84, 0x34, 0xfb, 0xd4, 0x46, 0xff, 0xe0, 0xe7,
	0x32, 0xfa, 0x69, 0x91, 0xfc, 0xbc, 0x0a, 0x5e, 0xea, 0x94, 0xd9, 0x0d, 0xdf, 0xca, 0x6e, 0x72,
	0x29, 0xb2, 0x5f, 0xfa, 0x8c, 0xff, 0x6c, 0x76, 0x32, 0xd8, 0xee, 0x64, 0x70, 0xbf, 0x93, 0xc1,
	0xf5, 0x5e, 0x6e, 0x6c, 0xf7, 0x72, 0xe3, 0x76, 0x2f, 0x37, 0xfe, 0xfe, 0x3c, 0x5b, 0xa5, 0xb6,
	0x1d, 0x06, 0xee, 0x82, 0x1f, 0x0a, 0x33, 0xad, 0x7e, 0x5a, 0xb1, 0xd3, 0xa2, 0x5d, 0xbe, 0xf6,
	0xef, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x5e, 0x9c, 0x3d, 0x85, 0x02, 0x00, 0x00,
}

func (m *TickLiquidity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickLiquidity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickLiquidity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LimitOrderTranche != nil {
		{
			size, err := m.LimitOrderTranche.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTickLiquidity(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.LPReserve != nil {
		{
			size := m.LPReserve.Size()
			i -= size
			if _, err := m.LPReserve.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTickLiquidity(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.LiquidityIndex != 0 {
		i = encodeVarintTickLiquidity(dAtA, i, uint64(m.LiquidityIndex))
		i--
		dAtA[i] = 0x28
	}
	if len(m.LiquidityType) > 0 {
		i -= len(m.LiquidityType)
		copy(dAtA[i:], m.LiquidityType)
		i = encodeVarintTickLiquidity(dAtA, i, uint64(len(m.LiquidityType)))
		i--
		dAtA[i] = 0x22
	}
	if m.TickIndex != 0 {
		i = encodeVarintTickLiquidity(dAtA, i, uint64(m.TickIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TokenIn) > 0 {
		i -= len(m.TokenIn)
		copy(dAtA[i:], m.TokenIn)
		i = encodeVarintTickLiquidity(dAtA, i, uint64(len(m.TokenIn)))
		i--
		dAtA[i] = 0x12
	}
	if m.PairId != nil {
		{
			size, err := m.PairId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTickLiquidity(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTickLiquidity(dAtA []byte, offset int, v uint64) int {
	offset -= sovTickLiquidity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TickLiquidity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PairId != nil {
		l = m.PairId.Size()
		n += 1 + l + sovTickLiquidity(uint64(l))
	}
	l = len(m.TokenIn)
	if l > 0 {
		n += 1 + l + sovTickLiquidity(uint64(l))
	}
	if m.TickIndex != 0 {
		n += 1 + sovTickLiquidity(uint64(m.TickIndex))
	}
	l = len(m.LiquidityType)
	if l > 0 {
		n += 1 + l + sovTickLiquidity(uint64(l))
	}
	if m.LiquidityIndex != 0 {
		n += 1 + sovTickLiquidity(uint64(m.LiquidityIndex))
	}
	if m.LPReserve != nil {
		l = m.LPReserve.Size()
		n += 1 + l + sovTickLiquidity(uint64(l))
	}
	if m.LimitOrderTranche != nil {
		l = m.LimitOrderTranche.Size()
		n += 1 + l + sovTickLiquidity(uint64(l))
	}
	return n
}

func sovTickLiquidity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTickLiquidity(x uint64) (n int) {
	return sovTickLiquidity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TickLiquidity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTickLiquidity
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
			return fmt.Errorf("proto: TickLiquidity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickLiquidity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickLiquidity
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
				return ErrInvalidLengthTickLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTickLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PairId == nil {
				m.PairId = &PairId{}
			}
			if err := m.PairId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickLiquidity
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
				return ErrInvalidLengthTickLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTickLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndex", wireType)
			}
			m.TickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickLiquidity
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
				return ErrInvalidLengthTickLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTickLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidityType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityIndex", wireType)
			}
			m.LiquidityIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LiquidityIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LPReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickLiquidity
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
				return ErrInvalidLengthTickLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTickLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.LPReserve = &v
			if err := m.LPReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderTranche", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickLiquidity
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
				return ErrInvalidLengthTickLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTickLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LimitOrderTranche == nil {
				m.LimitOrderTranche = &LimitOrderTranche{}
			}
			if err := m.LimitOrderTranche.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTickLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTickLiquidity
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
func skipTickLiquidity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTickLiquidity
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
					return 0, ErrIntOverflowTickLiquidity
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
					return 0, ErrIntOverflowTickLiquidity
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
				return 0, ErrInvalidLengthTickLiquidity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTickLiquidity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTickLiquidity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTickLiquidity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTickLiquidity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTickLiquidity = fmt.Errorf("proto: unexpected end of group")
)
