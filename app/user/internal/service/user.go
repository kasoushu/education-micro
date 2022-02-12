package service

import (
	"context"
	v1 "education/api/user/v1"
	"education/app/user/internal/model"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc  UserCase
	log *log.Helper
}

type UserCase interface {
	VerifyPassword(context.Context, model.UserCheck) (int32, error)
	CheckAndCreate(context.Context, model.UserModel, int) error //int wei id 0 student,1 teacher, 2 admin
	GetInfo(context.Context, int32) (model.UserInfo, error)
	DeleteUser(context.Context, int32) error
	UpdateUser(context.Context, model.UserModel, int32) error
}

func NewUserService(u UserCase, l log.Logger) *UserService {
	return &UserService{
		uc:  u,
		log: log.NewHelper(log.With(l, "module", "user-service")),
	}
}

//

func (s *UserService) LogIn(ctx context.Context, req *v1.UserLogInReq) (*v1.LogInReply, error) {
	// check phone & password
	id, err := s.uc.VerifyPassword(ctx, model.UserCheck{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.LogInReply{Id: id}, nil
}
func (s *UserService) AdminRegister(ctx context.Context, req *v1.UserRegisterReq) (*v1.UserReply, error) {
	if req.Password == "" {
		return nil, errors.New("password can not be empty")
	}
	err := s.uc.CheckAndCreate(ctx, model.UserModel{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}, 2)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{Message: "register successful!"}, nil
}
func (s *UserService) TeacherRegister(ctx context.Context, req *v1.UserRegisterReq) (*v1.UserReply, error) {
	if req.Password == "" {
		return nil, errors.New("password can not be empty")
	}
	err := s.uc.CheckAndCreate(ctx, model.UserModel{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}, 1)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{Message: "register successful!"}, nil

}
func (s *UserService) StudentRegister(ctx context.Context, req *v1.UserRegisterReq) (*v1.UserReply, error) {
	if req.Password == "" {
		return nil, errors.New("password can not be empty")
	}
	err := s.uc.CheckAndCreate(ctx, model.UserModel{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}, 0)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{Message: "register successful!"}, nil
}
func (s *UserService) UserInfo(ctx context.Context, req *v1.UserReq) (*v1.UserInfoReply, error) {

	info, err := s.uc.GetInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoReply{
		Name:      info.Name,
		Phone:     info.Phone,
		IsAdmin:   info.IsAdmin,
		IsTeacher: info.IsTeacher,
		IsStudent: info.IsStudent,
		Id:        info.Id,
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *v1.UserReq) (*v1.UserReply, error) {
	err := s.uc.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{Message: "delete successful!"}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UserReply, error) {
	err := s.uc.UpdateUser(ctx, model.UserModel{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{Message: "update successful!"}, nil
}
