// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/grpc_credential/v2alpha/aws_iam.proto

package v2alpha

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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

type AwsIamConfig struct {
	// The `service namespace
	// <https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces>`_
	// of the Grpc endpoint.
	//
	// Example: appmesh
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The `region <https://docs.aws.amazon.com/general/latest/gr/rande.html>`_ hosting the Grpc
	// endpoint. If unspecified, the extension will use the value in the ``AWS_REGION`` environment
	// variable.
	//
	// Example: us-west-2
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsIamConfig) Reset()         { *m = AwsIamConfig{} }
func (m *AwsIamConfig) String() string { return proto.CompactTextString(m) }
func (*AwsIamConfig) ProtoMessage()    {}
func (*AwsIamConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb99e742622a8430, []int{0}
}
func (m *AwsIamConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AwsIamConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AwsIamConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AwsIamConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsIamConfig.Merge(m, src)
}
func (m *AwsIamConfig) XXX_Size() int {
	return m.Size()
}
func (m *AwsIamConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsIamConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AwsIamConfig proto.InternalMessageInfo

func (m *AwsIamConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AwsIamConfig) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func init() {
	proto.RegisterType((*AwsIamConfig)(nil), "envoy.config.grpc_credential.v2alpha.AwsIamConfig")
}

func init() {
	proto.RegisterFile("envoy/config/grpc_credential/v2alpha/aws_iam.proto", fileDescriptor_fb99e742622a8430)
}

var fileDescriptor_fb99e742622a8430 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x4f, 0x2e,
	0x4a, 0x4d, 0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xcc, 0xd1, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48,
	0xd4, 0x4f, 0x2c, 0x2f, 0x8e, 0xcf, 0x4c, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52,
	0x01, 0xeb, 0xd1, 0x83, 0xe8, 0xd1, 0x43, 0xd3, 0xa3, 0x07, 0xd5, 0x23, 0x25, 0x5e, 0x96, 0x98,
	0x93, 0x99, 0x92, 0x58, 0x92, 0xaa, 0x0f, 0x63, 0x40, 0xb4, 0x2b, 0x85, 0x70, 0xf1, 0x38, 0x96,
	0x17, 0x7b, 0x26, 0xe6, 0x3a, 0x83, 0x0d, 0x10, 0xd2, 0xe1, 0xe2, 0x29, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0x8d, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0xe2, 0xdc,
	0xf5, 0xf2, 0x00, 0x33, 0x4b, 0x11, 0x93, 0x02, 0x63, 0x10, 0x37, 0x54, 0xda, 0x2f, 0x31, 0x37,
	0x55, 0x48, 0x8c, 0x8b, 0xad, 0x28, 0x35, 0x3d, 0x33, 0x3f, 0x4f, 0x82, 0x09, 0xa4, 0x2e, 0x08,
	0xca, 0x73, 0x8a, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0xb9, 0x8c, 0x32, 0xf3, 0xf5, 0xc0, 0xae, 0x2c, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x23, 0xc6, 0xc1,
	0x4e, 0xdc, 0x10, 0x57, 0x05, 0x80, 0x1c, 0x19, 0xc0, 0x18, 0xc5, 0x0e, 0x15, 0x4f, 0x62, 0x03,
	0x3b, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x83, 0xe5, 0x44, 0x4d, 0x2b, 0x01, 0x00, 0x00,
}

func (m *AwsIamConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AwsIamConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ServiceName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAwsIam(dAtA, i, uint64(len(m.ServiceName)))
		i += copy(dAtA[i:], m.ServiceName)
	}
	if len(m.Region) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAwsIam(dAtA, i, uint64(len(m.Region)))
		i += copy(dAtA[i:], m.Region)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAwsIam(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AwsIamConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovAwsIam(uint64(l))
	}
	l = len(m.Region)
	if l > 0 {
		n += 1 + l + sovAwsIam(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAwsIam(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAwsIam(x uint64) (n int) {
	return sovAwsIam(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AwsIamConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAwsIam
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
			return fmt.Errorf("proto: AwsIamConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AwsIamConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAwsIam
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
				return ErrInvalidLengthAwsIam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAwsIam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Region", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAwsIam
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
				return ErrInvalidLengthAwsIam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAwsIam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Region = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAwsIam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAwsIam
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAwsIam
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
func skipAwsIam(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAwsIam
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
					return 0, ErrIntOverflowAwsIam
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
					return 0, ErrIntOverflowAwsIam
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
				return 0, ErrInvalidLengthAwsIam
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAwsIam
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAwsIam
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
				next, err := skipAwsIam(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAwsIam
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
	ErrInvalidLengthAwsIam = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAwsIam   = fmt.Errorf("proto: integer overflow")
)
