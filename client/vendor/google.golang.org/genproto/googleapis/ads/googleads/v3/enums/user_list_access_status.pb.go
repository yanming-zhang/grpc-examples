// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/user_list_access_status.proto

package enums

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

// Enum containing possible user list access statuses.
type UserListAccessStatusEnum_UserListAccessStatus int32

const (
	// Not specified.
	UserListAccessStatusEnum_UNSPECIFIED UserListAccessStatusEnum_UserListAccessStatus = 0
	// Used for return value only. Represents value unknown in this version.
	UserListAccessStatusEnum_UNKNOWN UserListAccessStatusEnum_UserListAccessStatus = 1
	// The access is enabled.
	UserListAccessStatusEnum_ENABLED UserListAccessStatusEnum_UserListAccessStatus = 2
	// The access is disabled.
	UserListAccessStatusEnum_DISABLED UserListAccessStatusEnum_UserListAccessStatus = 3
)

var UserListAccessStatusEnum_UserListAccessStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "DISABLED",
}

var UserListAccessStatusEnum_UserListAccessStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"DISABLED":    3,
}

func (x UserListAccessStatusEnum_UserListAccessStatus) String() string {
	return proto.EnumName(UserListAccessStatusEnum_UserListAccessStatus_name, int32(x))
}

func (UserListAccessStatusEnum_UserListAccessStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4c9f8c364837617, []int{0, 0}
}

// Indicates if this client still has access to the list.
type UserListAccessStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListAccessStatusEnum) Reset()         { *m = UserListAccessStatusEnum{} }
func (m *UserListAccessStatusEnum) String() string { return proto.CompactTextString(m) }
func (*UserListAccessStatusEnum) ProtoMessage()    {}
func (*UserListAccessStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4c9f8c364837617, []int{0}
}

func (m *UserListAccessStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListAccessStatusEnum.Unmarshal(m, b)
}
func (m *UserListAccessStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListAccessStatusEnum.Marshal(b, m, deterministic)
}
func (m *UserListAccessStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListAccessStatusEnum.Merge(m, src)
}
func (m *UserListAccessStatusEnum) XXX_Size() int {
	return xxx_messageInfo_UserListAccessStatusEnum.Size(m)
}
func (m *UserListAccessStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListAccessStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListAccessStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.UserListAccessStatusEnum_UserListAccessStatus", UserListAccessStatusEnum_UserListAccessStatus_name, UserListAccessStatusEnum_UserListAccessStatus_value)
	proto.RegisterType((*UserListAccessStatusEnum)(nil), "google.ads.googleads.v3.enums.UserListAccessStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/user_list_access_status.proto", fileDescriptor_c4c9f8c364837617)
}

var fileDescriptor_c4c9f8c364837617 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0xb5, 0x5b, 0x50, 0x49, 0x05, 0xcb, 0xe2, 0x41, 0xc5, 0x1e, 0xda, 0x0f, 0x48, 0x0e, 0xb9,
	0xa5, 0xa7, 0xac, 0x5d, 0x4b, 0xb1, 0x6c, 0x0b, 0xa5, 0x15, 0x64, 0xa1, 0xc4, 0x6e, 0x08, 0x8b,
	0x6d, 0x52, 0x76, 0xb2, 0xfd, 0x20, 0x8f, 0x7e, 0x8a, 0x9f, 0xe2, 0xc1, 0x6f, 0x90, 0x4d, 0xba,
	0x7b, 0xaa, 0x5e, 0xc2, 0x9b, 0xbc, 0x79, 0x6f, 0xde, 0x0c, 0x1a, 0x2a, 0x63, 0xd4, 0x56, 0x12,
	0x91, 0x01, 0xf1, 0xb0, 0x42, 0x07, 0x4a, 0xa4, 0x2e, 0x77, 0x40, 0x4a, 0x90, 0xc5, 0x7a, 0x9b,
	0x83, 0x5d, 0x8b, 0xcd, 0x46, 0x02, 0xac, 0xc1, 0x0a, 0x5b, 0x02, 0xde, 0x17, 0xc6, 0x9a, 0xb0,
	0xe7, 0x15, 0x58, 0x64, 0x80, 0x1b, 0x31, 0x3e, 0x50, 0xec, 0xc4, 0xf7, 0x0f, 0xb5, 0xf7, 0x3e,
	0x27, 0x42, 0x6b, 0x63, 0x85, 0xcd, 0x8d, 0x3e, 0x8a, 0x07, 0xef, 0xe8, 0x76, 0x09, 0xb2, 0x98,
	0xe6, 0x60, 0xb9, 0xf3, 0x5e, 0x38, 0xeb, 0x58, 0x97, 0xbb, 0xc1, 0x0c, 0xdd, 0x9c, 0xe2, 0xc2,
	0x6b, 0xd4, 0x59, 0x26, 0x8b, 0x79, 0xfc, 0x38, 0x79, 0x9a, 0xc4, 0xa3, 0xee, 0x59, 0xd8, 0x41,
	0x17, 0xcb, 0xe4, 0x39, 0x99, 0xbd, 0x24, 0xdd, 0x56, 0x55, 0xc4, 0x09, 0x8f, 0xa6, 0xf1, 0xa8,
	0x1b, 0x84, 0x57, 0xe8, 0x72, 0x34, 0x59, 0xf8, 0xaa, 0x1d, 0xfd, 0xb4, 0x50, 0x7f, 0x63, 0x76,
	0xf8, 0xdf, 0xc0, 0xd1, 0xdd, 0xa9, 0xa1, 0xf3, 0x2a, 0xed, 0xbc, 0xf5, 0x1a, 0x1d, 0xb5, 0xca,
	0x6c, 0x85, 0x56, 0xd8, 0x14, 0x8a, 0x28, 0xa9, 0xdd, 0x2e, 0xf5, 0xe5, 0xf6, 0x39, 0xfc, 0x71,
	0xc8, 0xa1, 0x7b, 0x3f, 0x82, 0xf6, 0x98, 0xf3, 0xcf, 0xa0, 0x37, 0xf6, 0x56, 0x3c, 0x03, 0xec,
	0x61, 0x85, 0x56, 0x14, 0x57, 0xcb, 0xc3, 0x57, 0xcd, 0xa7, 0x3c, 0x83, 0xb4, 0xe1, 0xd3, 0x15,
	0x4d, 0x1d, 0xff, 0x1d, 0xf4, 0xfd, 0x27, 0x63, 0x3c, 0x03, 0xc6, 0x9a, 0x0e, 0xc6, 0x56, 0x94,
	0x31, 0xd7, 0xf3, 0x76, 0xee, 0x82, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x20, 0xb6,
	0xbb, 0xe0, 0x01, 0x00, 0x00,
}
