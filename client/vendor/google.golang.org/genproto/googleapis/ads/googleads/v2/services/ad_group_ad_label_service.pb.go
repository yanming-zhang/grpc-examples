// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/ad_group_ad_label_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [AdGroupAdLabelService.GetAdGroupAdLabel][google.ads.googleads.v2.services.AdGroupAdLabelService.GetAdGroupAdLabel].
type GetAdGroupAdLabelRequest struct {
	// Required. The resource name of the ad group ad label to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupAdLabelRequest) Reset()         { *m = GetAdGroupAdLabelRequest{} }
func (m *GetAdGroupAdLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupAdLabelRequest) ProtoMessage()    {}
func (*GetAdGroupAdLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c813ff3d26d3adbb, []int{0}
}

func (m *GetAdGroupAdLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupAdLabelRequest.Unmarshal(m, b)
}
func (m *GetAdGroupAdLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupAdLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupAdLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupAdLabelRequest.Merge(m, src)
}
func (m *GetAdGroupAdLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupAdLabelRequest.Size(m)
}
func (m *GetAdGroupAdLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupAdLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupAdLabelRequest proto.InternalMessageInfo

func (m *GetAdGroupAdLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupAdLabelService.MutateAdGroupAdLabels][google.ads.googleads.v2.services.AdGroupAdLabelService.MutateAdGroupAdLabels].
type MutateAdGroupAdLabelsRequest struct {
	// Required. ID of the customer whose ad group ad labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on ad group ad labels.
	Operations []*AdGroupAdLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupAdLabelsRequest) Reset()         { *m = MutateAdGroupAdLabelsRequest{} }
func (m *MutateAdGroupAdLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdLabelsRequest) ProtoMessage()    {}
func (*MutateAdGroupAdLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c813ff3d26d3adbb, []int{1}
}

func (m *MutateAdGroupAdLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdLabelsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupAdLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdLabelsRequest.Merge(m, src)
}
func (m *MutateAdGroupAdLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdLabelsRequest.Size(m)
}
func (m *MutateAdGroupAdLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdLabelsRequest proto.InternalMessageInfo

func (m *MutateAdGroupAdLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupAdLabelsRequest) GetOperations() []*AdGroupAdLabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupAdLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupAdLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an ad group ad label.
type AdGroupAdLabelOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupAdLabelOperation_Create
	//	*AdGroupAdLabelOperation_Remove
	Operation            isAdGroupAdLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *AdGroupAdLabelOperation) Reset()         { *m = AdGroupAdLabelOperation{} }
func (m *AdGroupAdLabelOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdLabelOperation) ProtoMessage()    {}
func (*AdGroupAdLabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c813ff3d26d3adbb, []int{2}
}

func (m *AdGroupAdLabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdLabelOperation.Unmarshal(m, b)
}
func (m *AdGroupAdLabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdLabelOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupAdLabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdLabelOperation.Merge(m, src)
}
func (m *AdGroupAdLabelOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdLabelOperation.Size(m)
}
func (m *AdGroupAdLabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdLabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdLabelOperation proto.InternalMessageInfo

type isAdGroupAdLabelOperation_Operation interface {
	isAdGroupAdLabelOperation_Operation()
}

type AdGroupAdLabelOperation_Create struct {
	Create *resources.AdGroupAdLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupAdLabelOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupAdLabelOperation_Create) isAdGroupAdLabelOperation_Operation() {}

func (*AdGroupAdLabelOperation_Remove) isAdGroupAdLabelOperation_Operation() {}

func (m *AdGroupAdLabelOperation) GetOperation() isAdGroupAdLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupAdLabelOperation) GetCreate() *resources.AdGroupAdLabel {
	if x, ok := m.GetOperation().(*AdGroupAdLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupAdLabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupAdLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupAdLabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupAdLabelOperation_Create)(nil),
		(*AdGroupAdLabelOperation_Remove)(nil),
	}
}

// Response message for an ad group ad labels mutate.
type MutateAdGroupAdLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupAdLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MutateAdGroupAdLabelsResponse) Reset()         { *m = MutateAdGroupAdLabelsResponse{} }
func (m *MutateAdGroupAdLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdLabelsResponse) ProtoMessage()    {}
func (*MutateAdGroupAdLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c813ff3d26d3adbb, []int{3}
}

func (m *MutateAdGroupAdLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdLabelsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupAdLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdLabelsResponse.Merge(m, src)
}
func (m *MutateAdGroupAdLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdLabelsResponse.Size(m)
}
func (m *MutateAdGroupAdLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdLabelsResponse proto.InternalMessageInfo

func (m *MutateAdGroupAdLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupAdLabelsResponse) GetResults() []*MutateAdGroupAdLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for an ad group ad label mutate.
type MutateAdGroupAdLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupAdLabelResult) Reset()         { *m = MutateAdGroupAdLabelResult{} }
func (m *MutateAdGroupAdLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdLabelResult) ProtoMessage()    {}
func (*MutateAdGroupAdLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c813ff3d26d3adbb, []int{4}
}

func (m *MutateAdGroupAdLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdLabelResult.Unmarshal(m, b)
}
func (m *MutateAdGroupAdLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdLabelResult.Merge(m, src)
}
func (m *MutateAdGroupAdLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdLabelResult.Size(m)
}
func (m *MutateAdGroupAdLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdLabelResult proto.InternalMessageInfo

func (m *MutateAdGroupAdLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupAdLabelRequest)(nil), "google.ads.googleads.v2.services.GetAdGroupAdLabelRequest")
	proto.RegisterType((*MutateAdGroupAdLabelsRequest)(nil), "google.ads.googleads.v2.services.MutateAdGroupAdLabelsRequest")
	proto.RegisterType((*AdGroupAdLabelOperation)(nil), "google.ads.googleads.v2.services.AdGroupAdLabelOperation")
	proto.RegisterType((*MutateAdGroupAdLabelsResponse)(nil), "google.ads.googleads.v2.services.MutateAdGroupAdLabelsResponse")
	proto.RegisterType((*MutateAdGroupAdLabelResult)(nil), "google.ads.googleads.v2.services.MutateAdGroupAdLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/ad_group_ad_label_service.proto", fileDescriptor_c813ff3d26d3adbb)
}

var fileDescriptor_c813ff3d26d3adbb = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xbd, 0x6f, 0xd3, 0x4e,
	0x18, 0xfe, 0xd9, 0xf9, 0xa9, 0xd0, 0x4b, 0x0b, 0xe2, 0x50, 0xa9, 0x09, 0x45, 0x44, 0xa6, 0x52,
	0x4b, 0x84, 0x6c, 0xd5, 0x65, 0xa9, 0xa1, 0x80, 0x23, 0xd1, 0xb4, 0xe2, 0xa3, 0x55, 0x8a, 0x22,
	0x84, 0x22, 0xac, 0xab, 0x7d, 0x0d, 0x96, 0x1c, 0x9f, 0xb9, 0x3b, 0x47, 0xaa, 0xaa, 0x0e, 0xb0,
	0x33, 0xf1, 0x1f, 0x30, 0xf2, 0x4f, 0xb0, 0x57, 0x62, 0x62, 0xeb, 0xd4, 0x81, 0x89, 0xb1, 0x1b,
	0x4c, 0xc8, 0x1f, 0x97, 0xc4, 0x21, 0x56, 0x44, 0xb7, 0xd7, 0xf7, 0x3e, 0xf7, 0x3c, 0xef, 0xe7,
	0x19, 0x3c, 0xee, 0x10, 0xd2, 0xf1, 0xb1, 0x8e, 0x5c, 0xa6, 0xa7, 0x66, 0x6c, 0xf5, 0x0c, 0x9d,
	0x61, 0xda, 0xf3, 0x1c, 0xcc, 0x74, 0xe4, 0xda, 0x1d, 0x4a, 0xa2, 0xd0, 0x46, 0xae, 0xed, 0xa3,
	0x3d, 0xec, 0xdb, 0x99, 0x4b, 0x0b, 0x29, 0xe1, 0x04, 0x56, 0xd3, 0x6b, 0x1a, 0x72, 0x99, 0xd6,
	0x67, 0xd0, 0x7a, 0x86, 0x26, 0x18, 0x2a, 0x6b, 0x45, 0x1a, 0x14, 0x33, 0x12, 0xd1, 0xb1, 0x22,
	0x29, 0x79, 0x65, 0x41, 0x5c, 0x0d, 0x3d, 0x1d, 0x05, 0x01, 0xe1, 0x88, 0x7b, 0x24, 0x60, 0x99,
	0x77, 0x7e, 0xc8, 0xeb, 0xf8, 0x1e, 0x0e, 0x78, 0xe6, 0xb8, 0x35, 0xe4, 0xd8, 0xf7, 0xb0, 0xef,
	0xda, 0x7b, 0xf8, 0x2d, 0xea, 0x79, 0x84, 0x66, 0x80, 0xeb, 0x43, 0x00, 0x11, 0xc5, 0x08, 0x29,
	0x0d, 0x1d, 0x9d, 0x71, 0xc4, 0xa3, 0x4c, 0x4d, 0x0d, 0x81, 0xd2, 0xc0, 0xdc, 0x72, 0x1b, 0x71,
	0xa0, 0x96, 0xfb, 0x2c, 0x0e, 0xb3, 0x89, 0xdf, 0x45, 0x98, 0x71, 0xf8, 0x12, 0xcc, 0x0a, 0x1a,
	0x3b, 0x40, 0x5d, 0xac, 0x48, 0x55, 0x69, 0x79, 0xba, 0xae, 0x9f, 0x5a, 0xf2, 0x6f, 0xeb, 0x0e,
	0x58, 0x1a, 0x14, 0x26, 0xb3, 0x42, 0x8f, 0x69, 0x0e, 0xe9, 0xea, 0x23, 0x74, 0x33, 0x82, 0xe5,
	0x05, 0xea, 0x62, 0xf5, 0x4c, 0x02, 0x0b, 0xcf, 0x23, 0x8e, 0x38, 0xce, 0xc3, 0x98, 0x90, 0x5d,
	0x04, 0x65, 0x27, 0x62, 0x9c, 0x74, 0x31, 0xb5, 0x3d, 0x37, 0x13, 0x2d, 0x9d, 0x5a, 0x72, 0x13,
	0x88, 0xf3, 0x2d, 0x17, 0xbe, 0x01, 0x80, 0x84, 0x98, 0xa6, 0xa5, 0x53, 0xe4, 0x6a, 0x69, 0xb9,
	0x6c, 0xac, 0x69, 0x93, 0xda, 0xa6, 0xe5, 0x35, 0xb7, 0x05, 0x43, 0xc6, 0x3f, 0x60, 0x84, 0x4b,
	0xe0, 0x72, 0x88, 0x28, 0xf7, 0x90, 0x6f, 0xef, 0x23, 0xcf, 0x8f, 0x28, 0x56, 0x4a, 0x55, 0x69,
	0xf9, 0x62, 0xf3, 0x52, 0x76, 0xbc, 0x91, 0x9e, 0xc2, 0xdb, 0x60, 0xb6, 0x87, 0x7c, 0xcf, 0x45,
	0x1c, 0xdb, 0x24, 0xf0, 0x0f, 0x94, 0xff, 0x13, 0xd8, 0x8c, 0x38, 0xdc, 0x0e, 0xfc, 0x03, 0xf5,
	0xa3, 0x04, 0xe6, 0x0b, 0xa4, 0xe1, 0x53, 0x30, 0xe5, 0x50, 0x8c, 0x78, 0x5a, 0xdf, 0xb2, 0xb1,
	0x52, 0x98, 0x45, 0x7f, 0xb4, 0x46, 0xd2, 0xd8, 0xfc, 0xaf, 0x99, 0x51, 0x40, 0x05, 0x4c, 0x51,
	0xdc, 0x25, 0x3d, 0xac, 0xc8, 0x71, 0xdd, 0x62, 0x4f, 0xfa, 0x5d, 0x2f, 0x83, 0xe9, 0x7e, 0x7a,
	0xea, 0x57, 0x09, 0xdc, 0x2c, 0x68, 0x02, 0x0b, 0x49, 0xc0, 0x30, 0xdc, 0x00, 0x73, 0x23, 0xf9,
	0xdb, 0x98, 0x52, 0x42, 0x93, 0x2a, 0x94, 0x0d, 0x28, 0x82, 0xa4, 0xa1, 0xa3, 0xed, 0x26, 0x13,
	0xd5, 0xbc, 0x9a, 0xaf, 0xcc, 0x93, 0x18, 0x0e, 0x5b, 0xe0, 0x02, 0xc5, 0x2c, 0xf2, 0xb9, 0x68,
	0xd2, 0x83, 0xc9, 0x4d, 0x1a, 0x17, 0x59, 0x33, 0x21, 0x69, 0x0a, 0x32, 0xd5, 0x02, 0x95, 0x62,
	0x58, 0xdc, 0x94, 0x31, 0xa3, 0x9b, 0x9f, 0x44, 0xe3, 0x57, 0x09, 0xcc, 0xe5, 0x6f, 0xef, 0xa6,
	0x11, 0xc0, 0x6f, 0x12, 0xb8, 0xf2, 0xd7, 0x5a, 0x40, 0x73, 0x72, 0xe4, 0x45, 0xbb, 0x54, 0xf9,
	0xf7, 0xa6, 0xaa, 0x5b, 0x27, 0x56, 0x3e, 0x89, 0x0f, 0xdf, 0x7f, 0x7c, 0x92, 0x57, 0xe1, 0x4a,
	0xfc, 0xca, 0x1c, 0xe6, 0x3c, 0xeb, 0x62, 0x31, 0x98, 0x5e, 0xd3, 0x51, 0xbe, 0xa3, 0x7a, 0xed,
	0x08, 0x9e, 0x49, 0x60, 0x6e, 0x6c, 0xbb, 0xe1, 0xc3, 0xf3, 0x75, 0x43, 0x2c, 0x6b, 0xe5, 0xd1,
	0xb9, 0xef, 0xa7, 0x73, 0xa6, 0xbe, 0x3a, 0xb1, 0xae, 0x0d, 0xad, 0xfb, 0xdd, 0xc1, 0x0a, 0x26,
	0xe9, 0xae, 0xa9, 0xf7, 0xe2, 0x74, 0x07, 0xf9, 0x1d, 0x0e, 0x81, 0xd7, 0x6b, 0x47, 0xa3, 0xd9,
	0x9a, 0xdd, 0x44, 0xcd, 0x94, 0x6a, 0x95, 0x1b, 0xc7, 0x96, 0x52, 0xf4, 0x44, 0xd5, 0xdf, 0xcb,
	0x60, 0xd1, 0x21, 0xdd, 0x89, 0xd1, 0xd7, 0x2b, 0x63, 0x27, 0x64, 0x27, 0x7e, 0x3c, 0x77, 0xa4,
	0xd7, 0x9b, 0xd9, 0xfd, 0x0e, 0xf1, 0x51, 0xd0, 0xd1, 0x08, 0xed, 0xe8, 0x1d, 0x1c, 0x24, 0x4f,
	0xab, 0x3e, 0x50, 0x2c, 0xfe, 0x11, 0xdd, 0x17, 0xc6, 0x67, 0xb9, 0xd4, 0xb0, 0xac, 0x2f, 0x72,
	0xb5, 0x91, 0x12, 0x5a, 0x2e, 0xd3, 0x52, 0x33, 0xb6, 0x5a, 0x86, 0x96, 0x09, 0xb3, 0x63, 0x01,
	0x69, 0x5b, 0x2e, 0x6b, 0xf7, 0x21, 0xed, 0x96, 0xd1, 0x16, 0x90, 0x9f, 0xf2, 0x62, 0x7a, 0x6e,
	0x9a, 0x96, 0xcb, 0x4c, 0xb3, 0x0f, 0x32, 0xcd, 0x96, 0x61, 0x9a, 0x02, 0xb6, 0x37, 0x95, 0xc4,
	0xb9, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x55, 0xc9, 0x2b, 0x2b, 0x2f, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupAdLabelServiceClient is the client API for AdGroupAdLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupAdLabelServiceClient interface {
	// Returns the requested ad group ad label in full detail.
	GetAdGroupAdLabel(ctx context.Context, in *GetAdGroupAdLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupAdLabel, error)
	// Creates and removes ad group ad labels.
	// Operation statuses are returned.
	MutateAdGroupAdLabels(ctx context.Context, in *MutateAdGroupAdLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdLabelsResponse, error)
}

type adGroupAdLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupAdLabelServiceClient(cc grpc.ClientConnInterface) AdGroupAdLabelServiceClient {
	return &adGroupAdLabelServiceClient{cc}
}

func (c *adGroupAdLabelServiceClient) GetAdGroupAdLabel(ctx context.Context, in *GetAdGroupAdLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupAdLabel, error) {
	out := new(resources.AdGroupAdLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupAdLabelService/GetAdGroupAdLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupAdLabelServiceClient) MutateAdGroupAdLabels(ctx context.Context, in *MutateAdGroupAdLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdLabelsResponse, error) {
	out := new(MutateAdGroupAdLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupAdLabelService/MutateAdGroupAdLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAdLabelServiceServer is the server API for AdGroupAdLabelService service.
type AdGroupAdLabelServiceServer interface {
	// Returns the requested ad group ad label in full detail.
	GetAdGroupAdLabel(context.Context, *GetAdGroupAdLabelRequest) (*resources.AdGroupAdLabel, error)
	// Creates and removes ad group ad labels.
	// Operation statuses are returned.
	MutateAdGroupAdLabels(context.Context, *MutateAdGroupAdLabelsRequest) (*MutateAdGroupAdLabelsResponse, error)
}

// UnimplementedAdGroupAdLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupAdLabelServiceServer struct {
}

func (*UnimplementedAdGroupAdLabelServiceServer) GetAdGroupAdLabel(ctx context.Context, req *GetAdGroupAdLabelRequest) (*resources.AdGroupAdLabel, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupAdLabel not implemented")
}
func (*UnimplementedAdGroupAdLabelServiceServer) MutateAdGroupAdLabels(ctx context.Context, req *MutateAdGroupAdLabelsRequest) (*MutateAdGroupAdLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupAdLabels not implemented")
}

func RegisterAdGroupAdLabelServiceServer(s *grpc.Server, srv AdGroupAdLabelServiceServer) {
	s.RegisterService(&_AdGroupAdLabelService_serviceDesc, srv)
}

func _AdGroupAdLabelService_GetAdGroupAdLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAdLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdLabelServiceServer).GetAdGroupAdLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupAdLabelService/GetAdGroupAdLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdLabelServiceServer).GetAdGroupAdLabel(ctx, req.(*GetAdGroupAdLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupAdLabelService_MutateAdGroupAdLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAdLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdLabelServiceServer).MutateAdGroupAdLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupAdLabelService/MutateAdGroupAdLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdLabelServiceServer).MutateAdGroupAdLabels(ctx, req.(*MutateAdGroupAdLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupAdLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AdGroupAdLabelService",
	HandlerType: (*AdGroupAdLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAdLabel",
			Handler:    _AdGroupAdLabelService_GetAdGroupAdLabel_Handler,
		},
		{
			MethodName: "MutateAdGroupAdLabels",
			Handler:    _AdGroupAdLabelService_MutateAdGroupAdLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/ad_group_ad_label_service.proto",
}
