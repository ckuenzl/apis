// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shareit/v1/service/shareit_service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_3022ae706155c48c, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "shareit.Empty")
}

func init() {
	proto.RegisterFile("shareit/v1/service/shareit_service.proto", fileDescriptor_3022ae706155c48c)
}

var fileDescriptor_3022ae706155c48c = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0xce, 0x48, 0x2c,
	0x4a, 0xcd, 0x2c, 0xd1, 0x2f, 0x33, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x87,
	0x0a, 0xc5, 0x43, 0xf9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0x61, 0x25, 0x76,
	0x2e, 0x56, 0xd7, 0xdc, 0x82, 0x92, 0x4a, 0x23, 0x0b, 0x2e, 0xbe, 0x60, 0x90, 0x98, 0x67, 0x49,
	0x30, 0x44, 0xa5, 0x90, 0x1a, 0x17, 0x4b, 0x48, 0x6a, 0x71, 0x89, 0x10, 0x9f, 0x1e, 0x54, 0xb1,
	0x1e, 0x58, 0xa5, 0x14, 0x1a, 0xdf, 0x89, 0x3b, 0x8a, 0x53, 0xcf, 0x1a, 0x6a, 0x7c, 0x12, 0x1b,
	0xd8, 0x7c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x7e, 0xef, 0xc3, 0x8b, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShareItServiceClient is the client API for ShareItService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShareItServiceClient interface {
	Test(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type shareItServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShareItServiceClient(cc grpc.ClientConnInterface) ShareItServiceClient {
	return &shareItServiceClient{cc}
}

func (c *shareItServiceClient) Test(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/shareit.ShareItService/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShareItServiceServer is the server API for ShareItService service.
type ShareItServiceServer interface {
	Test(context.Context, *Empty) (*Empty, error)
}

// UnimplementedShareItServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShareItServiceServer struct {
}

func (*UnimplementedShareItServiceServer) Test(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}

func RegisterShareItServiceServer(s *grpc.Server, srv ShareItServiceServer) {
	s.RegisterService(&_ShareItService_serviceDesc, srv)
}

func _ShareItService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareItServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shareit.ShareItService/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareItServiceServer).Test(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShareItService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shareit.ShareItService",
	HandlerType: (*ShareItServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _ShareItService_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shareit/v1/service/shareit_service.proto",
}
