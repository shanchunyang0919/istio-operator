// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1alpha1/istiocontrolplane.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	v1alpha1 "istio.io/api/mesh/v1alpha1"
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

// IstioControlPlane defines an Istio control plane
//
// <!-- crd generation tags
// +cue-gen:IstioControlPlane:groupName:servicemesh.cisco.com
// +cue-gen:IstioControlPlane:version:v1alpha1
// +cue-gen:IstioControlPlane:storageVersion
// +cue-gen:IstioControlPlane:annotations:helm.sh/resource-policy=keep
// +cue-gen:IstioControlPlane:subresource:status
// +cue-gen:IstioControlPlane:scope:Namespaced
// +cue-gen:IstioControlPlane:resource:shortNames=icp,istiocp
// +cue-gen:IstioControlPlane:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:IstioControlPlane:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type IstioControlPlaneSpec struct {
	// Contains the intended version for the Istio control plane.
	// +kubebuilder:validation:Pattern=^1.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Defines mesh-wide settings for the Istio control plane.
	MeshConfig           *v1alpha1.MeshConfig `protobuf:"bytes,2,opt,name=mesh_config,json=meshConfig,proto3" json:"mesh_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IstioControlPlaneSpec) Reset()         { *m = IstioControlPlaneSpec{} }
func (m *IstioControlPlaneSpec) String() string { return proto.CompactTextString(m) }
func (*IstioControlPlaneSpec) ProtoMessage()    {}
func (*IstioControlPlaneSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6817de833805cb8b, []int{0}
}
func (m *IstioControlPlaneSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioControlPlaneSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioControlPlaneSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioControlPlaneSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioControlPlaneSpec.Merge(m, src)
}
func (m *IstioControlPlaneSpec) XXX_Size() int {
	return m.Size()
}
func (m *IstioControlPlaneSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioControlPlaneSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioControlPlaneSpec proto.InternalMessageInfo

func (m *IstioControlPlaneSpec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *IstioControlPlaneSpec) GetMeshConfig() *v1alpha1.MeshConfig {
	if m != nil {
		return m.MeshConfig
	}
	return nil
}

// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type IstioControlPlaneStatus struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioControlPlaneStatus) Reset()         { *m = IstioControlPlaneStatus{} }
func (m *IstioControlPlaneStatus) String() string { return proto.CompactTextString(m) }
func (*IstioControlPlaneStatus) ProtoMessage()    {}
func (*IstioControlPlaneStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6817de833805cb8b, []int{1}
}
func (m *IstioControlPlaneStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioControlPlaneStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioControlPlaneStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioControlPlaneStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioControlPlaneStatus.Merge(m, src)
}
func (m *IstioControlPlaneStatus) XXX_Size() int {
	return m.Size()
}
func (m *IstioControlPlaneStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioControlPlaneStatus.DiscardUnknown(m)
}

var xxx_messageInfo_IstioControlPlaneStatus proto.InternalMessageInfo

func (m *IstioControlPlaneStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*IstioControlPlaneSpec)(nil), "istio_operator.v2.api.v1alpha1.IstioControlPlaneSpec")
	proto.RegisterType((*IstioControlPlaneStatus)(nil), "istio_operator.v2.api.v1alpha1.IstioControlPlaneStatus")
}

func init() {
	proto.RegisterFile("api/v1alpha1/istiocontrolplane.proto", fileDescriptor_6817de833805cb8b)
}

var fileDescriptor_6817de833805cb8b = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x1c, 0xc5, 0x89, 0x88, 0x62, 0x6e, 0x0b, 0x88, 0xe5, 0x86, 0x7a, 0x1c, 0x0e, 0xb7, 0x5c, 0xc2,
	0x55, 0xdc, 0xc5, 0x4e, 0x0e, 0x82, 0x9c, 0x9b, 0xcb, 0xf1, 0x6f, 0x8c, 0xd7, 0x40, 0x9b, 0x7f,
	0x48, 0xd2, 0x0e, 0x7e, 0x42, 0x47, 0x3f, 0x82, 0xf4, 0x93, 0x48, 0x13, 0x7b, 0x0a, 0xb7, 0xbd,
	0xf0, 0xf2, 0x7e, 0x2f, 0x79, 0xf4, 0x06, 0xac, 0x16, 0xfd, 0x06, 0x1a, 0x5b, 0xc3, 0x46, 0x68,
	0x1f, 0x34, 0x4a, 0x34, 0xc1, 0x61, 0x63, 0x1b, 0x30, 0x8a, 0x5b, 0x87, 0x01, 0x59, 0x1e, 0x8d,
	0x1d, 0x5a, 0xe5, 0x20, 0xa0, 0xe3, 0x7d, 0xc1, 0xc1, 0x6a, 0x3e, 0xe5, 0xe6, 0xf3, 0x56, 0xf9,
	0xfa, 0x0f, 0x23, 0xd1, 0xbc, 0xeb, 0x7d, 0xca, 0x2e, 0x3d, 0xbd, 0x7c, 0x1c, 0xd3, 0x65, 0xc2,
	0x3e, 0x8f, 0xd8, 0x17, 0xab, 0x24, 0xcb, 0xe8, 0x79, 0xaf, 0x9c, 0xd7, 0x68, 0x32, 0xb2, 0x20,
	0xab, 0x8b, 0xed, 0x74, 0x64, 0xf7, 0x74, 0x36, 0x02, 0x77, 0x89, 0x93, 0x9d, 0x2c, 0xc8, 0x6a,
	0x56, 0x5c, 0xf3, 0xf8, 0x08, 0x3e, 0x3a, 0x87, 0x66, 0xfe, 0xa4, 0x7c, 0x5d, 0xc6, 0x6b, 0x5b,
	0xda, 0x1e, 0xf4, 0x72, 0x4d, 0xaf, 0x8e, 0x4b, 0x03, 0x84, 0xce, 0x33, 0x46, 0x4f, 0x0d, 0xb4,
	0xea, 0xb7, 0x33, 0xea, 0x87, 0xf2, 0x73, 0xc8, 0xc9, 0xd7, 0x90, 0x93, 0xef, 0x21, 0x27, 0xaf,
	0x77, 0x7b, 0x1d, 0xea, 0xae, 0xe2, 0x12, 0x5b, 0x51, 0x81, 0xf9, 0x00, 0x2d, 0x1b, 0xec, 0xde,
	0xd2, 0x3a, 0xeb, 0x69, 0x04, 0xd1, 0x17, 0xe2, 0xff, 0x78, 0xd5, 0x59, 0xfc, 0xef, 0xed, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x44, 0xea, 0xb8, 0x53, 0x01, 0x00, 0x00,
}

func (m *IstioControlPlaneSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioControlPlaneSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioControlPlaneSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MeshConfig != nil {
		{
			size, err := m.MeshConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIstiocontrolplane(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintIstiocontrolplane(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IstioControlPlaneStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioControlPlaneStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioControlPlaneStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintIstiocontrolplane(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIstiocontrolplane(dAtA []byte, offset int, v uint64) int {
	offset -= sovIstiocontrolplane(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IstioControlPlaneSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovIstiocontrolplane(uint64(l))
	}
	if m.MeshConfig != nil {
		l = m.MeshConfig.Size()
		n += 1 + l + sovIstiocontrolplane(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IstioControlPlaneStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovIstiocontrolplane(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIstiocontrolplane(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIstiocontrolplane(x uint64) (n int) {
	return sovIstiocontrolplane(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IstioControlPlaneSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIstiocontrolplane
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
			return fmt.Errorf("proto: IstioControlPlaneSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioControlPlaneSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiocontrolplane
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
				return ErrInvalidLengthIstiocontrolplane
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeshConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiocontrolplane
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
				return ErrInvalidLengthIstiocontrolplane
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MeshConfig == nil {
				m.MeshConfig = &v1alpha1.MeshConfig{}
			}
			if err := m.MeshConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIstiocontrolplane(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIstiocontrolplane
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
func (m *IstioControlPlaneStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIstiocontrolplane
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
			return fmt.Errorf("proto: IstioControlPlaneStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioControlPlaneStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiocontrolplane
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
				return ErrInvalidLengthIstiocontrolplane
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIstiocontrolplane(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIstiocontrolplane
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
func skipIstiocontrolplane(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIstiocontrolplane
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
					return 0, ErrIntOverflowIstiocontrolplane
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
					return 0, ErrIntOverflowIstiocontrolplane
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
				return 0, ErrInvalidLengthIstiocontrolplane
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIstiocontrolplane
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIstiocontrolplane
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIstiocontrolplane        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIstiocontrolplane          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIstiocontrolplane = fmt.Errorf("proto: unexpected end of group")
)
