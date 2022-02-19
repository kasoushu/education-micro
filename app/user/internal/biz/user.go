package biz

import (
	"context"
	"education/app/user/internal/model"
	"education/app/user/internal/service"
	"github.com/go-kratos/kratos/v2/log"
)

type UserCase struct {
	u   UserRepoInt
	log *log.Helper
}

type UserRepoInt interface {
	CreateStudent(context.Context, model.UserModel) (int32, error)
	CreateAdmin(context.Context, model.UserModel) (int32, error)
	CreateTeacher(context.Context, model.UserModel) (int32, error)
	UpdateUser(context.Context, model.UserModel, int) error
	DeleteUser(context.Context, int) error
	VerifyPassword(context.Context, model.UserCheck) (int, error)
	CheckUserByPhone(context.Context, string) (int, error)
	GetUser(context.Context, int) (model.UserInfo, error)
}

func NewUserCase(repo UserRepoInt, logger log.Logger) service.UserCase {
	return &UserCase{
		u:   repo,
		log: log.NewHelper(log.With(logger, "module", "biz-interface")),
	}
}

func (uc *UserCase) VerifyPassword(ctx context.Context, check model.UserCheck) (int32, error) {
	id, err := uc.u.VerifyPassword(ctx, check)
	if err != nil {
		return -1, err
	}
	return int32(id), nil
}

func (uc *UserCase) CheckAndCreate(ctx context.Context, model model.UserModel, i int) error {
	//check
	var err error
	if i == 0 {
		_, err = uc.u.CreateStudent(ctx, model)
	} else if i == 1 {
		_, err = uc.u.CreateTeacher(ctx, model)
	} else if i == 2 {
		_, err = uc.u.CreateAdmin(ctx, model)
	}
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserCase) GetInfo(ctx context.Context, i int32) (model.UserInfo, error) {
	info, err := uc.u.GetUser(ctx, int(i))
	if err != nil {
		return model.UserInfo{}, err
	}
	return info, nil
}

func (uc *UserCase) DeleteUser(ctx context.Context, i int32) error {
	err := uc.u.DeleteUser(ctx, int(i))
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserCase) UpdateUser(ctx context.Context, model model.UserModel, i int32) error {
	err := uc.u.UpdateUser(ctx, model, int(i))
	if err != nil {
		return err
	}
	return nil
}
