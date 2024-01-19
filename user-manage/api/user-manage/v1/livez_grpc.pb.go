// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/user-manage/v1/livez.proto

package adminv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LivezService_Livez_FullMethodName = "/usermanage.admin.v1.LivezService/Livez"
)

// LivezServiceClient is the client API for LivezService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LivezServiceClient interface {
	Livez(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type livezServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLivezServiceClient(cc grpc.ClientConnInterface) LivezServiceClient {
	return &livezServiceClient{cc}
}

func (c *livezServiceClient) Livez(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LivezService_Livez_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LivezServiceServer is the server API for LivezService service.
// All implementations should embed UnimplementedLivezServiceServer
// for forward compatibility
type LivezServiceServer interface {
	Livez(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

// UnimplementedLivezServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLivezServiceServer struct {
}

func (UnimplementedLivezServiceServer) Livez(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Livez not implemented")
}

// UnsafeLivezServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LivezServiceServer will
// result in compilation errors.
type UnsafeLivezServiceServer interface {
	mustEmbedUnimplementedLivezServiceServer()
}

func RegisterLivezServiceServer(s grpc.ServiceRegistrar, srv LivezServiceServer) {
	s.RegisterService(&LivezService_ServiceDesc, srv)
}

func _LivezService_Livez_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivezServiceServer).Livez(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivezService_Livez_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivezServiceServer).Livez(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// LivezService_ServiceDesc is the grpc.ServiceDesc for LivezService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LivezService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermanage.admin.v1.LivezService",
	HandlerType: (*LivezServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Livez",
			Handler:    _LivezService_Livez_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user-manage/v1/livez.proto",
}