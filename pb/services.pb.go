// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() { proto.RegisterFile("services.proto", fileDescriptor_8e16ccb8c5307b32) }

var fileDescriptor_8e16ccb8c5307b32 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0x92,
	0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f,
	0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x83, 0xaa, 0x90, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x87,
	0xe9, 0x30, 0x3a, 0xcf, 0xcc, 0xc5, 0x1d, 0x99, 0x5f, 0x5a, 0x14, 0x0c, 0x31, 0x48, 0xc8, 0x93,
	0x8b, 0xc5, 0x35, 0x39, 0x23, 0x5f, 0x48, 0x50, 0xaf, 0x20, 0x49, 0x2f, 0xb8, 0xa4, 0x28, 0x33,
	0x2f, 0xdd, 0x17, 0xa2, 0x43, 0x0a, 0x53, 0x48, 0x49, 0xba, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xa2,
	0x4a, 0x02, 0xfa, 0x65, 0x86, 0xfa, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0xfa, 0xa9, 0xc9,
	0x19, 0xf9, 0x56, 0x8c, 0x5a, 0x42, 0x4e, 0x5c, 0x2c, 0x01, 0x99, 0x79, 0xe9, 0x42, 0xfc, 0x20,
	0x7d, 0x01, 0x48, 0x06, 0x41, 0x04, 0xf2, 0x11, 0xc6, 0x48, 0x80, 0x8d, 0x11, 0x12, 0x42, 0x31,
	0xa6, 0x00, 0xa4, 0x37, 0x95, 0x4b, 0x00, 0xe4, 0x1c, 0x90, 0xeb, 0x52, 0x8b, 0x82, 0x4b, 0x8a,
	0x52, 0x13, 0x73, 0x89, 0x74, 0x9a, 0x26, 0xd8, 0x4c, 0x65, 0x25, 0x39, 0x90, 0x99, 0xc5, 0x60,
	0xfd, 0xfa, 0xc5, 0x60, 0x03, 0xd0, 0x1d, 0x6a, 0xc0, 0x08, 0xb3, 0xc6, 0x39, 0x27, 0x33, 0x35,
	0xaf, 0x84, 0x7c, 0x6b, 0x92, 0xc1, 0xfa, 0x71, 0x58, 0xa3, 0xc1, 0x28, 0x94, 0xcc, 0xc5, 0x07,
	0xb2, 0xc6, 0x29, 0xd3, 0x25, 0x93, 0x24, 0x4b, 0xd4, 0xc1, 0x96, 0x28, 0x2a, 0xc9, 0x80, 0x2c,
	0x49, 0xca, 0x4c, 0xc9, 0xc4, 0x69, 0x85, 0x01, 0x63, 0x12, 0x1b, 0x38, 0x62, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x0b, 0x40, 0x6a, 0x9e, 0x1c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// YourServiceClient is the client API for YourService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type YourServiceClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	Ping(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PongMessage, error)
	EchoServerStream(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (YourService_EchoServerStreamClient, error)
	EchoClientStream(ctx context.Context, opts ...grpc.CallOption) (YourService_EchoClientStreamClient, error)
	EchoBiDiStream(ctx context.Context, opts ...grpc.CallOption) (YourService_EchoBiDiStreamClient, error)
}

type yourServiceClient struct {
	cc *grpc.ClientConn
}

func NewYourServiceClient(cc *grpc.ClientConn) YourServiceClient {
	return &yourServiceClient{cc}
}

func (c *yourServiceClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/pb.YourService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yourServiceClient) Ping(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PongMessage, error) {
	out := new(PongMessage)
	err := c.cc.Invoke(ctx, "/pb.YourService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yourServiceClient) EchoServerStream(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (YourService_EchoServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_YourService_serviceDesc.Streams[0], "/pb.YourService/EchoServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &yourServiceEchoServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type YourService_EchoServerStreamClient interface {
	Recv() (*StringMessage, error)
	grpc.ClientStream
}

type yourServiceEchoServerStreamClient struct {
	grpc.ClientStream
}

func (x *yourServiceEchoServerStreamClient) Recv() (*StringMessage, error) {
	m := new(StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *yourServiceClient) EchoClientStream(ctx context.Context, opts ...grpc.CallOption) (YourService_EchoClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_YourService_serviceDesc.Streams[1], "/pb.YourService/EchoClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &yourServiceEchoClientStreamClient{stream}
	return x, nil
}

type YourService_EchoClientStreamClient interface {
	Send(*StringMessage) error
	CloseAndRecv() (*StringMessage, error)
	grpc.ClientStream
}

type yourServiceEchoClientStreamClient struct {
	grpc.ClientStream
}

func (x *yourServiceEchoClientStreamClient) Send(m *StringMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *yourServiceEchoClientStreamClient) CloseAndRecv() (*StringMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *yourServiceClient) EchoBiDiStream(ctx context.Context, opts ...grpc.CallOption) (YourService_EchoBiDiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_YourService_serviceDesc.Streams[2], "/pb.YourService/EchoBiDiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &yourServiceEchoBiDiStreamClient{stream}
	return x, nil
}

type YourService_EchoBiDiStreamClient interface {
	Send(*StringMessage) error
	Recv() (*StringMessage, error)
	grpc.ClientStream
}

type yourServiceEchoBiDiStreamClient struct {
	grpc.ClientStream
}

func (x *yourServiceEchoBiDiStreamClient) Send(m *StringMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *yourServiceEchoBiDiStreamClient) Recv() (*StringMessage, error) {
	m := new(StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// YourServiceServer is the server API for YourService service.
type YourServiceServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	Ping(context.Context, *PingMessage) (*PongMessage, error)
	EchoServerStream(*StringMessage, YourService_EchoServerStreamServer) error
	EchoClientStream(YourService_EchoClientStreamServer) error
	EchoBiDiStream(YourService_EchoBiDiStreamServer) error
}

// UnimplementedYourServiceServer can be embedded to have forward compatible implementations.
type UnimplementedYourServiceServer struct {
}

func (*UnimplementedYourServiceServer) Echo(ctx context.Context, req *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedYourServiceServer) Ping(ctx context.Context, req *PingMessage) (*PongMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedYourServiceServer) EchoServerStream(req *StringMessage, srv YourService_EchoServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoServerStream not implemented")
}
func (*UnimplementedYourServiceServer) EchoClientStream(srv YourService_EchoClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoClientStream not implemented")
}
func (*UnimplementedYourServiceServer) EchoBiDiStream(srv YourService_EchoBiDiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoBiDiStream not implemented")
}

func RegisterYourServiceServer(s *grpc.Server, srv YourServiceServer) {
	s.RegisterService(&_YourService_serviceDesc, srv)
}

func _YourService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.YourService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _YourService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.YourService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).Ping(ctx, req.(*PingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _YourService_EchoServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StringMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(YourServiceServer).EchoServerStream(m, &yourServiceEchoServerStreamServer{stream})
}

type YourService_EchoServerStreamServer interface {
	Send(*StringMessage) error
	grpc.ServerStream
}

type yourServiceEchoServerStreamServer struct {
	grpc.ServerStream
}

func (x *yourServiceEchoServerStreamServer) Send(m *StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _YourService_EchoClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(YourServiceServer).EchoClientStream(&yourServiceEchoClientStreamServer{stream})
}

type YourService_EchoClientStreamServer interface {
	SendAndClose(*StringMessage) error
	Recv() (*StringMessage, error)
	grpc.ServerStream
}

type yourServiceEchoClientStreamServer struct {
	grpc.ServerStream
}

func (x *yourServiceEchoClientStreamServer) SendAndClose(m *StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *yourServiceEchoClientStreamServer) Recv() (*StringMessage, error) {
	m := new(StringMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _YourService_EchoBiDiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(YourServiceServer).EchoBiDiStream(&yourServiceEchoBiDiStreamServer{stream})
}

type YourService_EchoBiDiStreamServer interface {
	Send(*StringMessage) error
	Recv() (*StringMessage, error)
	grpc.ServerStream
}

type yourServiceEchoBiDiStreamServer struct {
	grpc.ServerStream
}

func (x *yourServiceEchoBiDiStreamServer) Send(m *StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *yourServiceEchoBiDiStreamServer) Recv() (*StringMessage, error) {
	m := new(StringMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _YourService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.YourService",
	HandlerType: (*YourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _YourService_Echo_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _YourService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EchoServerStream",
			Handler:       _YourService_EchoServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EchoClientStream",
			Handler:       _YourService_EchoClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "EchoBiDiStream",
			Handler:       _YourService_EchoBiDiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}
