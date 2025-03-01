// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc1
// source: service/service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MainService_GetMovies_FullMethodName = "/pb.MainService/GetMovies"
)

// MainServiceClient is the client API for MainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainServiceClient interface {
	GetMovies(ctx context.Context, in *GetMoviesRequest, opts ...grpc.CallOption) (*GetMoviesResponse, error)
}

type mainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMainServiceClient(cc grpc.ClientConnInterface) MainServiceClient {
	return &mainServiceClient{cc}
}

func (c *mainServiceClient) GetMovies(ctx context.Context, in *GetMoviesRequest, opts ...grpc.CallOption) (*GetMoviesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMoviesResponse)
	err := c.cc.Invoke(ctx, MainService_GetMovies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServiceServer is the server API for MainService service.
// All implementations must embed UnimplementedMainServiceServer
// for forward compatibility.
type MainServiceServer interface {
	GetMovies(context.Context, *GetMoviesRequest) (*GetMoviesResponse, error)
	mustEmbedUnimplementedMainServiceServer()
}

// UnimplementedMainServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMainServiceServer struct{}

func (UnimplementedMainServiceServer) GetMovies(context.Context, *GetMoviesRequest) (*GetMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
}
func (UnimplementedMainServiceServer) mustEmbedUnimplementedMainServiceServer() {}
func (UnimplementedMainServiceServer) testEmbeddedByValue()                     {}

// UnsafeMainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainServiceServer will
// result in compilation errors.
type UnsafeMainServiceServer interface {
	mustEmbedUnimplementedMainServiceServer()
}

func RegisterMainServiceServer(s grpc.ServiceRegistrar, srv MainServiceServer) {
	// If the following call pancis, it indicates UnimplementedMainServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MainService_ServiceDesc, srv)
}

func _MainService_GetMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).GetMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_GetMovies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).GetMovies(ctx, req.(*GetMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MainService_ServiceDesc is the grpc.ServiceDesc for MainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MainService",
	HandlerType: (*MainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovies",
			Handler:    _MainService_GetMovies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}
