// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto

package envoy_config_filter_http_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	v2alpha "github.com/datawire/ambassador/pkg/apis/envoy/config/common/dynamic_forward_proxy/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Configuration for the dynamic forward proxy HTTP filter. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
type FilterConfig struct {
	// The DNS cache configuration that the filter will attach to. Note this configuration must
	// match that of associated :ref:`dynamic forward proxy cluster configuration
	// <envoy_api_field_config.cluster.dynamic_forward_proxy.v2alpha.ClusterConfig.dns_cache_config>`.
	DnsCacheConfig       *v2alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a2356b260c47da, []int{0}
}
func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return m.Size()
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetDnsCacheConfig() *v2alpha.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v2alpha.FilterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto", fileDescriptor_85a2356b260c47da)
}

var fileDescriptor_85a2356b260c47da = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0x4a, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0xa9, 0xcc, 0x4b, 0xcc, 0xcd, 0x4c, 0x8e, 0x4f, 0xcb, 0x2f,
	0x2a, 0x4f, 0x2c, 0x4a, 0x89, 0x2f, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x2f, 0x33, 0x4a, 0xcc, 0x29,
	0xc8, 0x48, 0xc4, 0x2e, 0xab, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x64, 0x06, 0x36, 0x53, 0x0f,
	0x62, 0xa6, 0x1e, 0xc4, 0x4c, 0x3d, 0x90, 0x99, 0x7a, 0xd8, 0x75, 0x41, 0xcd, 0x94, 0x72, 0x44,
	0x71, 0x4b, 0x72, 0x7e, 0x6e, 0x6e, 0x7e, 0x1e, 0x21, 0x67, 0xe4, 0x15, 0xc7, 0x27, 0x27, 0x26,
	0x67, 0xa4, 0x42, 0xac, 0x96, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87,
	0x31, 0x20, 0x12, 0x4a, 0xed, 0x8c, 0x5c, 0x3c, 0x6e, 0x60, 0x97, 0x38, 0x83, 0x8d, 0x17, 0x2a,
	0xe7, 0x12, 0x80, 0x6b, 0x8e, 0x87, 0x58, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xe4, 0xa8,
	0x87, 0xe2, 0x7e, 0x88, 0x3b, 0xf0, 0x3b, 0x5d, 0xcf, 0x25, 0xaf, 0xd8, 0x19, 0x64, 0x12, 0xc4,
	0x70, 0x27, 0xae, 0x5d, 0x2f, 0x0f, 0x30, 0xb3, 0x76, 0x31, 0x32, 0x09, 0x30, 0x06, 0xf1, 0xa5,
	0xa0, 0xca, 0xe5, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x5c, 0x2e, 0x99, 0xf9, 0x10, 0xeb, 0x20, 0xe6, 0x91, 0x17, 0x72, 0x4e, 0x12, 0x2e, 0x10, 0x69,
	0x37, 0x88, 0x6c, 0x00, 0x48, 0x32, 0x00, 0xe4, 0xef, 0x00, 0xc6, 0x24, 0x36, 0x70, 0x00, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xc4, 0xc0, 0x4e, 0xea, 0x01, 0x00, 0x00,
}

func (m *FilterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilterConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.DnsCacheConfig != nil {
		{
			size, err := m.DnsCacheConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDynamicForwardProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
		}
	return len(dAtA) - i, nil
}

func encodeVarintDynamicForwardProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovDynamicForwardProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FilterConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DnsCacheConfig != nil {
		l = m.DnsCacheConfig.Size()
		n += 1 + l + sovDynamicForwardProxy(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDynamicForwardProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDynamicForwardProxy(x uint64) (n int) {
	return sovDynamicForwardProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FilterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicForwardProxy
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
			return fmt.Errorf("proto: FilterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DnsCacheConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicForwardProxy
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
				return ErrInvalidLengthDynamicForwardProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DnsCacheConfig == nil {
				m.DnsCacheConfig = &v2alpha.DnsCacheConfig{}
			}
			if err := m.DnsCacheConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicForwardProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicForwardProxy
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
func skipDynamicForwardProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDynamicForwardProxy
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
					return 0, ErrIntOverflowDynamicForwardProxy
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
					return 0, ErrIntOverflowDynamicForwardProxy
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
				return 0, ErrInvalidLengthDynamicForwardProxy
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDynamicForwardProxy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDynamicForwardProxy
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
				next, err := skipDynamicForwardProxy(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDynamicForwardProxy
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
	ErrInvalidLengthDynamicForwardProxy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDynamicForwardProxy   = fmt.Errorf("proto: integer overflow")
)
