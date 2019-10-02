// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/grpc_credential/v3alpha/file_based_metadata.proto

package envoy_config_grpc_credential_v3alpha

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

type FileBasedMetadataConfig struct {
	// Location or inline data of secret to use for authentication of the Google gRPC connection
	// this secret will be attached to a header of the gRPC connection
	SecretData *core.DataSource `protobuf:"bytes,1,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// Metadata header key to use for sending the secret data
	// if no header key is set, "authorization" header will be used
	HeaderKey string `protobuf:"bytes,2,opt,name=header_key,json=headerKey,proto3" json:"header_key,omitempty"`
	// Prefix to prepend to the secret in the metadata header
	// if no prefix is set, the default is to use no prefix
	HeaderPrefix         string   `protobuf:"bytes,3,opt,name=header_prefix,json=headerPrefix,proto3" json:"header_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileBasedMetadataConfig) Reset()         { *m = FileBasedMetadataConfig{} }
func (m *FileBasedMetadataConfig) String() string { return proto.CompactTextString(m) }
func (*FileBasedMetadataConfig) ProtoMessage()    {}
func (*FileBasedMetadataConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_63617b887cebff62, []int{0}
}
func (m *FileBasedMetadataConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileBasedMetadataConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileBasedMetadataConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileBasedMetadataConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileBasedMetadataConfig.Merge(m, src)
}
func (m *FileBasedMetadataConfig) XXX_Size() int {
	return m.Size()
}
func (m *FileBasedMetadataConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FileBasedMetadataConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FileBasedMetadataConfig proto.InternalMessageInfo

func (m *FileBasedMetadataConfig) GetSecretData() *core.DataSource {
	if m != nil {
		return m.SecretData
	}
	return nil
}

func (m *FileBasedMetadataConfig) GetHeaderKey() string {
	if m != nil {
		return m.HeaderKey
	}
	return ""
}

func (m *FileBasedMetadataConfig) GetHeaderPrefix() string {
	if m != nil {
		return m.HeaderPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*FileBasedMetadataConfig)(nil), "envoy.config.grpc_credential.v3alpha.FileBasedMetadataConfig")
}

func init() {
	proto.RegisterFile("envoy/config/grpc_credential/v3alpha/file_based_metadata.proto", fileDescriptor_63617b887cebff62)
}

var fileDescriptor_63617b887cebff62 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x65, 0x90, 0x90, 0xea, 0xc2, 0x92, 0x01, 0x22, 0x24, 0xa2, 0x52, 0x18, 0x3a, 0xd9,
	0x52, 0xbb, 0x33, 0xa4, 0x88, 0x05, 0x21, 0x45, 0x65, 0x63, 0xb1, 0xae, 0xce, 0xa5, 0xb5, 0x08,
	0xb1, 0xe5, 0x9a, 0x2a, 0x79, 0x1d, 0x9e, 0x86, 0x91, 0x47, 0x40, 0x79, 0x12, 0x64, 0x3b, 0x30,
	0xc0, 0xc2, 0xfa, 0xdb, 0xff, 0xf7, 0xdd, 0x1d, 0xbd, 0xc1, 0x66, 0xaf, 0x3b, 0x2e, 0x75, 0x53,
	0xa9, 0x0d, 0xdf, 0x58, 0x23, 0x85, 0xb4, 0x58, 0x62, 0xe3, 0x14, 0xd4, 0x7c, 0xbf, 0x80, 0xda,
	0x6c, 0x81, 0x57, 0xaa, 0x46, 0xb1, 0x86, 0x1d, 0x96, 0xe2, 0x05, 0x1d, 0x94, 0xe0, 0x80, 0x19,
	0xab, 0x9d, 0x4e, 0xae, 0x43, 0x9f, 0xc5, 0x3e, 0xfb, 0xd5, 0x67, 0x43, 0xff, 0xfc, 0x32, 0x5a,
	0xc0, 0xa8, 0x1f, 0xa4, 0xd4, 0x16, 0xb9, 0x47, 0x46, 0xd0, 0xf4, 0x8d, 0xd0, 0xb3, 0x3b, 0x55,
	0x63, 0xee, 0x2d, 0x0f, 0x83, 0x64, 0x19, 0xb0, 0xc9, 0x92, 0x8e, 0x77, 0x28, 0x2d, 0x3a, 0xe1,
	0xc3, 0x94, 0x4c, 0xc8, 0x6c, 0x3c, 0x9f, 0xb2, 0xa8, 0x06, 0xa3, 0xbe, 0x3d, 0xcc, 0x43, 0xd9,
	0x2d, 0x38, 0x78, 0xd4, 0xaf, 0x56, 0xe2, 0x8a, 0xc6, 0x9a, 0x4f, 0x92, 0x0b, 0x4a, 0xb7, 0x08,
	0x25, 0x5a, 0xf1, 0x8c, 0x5d, 0x7a, 0x30, 0x21, 0xb3, 0xd1, 0x6a, 0x14, 0x93, 0x7b, 0xec, 0x92,
	0x2b, 0x7a, 0x32, 0x3c, 0x1b, 0x8b, 0x95, 0x6a, 0xd3, 0xc3, 0xf0, 0xe3, 0x38, 0x86, 0x45, 0xc8,
	0xf2, 0xa7, 0xf7, 0x3e, 0x23, 0x1f, 0x7d, 0x46, 0x3e, 0xfb, 0x8c, 0xd0, 0xb9, 0xd2, 0x71, 0x06,
	0x63, 0x75, 0xdb, 0xb1, 0xff, 0x5c, 0x22, 0x3f, 0xfd, 0xb3, 0x63, 0xe1, 0xd7, 0x2f, 0xc8, 0xfa,
	0x28, 0xdc, 0x61, 0xf1, 0x15, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x1a, 0x2b, 0x04, 0x92, 0x01, 0x00,
	0x00,
}

func (m *FileBasedMetadataConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileBasedMetadataConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileBasedMetadataConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.HeaderPrefix) > 0 {
		i -= len(m.HeaderPrefix)
		copy(dAtA[i:], m.HeaderPrefix)
		i = encodeVarintFileBasedMetadata(dAtA, i, uint64(len(m.HeaderPrefix)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.HeaderKey) > 0 {
		i -= len(m.HeaderKey)
		copy(dAtA[i:], m.HeaderKey)
		i = encodeVarintFileBasedMetadata(dAtA, i, uint64(len(m.HeaderKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.SecretData != nil {
		{
			size, err := m.SecretData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFileBasedMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFileBasedMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovFileBasedMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FileBasedMetadataConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SecretData != nil {
		l = m.SecretData.Size()
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	l = len(m.HeaderKey)
	if l > 0 {
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	l = len(m.HeaderPrefix)
	if l > 0 {
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFileBasedMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFileBasedMetadata(x uint64) (n int) {
	return sovFileBasedMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileBasedMetadataConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileBasedMetadata
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
			return fmt.Errorf("proto: FileBasedMetadataConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileBasedMetadataConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecretData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
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
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SecretData == nil {
				m.SecretData = &core.DataSource{}
			}
			if err := m.SecretData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
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
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
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
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFileBasedMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFileBasedMetadata
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
func skipFileBasedMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFileBasedMetadata
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
					return 0, ErrIntOverflowFileBasedMetadata
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
					return 0, ErrIntOverflowFileBasedMetadata
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
				return 0, ErrInvalidLengthFileBasedMetadata
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFileBasedMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFileBasedMetadata
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
				next, err := skipFileBasedMetadata(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFileBasedMetadata
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
	ErrInvalidLengthFileBasedMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFileBasedMetadata   = fmt.Errorf("proto: integer overflow")
)
