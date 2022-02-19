package data

import (
	"context"
	userv1 "education/api/v1/user"
	"education/app/interface/internal/biz"
	"education/app/interface/internal/model"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	d   *Data
	log *log.Helper
}

func NewUserRepo(d *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		d:   d,
		log: log.NewHelper(logger),
	}
}
func (u *UserRepo) LogIn(ctx context.Context, um model.UserLogIn) (int32, error) {
	res, err := u.d.userClient.LogIn(ctx, &userv1.UserLogInReq{
		Phone:    um.Phone,
		Password: um.Password,
	})
	if err != nil {
		log.Error(err)
		return -1, model.ERROR_LOGIN_FAIL
	}
	return res.Id, nil
}

func (u *UserRepo) CreateStudent(ctx context.Context, um model.UserModel) error {
	//u.log.Info("passing")
	_, err := u.d.userClient.StudentRegister(ctx, &userv1.UserRegisterReq{
		Name:     um.Name,
		Phone:    um.Phone,
		Password: um.Password,
	})
	if err != nil {
		u.log.Error(err)
		return model.ERROR_REGISTER_FAIL
	}
	return nil
}

func (u *UserRepo) CreateAdmin(ctx context.Context, um model.UserModel) error {
	_, err := u.d.userClient.AdminRegister(ctx, &userv1.UserRegisterReq{
		Name:     um.Name,
		Phone:    um.Phone,
		Password: um.Password,
	})
	if err != nil {
		u.log.Error(err)
		return model.ERROR_REGISTER_FAIL
	}
	return nil
}

func (u *UserRepo) CreateTeacher(ctx context.Context, um model.UserModel) error {
	_, err := u.d.userClient.TeacherRegister(ctx, &userv1.UserRegisterReq{
		Name:     um.Name,
		Phone:    um.Phone,
		Password: um.Password,
	})
	if err != nil {
		u.log.Error(err)
		return model.ERROR_REGISTER_FAIL
	}
	return nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, um model.UserModel, id int) error {
	_, err := u.d.userClient.UpdateUser(ctx, &userv1.UpdateUserReq{
		Id:       int32(id),
		Name:     um.Name,
		Phone:    um.Phone,
		Password: um.Password,
	})
	if err != nil {
		u.log.Error(err)
		return model.ERROR_UPDATE_FAIL
	}
	return nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, id int) error {
	_, err := u.d.userClient.DeleteUser(ctx, &userv1.UserReq{
		Id: int32(id),
	})
	if err != nil {
		u.log.Error(err)
		return model.ERROR_DELETE_FAIL
	}
	return nil
}

func (u *UserRepo) GetUser(ctx context.Context, id int) (model.UserInfo, error) {
	info, err := u.d.userClient.UserInfo(ctx, &userv1.UserReq{
		Id: int32(id),
	})
	if err != nil {
		u.log.Error(err)
		return model.UserInfo{}, err
	}
	return model.UserInfo{
		Name:      info.Name,
		Phone:     info.Phone,
		IsAdmin:   info.IsAdmin,
		IsTeacher: info.IsTeacher,
		IsStudent: info.IsStudent,
		Id:        info.Id,
	}, nil
}
