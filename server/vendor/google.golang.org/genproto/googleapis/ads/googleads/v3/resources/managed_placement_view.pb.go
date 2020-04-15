// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/managed_placement_view.proto

package resources

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

// A managed placement view.
type ManagedPlacementView struct {
	// Output only. The resource name of the Managed Placement view.
	// Managed placement view resource names have the form:
	//
	// `customers/{customer_id}/managedPlacementViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManagedPlacementView) Reset()         { *m = ManagedPlacementView{} }
func (m *ManagedPlacementView) String() string { return proto.CompactTextString(m) }
func (*ManagedPlacementView) ProtoMessage()    {}
func (*ManagedPlacementView) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3083284e69e9877, []int{0}
}

func (m *ManagedPlacementView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagedPlacementView.Unmarshal(m, b)
}
func (m *ManagedPlacementView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagedPlacementView.Marshal(b, m, deterministic)
}
func (m *ManagedPlacementView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagedPlacementView.Merge(m, src)
}
func (m *ManagedPlacementView) XXX_Size() int {
	return xxx_messageInfo_ManagedPlacementView.Size(m)
}
func (m *ManagedPlacementView) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagedPlacementView.DiscardUnknown(m)
}

var xxx_messageInfo_ManagedPlacementView proto.InternalMessageInfo

func (m *ManagedPlacementView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ManagedPlacementView)(nil), "google.ads.googleads.v3.resources.ManagedPlacementView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/managed_placement_view.proto", fileDescriptor_e3083284e69e9877)
}

var fileDescriptor_e3083284e69e9877 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x6a, 0xeb, 0x30,
	0x18, 0x85, 0xb1, 0x03, 0x17, 0xae, 0xb9, 0x77, 0x09, 0x77, 0xb8, 0x09, 0x85, 0x36, 0x85, 0x40,
	0x97, 0x4a, 0x83, 0xe9, 0xa2, 0x42, 0x41, 0xe9, 0x10, 0x28, 0xb4, 0x84, 0x0c, 0x1e, 0x82, 0xc1,
	0x28, 0xf6, 0x5f, 0x57, 0x60, 0x49, 0xc6, 0x72, 0x9c, 0x21, 0x64, 0xed, 0x83, 0x74, 0xec, 0xa3,
	0xf4, 0x29, 0x3a, 0xa7, 0x6f, 0xd0, 0xa9, 0x24, 0x8a, 0x94, 0x14, 0x42, 0x4b, 0xb7, 0x03, 0xff,
	0x77, 0x8e, 0x8f, 0x0f, 0x0a, 0xae, 0x72, 0xa5, 0xf2, 0x02, 0x30, 0xcb, 0x34, 0x36, 0x72, 0xad,
	0x9a, 0x10, 0x57, 0xa0, 0xd5, 0xac, 0x4a, 0x41, 0x63, 0xc1, 0x24, 0xcb, 0x21, 0x4b, 0xca, 0x82,
	0xa5, 0x20, 0x40, 0xd6, 0x49, 0xc3, 0x61, 0x8e, 0xca, 0x4a, 0xd5, 0xaa, 0xdd, 0x33, 0x26, 0xc4,
	0x32, 0x8d, 0x9c, 0x1f, 0x35, 0x21, 0x72, 0xfe, 0xee, 0xb1, 0xfd, 0x44, 0xc9, 0xf1, 0x3d, 0x87,
	0x22, 0x4b, 0xa6, 0xf0, 0xc0, 0x1a, 0xae, 0x2a, 0x93, 0xd1, 0xed, 0xec, 0x01, 0xd6, 0xb6, 0x3d,
	0x1d, 0xed, 0x9d, 0x98, 0x94, 0xaa, 0x66, 0x35, 0x57, 0x52, 0x9b, 0xeb, 0xe9, 0x9b, 0x17, 0xfc,
	0xbb, 0x35, 0xed, 0x46, 0xb6, 0x5c, 0xc4, 0x61, 0xde, 0x9e, 0x04, 0x7f, 0x6d, 0x50, 0x22, 0x99,
	0x80, 0xff, 0xde, 0x89, 0x77, 0xf6, 0x7b, 0x70, 0xf1, 0x4a, 0x5b, 0xef, 0x14, 0x07, 0xe7, 0xbb,
	0xa6, 0x5b, 0x55, 0x72, 0x8d, 0x52, 0x25, 0xf0, 0xa1, 0xb4, 0xf1, 0x1f, 0x9b, 0x75, 0xc7, 0x04,
	0x90, 0xf9, 0x8a, 0xd6, 0x3f, 0x4c, 0x68, 0x5f, 0xa7, 0x33, 0x5d, 0x2b, 0x01, 0x95, 0xc6, 0x0b,
	0x2b, 0x97, 0x76, 0xd8, 0x4f, 0xa8, 0xc6, 0x8b, 0xc3, 0x7b, 0x2f, 0x07, 0x8f, 0x7e, 0xd0, 0x4f,
	0x95, 0x40, 0xdf, 0x2e, 0x3e, 0xe8, 0x1c, 0x2a, 0x31, 0x5a, 0x4f, 0x36, 0xf2, 0x26, 0x37, 0x5b,
	0x7f, 0xae, 0x0a, 0x26, 0x73, 0xa4, 0xaa, 0x1c, 0xe7, 0x20, 0x37, 0x83, 0xe2, 0xdd, 0xbf, 0x7c,
	0xf1, 0x20, 0x2e, 0x9d, 0x7a, 0xf2, 0x5b, 0x43, 0x4a, 0x9f, 0xfd, 0xde, 0xd0, 0x44, 0xd2, 0x4c,
	0x23, 0x23, 0xd7, 0x2a, 0x0a, 0xd1, 0xd8, 0x92, 0x2f, 0x96, 0x89, 0x69, 0xa6, 0x63, 0xc7, 0xc4,
	0x51, 0x18, 0x3b, 0x66, 0xe5, 0xf7, 0xcd, 0x81, 0x10, 0x9a, 0x69, 0x42, 0x1c, 0x45, 0x48, 0x14,
	0x12, 0xe2, 0xb8, 0xe9, 0xaf, 0x4d, 0xd9, 0xf0, 0x23, 0x00, 0x00, 0xff, 0xff, 0x87, 0x95, 0xc2,
	0xf1, 0xbc, 0x02, 0x00, 0x00,
}
