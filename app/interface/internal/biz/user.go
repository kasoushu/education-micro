package biz

import (
	"context"
	iv1 "education/api/v1/interface"
	"education/app/interface/internal/conf"
	"education/app/interface/internal/model"
	"education/app/interface/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserCase struct {
	ur  UserRepo
	log *log.Helper
	key string
}
type LogInClaim struct {
	Id int32
	jwt.RegisteredClaims
}
type UserRepo interface {
	CreateStudent(context.Context, model.UserModel) error
	CreateAdmin(context.Context, model.UserModel) error
	CreateTeacher(context.Context, model.UserModel) error
	UpdateUser(context.Context, model.UserModel, int) error
	DeleteUser(context.Context, int) error
	GetUser(context.Context, int) (model.UserInfo, error)
	LogIn(context.Context, model.UserLogIn) (int32, error)
}

func NewUserCase(ur UserRepo, c *conf.AppConfig, l log.Logger) service.UserCase {
	return &UserCase{
		ur:  ur,
		key: c.Auth.ServiceKey,
		log: log.NewHelper(l),
	}
}

func (u *UserCase) LogIn(ctx context.Context, in model.UserLogIn) (iv1.LogInReply, error) {

	id, err := u.ur.LogIn(ctx, in)
	if err != nil {
		log.Error(err)
		return iv1.LogInReply{}, model.ERROR_LOGIN_FAIL
	}
	// generate token
	claim := LogInClaim{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(u.key)
	if err != nil {
		u.log.Error(err)
		return iv1.LogInReply{}, model.ERROR_TOKEN_GENERATE_FAIL
	}
	return iv1.LogInReply{
		Id:    id,
		Token: tokenString,
	}, nil
}

func (u *UserCase) CreateUser(ctx context.Context, userModel model.UserModel, i int) error {
	var err error
	if i == 0 {
		err = u.ur.CreateStudent(ctx, userModel)
	} else if i == 1 {
		err = u.ur.CreateTeacher(ctx, userModel)
	} else if i == 2 {
		err = u.ur.CreateAdmin(ctx, userModel)
	}
	if err != nil {
		u.log.Error(err)
		return err
	}
	return nil
}

func (u *UserCase) GetInfo(ctx context.Context, id int32) (model.UserInfo, error) {
	info, err := u.ur.GetUser(ctx, int(id))
	if err != nil {
		u.log.Error(err)
		return model.UserInfo{}, model.ERROR_GET_INFO_FAIL
	}
	return info, nil
}

func (u *UserCase) DeleteUser(ctx context.Context, id int32) error {
	err := u.ur.DeleteUser(ctx, int(id))
	if err != nil {
		u.log.Error(err)
		return model.ERROR_DELETE_FAIL
	}
	return nil
}

func (u *UserCase) UpdateUser(ctx context.Context, userModel model.UserModel, i int32) error {
	err := u.ur.UpdateUser(ctx, userModel, int(i))
	if err != nil {
		u.log.Error(err)
		return model.ERROR_DELETE_FAIL
	}
	return nil
}
