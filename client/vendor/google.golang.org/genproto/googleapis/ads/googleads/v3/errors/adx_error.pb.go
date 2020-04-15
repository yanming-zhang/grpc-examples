// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/adx_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible adx errors.
type AdxErrorEnum_AdxError int32

const (
	// Enum unspecified.
	AdxErrorEnum_UNSPECIFIED AdxErrorEnum_AdxError = 0
	// The received error code is not known in this version.
	AdxErrorEnum_UNKNOWN AdxErrorEnum_AdxError = 1
	// Attempt to use non-AdX feature by AdX customer.
	AdxErrorEnum_UNSUPPORTED_FEATURE AdxErrorEnum_AdxError = 2
)

var AdxErrorEnum_AdxError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "UNSUPPORTED_FEATURE",
}

var AdxErrorEnum_AdxError_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"UNSUPPORTED_FEATURE": 2,
}

func (x AdxErrorEnum_AdxError) String() string {
	return proto.EnumName(AdxErrorEnum_AdxError_name, int32(x))
}

func (AdxErrorEnum_AdxError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8c9a39d67b90bc57, []int{0, 0}
}

// Container for enum describing possible adx errors.
type AdxErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdxErrorEnum) Reset()         { *m = AdxErrorEnum{} }
func (m *AdxErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdxErrorEnum) ProtoMessage()    {}
func (*AdxErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c9a39d67b90bc57, []int{0}
}

func (m *AdxErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdxErrorEnum.Unmarshal(m, b)
}
func (m *AdxErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdxErrorEnum.Marshal(b, m, deterministic)
}
func (m *AdxErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdxErrorEnum.Merge(m, src)
}
func (m *AdxErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdxErrorEnum.Size(m)
}
func (m *AdxErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdxErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdxErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.AdxErrorEnum_AdxError", AdxErrorEnum_AdxError_name, AdxErrorEnum_AdxError_value)
	proto.RegisterType((*AdxErrorEnum)(nil), "google.ads.googleads.v3.errors.AdxErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/adx_error.proto", fileDescriptor_8c9a39d67b90bc57)
}

var fileDescriptor_8c9a39d67b90bc57 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x6d, 0x05, 0x95, 0x4c, 0xb1, 0xd4, 0x83, 0x20, 0xb2, 0x43, 0x1f, 0x20, 0x39, 0xe4,
	0x16, 0x4f, 0x99, 0xcd, 0xc6, 0x10, 0xba, 0xba, 0xad, 0x15, 0xa4, 0x30, 0xa2, 0x29, 0xa1, 0xb0,
	0x25, 0xa5, 0xa9, 0x63, 0xcf, 0xe3, 0xd1, 0x47, 0xf1, 0x49, 0xc4, 0xa7, 0x90, 0x36, 0xa6, 0x37,
	0x3d, 0xe5, 0x97, 0x8f, 0xdf, 0xf7, 0x7d, 0x7f, 0x3e, 0x00, 0xa5, 0xd6, 0x72, 0x5b, 0x22, 0x2e,
	0x0c, 0xb2, 0xd8, 0xd1, 0x1e, 0xa3, 0xb2, 0x69, 0x74, 0x63, 0x10, 0x17, 0x87, 0x4d, 0x8f, 0xb0,
	0x6e, 0x74, 0xab, 0xc3, 0xb1, 0x95, 0x20, 0x17, 0x06, 0x0e, 0x3e, 0xdc, 0x63, 0x68, 0xfd, 0x9b,
	0x5b, 0x37, 0xaf, 0xae, 0x10, 0x57, 0x4a, 0xb7, 0xbc, 0xad, 0xb4, 0x32, 0xb6, 0x3b, 0x7a, 0x04,
	0xe7, 0x54, 0x1c, 0x58, 0xa7, 0x32, 0xf5, 0xb6, 0x8b, 0x28, 0x38, 0x73, 0xff, 0xf0, 0x12, 0x8c,
	0xb2, 0x64, 0x95, 0xb2, 0xfb, 0xf9, 0x74, 0xce, 0xe2, 0xe0, 0x28, 0x1c, 0x81, 0xd3, 0x2c, 0x79,
	0x48, 0x16, 0x4f, 0x49, 0xe0, 0x85, 0xd7, 0xe0, 0x2a, 0x4b, 0x56, 0x59, 0x9a, 0x2e, 0x96, 0x6b,
	0x16, 0x6f, 0xa6, 0x8c, 0xae, 0xb3, 0x25, 0x0b, 0xfc, 0xc9, 0x97, 0x07, 0xa2, 0x57, 0xbd, 0x83,
	0xff, 0xe7, 0x9a, 0x5c, 0xb8, 0x3d, 0x69, 0x17, 0x24, 0xf5, 0x9e, 0xe3, 0xdf, 0x06, 0xa9, 0xb7,
	0x5c, 0x49, 0xa8, 0x1b, 0x89, 0x64, 0xa9, 0xfa, 0x98, 0xee, 0x10, 0x75, 0x65, 0xfe, 0xba, 0xcb,
	0x9d, 0x7d, 0xde, 0xfd, 0xe3, 0x19, 0xa5, 0x1f, 0xfe, 0x78, 0x66, 0x87, 0x51, 0x61, 0xa0, 0xc5,
	0x8e, 0x72, 0x0c, 0xfb, 0x95, 0xe6, 0xd3, 0x09, 0x05, 0x15, 0xa6, 0x18, 0x84, 0x22, 0xc7, 0x85,
	0x15, 0xbe, 0xfd, 0xc8, 0x56, 0x09, 0xa1, 0xc2, 0x10, 0x32, 0x28, 0x84, 0xe4, 0x98, 0x10, 0x2b,
	0xbd, 0x9c, 0xf4, 0xe9, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x7e, 0x51, 0x7c, 0xb4,
	0x01, 0x00, 0x00,
}
