// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v3alpha/listeners.proto

package envoy_admin_v3alpha

import (
	fmt "fmt"
	core "github.com/datawire/ambassador/pkg/apis/envoy/api/v3alpha/core"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Admin endpoint uses this wrapper for `/listeners` to display listener status information.
// See :ref:`/listeners <operations_admin_interface_listeners>` for more information.
type Listeners struct {
	// List of listener statuses.
	ListenerStatuses     []*ListenerStatus `protobuf:"bytes,1,rep,name=listener_statuses,json=listenerStatuses,proto3" json:"listener_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Listeners) Reset()         { *m = Listeners{} }
func (m *Listeners) String() string { return proto.CompactTextString(m) }
func (*Listeners) ProtoMessage()    {}
func (*Listeners) Descriptor() ([]byte, []int) {
	return fileDescriptor_088657d81081dfcb, []int{0}
}
func (m *Listeners) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Listeners) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Listeners.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Listeners) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listeners.Merge(m, src)
}
func (m *Listeners) XXX_Size() int {
	return m.Size()
}
func (m *Listeners) XXX_DiscardUnknown() {
	xxx_messageInfo_Listeners.DiscardUnknown(m)
}

var xxx_messageInfo_Listeners proto.InternalMessageInfo

func (m *Listeners) GetListenerStatuses() []*ListenerStatus {
	if m != nil {
		return m.ListenerStatuses
	}
	return nil
}

// Details an individual listener's current status.
type ListenerStatus struct {
	// Name of the listener
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The actual local address that the listener is listening on. If a listener was configured
	// to listen on port 0, then this address has the port that was allocated by the OS.
	LocalAddress         *core.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListenerStatus) Reset()         { *m = ListenerStatus{} }
func (m *ListenerStatus) String() string { return proto.CompactTextString(m) }
func (*ListenerStatus) ProtoMessage()    {}
func (*ListenerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_088657d81081dfcb, []int{1}
}
func (m *ListenerStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListenerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListenerStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListenerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerStatus.Merge(m, src)
}
func (m *ListenerStatus) XXX_Size() int {
	return m.Size()
}
func (m *ListenerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerStatus proto.InternalMessageInfo

func (m *ListenerStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListenerStatus) GetLocalAddress() *core.Address {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*Listeners)(nil), "envoy.admin.v3alpha.Listeners")
	proto.RegisterType((*ListenerStatus)(nil), "envoy.admin.v3alpha.ListenerStatus")
}

func init() {
	proto.RegisterFile("envoy/admin/v3alpha/listeners.proto", fileDescriptor_088657d81081dfcb)
}

var fileDescriptor_088657d81081dfcb = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x2f, 0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48, 0xd4,
	0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x06, 0x2b, 0xd2, 0x03, 0x2b, 0xd2, 0x83, 0x2a, 0x92, 0x52, 0x81, 0xea, 0x2c, 0xc8, 0x84,
	0xeb, 0x4b, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x86, 0x6a, 0x55, 0x8a,
	0xe5, 0xe2, 0xf4, 0x81, 0x99, 0x26, 0x14, 0xc0, 0x25, 0x08, 0x33, 0x3a, 0xbe, 0xb8, 0x24, 0xb1,
	0xa4, 0xb4, 0x38, 0xb5, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x59, 0x0f, 0x8b, 0x1d,
	0x7a, 0x30, 0xad, 0xc1, 0x60, 0xc5, 0x41, 0x02, 0x39, 0x28, 0xfc, 0xd4, 0x62, 0xa5, 0x2c, 0x2e,
	0x3e, 0x54, 0x35, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x60, 0xb6, 0x90, 0x0b, 0x17, 0x6f, 0x4e, 0x7e, 0x72, 0x62, 0x4e, 0x3c, 0xd4, 0x6d,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xf2, 0x30, 0x3b, 0x0b, 0x32, 0xe1, 0x36, 0x82, 0xbc,
	0xa0, 0xe7, 0x08, 0x51, 0x16, 0xc4, 0x03, 0xd6, 0x05, 0xe5, 0x39, 0x39, 0x9e, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x5c, 0x8a, 0x99, 0xf9, 0x10, 0xed, 0x05,
	0x45, 0xf9, 0x15, 0x95, 0xd8, 0x5c, 0xef, 0x04, 0x77, 0x5a, 0x71, 0x00, 0x28, 0x2c, 0x02, 0x18,
	0x93, 0xd8, 0xc0, 0x81, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x15, 0x35, 0x4b, 0xc3, 0x76,
	0x01, 0x00, 0x00,
}

func (m *Listeners) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Listeners) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Listeners) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ListenerStatuses) > 0 {
		for iNdEx := len(m.ListenerStatuses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ListenerStatuses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintListeners(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListenerStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListenerStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListenerStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LocalAddress != nil {
		{
			size, err := m.LocalAddress.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListeners(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintListeners(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintListeners(dAtA []byte, offset int, v uint64) int {
	offset -= sovListeners(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Listeners) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ListenerStatuses) > 0 {
		for _, e := range m.ListenerStatuses {
			l = e.Size()
			n += 1 + l + sovListeners(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListenerStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovListeners(uint64(l))
	}
	if m.LocalAddress != nil {
		l = m.LocalAddress.Size()
		n += 1 + l + sovListeners(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovListeners(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListeners(x uint64) (n int) {
	return sovListeners(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Listeners) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListeners
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
			return fmt.Errorf("proto: Listeners: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Listeners: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenerStatuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListeners
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
				return ErrInvalidLengthListeners
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListeners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenerStatuses = append(m.ListenerStatuses, &ListenerStatus{})
			if err := m.ListenerStatuses[len(m.ListenerStatuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListeners(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListeners
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthListeners
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
func (m *ListenerStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListeners
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
			return fmt.Errorf("proto: ListenerStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListenerStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListeners
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
				return ErrInvalidLengthListeners
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListeners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListeners
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
				return ErrInvalidLengthListeners
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListeners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LocalAddress == nil {
				m.LocalAddress = &core.Address{}
			}
			if err := m.LocalAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListeners(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListeners
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthListeners
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
func skipListeners(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListeners
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
					return 0, ErrIntOverflowListeners
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
					return 0, ErrIntOverflowListeners
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
				return 0, ErrInvalidLengthListeners
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthListeners
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowListeners
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
				next, err := skipListeners(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthListeners
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
	ErrInvalidLengthListeners = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListeners   = fmt.Errorf("proto: integer overflow")
)
