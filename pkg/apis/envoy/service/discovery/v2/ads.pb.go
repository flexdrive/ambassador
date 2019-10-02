// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/discovery/v2/ads.proto

package envoy_service_discovery_v2

import (
	context "context"
	fmt "fmt"
	v2 "github.com/datawire/ambassador/pkg/apis/envoy/api/v2"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221
type AdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdsDummy) Reset()         { *m = AdsDummy{} }
func (m *AdsDummy) String() string { return proto.CompactTextString(m) }
func (*AdsDummy) ProtoMessage()    {}
func (*AdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_187fd5dcc2dab695, []int{0}
}
func (m *AdsDummy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdsDummy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdsDummy.Merge(m, src)
}
func (m *AdsDummy) XXX_Size() int {
	return m.Size()
}
func (m *AdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_AdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_AdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdsDummy)(nil), "envoy.service.discovery.v2.AdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v2/ads.proto", fileDescriptor_187fd5dcc2dab695)
}

var fileDescriptor_187fd5dcc2dab695 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4e, 0xc3, 0x40,
	0x10, 0x45, 0x99, 0x06, 0xa1, 0x2d, 0x5d, 0x81, 0x85, 0x8c, 0x14, 0x52, 0xa4, 0x1a, 0x23, 0x73,
	0x82, 0x44, 0xa6, 0x8f, 0x92, 0x13, 0x6c, 0xec, 0x91, 0xb5, 0x02, 0x67, 0x96, 0x9d, 0xf5, 0x0a,
	0xdf, 0x80, 0xa3, 0x51, 0x72, 0x04, 0xe4, 0x9e, 0x3b, 0x20, 0xdb, 0x90, 0x20, 0xc0, 0xd4, 0xf3,
	0xfe, 0xff, 0xa3, 0xa7, 0xe6, 0xb4, 0x0f, 0xdc, 0xa6, 0x42, 0x2e, 0x98, 0x82, 0xd2, 0xd2, 0x48,
	0xc1, 0x81, 0x5c, 0x9b, 0x86, 0x2c, 0xd5, 0xa5, 0xa0, 0x75, 0xec, 0x39, 0x8a, 0x07, 0x0a, 0x3f,
	0x29, 0x3c, 0x50, 0x18, 0xb2, 0xf8, 0x72, 0x6c, 0xd0, 0xd6, 0xf4, 0x99, 0xe3, 0x69, 0x48, 0xce,
	0x94, 0x3a, 0x5b, 0x96, 0x92, 0x37, 0x75, 0xdd, 0x66, 0xef, 0xa0, 0xe2, 0x65, 0x55, 0x39, 0xaa,
	0xb4, 0xa7, 0x32, 0xff, 0x22, 0xb7, 0x63, 0x6b, 0xb4, 0x53, 0x17, 0x5b, 0xef, 0x48, 0xd7, 0x47,
	0x66, 0x43, 0xc2, 0x8d, 0x2b, 0x48, 0xa2, 0x04, 0xc7, 0x17, 0xb4, 0x35, 0x18, 0x32, 0x3c, 0x84,
	0x37, 0xf4, 0xd8, 0x90, 0xf8, 0xf8, 0x6a, 0xf2, 0x2e, 0x96, 0xf7, 0x42, 0xb3, 0x93, 0x05, 0xdc,
	0x40, 0x74, 0xaf, 0xce, 0x73, 0x7a, 0xf0, 0xfa, 0xaf, 0x89, 0xeb, 0x1f, 0x15, 0x3d, 0xf7, 0x6b,
	0x67, 0xfe, 0x3f, 0xf4, 0x7d, 0x6c, 0x75, 0xf7, 0xd2, 0x25, 0xf0, 0xda, 0x25, 0xf0, 0xd6, 0x25,
	0xa0, 0x16, 0x86, 0xc7, 0xac, 0x75, 0xfc, 0xd4, 0xe2, 0xb4, 0xd1, 0x55, 0x6f, 0x6c, 0xdd, 0xdb,
	0x5b, 0xc3, 0x33, 0xc0, 0xee, 0x74, 0x30, 0x79, 0xfb, 0x11, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x25,
	0xf7, 0xf7, 0xab, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AggregatedDiscoveryServiceClient is the client API for AggregatedDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AggregatedDiscoveryServiceClient interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error)
	DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_DeltaAggregatedResourcesClient, error)
}

type aggregatedDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewAggregatedDiscoveryServiceClient(cc *grpc.ClientConn) AggregatedDiscoveryServiceClient {
	return &aggregatedDiscoveryServiceClient{cc}
}

func (c *aggregatedDiscoveryServiceClient) StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.AggregatedDiscoveryService/StreamAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_StreamAggregatedResourcesClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aggregatedDiscoveryServiceClient) DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_DeltaAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[1], "/envoy.service.discovery.v2.AggregatedDiscoveryService/DeltaAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceDeltaAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_DeltaAggregatedResourcesClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AggregatedDiscoveryServiceServer is the server API for AggregatedDiscoveryService service.
type AggregatedDiscoveryServiceServer interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(AggregatedDiscoveryService_StreamAggregatedResourcesServer) error
	DeltaAggregatedResources(AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error
}

// UnimplementedAggregatedDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAggregatedDiscoveryServiceServer struct {
}

func (*UnimplementedAggregatedDiscoveryServiceServer) StreamAggregatedResources(srv AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAggregatedResources not implemented")
}
func (*UnimplementedAggregatedDiscoveryServiceServer) DeltaAggregatedResources(srv AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaAggregatedResources not implemented")
}

func RegisterAggregatedDiscoveryServiceServer(s *grpc.Server, srv AggregatedDiscoveryServiceServer) {
	s.RegisterService(&_AggregatedDiscoveryService_serviceDesc, srv)
}

func _AggregatedDiscoveryService_StreamAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).StreamAggregatedResources(&aggregatedDiscoveryServiceStreamAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_StreamAggregatedResourcesServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AggregatedDiscoveryService_DeltaAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).DeltaAggregatedResources(&aggregatedDiscoveryServiceDeltaAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_DeltaAggregatedResourcesServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AggregatedDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.AggregatedDiscoveryService",
	HandlerType: (*AggregatedDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAggregatedResources",
			Handler:       _AggregatedDiscoveryService_StreamAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaAggregatedResources",
			Handler:       _AggregatedDiscoveryService_DeltaAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/ads.proto",
}

func (m *AdsDummy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdsDummy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AdsDummy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintAds(dAtA []byte, offset int, v uint64) int {
	offset -= sovAds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AdsDummy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAds(x uint64) (n int) {
	return sovAds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AdsDummy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAds
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
			return fmt.Errorf("proto: AdsDummy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdsDummy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAds
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
func skipAds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAds
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
					return 0, ErrIntOverflowAds
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
					return 0, ErrIntOverflowAds
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
				return 0, ErrInvalidLengthAds
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAds
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAds
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
				next, err := skipAds(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAds
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
	ErrInvalidLengthAds = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAds   = fmt.Errorf("proto: integer overflow")
)
