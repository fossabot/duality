// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/genesis.proto

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

// GenesisState defines the dex module's genesis state.
type GenesisState struct {
	Params                             Params                           `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TickMapList                        []TickMap                        `protobuf:"bytes,2,rep,name=tickMapList,proto3" json:"tickMapList"`
	PairMapList                        []PairMap                        `protobuf:"bytes,3,rep,name=pairMapList,proto3" json:"pairMapList"`
	TokensList                         []Tokens                         `protobuf:"bytes,4,rep,name=tokensList,proto3" json:"tokensList"`
	TokensCount                        uint64                           `protobuf:"varint,5,opt,name=tokensCount,proto3" json:"tokensCount,omitempty"`
	TokenMapList                       []TokenMap                       `protobuf:"bytes,6,rep,name=tokenMapList,proto3" json:"tokenMapList"`
	SharesList                         []Shares                         `protobuf:"bytes,7,rep,name=sharesList,proto3" json:"sharesList"`
	FeeListList                        []FeeList                        `protobuf:"bytes,8,rep,name=feeListList,proto3" json:"feeListList"`
	FeeListCount                       uint64                           `protobuf:"varint,9,opt,name=feeListCount,proto3" json:"feeListCount,omitempty"`
	EdgeRowList                        []EdgeRow                        `protobuf:"bytes,10,rep,name=edgeRowList,proto3" json:"edgeRowList"`
	EdgeRowCount                       uint64                           `protobuf:"varint,11,opt,name=edgeRowCount,proto3" json:"edgeRowCount,omitempty"`
	AdjanceyMatrixList                 []AdjanceyMatrix                 `protobuf:"bytes,12,rep,name=adjanceyMatrixList,proto3" json:"adjanceyMatrixList"`
	AdjanceyMatrixCount                uint64                           `protobuf:"varint,13,opt,name=adjanceyMatrixCount,proto3" json:"adjanceyMatrixCount,omitempty"`
	LimitOrderPoolUserShareMapList     []LimitOrderPoolUserShareMap     `protobuf:"bytes,14,rep,name=limitOrderPoolUserShareMapList,proto3" json:"limitOrderPoolUserShareMapList"`
	LimitOrderPoolUserSharesFilledList []LimitOrderPoolUserSharesFilled `protobuf:"bytes,15,rep,name=limitOrderPoolUserSharesFilledList,proto3" json:"limitOrderPoolUserSharesFilledList"`
	LimitOrderPoolTotalSharesMapList   []LimitOrderPoolTotalSharesMap   `protobuf:"bytes,16,rep,name=limitOrderPoolTotalSharesMapList,proto3" json:"limitOrderPoolTotalSharesMapList"`
	LimitOrderPoolReserveMapList       []LimitOrderPoolReserveMap       `protobuf:"bytes,17,rep,name=limitOrderPoolReserveMapList,proto3" json:"limitOrderPoolReserveMapList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a803aaabd08db59d, []int{0}
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

func (m *GenesisState) GetTickMapList() []TickMap {
	if m != nil {
		return m.TickMapList
	}
	return nil
}

func (m *GenesisState) GetPairMapList() []PairMap {
	if m != nil {
		return m.PairMapList
	}
	return nil
}

func (m *GenesisState) GetTokensList() []Tokens {
	if m != nil {
		return m.TokensList
	}
	return nil
}

func (m *GenesisState) GetTokensCount() uint64 {
	if m != nil {
		return m.TokensCount
	}
	return 0
}

func (m *GenesisState) GetTokenMapList() []TokenMap {
	if m != nil {
		return m.TokenMapList
	}
	return nil
}

func (m *GenesisState) GetSharesList() []Shares {
	if m != nil {
		return m.SharesList
	}
	return nil
}

func (m *GenesisState) GetFeeListList() []FeeList {
	if m != nil {
		return m.FeeListList
	}
	return nil
}

func (m *GenesisState) GetFeeListCount() uint64 {
	if m != nil {
		return m.FeeListCount
	}
	return 0
}

func (m *GenesisState) GetEdgeRowList() []EdgeRow {
	if m != nil {
		return m.EdgeRowList
	}
	return nil
}

func (m *GenesisState) GetEdgeRowCount() uint64 {
	if m != nil {
		return m.EdgeRowCount
	}
	return 0
}

func (m *GenesisState) GetAdjanceyMatrixList() []AdjanceyMatrix {
	if m != nil {
		return m.AdjanceyMatrixList
	}
	return nil
}

func (m *GenesisState) GetAdjanceyMatrixCount() uint64 {
	if m != nil {
		return m.AdjanceyMatrixCount
	}
	return 0
}

func (m *GenesisState) GetLimitOrderPoolUserShareMapList() []LimitOrderPoolUserShareMap {
	if m != nil {
		return m.LimitOrderPoolUserShareMapList
	}
	return nil
}

func (m *GenesisState) GetLimitOrderPoolUserSharesFilledList() []LimitOrderPoolUserSharesFilled {
	if m != nil {
		return m.LimitOrderPoolUserSharesFilledList
	}
	return nil
}

func (m *GenesisState) GetLimitOrderPoolTotalSharesMapList() []LimitOrderPoolTotalSharesMap {
	if m != nil {
		return m.LimitOrderPoolTotalSharesMapList
	}
	return nil
}

func (m *GenesisState) GetLimitOrderPoolReserveMapList() []LimitOrderPoolReserveMap {
	if m != nil {
		return m.LimitOrderPoolReserveMapList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nicholasdotsol.duality.dex.GenesisState")
}

func init() { proto.RegisterFile("dex/genesis.proto", fileDescriptor_a803aaabd08db59d) }

var fileDescriptor_a803aaabd08db59d = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdf, 0x4f, 0x13, 0x31,
	0x1c, 0xdf, 0x09, 0xa2, 0x76, 0x53, 0xa1, 0xf8, 0x80, 0x8b, 0x39, 0x97, 0x69, 0x0c, 0x6a, 0xb8,
	0x33, 0x68, 0x8c, 0xf1, 0x49, 0x51, 0xc1, 0x44, 0x40, 0x32, 0xf0, 0xc5, 0x97, 0xb3, 0xec, 0xca,
	0xa8, 0xdc, 0xd6, 0x4b, 0xdb, 0xc9, 0x78, 0xf1, 0xc9, 0x27, 0x13, 0x13, 0xfe, 0x2c, 0x1e, 0x79,
	0xf4, 0xc9, 0x18, 0xf8, 0x47, 0x4c, 0xbf, 0x6d, 0x5d, 0xcf, 0x8c, 0x1d, 0xf1, 0xed, 0xfa, 0xb9,
	0xcf, 0xaf, 0x7d, 0x7b, 0xed, 0xd0, 0x4c, 0x4a, 0x07, 0x71, 0x87, 0xf6, 0xa8, 0x64, 0x32, 0xca,
	0x05, 0x57, 0x1c, 0xd7, 0x7b, 0xac, 0xbd, 0xcb, 0x33, 0x22, 0x53, 0xae, 0x24, 0xcf, 0xa2, 0xb4,
	0x4f, 0x32, 0xa6, 0x0e, 0xa2, 0x94, 0x0e, 0xea, 0x37, 0x3a, 0xbc, 0xc3, 0x81, 0x16, 0xeb, 0x27,
	0xa3, 0xa8, 0x4f, 0x6b, 0x93, 0x9c, 0x08, 0xd2, 0xb5, 0x1e, 0x75, 0xac, 0x11, 0xc5, 0xda, 0x7b,
	0x49, 0x97, 0xe4, 0x3e, 0x96, 0x13, 0x26, 0x3c, 0x0c, 0x94, 0x8a, 0xef, 0xd1, 0x9e, 0x53, 0xce,
	0xfe, 0x45, 0xfe, 0xa5, 0xc9, 0x5d, 0x22, 0x68, 0x21, 0x60, 0x87, 0xd2, 0x24, 0x63, 0x52, 0xf9,
	0x18, 0x4d, 0x3b, 0x34, 0x11, 0x7c, 0xdf, 0x62, 0x37, 0x35, 0x46, 0xd2, 0xcf, 0xa4, 0xd7, 0xa6,
	0x07, 0x49, 0x97, 0x28, 0xc1, 0x06, 0xf6, 0xd5, 0x7d, 0xfd, 0x2a, 0x63, 0x5d, 0xa6, 0x12, 0x2e,
	0x52, 0x2a, 0x92, 0x9c, 0xf3, 0x2c, 0xe9, 0x4b, 0x2a, 0x12, 0x88, 0xf2, 0xf2, 0x17, 0x4a, 0xa8,
	0x32, 0xd9, 0x61, 0x59, 0x46, 0x53, 0x4b, 0x7f, 0x38, 0x92, 0xae, 0xb8, 0x22, 0x99, 0xe3, 0x0f,
	0xbd, 0xef, 0x8d, 0x24, 0x0b, 0x2a, 0xa9, 0xf8, 0xe2, 0x75, 0x68, 0xfe, 0xa8, 0xa2, 0xda, 0x8a,
	0xd9, 0xa8, 0x4d, 0x45, 0x14, 0xc5, 0x2f, 0xd0, 0x94, 0x99, 0xf9, 0x5c, 0xd0, 0x08, 0xe6, 0xab,
	0x8b, 0xcd, 0xe8, 0xec, 0x8d, 0x8b, 0x36, 0x80, 0xb9, 0x34, 0x79, 0xf4, 0xeb, 0x76, 0xa5, 0x65,
	0x75, 0xf8, 0x1d, 0xaa, 0xea, 0x3d, 0x5a, 0x23, 0xf9, 0x2a, 0x93, 0x6a, 0xee, 0x42, 0x63, 0x62,
	0xbe, 0xba, 0x78, 0x67, 0x9c, 0xcd, 0x96, 0xa1, 0x5b, 0x1f, 0x5f, 0xad, 0xcd, 0xf4, 0xe6, 0x3a,
	0xb3, 0x89, 0x72, 0xb3, 0x0d, 0x43, 0x77, 0x66, 0x9e, 0x1a, 0xbf, 0x45, 0xc8, 0x7c, 0x15, 0xe0,
	0x35, 0x09, 0x5e, 0x63, 0x7f, 0xdf, 0x16, 0xb0, 0xad, 0x95, 0xa7, 0xc5, 0x0d, 0x54, 0x35, 0xab,
	0x57, 0xbc, 0xdf, 0x53, 0x73, 0x17, 0x1b, 0xc1, 0xfc, 0x64, 0xcb, 0x87, 0xf0, 0x3a, 0xaa, 0xc1,
	0xd2, 0x35, 0x9f, 0x82, 0xb4, 0xbb, 0xa5, 0x69, 0xc3, 0xea, 0x05, 0xbd, 0xee, 0x6e, 0x36, 0x19,
	0xdc, 0x2e, 0x95, 0x77, 0xdf, 0x04, 0xb6, 0xeb, 0x3e, 0xd4, 0xea, 0x91, 0xee, 0x50, 0xaa, 0x1f,
	0xc1, 0xea, 0x72, 0xf9, 0x48, 0x97, 0x0d, 0xdd, 0x8d, 0xd4, 0x53, 0xe3, 0x26, 0xaa, 0xd9, 0xa5,
	0x99, 0xc4, 0x15, 0x98, 0x44, 0x01, 0xd3, 0x81, 0xfa, 0xfc, 0xb4, 0xf8, 0x3e, 0x04, 0xa2, 0xf2,
	0xc0, 0x37, 0x86, 0xee, 0x02, 0x3d, 0xb5, 0x0e, 0xb4, 0x4b, 0x13, 0x58, 0x35, 0x81, 0x3e, 0x86,
	0x3f, 0x21, 0xec, 0x0e, 0xe7, 0x1a, 0x9c, 0x4d, 0xc8, 0xad, 0x41, 0xee, 0x83, 0x71, 0xb9, 0x2f,
	0x0b, 0x2a, 0x1b, 0x3f, 0xc2, 0x0b, 0x3f, 0x42, 0xb3, 0x45, 0xd4, 0x94, 0xb9, 0x0a, 0x65, 0x46,
	0xbd, 0xc2, 0xdf, 0x02, 0x14, 0xc2, 0x79, 0x7c, 0xaf, 0x8f, 0xe3, 0x06, 0xe7, 0xd9, 0x07, 0x49,
	0x05, 0x6c, 0x92, 0xfb, 0x44, 0xae, 0x41, 0xc1, 0xa7, 0xe3, 0x0a, 0xae, 0x9e, 0xe9, 0x60, 0xcb,
	0x96, 0x64, 0xe0, 0xc3, 0x00, 0x35, 0xcf, 0xa0, 0xc8, 0x65, 0xb8, 0x6d, 0xa0, 0xca, 0x75, 0xa8,
	0xf2, 0xfc, 0x3f, 0xaa, 0x58, 0x17, 0x5b, 0xe7, 0x1c, 0x59, 0xf8, 0x7b, 0x80, 0x1a, 0x45, 0xda,
	0x96, 0xbe, 0xd3, 0x0c, 0xcf, 0xcd, 0x66, 0x1a, 0x0a, 0x3d, 0x3b, 0x7f, 0xa1, 0xa2, 0x87, 0xad,
	0x53, 0x9a, 0x83, 0xbf, 0xa2, 0x5b, 0x45, 0x4e, 0xcb, 0x5c, 0x99, 0xae, 0xc7, 0x0c, 0xf4, 0x78,
	0x72, 0xfe, 0x1e, 0x43, 0xbd, 0xed, 0x30, 0xd6, 0x7f, 0x69, 0xe5, 0xe8, 0x24, 0x0c, 0x8e, 0x4f,
	0xc2, 0xe0, 0xf7, 0x49, 0x18, 0x1c, 0x9e, 0x86, 0x95, 0xe3, 0xd3, 0xb0, 0xf2, 0xf3, 0x34, 0xac,
	0x7c, 0x5c, 0xe8, 0x30, 0xb5, 0xdb, 0xdf, 0x8e, 0xda, 0xbc, 0x1b, 0xaf, 0xdb, 0xf4, 0xd7, 0x5c,
	0x6d, 0xf2, 0x2c, 0xb6, 0xe9, 0xf1, 0x20, 0x86, 0xbf, 0xb9, 0x83, 0x9c, 0xca, 0xed, 0x29, 0xb8,
	0xdf, 0x1f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xde, 0xf5, 0x5f, 0xb9, 0x8b, 0x07, 0x00, 0x00,
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
	if len(m.LimitOrderPoolReserveMapList) > 0 {
		for iNdEx := len(m.LimitOrderPoolReserveMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderPoolReserveMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x8a
		}
	}
	if len(m.LimitOrderPoolTotalSharesMapList) > 0 {
		for iNdEx := len(m.LimitOrderPoolTotalSharesMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderPoolTotalSharesMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if len(m.LimitOrderPoolUserSharesFilledList) > 0 {
		for iNdEx := len(m.LimitOrderPoolUserSharesFilledList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderPoolUserSharesFilledList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.LimitOrderPoolUserShareMapList) > 0 {
		for iNdEx := len(m.LimitOrderPoolUserShareMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderPoolUserShareMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if m.AdjanceyMatrixCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AdjanceyMatrixCount))
		i--
		dAtA[i] = 0x68
	}
	if len(m.AdjanceyMatrixList) > 0 {
		for iNdEx := len(m.AdjanceyMatrixList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AdjanceyMatrixList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if m.EdgeRowCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EdgeRowCount))
		i--
		dAtA[i] = 0x58
	}
	if len(m.EdgeRowList) > 0 {
		for iNdEx := len(m.EdgeRowList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EdgeRowList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.FeeListCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.FeeListCount))
		i--
		dAtA[i] = 0x48
	}
	if len(m.FeeListList) > 0 {
		for iNdEx := len(m.FeeListList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeListList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.SharesList) > 0 {
		for iNdEx := len(m.SharesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SharesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.TokenMapList) > 0 {
		for iNdEx := len(m.TokenMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.TokensCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TokensCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.TokensList) > 0 {
		for iNdEx := len(m.TokensList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokensList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PairMapList) > 0 {
		for iNdEx := len(m.PairMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PairMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.TickMapList) > 0 {
		for iNdEx := len(m.TickMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TickMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.TickMapList) > 0 {
		for _, e := range m.TickMapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PairMapList) > 0 {
		for _, e := range m.PairMapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TokensList) > 0 {
		for _, e := range m.TokensList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TokensCount != 0 {
		n += 1 + sovGenesis(uint64(m.TokensCount))
	}
	if len(m.TokenMapList) > 0 {
		for _, e := range m.TokenMapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SharesList) > 0 {
		for _, e := range m.SharesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FeeListList) > 0 {
		for _, e := range m.FeeListList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.FeeListCount != 0 {
		n += 1 + sovGenesis(uint64(m.FeeListCount))
	}
	if len(m.EdgeRowList) > 0 {
		for _, e := range m.EdgeRowList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.EdgeRowCount != 0 {
		n += 1 + sovGenesis(uint64(m.EdgeRowCount))
	}
	if len(m.AdjanceyMatrixList) > 0 {
		for _, e := range m.AdjanceyMatrixList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.AdjanceyMatrixCount != 0 {
		n += 1 + sovGenesis(uint64(m.AdjanceyMatrixCount))
	}
	if len(m.LimitOrderPoolUserShareMapList) > 0 {
		for _, e := range m.LimitOrderPoolUserShareMapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LimitOrderPoolUserSharesFilledList) > 0 {
		for _, e := range m.LimitOrderPoolUserSharesFilledList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LimitOrderPoolTotalSharesMapList) > 0 {
		for _, e := range m.LimitOrderPoolTotalSharesMapList {
			l = e.Size()
			n += 2 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LimitOrderPoolReserveMapList) > 0 {
		for _, e := range m.LimitOrderPoolReserveMapList {
			l = e.Size()
			n += 2 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field TickMapList", wireType)
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
			m.TickMapList = append(m.TickMapList, TickMap{})
			if err := m.TickMapList[len(m.TickMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairMapList", wireType)
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
			m.PairMapList = append(m.PairMapList, PairMap{})
			if err := m.PairMapList[len(m.PairMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensList", wireType)
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
			m.TokensList = append(m.TokensList, Tokens{})
			if err := m.TokensList[len(m.TokensList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensCount", wireType)
			}
			m.TokensCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokensCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMapList", wireType)
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
			m.TokenMapList = append(m.TokenMapList, TokenMap{})
			if err := m.TokenMapList[len(m.TokenMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesList", wireType)
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
			m.SharesList = append(m.SharesList, Shares{})
			if err := m.SharesList[len(m.SharesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeListList", wireType)
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
			m.FeeListList = append(m.FeeListList, FeeList{})
			if err := m.FeeListList[len(m.FeeListList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeListCount", wireType)
			}
			m.FeeListCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeeListCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdgeRowList", wireType)
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
			m.EdgeRowList = append(m.EdgeRowList, EdgeRow{})
			if err := m.EdgeRowList[len(m.EdgeRowList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdgeRowCount", wireType)
			}
			m.EdgeRowCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EdgeRowCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjanceyMatrixList", wireType)
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
			m.AdjanceyMatrixList = append(m.AdjanceyMatrixList, AdjanceyMatrix{})
			if err := m.AdjanceyMatrixList[len(m.AdjanceyMatrixList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjanceyMatrixCount", wireType)
			}
			m.AdjanceyMatrixCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdjanceyMatrixCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderPoolUserShareMapList", wireType)
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
			m.LimitOrderPoolUserShareMapList = append(m.LimitOrderPoolUserShareMapList, LimitOrderPoolUserShareMap{})
			if err := m.LimitOrderPoolUserShareMapList[len(m.LimitOrderPoolUserShareMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderPoolUserSharesFilledList", wireType)
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
			m.LimitOrderPoolUserSharesFilledList = append(m.LimitOrderPoolUserSharesFilledList, LimitOrderPoolUserSharesFilled{})
			if err := m.LimitOrderPoolUserSharesFilledList[len(m.LimitOrderPoolUserSharesFilledList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderPoolTotalSharesMapList", wireType)
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
			m.LimitOrderPoolTotalSharesMapList = append(m.LimitOrderPoolTotalSharesMapList, LimitOrderPoolTotalSharesMap{})
			if err := m.LimitOrderPoolTotalSharesMapList[len(m.LimitOrderPoolTotalSharesMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderPoolReserveMapList", wireType)
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
			m.LimitOrderPoolReserveMapList = append(m.LimitOrderPoolReserveMapList, LimitOrderPoolReserveMap{})
			if err := m.LimitOrderPoolReserveMapList[len(m.LimitOrderPoolReserveMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
