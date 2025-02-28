// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stablecoin/v1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	// collRatio is the ratio needed as collateral to exchange for stables
	CollRatio int64 `protobuf:"varint,1,opt,name=coll_ratio,json=collRatio,proto3" json:"coll_ratio,omitempty"`
	// feeRatio is the ratio taken as fees when minting or burning stables
	FeeRatio int64 `protobuf:"varint,2,opt,name=fee_ratio,json=feeRatio,proto3" json:"fee_ratio,omitempty"`
	// efFeeRatio is the ratio taken from the fees that goes to Ecosystem Fund
	EfFeeRatio int64 `protobuf:"varint,3,opt,name=ef_fee_ratio,json=efFeeRatio,proto3" json:"ef_fee_ratio,omitempty"`
	// BonusRateRecoll is the percentage of extra stablecoin value given to the
	// caller of 'Recollateralize' in units of governance tokens.
	BonusRateRecoll int64 `protobuf:"varint,4,opt,name=bonus_rate_recoll,json=bonusRateRecoll,proto3" json:"bonus_rate_recoll,omitempty"`
	// distr_epoch_identifier defines the frequnecy of update for the collateral
	// ratio
	DistrEpochIdentifier string `protobuf:"bytes,5,opt,name=distr_epoch_identifier,json=distrEpochIdentifier,proto3" json:"distr_epoch_identifier,omitempty" yaml:"distr_epoch_identifier"`
	// adjustmentStep is the size of the step taken when updating the collateral
	// ratio
	AdjustmentStep int64 `protobuf:"varint,6,opt,name=adjustment_step,json=adjustmentStep,proto3" json:"adjustment_step,omitempty"`
	// priceLowerBound is the lower bound for the stable coin to trigger a
	// collateral ratio update
	PriceLowerBound int64 `protobuf:"varint,7,opt,name=price_lower_bound,json=priceLowerBound,proto3" json:"price_lower_bound,omitempty"`
	// priceUpperBound is the upper bound for the stable coin to trigger a
	// collateral ratio update
	PriceUpperBound int64 `protobuf:"varint,8,opt,name=price_upper_bound,json=priceUpperBound,proto3" json:"price_upper_bound,omitempty"`
	// isCollateralRatioValid checks if the collateral ratio is correctly updated
	IsCollateralRatioValid bool `protobuf:"varint,9,opt,name=is_collateral_ratio_valid,json=isCollateralRatioValid,proto3" json:"is_collateral_ratio_valid,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bfa1f96ac87927, []int{0}
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

func (m *Params) GetCollRatio() int64 {
	if m != nil {
		return m.CollRatio
	}
	return 0
}

func (m *Params) GetFeeRatio() int64 {
	if m != nil {
		return m.FeeRatio
	}
	return 0
}

func (m *Params) GetEfFeeRatio() int64 {
	if m != nil {
		return m.EfFeeRatio
	}
	return 0
}

func (m *Params) GetBonusRateRecoll() int64 {
	if m != nil {
		return m.BonusRateRecoll
	}
	return 0
}

func (m *Params) GetDistrEpochIdentifier() string {
	if m != nil {
		return m.DistrEpochIdentifier
	}
	return ""
}

func (m *Params) GetAdjustmentStep() int64 {
	if m != nil {
		return m.AdjustmentStep
	}
	return 0
}

func (m *Params) GetPriceLowerBound() int64 {
	if m != nil {
		return m.PriceLowerBound
	}
	return 0
}

func (m *Params) GetPriceUpperBound() int64 {
	if m != nil {
		return m.PriceUpperBound
	}
	return 0
}

func (m *Params) GetIsCollateralRatioValid() bool {
	if m != nil {
		return m.IsCollateralRatioValid
	}
	return false
}

func init() {
	proto.RegisterType((*Params)(nil), "nibiru.stablecoin.v1.Params")
}

func init() { proto.RegisterFile("stablecoin/v1/params.proto", fileDescriptor_f9bfa1f96ac87927) }

var fileDescriptor_f9bfa1f96ac87927 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0x80, 0x1b, 0xab, 0xb5, 0x19, 0xc4, 0x62, 0x28, 0x25, 0x56, 0x1a, 0x63, 0x2f, 0x16, 0x0f,
	0x89, 0xc5, 0x93, 0x1e, 0x5b, 0x14, 0x14, 0x11, 0x89, 0xa8, 0xb0, 0x97, 0x61, 0x92, 0xfc, 0x69,
	0x67, 0x49, 0x32, 0xc3, 0xcc, 0xa4, 0xbb, 0x7d, 0x8b, 0x7d, 0x8b, 0x7d, 0x95, 0x3d, 0xf6, 0xb8,
	0xa7, 0x65, 0x69, 0xdf, 0x60, 0x9f, 0x60, 0x99, 0x49, 0xb7, 0xe9, 0x61, 0x6f, 0xe1, 0xfb, 0xbe,
	0x84, 0x3f, 0x33, 0x3f, 0x1a, 0x4a, 0x45, 0xe2, 0x1c, 0x12, 0x46, 0xcb, 0x70, 0x35, 0x0d, 0x39,
	0x11, 0xa4, 0x90, 0x01, 0x17, 0x4c, 0x31, 0xa7, 0x5f, 0xd2, 0x98, 0x8a, 0x2a, 0x68, 0x92, 0x60,
	0x35, 0x1d, 0xf6, 0x17, 0x6c, 0xc1, 0x4c, 0x10, 0xea, 0xa7, 0xba, 0x1d, 0x5f, 0xb6, 0x51, 0xe7,
	0xb7, 0x79, 0xd9, 0x19, 0x21, 0x94, 0xb0, 0x3c, 0xc7, 0x82, 0x28, 0xca, 0x5c, 0xcb, 0xb7, 0x26,
	0xed, 0xc8, 0xd6, 0x24, 0xd2, 0xc0, 0x79, 0x83, 0xec, 0x0c, 0x60, 0x6f, 0x9f, 0x18, 0xdb, 0xcd,
	0x00, 0x6a, 0xe9, 0xa3, 0x17, 0x90, 0xe1, 0xc6, 0xb7, 0x8d, 0x47, 0x90, 0x7d, 0x7b, 0x28, 0x3e,
	0xa0, 0x57, 0x31, 0x2b, 0x2b, 0xa9, 0x03, 0xc0, 0x02, 0xf4, 0x87, 0xdd, 0xa7, 0x26, 0xeb, 0x19,
	0x11, 0x11, 0x05, 0x91, 0xc1, 0xce, 0x7f, 0x34, 0x48, 0xa9, 0x54, 0x02, 0x03, 0x67, 0xc9, 0x12,
	0xd3, 0x14, 0x4a, 0x45, 0x33, 0x0a, 0xc2, 0x7d, 0xe6, 0x5b, 0x13, 0x7b, 0xf6, 0xee, 0xee, 0xe6,
	0xed, 0x68, 0x4d, 0x8a, 0xfc, 0xcb, 0xf8, 0xf1, 0x6e, 0x1c, 0xf5, 0x8d, 0xf8, 0xaa, 0xf9, 0xf7,
	0x03, 0x76, 0xde, 0xa3, 0x1e, 0x49, 0x4f, 0x2b, 0xa9, 0x0a, 0x28, 0x15, 0x96, 0x0a, 0xb8, 0xdb,
	0x31, 0x23, 0xbc, 0x6c, 0xf0, 0x1f, 0x05, 0x5c, 0x4f, 0xcb, 0x05, 0x4d, 0x00, 0xe7, 0xec, 0x0c,
	0x04, 0x8e, 0x59, 0x55, 0xa6, 0xee, 0xf3, 0x7a, 0x5a, 0x23, 0x7e, 0x6a, 0x3e, 0xd3, 0xb8, 0x69,
	0x2b, 0xce, 0x0f, 0x6d, 0xf7, 0xa8, 0xfd, 0xab, 0x79, 0xdd, 0x7e, 0x46, 0xaf, 0xa9, 0xc4, 0xfa,
	0x27, 0x89, 0x02, 0x41, 0xf6, 0x87, 0x8d, 0x57, 0x24, 0xa7, 0xa9, 0x6b, 0xfb, 0xd6, 0xa4, 0x1b,
	0x0d, 0xa8, 0x9c, 0x1f, 0xbc, 0x39, 0xbb, 0x7f, 0xda, 0xce, 0x7e, 0x5c, 0x6d, 0x3d, 0x6b, 0xb3,
	0xf5, 0xac, 0xdb, 0xad, 0x67, 0x5d, 0xec, 0xbc, 0xd6, 0x66, 0xe7, 0xb5, 0xae, 0x77, 0x5e, 0xeb,
	0xe4, 0xe3, 0x82, 0xaa, 0x65, 0x15, 0x07, 0x09, 0x2b, 0xc2, 0x5f, 0xe6, 0xea, 0xe7, 0x4b, 0x42,
	0xcb, 0xb0, 0x5e, 0x83, 0xf0, 0x3c, 0x3c, 0xda, 0x15, 0xb5, 0xe6, 0x20, 0xe3, 0x8e, 0xb9, 0xfc,
	0x4f, 0xf7, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x13, 0xe2, 0x0f, 0x46, 0x02, 0x00, 0x00,
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
	if m.IsCollateralRatioValid {
		i--
		if m.IsCollateralRatioValid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.PriceUpperBound != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PriceUpperBound))
		i--
		dAtA[i] = 0x40
	}
	if m.PriceLowerBound != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PriceLowerBound))
		i--
		dAtA[i] = 0x38
	}
	if m.AdjustmentStep != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AdjustmentStep))
		i--
		dAtA[i] = 0x30
	}
	if len(m.DistrEpochIdentifier) > 0 {
		i -= len(m.DistrEpochIdentifier)
		copy(dAtA[i:], m.DistrEpochIdentifier)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DistrEpochIdentifier)))
		i--
		dAtA[i] = 0x2a
	}
	if m.BonusRateRecoll != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BonusRateRecoll))
		i--
		dAtA[i] = 0x20
	}
	if m.EfFeeRatio != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EfFeeRatio))
		i--
		dAtA[i] = 0x18
	}
	if m.FeeRatio != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FeeRatio))
		i--
		dAtA[i] = 0x10
	}
	if m.CollRatio != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CollRatio))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CollRatio != 0 {
		n += 1 + sovParams(uint64(m.CollRatio))
	}
	if m.FeeRatio != 0 {
		n += 1 + sovParams(uint64(m.FeeRatio))
	}
	if m.EfFeeRatio != 0 {
		n += 1 + sovParams(uint64(m.EfFeeRatio))
	}
	if m.BonusRateRecoll != 0 {
		n += 1 + sovParams(uint64(m.BonusRateRecoll))
	}
	l = len(m.DistrEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.AdjustmentStep != 0 {
		n += 1 + sovParams(uint64(m.AdjustmentStep))
	}
	if m.PriceLowerBound != 0 {
		n += 1 + sovParams(uint64(m.PriceLowerBound))
	}
	if m.PriceUpperBound != 0 {
		n += 1 + sovParams(uint64(m.PriceUpperBound))
	}
	if m.IsCollateralRatioValid {
		n += 2
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field CollRatio", wireType)
			}
			m.CollRatio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CollRatio |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRatio", wireType)
			}
			m.FeeRatio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeeRatio |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EfFeeRatio", wireType)
			}
			m.EfFeeRatio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EfFeeRatio |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BonusRateRecoll", wireType)
			}
			m.BonusRateRecoll = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BonusRateRecoll |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistrEpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistrEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjustmentStep", wireType)
			}
			m.AdjustmentStep = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdjustmentStep |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceLowerBound", wireType)
			}
			m.PriceLowerBound = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PriceLowerBound |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceUpperBound", wireType)
			}
			m.PriceUpperBound = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PriceUpperBound |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsCollateralRatioValid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.IsCollateralRatioValid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
