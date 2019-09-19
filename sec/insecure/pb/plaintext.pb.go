// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plaintext.proto

package plaintext_pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	pb "github.com/adam-hanna/go-libp2p-core/crypto/pb"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Exchange struct {
	Id     []byte        `protobuf:"bytes,1,opt,name=id" json:"id"`
	Pubkey *pb.PublicKey `protobuf:"bytes,2,opt,name=pubkey" json:"pubkey,omitempty"`
}

func (m *Exchange) Reset()         { *m = Exchange{} }
func (m *Exchange) String() string { return proto.CompactTextString(m) }
func (*Exchange) ProtoMessage()    {}
func (*Exchange) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba144f73931b711, []int{0}
}
func (m *Exchange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Exchange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Exchange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Exchange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exchange.Merge(m, src)
}
func (m *Exchange) XXX_Size() int {
	return m.Size()
}
func (m *Exchange) XXX_DiscardUnknown() {
	xxx_messageInfo_Exchange.DiscardUnknown(m)
}

var xxx_messageInfo_Exchange proto.InternalMessageInfo

func (m *Exchange) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Exchange) GetPubkey() *pb.PublicKey {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func init() {
	proto.RegisterType((*Exchange)(nil), "plaintext.pb.Exchange")
}

func init() { proto.RegisterFile("plaintext.proto", fileDescriptor_aba144f73931b711) }

var fileDescriptor_aba144f73931b711 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc8, 0x49, 0xcc,
	0xcc, 0x2b, 0x49, 0xad, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x41, 0x12, 0x48,
	0x92, 0x32, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0x4c,
	0x2a, 0x30, 0x2a, 0xd0, 0x4f, 0xcf, 0xd7, 0x85, 0xb0, 0x74, 0x93, 0xf3, 0x8b, 0x52, 0xf5, 0x93,
	0x8b, 0x2a, 0x0b, 0x4a, 0xf2, 0xf5, 0x0b, 0x92, 0xa0, 0x2c, 0x88, 0x31, 0x4a, 0x7e, 0x5c, 0x1c,
	0xae, 0x15, 0xc9, 0x19, 0x89, 0x79, 0xe9, 0xa9, 0x42, 0x22, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x3c, 0x4e, 0x2c, 0x27, 0xee, 0xc9, 0x33, 0x04, 0x31, 0x65, 0xa6, 0x08, 0xe9,
	0x70, 0xb1, 0x15, 0x94, 0x26, 0x65, 0xa7, 0x56, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x89,
	0xe8, 0xc1, 0x0c, 0x48, 0xd2, 0x0b, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xf6, 0x4e, 0xad, 0x0c, 0x82,
	0xaa, 0x71, 0x92, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x06, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x40, 0xde, 0x90, 0x0b, 0xc2, 0x00, 0x00, 0x00,
}

func (m *Exchange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Exchange) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlaintext(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Pubkey != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPlaintext(dAtA, i, uint64(m.Pubkey.Size()))
		n1, err := m.Pubkey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintPlaintext(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Exchange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		l = len(m.Id)
		n += 1 + l + sovPlaintext(uint64(l))
	}
	if m.Pubkey != nil {
		l = m.Pubkey.Size()
		n += 1 + l + sovPlaintext(uint64(l))
	}
	return n
}

func sovPlaintext(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPlaintext(x uint64) (n int) {
	return sovPlaintext(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Exchange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlaintext
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
			return fmt.Errorf("proto: Exchange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Exchange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlaintext
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
				return ErrInvalidLengthPlaintext
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPlaintext
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlaintext
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
				return ErrInvalidLengthPlaintext
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlaintext
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pubkey == nil {
				m.Pubkey = &pb.PublicKey{}
			}
			if err := m.Pubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlaintext(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlaintext
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlaintext
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
func skipPlaintext(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlaintext
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
					return 0, ErrIntOverflowPlaintext
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlaintext
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
				return 0, ErrInvalidLengthPlaintext
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPlaintext
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPlaintext
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPlaintext(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPlaintext
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPlaintext = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlaintext   = fmt.Errorf("proto: integer overflow")
)
