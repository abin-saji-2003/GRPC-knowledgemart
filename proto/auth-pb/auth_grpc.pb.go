// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: auth-pb/auth.proto

package auth_pb

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
	AuthService_EmailSignup_FullMethodName    = "/auth.AuthService/EmailSignup"
	AuthService_EmailLogin_FullMethodName     = "/auth.AuthService/EmailLogin"
	AuthService_VerifyEmail_FullMethodName    = "/auth.AuthService/VerifyEmail"
	AuthService_ResendOTP_FullMethodName      = "/auth.AuthService/ResendOTP"
	AuthService_SellerRegister_FullMethodName = "/auth.AuthService/SellerRegister"
	AuthService_SellerLogin_FullMethodName    = "/auth.AuthService/SellerLogin"
	AuthService_AdminLogin_FullMethodName     = "/auth.AuthService/AdminLogin"
	AuthService_ValidateToken_FullMethodName  = "/auth.AuthService/ValidateToken"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	EmailSignup(ctx context.Context, in *EmailSignupRequest, opts ...grpc.CallOption) (*EmailSignupResponse, error)
	EmailLogin(ctx context.Context, in *EmailLoginRequest, opts ...grpc.CallOption) (*EmailLoginResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
	ResendOTP(ctx context.Context, in *ResendOTPRequest, opts ...grpc.CallOption) (*ResendOTPResponse, error)
	SellerRegister(ctx context.Context, in *SellerRegisterRequest, opts ...grpc.CallOption) (*SellerRegisterResponse, error)
	SellerLogin(ctx context.Context, in *SellerLoginRequest, opts ...grpc.CallOption) (*SellerLoginResponse, error)
	AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) EmailSignup(ctx context.Context, in *EmailSignupRequest, opts ...grpc.CallOption) (*EmailSignupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmailSignupResponse)
	err := c.cc.Invoke(ctx, AuthService_EmailSignup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) EmailLogin(ctx context.Context, in *EmailLoginRequest, opts ...grpc.CallOption) (*EmailLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmailLoginResponse)
	err := c.cc.Invoke(ctx, AuthService_EmailLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, AuthService_VerifyEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResendOTP(ctx context.Context, in *ResendOTPRequest, opts ...grpc.CallOption) (*ResendOTPResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResendOTPResponse)
	err := c.cc.Invoke(ctx, AuthService_ResendOTP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SellerRegister(ctx context.Context, in *SellerRegisterRequest, opts ...grpc.CallOption) (*SellerRegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SellerRegisterResponse)
	err := c.cc.Invoke(ctx, AuthService_SellerRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SellerLogin(ctx context.Context, in *SellerLoginRequest, opts ...grpc.CallOption) (*SellerLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SellerLoginResponse)
	err := c.cc.Invoke(ctx, AuthService_SellerLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, AuthService_AdminLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_ValidateToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	EmailSignup(context.Context, *EmailSignupRequest) (*EmailSignupResponse, error)
	EmailLogin(context.Context, *EmailLoginRequest) (*EmailLoginResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	ResendOTP(context.Context, *ResendOTPRequest) (*ResendOTPResponse, error)
	SellerRegister(context.Context, *SellerRegisterRequest) (*SellerRegisterResponse, error)
	SellerLogin(context.Context, *SellerLoginRequest) (*SellerLoginResponse, error)
	AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) EmailSignup(context.Context, *EmailSignupRequest) (*EmailSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailSignup not implemented")
}
func (UnimplementedAuthServiceServer) EmailLogin(context.Context, *EmailLoginRequest) (*EmailLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailLogin not implemented")
}
func (UnimplementedAuthServiceServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedAuthServiceServer) ResendOTP(context.Context, *ResendOTPRequest) (*ResendOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendOTP not implemented")
}
func (UnimplementedAuthServiceServer) SellerRegister(context.Context, *SellerRegisterRequest) (*SellerRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerRegister not implemented")
}
func (UnimplementedAuthServiceServer) SellerLogin(context.Context, *SellerLoginRequest) (*SellerLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerLogin not implemented")
}
func (UnimplementedAuthServiceServer) AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAuthServiceServer) ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_EmailSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).EmailSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_EmailSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).EmailSignup(ctx, req.(*EmailSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_EmailLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).EmailLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_EmailLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).EmailLogin(ctx, req.(*EmailLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResendOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResendOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ResendOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResendOTP(ctx, req.(*ResendOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SellerRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SellerRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SellerRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SellerRegister(ctx, req.(*SellerRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SellerLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SellerLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SellerLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SellerLogin(ctx, req.(*SellerLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AdminLogin(ctx, req.(*AdminLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ValidateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmailSignup",
			Handler:    _AuthService_EmailSignup_Handler,
		},
		{
			MethodName: "EmailLogin",
			Handler:    _AuthService_EmailLogin_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _AuthService_VerifyEmail_Handler,
		},
		{
			MethodName: "ResendOTP",
			Handler:    _AuthService_ResendOTP_Handler,
		},
		{
			MethodName: "SellerRegister",
			Handler:    _AuthService_SellerRegister_Handler,
		},
		{
			MethodName: "SellerLogin",
			Handler:    _AuthService_SellerLogin_Handler,
		},
		{
			MethodName: "AdminLogin",
			Handler:    _AuthService_AdminLogin_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _AuthService_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth-pb/auth.proto",
}
