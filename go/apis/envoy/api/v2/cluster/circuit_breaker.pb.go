// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/cluster/circuit_breaker.proto

package cluster

import (
	bytes "bytes"
	fmt "fmt"
	core "github.com/datawire/ambassador/go/apis/envoy/api/v2/core"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// :ref:`Circuit breaking<arch_overview_circuit_break>` settings can be
// specified individually for each defined priority.
type CircuitBreakers struct {
	// If multiple :ref:`Thresholds<envoy_api_msg_cluster.CircuitBreakers.Thresholds>`
	// are defined with the same :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`,
	// the first one in the list is used. If no Thresholds is defined for a given
	// :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`, the default values
	// are used.
	Thresholds           []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CircuitBreakers) Reset()         { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()    {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_89bc8d4e21efdd79, []int{0}
}
func (m *CircuitBreakers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CircuitBreakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CircuitBreakers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CircuitBreakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers.Merge(m, src)
}
func (m *CircuitBreakers) XXX_Size() int {
	return m.Size()
}
func (m *CircuitBreakers) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers proto.InternalMessageInfo

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

// A Thresholds defines CircuitBreaker settings for a
// :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`.
type CircuitBreakers_Thresholds struct {
	// The :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`
	// the specified CircuitBreaker settings apply to.
	// [#comment:TODO(htuch): add (validate.rules).enum.defined_only = true once
	// https://github.com/lyft/protoc-gen-validate/issues/42 is resolved.]
	Priority core.RoutingPriority `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.api.v2.core.RoutingPriority" json:"priority,omitempty"`
	// The maximum number of connections that Envoy will make to the upstream
	// cluster. If not specified, the default is 1024.
	MaxConnections *types.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// The maximum number of pending requests that Envoy will allow to the
	// upstream cluster. If not specified, the default is 1024.
	MaxPendingRequests *types.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	// The maximum number of parallel requests that Envoy will make to the
	// upstream cluster. If not specified, the default is 1024.
	MaxRequests *types.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	// The maximum number of parallel retries that Envoy will allow to the
	// upstream cluster. If not specified, the default is 3.
	MaxRetries *types.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// If track_remaining is true, then stats will be published that expose
	// the number of resources remaining until the circuit breakers open. If
	// not specified, the default is false.
	TrackRemaining bool `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	// The maximum number of connection pools per cluster that Envoy will concurrently support at
	// once. If not specified, the default is unlimited. Set this for clusters which create a
	// large number of connection pools. See
	// :ref:`Circuit Breaking <arch_overview_circuit_break_cluster_maximum_connection_pools>` for
	// more details.
	MaxConnectionPools   *types.UInt32Value `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_89bc8d4e21efdd79, []int{0, 0}
}
func (m *CircuitBreakers_Thresholds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CircuitBreakers_Thresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CircuitBreakers_Thresholds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CircuitBreakers_Thresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds) XXX_Size() int {
	return m.Size()
}
func (m *CircuitBreakers_Thresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds) GetPriority() core.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return core.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *types.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *types.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if m != nil {
		return m.TrackRemaining
	}
	return false
}

func (m *CircuitBreakers_Thresholds) GetMaxConnectionPools() *types.UInt32Value {
	if m != nil {
		return m.MaxConnectionPools
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.api.v2.cluster.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.api.v2.cluster.CircuitBreakers.Thresholds")
}

func init() {
	proto.RegisterFile("envoy/api/v2/cluster/circuit_breaker.proto", fileDescriptor_89bc8d4e21efdd79)
}

var fileDescriptor_89bc8d4e21efdd79 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x40, 0xb5, 0x4d, 0x69, 0xab, 0x0d, 0x4a, 0x24, 0x13, 0x21, 0x2b, 0xaa, 0xa2, 0x28, 0x17,
	0x22, 0x0e, 0x6b, 0xe4, 0x9e, 0x01, 0x91, 0xa8, 0x07, 0x2e, 0x95, 0x65, 0xa0, 0x07, 0x2e, 0xd6,
	0xc6, 0x5d, 0xdc, 0x55, 0xed, 0x9d, 0x65, 0x76, 0x1d, 0x9c, 0x1f, 0xe0, 0x5b, 0x10, 0x9f, 0xc1,
	0x89, 0x23, 0x1f, 0xc0, 0x01, 0xe5, 0xc8, 0x57, 0x20, 0x7b, 0x9d, 0x84, 0x56, 0xad, 0x94, 0xdb,
	0x7a, 0x66, 0xde, 0xf3, 0xec, 0xec, 0xd0, 0xe7, 0x42, 0x2d, 0x61, 0x15, 0x70, 0x2d, 0x83, 0x65,
	0x18, 0xa4, 0x79, 0x69, 0xac, 0xc0, 0x20, 0x95, 0x98, 0x96, 0xd2, 0x26, 0x0b, 0x14, 0xfc, 0x46,
	0x20, 0xd3, 0x08, 0x16, 0xbc, 0x41, 0x53, 0xcb, 0xb8, 0x96, 0x6c, 0x19, 0xb2, 0xb6, 0x76, 0x78,
	0x7a, 0xdb, 0x00, 0x28, 0x82, 0x05, 0x37, 0xc2, 0x31, 0xc3, 0x51, 0x06, 0x90, 0xe5, 0x22, 0x68,
	0xbe, 0x16, 0xe5, 0xa7, 0xe0, 0x0b, 0x72, 0xad, 0x05, 0x9a, 0x36, 0x3f, 0xc8, 0x20, 0x83, 0xe6,
	0x18, 0xd4, 0x27, 0x17, 0x9d, 0xfc, 0x38, 0xa4, 0xfd, 0xb9, 0xeb, 0x61, 0xe6, 0x5a, 0x30, 0x5e,
	0x44, 0xa9, 0xbd, 0x46, 0x61, 0xae, 0x21, 0xbf, 0x32, 0x3e, 0x19, 0x77, 0xa6, 0xdd, 0xf0, 0x05,
	0xbb, 0xaf, 0x25, 0x76, 0x07, 0x65, 0xef, 0xb7, 0x5c, 0xfc, 0x9f, 0x63, 0xf8, 0xbb, 0x43, 0xe9,
	0x2e, 0xe5, 0xbd, 0xa2, 0x27, 0x1a, 0x25, 0xa0, 0xb4, 0x2b, 0x9f, 0x8c, 0xc9, 0xb4, 0x17, 0x4e,
	0xee, 0xe8, 0x01, 0x05, 0x8b, 0xa1, 0xb4, 0x52, 0x65, 0x51, 0x5b, 0x19, 0x6f, 0x19, 0xef, 0x9c,
	0xf6, 0x0b, 0x5e, 0x25, 0x29, 0x28, 0x25, 0x52, 0x2b, 0x41, 0x19, 0xff, 0x60, 0x4c, 0xa6, 0xdd,
	0xf0, 0x94, 0xb9, 0x21, 0xb0, 0xcd, 0x10, 0xd8, 0x87, 0xb7, 0xca, 0x9e, 0x85, 0x97, 0x3c, 0x2f,
	0x45, 0xdc, 0x2b, 0x78, 0x35, 0xdf, 0x31, 0xde, 0x05, 0x1d, 0xd4, 0x1a, 0x2d, 0xd4, 0x95, 0x54,
	0x59, 0x82, 0xe2, 0x73, 0x29, 0x8c, 0x35, 0x7e, 0x67, 0x0f, 0x97, 0x57, 0xf0, 0x2a, 0x72, 0x60,
	0xdc, 0x72, 0xde, 0x6b, 0xfa, 0xb8, 0xf6, 0x6d, 0x3d, 0x87, 0x7b, 0x78, 0xba, 0x05, 0xaf, 0xb6,
	0x82, 0x97, 0xb4, 0xeb, 0x04, 0x16, 0xa5, 0x30, 0xfe, 0xa3, 0x3d, 0x78, 0xda, 0xf0, 0x4d, 0xbd,
	0xf7, 0x8c, 0xf6, 0x2d, 0xf2, 0xf4, 0x26, 0x41, 0x51, 0x70, 0xa9, 0xa4, 0xca, 0xfc, 0xa3, 0x31,
	0x99, 0x9e, 0xc4, 0xbd, 0x26, 0x1c, 0x6f, 0xa2, 0x9b, 0x8b, 0xef, 0xe6, 0x97, 0x68, 0x80, 0xdc,
	0xf8, 0xc7, 0x7b, 0x5e, 0x7c, 0x37, 0xc4, 0xa8, 0xe6, 0x66, 0x5f, 0xc9, 0xb7, 0xf5, 0x88, 0xfc,
	0x5c, 0x8f, 0xc8, 0xaf, 0xf5, 0x88, 0xfc, 0x59, 0x8f, 0x08, 0x9d, 0x48, 0x70, 0x4f, 0xaa, 0x11,
	0xaa, 0xd5, 0xbd, 0xcb, 0x33, 0x7b, 0x72, 0x7b, 0x7b, 0xa2, 0xfa, 0x77, 0x11, 0xf9, 0x78, 0xdc,
	0xe6, 0xbf, 0x1f, 0x3c, 0x3d, 0x6f, 0xb0, 0x37, 0x5a, 0xb2, 0xcb, 0x90, 0xcd, 0x5d, 0xf8, 0xe2,
	0xdd, 0xdf, 0x87, 0x12, 0x8b, 0xa3, 0xa6, 0xe5, 0xb3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3e,
	0x8b, 0xe2, 0xe4, 0x6c, 0x03, 0x00, 0x00,
}

func (this *CircuitBreakers) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CircuitBreakers)
	if !ok {
		that2, ok := that.(CircuitBreakers)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Thresholds) != len(that1.Thresholds) {
		return false
	}
	for i := range this.Thresholds {
		if !this.Thresholds[i].Equal(that1.Thresholds[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CircuitBreakers_Thresholds) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CircuitBreakers_Thresholds)
	if !ok {
		that2, ok := that.(CircuitBreakers_Thresholds)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Priority != that1.Priority {
		return false
	}
	if !this.MaxConnections.Equal(that1.MaxConnections) {
		return false
	}
	if !this.MaxPendingRequests.Equal(that1.MaxPendingRequests) {
		return false
	}
	if !this.MaxRequests.Equal(that1.MaxRequests) {
		return false
	}
	if !this.MaxRetries.Equal(that1.MaxRetries) {
		return false
	}
	if this.TrackRemaining != that1.TrackRemaining {
		return false
	}
	if !this.MaxConnectionPools.Equal(that1.MaxConnectionPools) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *CircuitBreakers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CircuitBreakers) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Thresholds) > 0 {
		for _, msg := range m.Thresholds {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CircuitBreakers_Thresholds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CircuitBreakers_Thresholds) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Priority != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCircuitBreaker(dAtA, i, uint64(m.Priority))
	}
	if m.MaxConnections != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCircuitBreaker(dAtA, i, uint64(m.MaxConnections.Size()))
		n1, err := m.MaxConnections.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.MaxPendingRequests != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCircuitBreaker(dAtA, i, uint64(m.MaxPendingRequests.Size()))
		n2, err := m.MaxPendingRequests.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.MaxRequests != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCircuitBreaker(dAtA, i, uint64(m.MaxRequests.Size()))
		n3, err := m.MaxRequests.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.MaxRetries != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCircuitBreaker(dAtA, i, uint64(m.MaxRetries.Size()))
		n4, err := m.MaxRetries.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.TrackRemaining {
		dAtA[i] = 0x30
		i++
		if m.TrackRemaining {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.MaxConnectionPools != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintCircuitBreaker(dAtA, i, uint64(m.MaxConnectionPools.Size()))
		n5, err := m.MaxConnectionPools.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintCircuitBreaker(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CircuitBreakers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Thresholds) > 0 {
		for _, e := range m.Thresholds {
			l = e.Size()
			n += 1 + l + sovCircuitBreaker(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CircuitBreakers_Thresholds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Priority != 0 {
		n += 1 + sovCircuitBreaker(uint64(m.Priority))
	}
	if m.MaxConnections != nil {
		l = m.MaxConnections.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MaxPendingRequests != nil {
		l = m.MaxPendingRequests.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MaxRequests != nil {
		l = m.MaxRequests.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MaxRetries != nil {
		l = m.MaxRetries.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.TrackRemaining {
		n += 2
	}
	if m.MaxConnectionPools != nil {
		l = m.MaxConnectionPools.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCircuitBreaker(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCircuitBreaker(x uint64) (n int) {
	return sovCircuitBreaker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CircuitBreakers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCircuitBreaker
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
			return fmt.Errorf("proto: CircuitBreakers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CircuitBreakers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Thresholds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Thresholds = append(m.Thresholds, &CircuitBreakers_Thresholds{})
			if err := m.Thresholds[len(m.Thresholds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCircuitBreaker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCircuitBreaker
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
func (m *CircuitBreakers_Thresholds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCircuitBreaker
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
			return fmt.Errorf("proto: Thresholds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Thresholds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= core.RoutingPriority(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxConnections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxConnections == nil {
				m.MaxConnections = &types.UInt32Value{}
			}
			if err := m.MaxConnections.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPendingRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxPendingRequests == nil {
				m.MaxPendingRequests = &types.UInt32Value{}
			}
			if err := m.MaxPendingRequests.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRequests == nil {
				m.MaxRequests = &types.UInt32Value{}
			}
			if err := m.MaxRequests.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRetries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRetries == nil {
				m.MaxRetries = &types.UInt32Value{}
			}
			if err := m.MaxRetries.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackRemaining", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
			m.TrackRemaining = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxConnectionPools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxConnectionPools == nil {
				m.MaxConnectionPools = &types.UInt32Value{}
			}
			if err := m.MaxConnectionPools.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCircuitBreaker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCircuitBreaker
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
func skipCircuitBreaker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCircuitBreaker
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
					return 0, ErrIntOverflowCircuitBreaker
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
					return 0, ErrIntOverflowCircuitBreaker
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
				return 0, ErrInvalidLengthCircuitBreaker
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCircuitBreaker
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCircuitBreaker
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
				next, err := skipCircuitBreaker(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCircuitBreaker
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
	ErrInvalidLengthCircuitBreaker = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCircuitBreaker   = fmt.Errorf("proto: integer overflow")
)
