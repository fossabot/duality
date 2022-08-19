// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/ticks.proto

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

type Ticks struct {
	Price       string                                 `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	Fee         string                                 `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	OrderType   string                                 `protobuf:"bytes,3,opt,name=orderType,proto3" json:"orderType,omitempty"`
	Reserve0    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=reserve0,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reserve0" yaml:"reserve0"`
	Reserve1    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=reserve1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reserve1" yaml:"reserve1"`
	PairPrice   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=pairPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"pairprice" yaml:"pairprice"`
	PairFee     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=pairFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"pairfee" yaml:"pairfee"`
	TotalShares github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=totalShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"totalShares" yaml:"totalShares"`
	Orderparams []*OrderParams                         `protobuf:"bytes,9,rep,name=orderparams,proto3" json:"orderparams,omitempty"`
}

func (m *Ticks) Reset()         { *m = Ticks{} }
func (m *Ticks) String() string { return proto.CompactTextString(m) }
func (*Ticks) ProtoMessage()    {}
func (*Ticks) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2fc2226c4509899, []int{0}
}
func (m *Ticks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ticks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ticks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ticks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticks.Merge(m, src)
}
func (m *Ticks) XXX_Size() int {
	return m.Size()
}
func (m *Ticks) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticks.DiscardUnknown(m)
}

var xxx_messageInfo_Ticks proto.InternalMessageInfo

func (m *Ticks) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Ticks) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *Ticks) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *Ticks) GetOrderparams() []*OrderParams {
	if m != nil {
		return m.Orderparams
	}
	return nil
}

func init() {
	proto.RegisterType((*Ticks)(nil), "nicholasdotsol.duality.dex.Ticks")
}

func init() { proto.RegisterFile("dex/ticks.proto", fileDescriptor_f2fc2226c4509899) }

var fileDescriptor_f2fc2226c4509899 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x13, 0x4a, 0xb6, 0xc5, 0x95, 0xd8, 0x64, 0x4d, 0xc8, 0xaa, 0x50, 0x32, 0xe5, 0x00,
	0xbb, 0x2c, 0x21, 0x70, 0xe3, 0x38, 0x4d, 0x20, 0x84, 0x60, 0x53, 0xba, 0x13, 0x17, 0xe4, 0x26,
	0x5f, 0xdb, 0xa8, 0x09, 0x8e, 0x6c, 0x17, 0x35, 0x6f, 0xc1, 0xfb, 0xf0, 0x02, 0x3d, 0xf6, 0x88,
	0x38, 0x44, 0xa8, 0xbd, 0x71, 0xec, 0x13, 0x20, 0x3b, 0x69, 0x13, 0x09, 0x71, 0xa8, 0x76, 0x8a,
	0xfd, 0xf7, 0xe7, 0xdf, 0x2f, 0xb6, 0x3e, 0xa3, 0xd3, 0x04, 0x16, 0x81, 0x4c, 0xe3, 0x99, 0xf0,
	0x0b, 0xce, 0x24, 0xc3, 0x83, 0xaf, 0x69, 0x3c, 0x65, 0x19, 0x15, 0x09, 0x93, 0x82, 0x65, 0x7e,
	0x32, 0xa7, 0x59, 0x2a, 0x4b, 0x3f, 0x81, 0xc5, 0xe0, 0x7c, 0xc2, 0x26, 0x4c, 0x97, 0x05, 0x6a,
	0x54, 0xef, 0x18, 0x3c, 0x55, 0x08, 0xc6, 0x13, 0xe0, 0x5f, 0x0a, 0xca, 0x69, 0xde, 0x90, 0xbc,
	0x1f, 0x16, 0xb2, 0xee, 0x15, 0x19, 0x9f, 0x23, 0xab, 0xe0, 0x69, 0x0c, 0xc4, 0xbc, 0x30, 0x2f,
	0xed, 0xa8, 0x9e, 0xe0, 0x33, 0xd4, 0x1b, 0x03, 0x90, 0x47, 0x3a, 0x53, 0x43, 0xfc, 0x0c, 0xd9,
	0x9a, 0x73, 0x5f, 0x16, 0x40, 0x7a, 0x3a, 0x6f, 0x03, 0x9c, 0xa2, 0x13, 0x0e, 0x02, 0xf8, 0x37,
	0x78, 0x49, 0x1e, 0xab, 0xc5, 0xeb, 0x8f, 0xcb, 0xca, 0x35, 0x7e, 0x55, 0xee, 0xf3, 0x49, 0x2a,
	0xa7, 0xf3, 0x91, 0x1f, 0xb3, 0x3c, 0x88, 0x99, 0xc8, 0x99, 0x68, 0x3e, 0x57, 0x22, 0x99, 0x05,
	0xb2, 0x2c, 0x40, 0xf8, 0x37, 0x10, 0xff, 0xa9, 0xdc, 0x3d, 0x61, 0x5b, 0xb9, 0xa7, 0x25, 0xcd,
	0xb3, 0x37, 0xde, 0x2e, 0xf1, 0xa2, 0xfd, 0x62, 0x47, 0x15, 0x12, 0xeb, 0x81, 0xaa, 0xf0, 0x1f,
	0x55, 0xd8, 0xaa, 0x42, 0x9c, 0x23, 0xbb, 0xa0, 0x29, 0xbf, 0xd3, 0xf7, 0x73, 0xa4, 0x5d, 0xb7,
	0x07, 0xbb, 0x34, 0x42, 0xdf, 0xea, 0xb6, 0x72, 0xcf, 0x6a, 0xd9, 0x3e, 0xf2, 0xa2, 0xd6, 0x80,
	0x01, 0x1d, 0xab, 0xc9, 0x5b, 0x00, 0x72, 0xac, 0x65, 0x1f, 0x0e, 0x96, 0x69, 0xc0, 0x18, 0x94,
	0xea, 0x49, 0xab, 0x1a, 0x03, 0x78, 0xd1, 0x8e, 0x8d, 0xe7, 0xa8, 0x2f, 0x99, 0xa4, 0xd9, 0x70,
	0x4a, 0x39, 0x08, 0x72, 0xa2, 0x55, 0xc3, 0x83, 0x55, 0x5d, 0xc8, 0xb6, 0x72, 0x71, 0xad, 0xeb,
	0x84, 0x5e, 0xd4, 0x2d, 0xc1, 0xef, 0x51, 0x5f, 0xf7, 0x4b, 0xdd, 0x87, 0xc4, 0xbe, 0xe8, 0x5d,
	0xf6, 0x5f, 0xbd, 0xf0, 0xff, 0xdf, 0xd2, 0xfe, 0xad, 0x2a, 0xbf, 0xd3, 0xe5, 0x51, 0x77, 0xef,
	0xf5, 0xbb, 0xe5, 0xda, 0x31, 0x57, 0x6b, 0xc7, 0xfc, 0xbd, 0x76, 0xcc, 0xef, 0x1b, 0xc7, 0x58,
	0x6d, 0x1c, 0xe3, 0xe7, 0xc6, 0x31, 0x3e, 0x5f, 0x75, 0x7e, 0xff, 0x53, 0x43, 0xbe, 0x61, 0x72,
	0xc8, 0xb2, 0xa0, 0x21, 0x07, 0x8b, 0x40, 0x3f, 0x2b, 0x75, 0x92, 0xd1, 0x91, 0x7e, 0x0d, 0xaf,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xb5, 0x51, 0xd9, 0x6a, 0x03, 0x00, 0x00,
}

func (m *Ticks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ticks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ticks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Orderparams) > 0 {
		for iNdEx := len(m.Orderparams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Orderparams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTicks(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	{
		size := m.TotalShares.Size()
		i -= size
		if _, err := m.TotalShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.PairFee.Size()
		i -= size
		if _, err := m.PairFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.PairPrice.Size()
		i -= size
		if _, err := m.PairPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Reserve1.Size()
		i -= size
		if _, err := m.Reserve1.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Reserve0.Size()
		i -= size
		if _, err := m.Reserve0.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.OrderType) > 0 {
		i -= len(m.OrderType)
		copy(dAtA[i:], m.OrderType)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.OrderType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicks(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Ticks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = len(m.OrderType)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = m.Reserve0.Size()
	n += 1 + l + sovTicks(uint64(l))
	l = m.Reserve1.Size()
	n += 1 + l + sovTicks(uint64(l))
	l = m.PairPrice.Size()
	n += 1 + l + sovTicks(uint64(l))
	l = m.PairFee.Size()
	n += 1 + l + sovTicks(uint64(l))
	l = m.TotalShares.Size()
	n += 1 + l + sovTicks(uint64(l))
	if len(m.Orderparams) > 0 {
		for _, e := range m.Orderparams {
			l = e.Size()
			n += 1 + l + sovTicks(uint64(l))
		}
	}
	return n
}

func sovTicks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicks(x uint64) (n int) {
	return sovTicks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ticks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicks
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
			return fmt.Errorf("proto: Ticks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ticks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserve0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reserve0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserve1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reserve1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PairPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PairFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orderparams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orderparams = append(m.Orderparams, &OrderParams{})
			if err := m.Orderparams[len(m.Orderparams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicks
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
func skipTicks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicks
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
					return 0, ErrIntOverflowTicks
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
					return 0, ErrIntOverflowTicks
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
				return 0, ErrInvalidLengthTicks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicks = fmt.Errorf("proto: unexpected end of group")
)
