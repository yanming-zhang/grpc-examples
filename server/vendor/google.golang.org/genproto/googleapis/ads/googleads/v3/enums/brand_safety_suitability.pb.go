// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/brand_safety_suitability.proto

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

// 3-Tier brand safety suitability control.
type BrandSafetySuitabilityEnum_BrandSafetySuitability int32

const (
	// Not specified.
	BrandSafetySuitabilityEnum_UNSPECIFIED BrandSafetySuitabilityEnum_BrandSafetySuitability = 0
	// Used for return value only. Represents value unknown in this version.
	BrandSafetySuitabilityEnum_UNKNOWN BrandSafetySuitabilityEnum_BrandSafetySuitability = 1
	// This option lets you show ads across all inventory on YouTube and video
	// partners that meet our standards for monetization. This option may be an
	// appropriate choice for brands that want maximum access to the full
	// breadth of videos eligible for ads, including, for example, videos that
	// have strong profanity in the context of comedy or a documentary, or
	// excessive violence as featured in video games.
	BrandSafetySuitabilityEnum_EXPANDED_INVENTORY BrandSafetySuitabilityEnum_BrandSafetySuitability = 2
	// This option lets you show ads across a wide range of content that's
	// appropriate for most brands, such as popular music videos, documentaries,
	// and movie trailers. The content you can show ads on is based on YouTube's
	// advertiser-friendly content guidelines that take into account, for
	// example, the strength or frequency of profanity, or the appropriateness
	// of subject matter like sensitive events. Ads won't show, for example, on
	// content with repeated strong profanity, strong sexual content, or graphic
	// violence.
	BrandSafetySuitabilityEnum_STANDARD_INVENTORY BrandSafetySuitabilityEnum_BrandSafetySuitability = 3
	// This option lets you show ads on a reduced range of content that's
	// appropriate for brands with particularly strict guidelines around
	// inappropriate language and sexual suggestiveness; above and beyond what
	// YouTube's advertiser-friendly content guidelines address. The videos
	// accessible in this sensitive category meet heightened requirements,
	// especially for inappropriate language and sexual suggestiveness. For
	// example, your ads will be excluded from showing on some of YouTube's most
	// popular music videos and other pop culture content across YouTube and
	// Google video partners.
	BrandSafetySuitabilityEnum_LIMITED_INVENTORY BrandSafetySuitabilityEnum_BrandSafetySuitability = 4
)

var BrandSafetySuitabilityEnum_BrandSafetySuitability_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "EXPANDED_INVENTORY",
	3: "STANDARD_INVENTORY",
	4: "LIMITED_INVENTORY",
}

var BrandSafetySuitabilityEnum_BrandSafetySuitability_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"EXPANDED_INVENTORY": 2,
	"STANDARD_INVENTORY": 3,
	"LIMITED_INVENTORY":  4,
}

func (x BrandSafetySuitabilityEnum_BrandSafetySuitability) String() string {
	return proto.EnumName(BrandSafetySuitabilityEnum_BrandSafetySuitability_name, int32(x))
}

func (BrandSafetySuitabilityEnum_BrandSafetySuitability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_86b8ec423673ab58, []int{0, 0}
}

// Container for enum with 3-Tier brand safety suitability control.
type BrandSafetySuitabilityEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrandSafetySuitabilityEnum) Reset()         { *m = BrandSafetySuitabilityEnum{} }
func (m *BrandSafetySuitabilityEnum) String() string { return proto.CompactTextString(m) }
func (*BrandSafetySuitabilityEnum) ProtoMessage()    {}
func (*BrandSafetySuitabilityEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_86b8ec423673ab58, []int{0}
}

func (m *BrandSafetySuitabilityEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrandSafetySuitabilityEnum.Unmarshal(m, b)
}
func (m *BrandSafetySuitabilityEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrandSafetySuitabilityEnum.Marshal(b, m, deterministic)
}
func (m *BrandSafetySuitabilityEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrandSafetySuitabilityEnum.Merge(m, src)
}
func (m *BrandSafetySuitabilityEnum) XXX_Size() int {
	return xxx_messageInfo_BrandSafetySuitabilityEnum.Size(m)
}
func (m *BrandSafetySuitabilityEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BrandSafetySuitabilityEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BrandSafetySuitabilityEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.BrandSafetySuitabilityEnum_BrandSafetySuitability", BrandSafetySuitabilityEnum_BrandSafetySuitability_name, BrandSafetySuitabilityEnum_BrandSafetySuitability_value)
	proto.RegisterType((*BrandSafetySuitabilityEnum)(nil), "google.ads.googleads.v3.enums.BrandSafetySuitabilityEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/brand_safety_suitability.proto", fileDescriptor_86b8ec423673ab58)
}

var fileDescriptor_86b8ec423673ab58 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4e, 0xc2, 0x30,
	0x1c, 0x76, 0xc3, 0x68, 0x52, 0x0e, 0xe2, 0x12, 0x39, 0xa0, 0x1c, 0xe0, 0x01, 0xba, 0xc3, 0x6e,
	0xd5, 0x4b, 0xe7, 0x26, 0x59, 0xd4, 0xb2, 0x30, 0x98, 0x7f, 0xb2, 0x84, 0x14, 0x37, 0x97, 0x25,
	0xd0, 0x12, 0x3a, 0x48, 0x38, 0xf8, 0x24, 0xde, 0x3c, 0xfa, 0x28, 0x3e, 0x8a, 0x27, 0x1f, 0xc1,
	0xb4, 0x95, 0xa9, 0x09, 0x7a, 0x69, 0xbe, 0xfc, 0xbe, 0x3f, 0xe9, 0xf7, 0x81, 0xb3, 0x9c, 0xf3,
	0x7c, 0x9a, 0xd9, 0x34, 0x15, 0xb6, 0x86, 0x12, 0xad, 0x1c, 0x3b, 0x63, 0xcb, 0x99, 0xb0, 0x27,
	0x0b, 0xca, 0xd2, 0xb1, 0xa0, 0x8f, 0x59, 0xb9, 0x1e, 0x8b, 0x65, 0x51, 0xd2, 0x49, 0x31, 0x2d,
	0xca, 0x35, 0x9c, 0x2f, 0x78, 0xc9, 0xad, 0xb6, 0xb6, 0x40, 0x9a, 0x0a, 0x58, 0xb9, 0xe1, 0xca,
	0x81, 0xca, 0xdd, 0x3a, 0xd9, 0x84, 0xcf, 0x0b, 0x9b, 0x32, 0xc6, 0x4b, 0x5a, 0x16, 0x9c, 0x09,
	0x6d, 0xee, 0x3e, 0x1b, 0xa0, 0xe5, 0xca, 0xfc, 0x48, 0xc5, 0x47, 0xdf, 0xe9, 0x3e, 0x5b, 0xce,
	0xba, 0x4f, 0xa0, 0xb9, 0x9d, 0xb5, 0x0e, 0x40, 0x7d, 0x44, 0xa2, 0xd0, 0x3f, 0x0f, 0x2e, 0x02,
	0xdf, 0x6b, 0xec, 0x58, 0x75, 0xb0, 0x3f, 0x22, 0x97, 0xa4, 0x7f, 0x43, 0x1a, 0x86, 0xd5, 0x04,
	0x96, 0x7f, 0x1b, 0x62, 0xe2, 0xf9, 0xde, 0x38, 0x20, 0xb1, 0x4f, 0x86, 0xfd, 0xc1, 0x5d, 0xc3,
	0x94, 0xf7, 0x68, 0x88, 0x89, 0x87, 0x07, 0x3f, 0xef, 0x35, 0xeb, 0x08, 0x1c, 0x5e, 0x05, 0xd7,
	0xc1, 0xf0, 0x97, 0x7c, 0xd7, 0xfd, 0x30, 0x40, 0xe7, 0x81, 0xcf, 0xe0, 0xbf, 0x0d, 0xdd, 0xe3,
	0xed, 0x5f, 0x0c, 0x65, 0xc1, 0xd0, 0xb8, 0x77, 0xbf, 0xdc, 0x39, 0x9f, 0x52, 0x96, 0x43, 0xbe,
	0xc8, 0xed, 0x3c, 0x63, 0xaa, 0xfe, 0x66, 0xed, 0x79, 0x21, 0xfe, 0x18, 0xff, 0x54, 0xbd, 0x2f,
	0x66, 0xad, 0x87, 0xf1, 0xab, 0xd9, 0xee, 0xe9, 0x28, 0x9c, 0x0a, 0xa8, 0xa1, 0x44, 0xb1, 0x03,
	0xe5, 0x58, 0xe2, 0x6d, 0xc3, 0x27, 0x38, 0x15, 0x49, 0xc5, 0x27, 0xb1, 0x93, 0x28, 0xfe, 0xdd,
	0xec, 0xe8, 0x23, 0x42, 0x38, 0x15, 0x08, 0x55, 0x0a, 0x84, 0x62, 0x07, 0x21, 0xa5, 0x99, 0xec,
	0xa9, 0x8f, 0x39, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x8d, 0x6d, 0x8c, 0x14, 0x02, 0x00,
	0x00,
}
