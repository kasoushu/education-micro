// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type UserHTTPServer interface {
	AdminRegister(context.Context, *UserRegisterReq) (*UserReply, error)
	DeleteUser(context.Context, *UserReq) (*UserReply, error)
	LogIn(context.Context, *UserLogInReq) (*LogInReply, error)
	StudentRegister(context.Context, *UserRegisterReq) (*UserReply, error)
	TeacherRegister(context.Context, *UserRegisterReq) (*UserReply, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UserReply, error)
	UserInfo(context.Context, *UserReq) (*UserInfoReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/login", _User_LogIn0_HTTP_Handler(srv))
	r.POST("/register/admin", _User_AdminRegister0_HTTP_Handler(srv))
	r.POST("/register/teacher", _User_TeacherRegister0_HTTP_Handler(srv))
	r.POST("/register/student", _User_StudentRegister0_HTTP_Handler(srv))
	r.GET("/user/{id}", _User_UserInfo0_HTTP_Handler(srv))
	r.DELETE("/user/{id}", _User_DeleteUser0_HTTP_Handler(srv))
	r.PUT("/user/{id}", _User_UpdateUser0_HTTP_Handler(srv))
}

func _User_LogIn0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserLogInReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/LogIn")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LogIn(ctx, req.(*UserLogInReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogInReply)
		return ctx.Result(200, reply)
	}
}

func _User_AdminRegister0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/AdminRegister")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminRegister(ctx, req.(*UserRegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _User_TeacherRegister0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/TeacherRegister")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TeacherRegister(ctx, req.(*UserRegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _User_StudentRegister0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/StudentRegister")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StudentRegister(ctx, req.(*UserRegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UserInfo0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/UserInfo")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserInfo(ctx, req.(*UserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/DeleteUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*UserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/UpdateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	AdminRegister(ctx context.Context, req *UserRegisterReq, opts ...http.CallOption) (rsp *UserReply, err error)
	DeleteUser(ctx context.Context, req *UserReq, opts ...http.CallOption) (rsp *UserReply, err error)
	LogIn(ctx context.Context, req *UserLogInReq, opts ...http.CallOption) (rsp *LogInReply, err error)
	StudentRegister(ctx context.Context, req *UserRegisterReq, opts ...http.CallOption) (rsp *UserReply, err error)
	TeacherRegister(ctx context.Context, req *UserRegisterReq, opts ...http.CallOption) (rsp *UserReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserReq, opts ...http.CallOption) (rsp *UserReply, err error)
	UserInfo(ctx context.Context, req *UserReq, opts ...http.CallOption) (rsp *UserInfoReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) AdminRegister(ctx context.Context, in *UserRegisterReq, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/register/admin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/AdminRegister"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) DeleteUser(ctx context.Context, in *UserReq, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/DeleteUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) LogIn(ctx context.Context, in *UserLogInReq, opts ...http.CallOption) (*LogInReply, error) {
	var out LogInReply
	pattern := "/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/LogIn"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) StudentRegister(ctx context.Context, in *UserRegisterReq, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/register/student"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/StudentRegister"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) TeacherRegister(ctx context.Context, in *UserRegisterReq, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/register/teacher"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/TeacherRegister"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/user/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/UpdateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UserInfo(ctx context.Context, in *UserReq, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/UserInfo"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
