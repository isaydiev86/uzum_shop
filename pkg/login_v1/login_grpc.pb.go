// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: login.proto

package login_v1

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

// LoginV1Client is the client API for LoginV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginV1Client interface {
	Login(ctx context.Context, in *Login_Request, opts ...grpc.CallOption) (*Login_Response, error)
	GetRefreshToken(ctx context.Context, in *GetRefreshToken_Request, opts ...grpc.CallOption) (*GetRefreshToken_Response, error)
	GetAccessToken(ctx context.Context, in *GetAccessToken_Request, opts ...grpc.CallOption) (*GetAccessToken_Response, error)
	Check(ctx context.Context, in *Check_Request, opts ...grpc.CallOption) (*Check_Respond, error)
}

type loginV1Client struct {
	cc grpc.ClientConnInterface
}

func NewLoginV1Client(cc grpc.ClientConnInterface) LoginV1Client {
	return &loginV1Client{cc}
}

func (c *loginV1Client) Login(ctx context.Context, in *Login_Request, opts ...grpc.CallOption) (*Login_Response, error) {
	out := new(Login_Response)
	err := c.cc.Invoke(ctx, "/login_v1.LoginV1/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginV1Client) GetRefreshToken(ctx context.Context, in *GetRefreshToken_Request, opts ...grpc.CallOption) (*GetRefreshToken_Response, error) {
	out := new(GetRefreshToken_Response)
	err := c.cc.Invoke(ctx, "/login_v1.LoginV1/GetRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginV1Client) GetAccessToken(ctx context.Context, in *GetAccessToken_Request, opts ...grpc.CallOption) (*GetAccessToken_Response, error) {
	out := new(GetAccessToken_Response)
	err := c.cc.Invoke(ctx, "/login_v1.LoginV1/GetAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginV1Client) Check(ctx context.Context, in *Check_Request, opts ...grpc.CallOption) (*Check_Respond, error) {
	out := new(Check_Respond)
	err := c.cc.Invoke(ctx, "/login_v1.LoginV1/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginV1Server is the server API for LoginV1 service.
// All implementations must embed UnimplementedLoginV1Server
// for forward compatibility
type LoginV1Server interface {
	Login(context.Context, *Login_Request) (*Login_Response, error)
	GetRefreshToken(context.Context, *GetRefreshToken_Request) (*GetRefreshToken_Response, error)
	GetAccessToken(context.Context, *GetAccessToken_Request) (*GetAccessToken_Response, error)
	Check(context.Context, *Check_Request) (*Check_Respond, error)
	mustEmbedUnimplementedLoginV1Server()
}

// UnimplementedLoginV1Server must be embedded to have forward compatible implementations.
type UnimplementedLoginV1Server struct {
}

func (UnimplementedLoginV1Server) Login(context.Context, *Login_Request) (*Login_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginV1Server) GetRefreshToken(context.Context, *GetRefreshToken_Request) (*GetRefreshToken_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefreshToken not implemented")
}
func (UnimplementedLoginV1Server) GetAccessToken(context.Context, *GetAccessToken_Request) (*GetAccessToken_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedLoginV1Server) Check(context.Context, *Check_Request) (*Check_Respond, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedLoginV1Server) mustEmbedUnimplementedLoginV1Server() {}

// UnsafeLoginV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginV1Server will
// result in compilation errors.
type UnsafeLoginV1Server interface {
	mustEmbedUnimplementedLoginV1Server()
}

func RegisterLoginV1Server(s grpc.ServiceRegistrar, srv LoginV1Server) {
	s.RegisterService(&LoginV1_ServiceDesc, srv)
}

func _LoginV1_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Login_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginV1Server).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login_v1.LoginV1/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginV1Server).Login(ctx, req.(*Login_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginV1_GetRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRefreshToken_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginV1Server).GetRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login_v1.LoginV1/GetRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginV1Server).GetRefreshToken(ctx, req.(*GetRefreshToken_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginV1_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessToken_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginV1Server).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login_v1.LoginV1/GetAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginV1Server).GetAccessToken(ctx, req.(*GetAccessToken_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginV1_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Check_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginV1Server).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login_v1.LoginV1/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginV1Server).Check(ctx, req.(*Check_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginV1_ServiceDesc is the grpc.ServiceDesc for LoginV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "login_v1.LoginV1",
	HandlerType: (*LoginV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginV1_Login_Handler,
		},
		{
			MethodName: "GetRefreshToken",
			Handler:    _LoginV1_GetRefreshToken_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _LoginV1_GetAccessToken_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _LoginV1_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login.proto",
}
