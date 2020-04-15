// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/asset/v1p1beta1/asset_service.proto

package asset

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Search all resources request.
type SearchAllResourcesRequest struct {
	// Required. The relative name of an asset. The search is limited to the resources
	// within the `scope`. The allowed value must be:
	// * Organization number (such as "organizations/123")
	// * Folder number(such as "folders/1234")
	// * Project number (such as "projects/12345")
	Scope string `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
	// Optional. The query statement.
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// Optional. A list of asset types that this request searches for. If empty, it will
	// search all the supported asset types.
	AssetTypes []string `protobuf:"bytes,3,rep,name=asset_types,json=assetTypes,proto3" json:"asset_types,omitempty"`
	// Optional. The page size for search result pagination. Page size is capped at 500 even
	// if a larger value is given. If set to zero, server will pick an appropriate
	// default. Returned results may be fewer than requested. When this happens,
	// there could be more results as long as `next_page_token` is returned.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the preceding call
	// to this method.  `page_token` must be the value of `next_page_token` from
	// the previous response. The values of all other method parameters, must be
	// identical to those in the previous call.
	PageToken            string   `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchAllResourcesRequest) Reset()         { *m = SearchAllResourcesRequest{} }
func (m *SearchAllResourcesRequest) String() string { return proto.CompactTextString(m) }
func (*SearchAllResourcesRequest) ProtoMessage()    {}
func (*SearchAllResourcesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ebba5b73acb095, []int{0}
}

func (m *SearchAllResourcesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAllResourcesRequest.Unmarshal(m, b)
}
func (m *SearchAllResourcesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAllResourcesRequest.Marshal(b, m, deterministic)
}
func (m *SearchAllResourcesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAllResourcesRequest.Merge(m, src)
}
func (m *SearchAllResourcesRequest) XXX_Size() int {
	return xxx_messageInfo_SearchAllResourcesRequest.Size(m)
}
func (m *SearchAllResourcesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAllResourcesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAllResourcesRequest proto.InternalMessageInfo

func (m *SearchAllResourcesRequest) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *SearchAllResourcesRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchAllResourcesRequest) GetAssetTypes() []string {
	if m != nil {
		return m.AssetTypes
	}
	return nil
}

func (m *SearchAllResourcesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchAllResourcesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Search all resources response.
type SearchAllResourcesResponse struct {
	// A list of resource that match the search query.
	Results []*StandardResourceMetadata `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// If there are more results than those appearing in this response, then
	// `next_page_token` is included.  To get the next set of results, call this
	// method again using the value of `next_page_token` as `page_token`.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchAllResourcesResponse) Reset()         { *m = SearchAllResourcesResponse{} }
func (m *SearchAllResourcesResponse) String() string { return proto.CompactTextString(m) }
func (*SearchAllResourcesResponse) ProtoMessage()    {}
func (*SearchAllResourcesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ebba5b73acb095, []int{1}
}

func (m *SearchAllResourcesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAllResourcesResponse.Unmarshal(m, b)
}
func (m *SearchAllResourcesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAllResourcesResponse.Marshal(b, m, deterministic)
}
func (m *SearchAllResourcesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAllResourcesResponse.Merge(m, src)
}
func (m *SearchAllResourcesResponse) XXX_Size() int {
	return xxx_messageInfo_SearchAllResourcesResponse.Size(m)
}
func (m *SearchAllResourcesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAllResourcesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAllResourcesResponse proto.InternalMessageInfo

func (m *SearchAllResourcesResponse) GetResults() []*StandardResourceMetadata {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *SearchAllResourcesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Search all IAM policies request.
type SearchAllIamPoliciesRequest struct {
	// Required. The relative name of an asset. The search is limited to the resources
	// within the `scope`. The allowed value must be:
	// * Organization number (such as "organizations/123")
	// * Folder number(such as "folders/1234")
	// * Project number (such as "projects/12345")
	Scope string `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
	// Optional. The query statement.
	// Examples:
	// * "policy:myuser@mydomain.com"
	// * "policy:(myuser@mydomain.com viewer)"
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// Optional. The page size for search result pagination. Page size is capped at 500 even
	// if a larger value is given. If set to zero, server will pick an appropriate
	// default. Returned results may be fewer than requested. When this happens,
	// there could be more results as long as `next_page_token` is returned.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. If present, retrieve the next batch of results from the preceding call to
	// this method. `page_token` must be the value of `next_page_token` from the
	// previous response. The values of all other method parameters must be
	// identical to those in the previous call.
	PageToken            string   `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchAllIamPoliciesRequest) Reset()         { *m = SearchAllIamPoliciesRequest{} }
func (m *SearchAllIamPoliciesRequest) String() string { return proto.CompactTextString(m) }
func (*SearchAllIamPoliciesRequest) ProtoMessage()    {}
func (*SearchAllIamPoliciesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ebba5b73acb095, []int{2}
}

func (m *SearchAllIamPoliciesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAllIamPoliciesRequest.Unmarshal(m, b)
}
func (m *SearchAllIamPoliciesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAllIamPoliciesRequest.Marshal(b, m, deterministic)
}
func (m *SearchAllIamPoliciesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAllIamPoliciesRequest.Merge(m, src)
}
func (m *SearchAllIamPoliciesRequest) XXX_Size() int {
	return xxx_messageInfo_SearchAllIamPoliciesRequest.Size(m)
}
func (m *SearchAllIamPoliciesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAllIamPoliciesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAllIamPoliciesRequest proto.InternalMessageInfo

func (m *SearchAllIamPoliciesRequest) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *SearchAllIamPoliciesRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchAllIamPoliciesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchAllIamPoliciesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Search all IAM policies response.
type SearchAllIamPoliciesResponse struct {
	// A list of IamPolicy that match the search query. Related information such
	// as the associated resource is returned along with the policy.
	Results []*IamPolicySearchResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// Set if there are more results than those appearing in this response; to get
	// the next set of results, call this method again, using this value as the
	// `page_token`.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchAllIamPoliciesResponse) Reset()         { *m = SearchAllIamPoliciesResponse{} }
func (m *SearchAllIamPoliciesResponse) String() string { return proto.CompactTextString(m) }
func (*SearchAllIamPoliciesResponse) ProtoMessage()    {}
func (*SearchAllIamPoliciesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ebba5b73acb095, []int{3}
}

func (m *SearchAllIamPoliciesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAllIamPoliciesResponse.Unmarshal(m, b)
}
func (m *SearchAllIamPoliciesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAllIamPoliciesResponse.Marshal(b, m, deterministic)
}
func (m *SearchAllIamPoliciesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAllIamPoliciesResponse.Merge(m, src)
}
func (m *SearchAllIamPoliciesResponse) XXX_Size() int {
	return xxx_messageInfo_SearchAllIamPoliciesResponse.Size(m)
}
func (m *SearchAllIamPoliciesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAllIamPoliciesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAllIamPoliciesResponse proto.InternalMessageInfo

func (m *SearchAllIamPoliciesResponse) GetResults() []*IamPolicySearchResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *SearchAllIamPoliciesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchAllResourcesRequest)(nil), "google.cloud.asset.v1p1beta1.SearchAllResourcesRequest")
	proto.RegisterType((*SearchAllResourcesResponse)(nil), "google.cloud.asset.v1p1beta1.SearchAllResourcesResponse")
	proto.RegisterType((*SearchAllIamPoliciesRequest)(nil), "google.cloud.asset.v1p1beta1.SearchAllIamPoliciesRequest")
	proto.RegisterType((*SearchAllIamPoliciesResponse)(nil), "google.cloud.asset.v1p1beta1.SearchAllIamPoliciesResponse")
}

func init() {
	proto.RegisterFile("google/cloud/asset/v1p1beta1/asset_service.proto", fileDescriptor_a3ebba5b73acb095)
}

var fileDescriptor_a3ebba5b73acb095 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x67, 0x76, 0x5b, 0xb5, 0x53, 0x45, 0x1c, 0x84, 0x6e, 0xd7, 0x82, 0x61, 0x11, 0xa9, 0x65,
	0xcd, 0xb8, 0x2d, 0xf8, 0xa7, 0xe2, 0x21, 0xeb, 0x41, 0x04, 0x0b, 0x21, 0x5b, 0x7a, 0x90, 0xc2,
	0x32, 0xcd, 0xbe, 0x66, 0x07, 0xb3, 0x99, 0x34, 0x33, 0x69, 0x6d, 0xc5, 0x8b, 0x37, 0xcf, 0xa2,
	0xe0, 0x57, 0xf0, 0x13, 0x78, 0xf1, 0x0b, 0xf4, 0xa6, 0xe2, 0xa5, 0x27, 0x0f, 0x7e, 0x10, 0xc9,
	0x4c, 0xd2, 0x8d, 0xed, 0xb6, 0x6b, 0xf5, 0x98, 0xf7, 0x7b, 0xbf, 0xdf, 0x7b, 0xef, 0xf7, 0x26,
	0x0f, 0xdf, 0x09, 0x84, 0x08, 0x42, 0xa0, 0x7e, 0x28, 0xd2, 0x1e, 0x65, 0x52, 0x82, 0xa2, 0xdb,
	0xad, 0xb8, 0xb5, 0x01, 0x8a, 0xb5, 0xcc, 0x77, 0x57, 0x42, 0xb2, 0xcd, 0x7d, 0xb0, 0xe3, 0x44,
	0x28, 0x41, 0xe6, 0x0c, 0xc3, 0xd6, 0x0c, 0x5b, 0x67, 0xd8, 0x87, 0x8c, 0x7a, 0x8e, 0x52, 0x16,
	0x73, 0xca, 0xa2, 0x48, 0x28, 0xa6, 0xb8, 0x88, 0xa4, 0xe1, 0xd6, 0x67, 0x4a, 0xa8, 0x1f, 0x72,
	0x88, 0x54, 0x0e, 0x5c, 0x2f, 0x01, 0x9b, 0x1c, 0xc2, 0x5e, 0x77, 0x03, 0xfa, 0x6c, 0x9b, 0x8b,
	0x24, 0x4f, 0xb8, 0x35, 0xbe, 0xcf, 0xbc, 0x48, 0xe3, 0x0b, 0xc2, 0xb3, 0x1d, 0x60, 0x89, 0xdf,
	0x77, 0xc2, 0xd0, 0x03, 0x29, 0xd2, 0xc4, 0x07, 0xe9, 0xc1, 0x56, 0x0a, 0x52, 0x91, 0x59, 0x3c,
	0x29, 0x7d, 0x11, 0x43, 0x0d, 0x59, 0x68, 0x7e, 0xaa, 0x5d, 0xfd, 0xe9, 0x54, 0x3c, 0x13, 0xc9,
	0xa0, 0xad, 0x14, 0x92, 0xdd, 0x5a, 0xa5, 0x80, 0x90, 0x67, 0x22, 0xe4, 0x06, 0x9e, 0x36, 0x5e,
	0xa8, 0xdd, 0x18, 0x64, 0xad, 0x6a, 0x55, 0x8b, 0x04, 0xac, 0xe3, 0xab, 0x59, 0x98, 0x58, 0x78,
	0x2a, 0x66, 0x01, 0x74, 0x25, 0xdf, 0x83, 0xda, 0x84, 0x85, 0xe6, 0x27, 0x4d, 0xce, 0x85, 0x2c,
	0xda, 0xe1, 0x7b, 0x40, 0x1a, 0x18, 0xeb, 0x0c, 0x25, 0x5e, 0x40, 0x54, 0x9b, 0x1c, 0xd6, 0xd1,
	0xc4, 0xd5, 0x2c, 0xda, 0xf8, 0x80, 0x70, 0x7d, 0x54, 0xff, 0x32, 0x16, 0x91, 0x04, 0xe2, 0xe2,
	0xf3, 0x09, 0xc8, 0x34, 0x54, 0xb2, 0x86, 0xac, 0xea, 0xfc, 0xf4, 0xe2, 0x5d, 0xfb, 0xb4, 0x8d,
	0xd8, 0x1d, 0xc5, 0xa2, 0x1e, 0x4b, 0x7a, 0x85, 0xd2, 0x0a, 0x28, 0xd6, 0x63, 0x8a, 0x79, 0x85,
	0x0c, 0xb9, 0x89, 0x2f, 0x47, 0xf0, 0x52, 0x75, 0x4b, 0x9d, 0x69, 0x07, 0xbc, 0x4b, 0x59, 0xd8,
	0x3d, 0x6c, 0xec, 0x23, 0xc2, 0xd7, 0x0e, 0x1b, 0x7b, 0xca, 0x06, 0xae, 0x08, 0xb9, 0xcf, 0xff,
	0xd7, 0xda, 0x3f, 0x4c, 0xab, 0x8e, 0x37, 0x6d, 0x62, 0xa4, 0x69, 0xef, 0x11, 0x9e, 0x1b, 0xdd,
	0x5b, 0x6e, 0xdb, 0xca, 0x51, 0xdb, 0x96, 0x4e, 0xb7, 0xad, 0xd0, 0xd8, 0x35, 0xaa, 0x9e, 0xe6,
	0x9e, 0xd9, 0xb3, 0xc5, 0xb7, 0x13, 0xf8, 0xa2, 0x93, 0x49, 0x77, 0xcc, 0x4f, 0x44, 0x7e, 0x20,
	0x4c, 0x8e, 0x6f, 0x97, 0xdc, 0x1b, 0xb3, 0xc4, 0x93, 0xde, 0x73, 0xfd, 0xfe, 0xd9, 0x89, 0xc6,
	0x91, 0xc6, 0xb3, 0x03, 0x67, 0x46, 0x6f, 0xa7, 0xa9, 0x17, 0xd1, 0x2c, 0x3d, 0xf0, 0x37, 0xdf,
	0x7f, 0xbd, 0xab, 0x34, 0xc9, 0x42, 0xe9, 0x17, 0x7b, 0xa5, 0x53, 0x1f, 0x2d, 0xd0, 0x85, 0xd7,
	0x34, 0x29, 0xb4, 0x96, 0x65, 0x21, 0x4f, 0xbe, 0x22, 0x7c, 0x75, 0xd4, 0x02, 0xc8, 0x83, 0xbf,
	0x6c, 0xf0, 0xf8, 0x83, 0xaa, 0x2f, 0xff, 0x0b, 0x35, 0x9f, 0xae, 0x7d, 0xe0, 0x4c, 0x97, 0xa6,
	0xd3, 0x13, 0xd9, 0xa4, 0x79, 0xc2, 0x44, 0x7c, 0xa8, 0x30, 0x9c, 0xa9, 0xbe, 0xb2, 0xef, 0xcc,
	0xea, 0xd2, 0xa6, 0xb2, 0x69, 0x86, 0xc5, 0x5c, 0xda, 0xbe, 0x18, 0x7c, 0x73, 0xec, 0xbe, 0x52,
	0xb1, 0x5c, 0xa6, 0x74, 0x67, 0x67, 0xe7, 0x08, 0x48, 0x59, 0xaa, 0xfa, 0xe6, 0x58, 0xdd, 0x8e,
	0x43, 0xa6, 0x36, 0x45, 0x32, 0x68, 0x7f, 0x46, 0xd8, 0xf2, 0xc5, 0xe0, 0xd4, 0xa1, 0xda, 0x57,
	0xca, 0xaf, 0xc5, 0xcd, 0x0e, 0x9a, 0x8b, 0x9e, 0x3b, 0x39, 0x25, 0x10, 0x21, 0x8b, 0x02, 0x5b,
	0x24, 0x01, 0x0d, 0x20, 0xd2, 0xe7, 0x8e, 0x0e, 0x0b, 0x8f, 0x3e, 0x8e, 0x0f, 0xf5, 0xf7, 0xa7,
	0xca, 0xdc, 0x13, 0xa3, 0xf1, 0x58, 0x97, 0xd5, 0x55, 0xec, 0xb5, 0x96, 0xdb, 0x6a, 0x67, 0x49,
	0xfb, 0x05, 0xbc, 0xae, 0xe1, 0x75, 0x0d, 0xaf, 0xaf, 0x15, 0x1a, 0x1b, 0xe7, 0x74, 0xad, 0xa5,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0x76, 0xbc, 0xbd, 0x2f, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssetServiceClient interface {
	// Searches all the resources under a given accessible CRM scope
	// (project/folder/organization). This RPC gives callers
	// especially admins the ability to search all the resources under a scope,
	// even if they don't have .get permission of all the resources. Callers
	// should have cloud.assets.SearchAllResources permission on the requested
	// scope, otherwise it will be rejected.
	SearchAllResources(ctx context.Context, in *SearchAllResourcesRequest, opts ...grpc.CallOption) (*SearchAllResourcesResponse, error)
	// Searches all the IAM policies under a given accessible CRM scope
	// (project/folder/organization). This RPC gives callers
	// especially admins the ability to search all the IAM policies under a scope,
	// even if they don't have .getIamPolicy permission of all the IAM policies.
	// Callers should have cloud.assets.SearchAllIamPolicies permission on the
	// requested scope, otherwise it will be rejected.
	SearchAllIamPolicies(ctx context.Context, in *SearchAllIamPoliciesRequest, opts ...grpc.CallOption) (*SearchAllIamPoliciesResponse, error)
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) SearchAllResources(ctx context.Context, in *SearchAllResourcesRequest, opts ...grpc.CallOption) (*SearchAllResourcesResponse, error) {
	out := new(SearchAllResourcesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.asset.v1p1beta1.AssetService/SearchAllResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) SearchAllIamPolicies(ctx context.Context, in *SearchAllIamPoliciesRequest, opts ...grpc.CallOption) (*SearchAllIamPoliciesResponse, error) {
	out := new(SearchAllIamPoliciesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.asset.v1p1beta1.AssetService/SearchAllIamPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
type AssetServiceServer interface {
	// Searches all the resources under a given accessible CRM scope
	// (project/folder/organization). This RPC gives callers
	// especially admins the ability to search all the resources under a scope,
	// even if they don't have .get permission of all the resources. Callers
	// should have cloud.assets.SearchAllResources permission on the requested
	// scope, otherwise it will be rejected.
	SearchAllResources(context.Context, *SearchAllResourcesRequest) (*SearchAllResourcesResponse, error)
	// Searches all the IAM policies under a given accessible CRM scope
	// (project/folder/organization). This RPC gives callers
	// especially admins the ability to search all the IAM policies under a scope,
	// even if they don't have .getIamPolicy permission of all the IAM policies.
	// Callers should have cloud.assets.SearchAllIamPolicies permission on the
	// requested scope, otherwise it will be rejected.
	SearchAllIamPolicies(context.Context, *SearchAllIamPoliciesRequest) (*SearchAllIamPoliciesResponse, error)
}

// UnimplementedAssetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAssetServiceServer struct {
}

func (*UnimplementedAssetServiceServer) SearchAllResources(ctx context.Context, req *SearchAllResourcesRequest) (*SearchAllResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAllResources not implemented")
}
func (*UnimplementedAssetServiceServer) SearchAllIamPolicies(ctx context.Context, req *SearchAllIamPoliciesRequest) (*SearchAllIamPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAllIamPolicies not implemented")
}

func RegisterAssetServiceServer(s *grpc.Server, srv AssetServiceServer) {
	s.RegisterService(&_AssetService_serviceDesc, srv)
}

func _AssetService_SearchAllResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAllResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).SearchAllResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.asset.v1p1beta1.AssetService/SearchAllResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).SearchAllResources(ctx, req.(*SearchAllResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_SearchAllIamPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAllIamPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).SearchAllIamPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.asset.v1p1beta1.AssetService/SearchAllIamPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).SearchAllIamPolicies(ctx, req.(*SearchAllIamPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AssetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.asset.v1p1beta1.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchAllResources",
			Handler:    _AssetService_SearchAllResources_Handler,
		},
		{
			MethodName: "SearchAllIamPolicies",
			Handler:    _AssetService_SearchAllIamPolicies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/asset/v1p1beta1/asset_service.proto",
}
