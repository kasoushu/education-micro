package service

import (
	"context"
	iv1 "education/api/v1/interface"
	"education/app/interface/internal/model"
)

type UserCase interface {
	LogIn(context.Context, model.UserLogIn) (iv1.LogInReply, error)
	CreateUser(context.Context, model.UserModel, int) error
	GetInfo(context.Context, int32) (model.UserInfo, error)
	DeleteUser(context.Context, int32) error
	UpdateUser(context.Context, model.UserModel, int32) error
}

//

func (s *InterfaceService) LogIn(ctx context.Context, req *iv1.UserLogInReq) (*iv1.LogInReply, error) {
	// check phone & password
	res, err := s.userCase.LogIn(ctx, model.UserLogIn{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &res, nil
}
func (s *InterfaceService) AdminRegister(ctx context.Context, req *iv1.UserRegisterReq) (*iv1.UserReply, error) {
	err := s.userCase.CreateUser(ctx, model.UserModel{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}, 2)
	if err != nil {
		return nil, err
	}
	return &iv1.UserReply{Message: "register successful!"}, nil
}
func (s *InterfaceService) TeacherRegister(ctx context.Context, req *iv1.UserRegisterReq) (*iv1.UserReply, error) {
	err := s.userCase.CreateUser(ctx, model.UserModel{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}, 1)
	if err != nil {
		return nil, err
	}
	return &iv1.UserReply{Message: "register successful!"}, nil

}
func (s *InterfaceService) StudentRegister(ctx context.Context, req *iv1.UserRegisterReq) (*iv1.UserReply, error) {
	err := s.userCase.CreateUser(ctx, model.UserModel{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}, 0)
	if err != nil {
		return nil, err
	}
	return &iv1.UserReply{Message: "register successful!"}, nil
}
func (s *InterfaceService) UserInfo(ctx context.Context, req *iv1.UserReq) (*iv1.UserInfoReply, error) {

	info, err := s.userCase.GetInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &iv1.UserInfoReply{
		Name:      info.Name,
		Phone:     info.Phone,
		IsAdmin:   info.IsAdmin,
		IsTeacher: info.IsTeacher,
		IsStudent: info.IsStudent,
		Id:        info.Id,
	}, nil
}
func (s *InterfaceService) DeleteUser(ctx context.Context, req *iv1.UserReq) (*iv1.UserReply, error) {
	err := s.userCase.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &iv1.UserReply{Message: "delete successful!"}, nil
}
func (s *InterfaceService) UpdateUser(ctx context.Context, req *iv1.UpdateUserReq) (*iv1.UserReply, error) {
	err := s.userCase.UpdateUser(ctx, model.UserModel{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}, req.Id)
	if err != nil {
		return nil, err
	}
	return &iv1.UserReply{Message: "update successful!"}, nil
}
