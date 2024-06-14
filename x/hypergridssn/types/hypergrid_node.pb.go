// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hypergridssn/hypergridssn/hypergrid_node.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type HypergridNode struct {
	Pubkey    string `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Rpc       string `protobuf:"bytes,3,opt,name=rpc,proto3" json:"rpc,omitempty"`
	Role      int32  `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	Starttime int32  `protobuf:"varint,5,opt,name=starttime,proto3" json:"starttime,omitempty"`
	Creator   string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *HypergridNode) Reset()         { *m = HypergridNode{} }
func (m *HypergridNode) String() string { return proto.CompactTextString(m) }
func (*HypergridNode) ProtoMessage()    {}
func (*HypergridNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_57288d5cabe9c27d, []int{0}
}
func (m *HypergridNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HypergridNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HypergridNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HypergridNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HypergridNode.Merge(m, src)
}
func (m *HypergridNode) XXX_Size() int {
	return m.Size()
}
func (m *HypergridNode) XXX_DiscardUnknown() {
	xxx_messageInfo_HypergridNode.DiscardUnknown(m)
}

var xxx_messageInfo_HypergridNode proto.InternalMessageInfo

func (m *HypergridNode) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *HypergridNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HypergridNode) GetRpc() string {
	if m != nil {
		return m.Rpc
	}
	return ""
}

func (m *HypergridNode) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *HypergridNode) GetStarttime() int32 {
	if m != nil {
		return m.Starttime
	}
	return 0
}

func (m *HypergridNode) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*HypergridNode)(nil), "hypergridssn.hypergridssn.HypergridNode")
}

func init() {
	proto.RegisterFile("hypergridssn/hypergridssn/hypergrid_node.proto", fileDescriptor_57288d5cabe9c27d)
}

var fileDescriptor_57288d5cabe9c27d = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcb, 0xa8, 0x2c, 0x48,
	0x2d, 0x4a, 0x2f, 0xca, 0x4c, 0x29, 0x2e, 0xce, 0xd3, 0xc7, 0xce, 0x89, 0xcf, 0xcb, 0x4f, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x44, 0x56, 0x82, 0xa2, 0x59, 0x69, 0x26, 0x23,
	0x17, 0xaf, 0x07, 0x4c, 0xc0, 0x2f, 0x3f, 0x25, 0x55, 0x48, 0x8c, 0x8b, 0xad, 0xa0, 0x34, 0x29,
	0x3b, 0xb5, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xca, 0x13, 0x12, 0xe2, 0x62, 0xc9,
	0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x82, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0x45, 0x05, 0xc9,
	0x12, 0xcc, 0x60, 0x21, 0x10, 0x13, 0xa4, 0xaa, 0x28, 0x3f, 0x27, 0x55, 0x82, 0x45, 0x81, 0x51,
	0x83, 0x35, 0x08, 0xcc, 0x16, 0x92, 0xe1, 0xe2, 0x2c, 0x2e, 0x49, 0x2c, 0x2a, 0x29, 0xc9, 0xcc,
	0x4d, 0x95, 0x60, 0x05, 0x4b, 0x20, 0x04, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b,
	0xf2, 0x8b, 0x24, 0xd8, 0xc0, 0xe6, 0xc0, 0xb8, 0x4e, 0x36, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc,
	0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x04, 0xf7, 0x83, 0x2e, 0xc8, 0xd3, 0x15, 0xa8, 0x61, 0x50, 0x52,
	0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0xbb, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x07, 0x19,
	0x7c, 0xf4, 0x2d, 0x01, 0x00, 0x00,
}

func (m *HypergridNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HypergridNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HypergridNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintHypergridNode(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if m.Starttime != 0 {
		i = encodeVarintHypergridNode(dAtA, i, uint64(m.Starttime))
		i--
		dAtA[i] = 0x28
	}
	if m.Role != 0 {
		i = encodeVarintHypergridNode(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Rpc) > 0 {
		i -= len(m.Rpc)
		copy(dAtA[i:], m.Rpc)
		i = encodeVarintHypergridNode(dAtA, i, uint64(len(m.Rpc)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintHypergridNode(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintHypergridNode(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHypergridNode(dAtA []byte, offset int, v uint64) int {
	offset -= sovHypergridNode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HypergridNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovHypergridNode(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovHypergridNode(uint64(l))
	}
	l = len(m.Rpc)
	if l > 0 {
		n += 1 + l + sovHypergridNode(uint64(l))
	}
	if m.Role != 0 {
		n += 1 + sovHypergridNode(uint64(m.Role))
	}
	if m.Starttime != 0 {
		n += 1 + sovHypergridNode(uint64(m.Starttime))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovHypergridNode(uint64(l))
	}
	return n
}

func sovHypergridNode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHypergridNode(x uint64) (n int) {
	return sovHypergridNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HypergridNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHypergridNode
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
			return fmt.Errorf("proto: HypergridNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HypergridNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHypergridNode
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
				return ErrInvalidLengthHypergridNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHypergridNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHypergridNode
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
				return ErrInvalidLengthHypergridNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHypergridNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rpc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHypergridNode
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
				return ErrInvalidLengthHypergridNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHypergridNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rpc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHypergridNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Starttime", wireType)
			}
			m.Starttime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHypergridNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Starttime |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHypergridNode
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
				return ErrInvalidLengthHypergridNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHypergridNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHypergridNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHypergridNode
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
func skipHypergridNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHypergridNode
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
					return 0, ErrIntOverflowHypergridNode
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
					return 0, ErrIntOverflowHypergridNode
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
				return 0, ErrInvalidLengthHypergridNode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHypergridNode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHypergridNode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHypergridNode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHypergridNode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHypergridNode = fmt.Errorf("proto: unexpected end of group")
)
