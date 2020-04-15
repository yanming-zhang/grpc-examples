// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/source.proto

package securitycenter

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

// Cloud Security Command Center's (Cloud SCC) finding source. A finding source
// is an entity or a mechanism that can produce a finding. A source is like a
// container of findings that come from the same scanner, logger, monitor, and
// other tools.
type Source struct {
	// The relative resource name of this source. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/{organization_id}/sources/{source_id}"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The source's display name.
	// A source's display name must be unique amongst its siblings, for example,
	// two sources with the same parent can't share the same display name.
	// The display name must have a length between 1 and 64 characters
	// (inclusive).
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The description of the source (max of 1024 characters).
	// Example:
	// "Web Security Scanner is a web security scanner for common
	// vulnerabilities in App Engine applications. It can automatically
	// scan and detect four common vulnerabilities, including cross-site-scripting
	// (XSS), Flash injection, mixed content (HTTP in HTTPS), and
	// outdated or insecure libraries."
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f5cbeb6a51b224, []int{0}
}

func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (m *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(m, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Source) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Source) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Source)(nil), "google.cloud.securitycenter.v1.Source")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/source.proto", fileDescriptor_43f5cbeb6a51b224)
}

var fileDescriptor_43f5cbeb6a51b224 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0xe9, 0x94, 0x81, 0x99, 0x57, 0xbd, 0x9a, 0x43, 0xc6, 0x36, 0xbc, 0x10, 0xc4, 0x84,
	0xe2, 0x5d, 0xbc, 0xd2, 0x5d, 0x78, 0x23, 0x22, 0x0e, 0x87, 0x48, 0x41, 0x62, 0x1a, 0x42, 0xa0,
	0xcd, 0x29, 0x49, 0x3b, 0x98, 0x63, 0x2f, 0xe4, 0x2b, 0xf8, 0x06, 0x3e, 0x82, 0x8f, 0xb0, 0xa7,
	0x90, 0x25, 0x91, 0xb5, 0x82, 0xde, 0x9d, 0xfe, 0xff, 0xd7, 0x73, 0x7e, 0xfe, 0xa0, 0x33, 0x09,
	0x20, 0x73, 0x41, 0x78, 0x0e, 0x75, 0x46, 0xac, 0xe0, 0xb5, 0x51, 0xd5, 0x92, 0x0b, 0x5d, 0x09,
	0x43, 0x16, 0x09, 0xb1, 0x50, 0x1b, 0x2e, 0x70, 0x69, 0xa0, 0x82, 0x78, 0xe8, 0x61, 0xec, 0x60,
	0xdc, 0x86, 0xf1, 0x22, 0x19, 0x1c, 0x87, 0x65, 0xac, 0x54, 0x84, 0x69, 0x0d, 0x15, 0xab, 0x14,
	0x68, 0xeb, 0xff, 0x1e, 0x1c, 0x35, 0x5c, 0x23, 0x9a, 0x8b, 0x27, 0x1f, 0x11, 0xea, 0xce, 0x9c,
	0x10, 0xc7, 0x68, 0x5f, 0xb3, 0x42, 0xf4, 0xa3, 0x51, 0x74, 0x7a, 0xf0, 0xe0, 0xe6, 0x78, 0x8c,
	0x0e, 0x33, 0x65, 0xcb, 0x9c, 0x2d, 0x5f, 0x9c, 0xd7, 0x71, 0x5e, 0x2f, 0x68, 0x77, 0x5b, 0x64,
	0x84, 0x7a, 0x99, 0xb0, 0xdc, 0xa8, 0x72, 0x7b, 0xb2, 0xbf, 0x17, 0x88, 0x9d, 0x44, 0x9f, 0x36,
	0x57, 0x8f, 0xe8, 0xe4, 0x57, 0x68, 0x1f, 0x89, 0x95, 0xca, 0x62, 0x0e, 0x05, 0x09, 0x19, 0xce,
	0xc1, 0x48, 0xa6, 0xd5, 0x9b, 0x8f, 0x4f, 0x56, 0xcd, 0xcf, 0x75, 0xa8, 0xc4, 0x92, 0x95, 0x1f,
	0xd6, 0xd7, 0x5f, 0x11, 0x9a, 0x70, 0x28, 0xf0, 0xff, 0xed, 0xdc, 0x47, 0xcf, 0xb7, 0x81, 0x90,
	0x90, 0x33, 0x2d, 0x31, 0x18, 0x49, 0xa4, 0xd0, 0xae, 0x02, 0xb2, 0x4b, 0xf2, 0xd7, 0x5b, 0x5c,
	0xb6, 0x95, 0xf7, 0xce, 0xf0, 0xc6, 0xaf, 0x9b, 0xba, 0x83, 0xb3, 0xe0, 0x4e, 0xfd, 0xc1, 0x79,
	0xf2, 0xf9, 0x03, 0xa4, 0x0e, 0x48, 0xdb, 0x40, 0x3a, 0x4f, 0x36, 0x9d, 0xb1, 0x07, 0x28, 0x75,
	0x04, 0xa5, 0x6d, 0x84, 0xd2, 0x79, 0xf2, 0xda, 0x75, 0xf1, 0x2e, 0xbe, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0x2b, 0x23, 0xf9, 0x29, 0x02, 0x00, 0x00,
}
