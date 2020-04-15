// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/shopping_performance_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for
// [ShoppingPerformanceViewService.GetShoppingPerformanceView][google.ads.googleads.v1.services.ShoppingPerformanceViewService.GetShoppingPerformanceView].
type GetShoppingPerformanceViewRequest struct {
	// Required. The resource name of the Shopping performance view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShoppingPerformanceViewRequest) Reset()         { *m = GetShoppingPerformanceViewRequest{} }
func (m *GetShoppingPerformanceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetShoppingPerformanceViewRequest) ProtoMessage()    {}
func (*GetShoppingPerformanceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6c4cf22050d6d3, []int{0}
}

func (m *GetShoppingPerformanceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Unmarshal(m, b)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShoppingPerformanceViewRequest.Merge(m, src)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Size(m)
}
func (m *GetShoppingPerformanceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShoppingPerformanceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShoppingPerformanceViewRequest proto.InternalMessageInfo

func (m *GetShoppingPerformanceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetShoppingPerformanceViewRequest)(nil), "google.ads.googleads.v1.services.GetShoppingPerformanceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/shopping_performance_view_service.proto", fileDescriptor_4d6c4cf22050d6d3)
}

var fileDescriptor_4d6c4cf22050d6d3 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x31, 0x6f, 0xd4, 0x30,
	0x18, 0x55, 0x82, 0x84, 0x44, 0x04, 0x4b, 0x16, 0x4a, 0x40, 0x70, 0x94, 0x0e, 0x88, 0xc1, 0x26,
	0x65, 0x00, 0x8c, 0x18, 0x7c, 0x0c, 0xc7, 0x04, 0x27, 0x2a, 0xdd, 0x80, 0x82, 0x22, 0x37, 0xf9,
	0x9a, 0x5a, 0x4a, 0xec, 0x60, 0xe7, 0xd2, 0x01, 0xb1, 0x74, 0x64, 0xe5, 0x1f, 0x30, 0x32, 0xf0,
	0x43, 0xba, 0xb2, 0x31, 0x31, 0x30, 0x31, 0xf0, 0x03, 0x98, 0x50, 0x6a, 0x3b, 0x77, 0x45, 0xf2,
	0xdd, 0xf6, 0x94, 0xf7, 0xf2, 0xbe, 0xe7, 0xe7, 0xcf, 0xd1, 0xcb, 0x4a, 0xca, 0xaa, 0x06, 0xcc,
	0x4a, 0x8d, 0x0d, 0x1c, 0x50, 0x9f, 0x62, 0x0d, 0xaa, 0xe7, 0x05, 0x68, 0xac, 0x8f, 0x65, 0xdb,
	0x72, 0x51, 0xe5, 0x2d, 0xa8, 0x23, 0xa9, 0x1a, 0x26, 0x0a, 0xc8, 0x7b, 0x0e, 0x27, 0xb9, 0x95,
	0xa0, 0x56, 0xc9, 0x4e, 0xc6, 0x13, 0xf3, 0x3b, 0x62, 0xa5, 0x46, 0xa3, 0x13, 0xea, 0x53, 0xe4,
	0x9c, 0x12, 0xea, 0x9b, 0xa5, 0x40, 0xcb, 0xa5, 0xda, 0x38, 0xcc, 0x0c, 0x49, 0x6e, 0x39, 0x8b,
	0x96, 0x63, 0x26, 0x84, 0xec, 0x58, 0xc7, 0xa5, 0xd0, 0x96, 0xbd, 0xbe, 0xc6, 0x16, 0x35, 0x07,
	0xd1, 0x59, 0xe2, 0xce, 0x1a, 0x71, 0xc4, 0xa1, 0x2e, 0xf3, 0x43, 0x38, 0x66, 0x3d, 0x97, 0xca,
	0x0a, 0x6e, 0xac, 0x09, 0x5c, 0x1a, 0x43, 0xed, 0x9e, 0x06, 0xd1, 0xdd, 0x19, 0x74, 0x07, 0x36,
	0xd9, 0x7c, 0x15, 0x6c, 0xc1, 0xe1, 0xe4, 0x0d, 0xbc, 0x5f, 0x82, 0xee, 0xe2, 0x77, 0xd1, 0x35,
	0xf7, 0x5f, 0x2e, 0x58, 0x03, 0x3b, 0xc1, 0x24, 0xb8, 0x7f, 0x65, 0xfa, 0xe4, 0x27, 0x0d, 0xff,
	0xd2, 0xfd, 0xe8, 0xe1, 0xaa, 0x11, 0x8b, 0x5a, 0xae, 0x51, 0x21, 0x1b, 0xec, 0xf3, 0xbd, 0xea,
	0xec, 0x5e, 0xb1, 0x06, 0xf6, 0xbf, 0x85, 0xd1, 0x6d, 0x8f, 0xf2, 0xc0, 0xd4, 0x1b, 0xff, 0x09,
	0xa2, 0xc4, 0x9f, 0x33, 0x7e, 0x81, 0xb6, 0xdd, 0x0f, 0xda, 0x7a, 0xca, 0x84, 0x78, 0x4d, 0xc6,
	0x2b, 0x44, 0x1e, 0x8b, 0xdd, 0xd7, 0x3f, 0xe8, 0xc5, 0x8a, 0x4e, 0xbf, 0xff, 0xfa, 0x1c, 0x3e,
	0x8d, 0x1f, 0x0f, 0x1b, 0xf0, 0xe1, 0x02, 0xf3, 0xbc, 0x58, 0xea, 0x4e, 0x36, 0xa0, 0x34, 0x7e,
	0x30, 0xae, 0xc4, 0x7f, 0x7e, 0x1f, 0x93, 0x9b, 0x67, 0x74, 0xc7, 0x57, 0xeb, 0xf4, 0x53, 0x18,
	0xed, 0x15, 0xb2, 0xd9, 0x7a, 0xe8, 0xe9, 0xbd, 0xcd, 0xb5, 0xce, 0x87, 0x1d, 0x98, 0x07, 0x6f,
	0xed, 0x3b, 0x41, 0x95, 0xac, 0x99, 0xa8, 0x90, 0x54, 0x15, 0xae, 0x40, 0x9c, 0x6f, 0x08, 0x5e,
	0x8d, 0xf6, 0x3f, 0xa3, 0x67, 0x0e, 0x7c, 0x09, 0x2f, 0xcd, 0x28, 0xfd, 0x1a, 0x4e, 0x66, 0xc6,
	0x90, 0x96, 0x1a, 0x19, 0x38, 0xa0, 0x45, 0x8a, 0xec, 0x60, 0x7d, 0xe6, 0x24, 0x19, 0x2d, 0x75,
	0x36, 0x4a, 0xb2, 0x45, 0x9a, 0x39, 0xc9, 0xef, 0x70, 0xcf, 0x7c, 0x27, 0x84, 0x96, 0x9a, 0x90,
	0x51, 0x44, 0xc8, 0x22, 0x25, 0xc4, 0xc9, 0x0e, 0x2f, 0x9f, 0xe7, 0x7c, 0xf4, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x2d, 0x96, 0xa3, 0xbf, 0xed, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShoppingPerformanceViewServiceClient is the client API for ShoppingPerformanceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShoppingPerformanceViewServiceClient interface {
	// Returns the requested Shopping performance view in full detail.
	GetShoppingPerformanceView(ctx context.Context, in *GetShoppingPerformanceViewRequest, opts ...grpc.CallOption) (*resources.ShoppingPerformanceView, error)
}

type shoppingPerformanceViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShoppingPerformanceViewServiceClient(cc grpc.ClientConnInterface) ShoppingPerformanceViewServiceClient {
	return &shoppingPerformanceViewServiceClient{cc}
}

func (c *shoppingPerformanceViewServiceClient) GetShoppingPerformanceView(ctx context.Context, in *GetShoppingPerformanceViewRequest, opts ...grpc.CallOption) (*resources.ShoppingPerformanceView, error) {
	out := new(resources.ShoppingPerformanceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ShoppingPerformanceViewService/GetShoppingPerformanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingPerformanceViewServiceServer is the server API for ShoppingPerformanceViewService service.
type ShoppingPerformanceViewServiceServer interface {
	// Returns the requested Shopping performance view in full detail.
	GetShoppingPerformanceView(context.Context, *GetShoppingPerformanceViewRequest) (*resources.ShoppingPerformanceView, error)
}

// UnimplementedShoppingPerformanceViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShoppingPerformanceViewServiceServer struct {
}

func (*UnimplementedShoppingPerformanceViewServiceServer) GetShoppingPerformanceView(ctx context.Context, req *GetShoppingPerformanceViewRequest) (*resources.ShoppingPerformanceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShoppingPerformanceView not implemented")
}

func RegisterShoppingPerformanceViewServiceServer(s *grpc.Server, srv ShoppingPerformanceViewServiceServer) {
	s.RegisterService(&_ShoppingPerformanceViewService_serviceDesc, srv)
}

func _ShoppingPerformanceViewService_GetShoppingPerformanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShoppingPerformanceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingPerformanceViewServiceServer).GetShoppingPerformanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ShoppingPerformanceViewService/GetShoppingPerformanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingPerformanceViewServiceServer).GetShoppingPerformanceView(ctx, req.(*GetShoppingPerformanceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShoppingPerformanceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ShoppingPerformanceViewService",
	HandlerType: (*ShoppingPerformanceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShoppingPerformanceView",
			Handler:    _ShoppingPerformanceViewService_GetShoppingPerformanceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/shopping_performance_view_service.proto",
}
