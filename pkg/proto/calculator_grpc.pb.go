// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pkg/proto/calculator.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ExpressionService_Register_FullMethodName         = "/calculatorpc.ExpressionService/Register"
	ExpressionService_CreateExpression_FullMethodName = "/calculatorpc.ExpressionService/CreateExpression"
)

// ExpressionServiceClient is the client API for ExpressionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpressionServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	CreateExpression(ctx context.Context, in *CreateExpressionRequest, opts ...grpc.CallOption) (*CreateExpressionResponse, error)
}

type expressionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExpressionServiceClient(cc grpc.ClientConnInterface) ExpressionServiceClient {
	return &expressionServiceClient{cc}
}

func (c *expressionServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, ExpressionService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressionServiceClient) CreateExpression(ctx context.Context, in *CreateExpressionRequest, opts ...grpc.CallOption) (*CreateExpressionResponse, error) {
	out := new(CreateExpressionResponse)
	err := c.cc.Invoke(ctx, ExpressionService_CreateExpression_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpressionServiceServer is the server API for ExpressionService service.
// All implementations must embed UnimplementedExpressionServiceServer
// for forward compatibility
type ExpressionServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	CreateExpression(context.Context, *CreateExpressionRequest) (*CreateExpressionResponse, error)
	mustEmbedUnimplementedExpressionServiceServer()
}

// UnimplementedExpressionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExpressionServiceServer struct {
}

func (UnimplementedExpressionServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedExpressionServiceServer) CreateExpression(context.Context, *CreateExpressionRequest) (*CreateExpressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpression not implemented")
}
func (UnimplementedExpressionServiceServer) mustEmbedUnimplementedExpressionServiceServer() {}

// UnsafeExpressionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpressionServiceServer will
// result in compilation errors.
type UnsafeExpressionServiceServer interface {
	mustEmbedUnimplementedExpressionServiceServer()
}

func RegisterExpressionServiceServer(s grpc.ServiceRegistrar, srv ExpressionServiceServer) {
	s.RegisterService(&ExpressionService_ServiceDesc, srv)
}

func _ExpressionService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressionServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpressionService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressionServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressionService_CreateExpression_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExpressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressionServiceServer).CreateExpression(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpressionService_CreateExpression_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressionServiceServer).CreateExpression(ctx, req.(*CreateExpressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExpressionService_ServiceDesc is the grpc.ServiceDesc for ExpressionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExpressionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculatorpc.ExpressionService",
	HandlerType: (*ExpressionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ExpressionService_Register_Handler,
		},
		{
			MethodName: "CreateExpression",
			Handler:    _ExpressionService_CreateExpression_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/calculator.proto",
}
