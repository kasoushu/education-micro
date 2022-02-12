// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user.proto

package v1

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	LogIn(ctx context.Context, in *UserLogInReq, opts ...grpc.CallOption) (*LogInReply, error)
	AdminRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserReply, error)
	TeacherRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserReply, error)
	StudentRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserReply, error)
	UserInfo(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserInfoReply, error)
	DeleteUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserReply, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UserReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) LogIn(ctx context.Context, in *UserLogInReq, opts ...grpc.CallOption) (*LogInReply, error) {
	out := new(LogInReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AdminRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/AdminRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) TeacherRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/TeacherRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) StudentRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/StudentRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserInfo(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	LogIn(context.Context, *UserLogInReq) (*LogInReply, error)
	AdminRegister(context.Context, *UserRegisterReq) (*UserReply, error)
	TeacherRegister(context.Context, *UserRegisterReq) (*UserReply, error)
	StudentRegister(context.Context, *UserRegisterReq) (*UserReply, error)
	UserInfo(context.Context, *UserReq) (*UserInfoReply, error)
	DeleteUser(context.Context, *UserReq) (*UserReply, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UserReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) LogIn(context.Context, *UserLogInReq) (*LogInReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedUserServer) AdminRegister(context.Context, *UserRegisterReq) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminRegister not implemented")
}
func (UnimplementedUserServer) TeacherRegister(context.Context, *UserRegisterReq) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeacherRegister not implemented")
}
func (UnimplementedUserServer) StudentRegister(context.Context, *UserRegisterReq) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StudentRegister not implemented")
}
func (UnimplementedUserServer) UserInfo(context.Context, *UserReq) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserServer) DeleteUser(context.Context, *UserReq) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UpdateUserReq) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LogIn(ctx, req.(*UserLogInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AdminRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AdminRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/AdminRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AdminRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_TeacherRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).TeacherRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/TeacherRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).TeacherRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_StudentRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).StudentRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/StudentRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).StudentRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserInfo(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogIn",
			Handler:    _User_LogIn_Handler,
		},
		{
			MethodName: "AdminRegister",
			Handler:    _User_AdminRegister_Handler,
		},
		{
			MethodName: "TeacherRegister",
			Handler:    _User_TeacherRegister_Handler,
		},
		{
			MethodName: "StudentRegister",
			Handler:    _User_StudentRegister_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _User_UserInfo_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
