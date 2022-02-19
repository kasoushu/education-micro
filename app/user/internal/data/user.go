package data

import (
	"context"
	"education/app/user/internal/biz"
	"education/app/user/internal/data/ent"
	"education/app/user/internal/data/ent/user"
	"education/app/user/internal/model"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

//
//type UserRepoInt interface {
//	CreateStudent(context.Context, UserModel) (int32, error)
//	CreateAdmin(context.Context, UserModel) (int32, error)
//	CreateTeacher(context.Context, UserModel) (int32, error)
//	UpdateUser(context.Context, UserModel, int) error
//	DeleteUser(context.Context, int) error
//	VerifyPassword(context.Context, UserCheck) (int, error)
//	CheckUserByPhone(context.Context, string) (int, error)
//	GetUser(context.Context, int) (UserInfo, error)
//}

func NewUserRepo(d *Data, logger log.Logger) biz.UserRepoInt {
	return &UserRepo{
		data: d,
		log:  log.NewHelper(log.With(logger, "module", "data")),
	}
}

func (u *UserRepo) GetUser(ctx context.Context, id int) (model.UserInfo, error) {
	p, err := u.data.db.User.Get(ctx, id)
	if err != nil {
		return model.UserInfo{}, err
	}
	return model.UserInfo{
		Name:      p.Name,
		Phone:     p.Phone,
		IsAdmin:   p.IsAdmin,
		IsTeacher: p.IsTeacher,
		IsStudent: p.IsStudent,
		Id:        int32(p.ID),
	}, nil
}
func (u *UserRepo) CheckUserByPhone(ctx context.Context, phone string) (int, error) {
	us, err := u.data.db.User.Query().Where(user.Phone(phone)).Only(ctx)
	if err != nil {
		return -1, err
	}
	return us.ID, nil
}

func (u *UserRepo) CreateStudent(ctx context.Context, um model.UserModel) (int32, error) {

	_, err := u.data.db.User.Query().Where(user.Phone(um.Phone)).Only(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {

			us, err := u.data.db.User.Create().SetName(um.Name).SetPassword(um.Password).SetPhone(um.Phone).SetIsStudent(true).Save(ctx)
			if err != nil {
				return -1, err
			}
			return int32(us.ID), nil
		}
		return -1, err
	}
	return -1, errors.New("interface have existed")
}

func (u *UserRepo) CreateAdmin(ctx context.Context, um model.UserModel) (int32, error) {
	_, err := u.data.db.User.Query().Where(user.Phone(um.Phone)).Only(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			us, err := u.data.db.User.Create().SetName(um.Name).SetPassword(um.Password).SetPhone(um.Phone).SetIsAdmin(true).Save(ctx)
			if err != nil {
				return -1, err
			}
			return int32(us.ID), nil
		}
		return -1, err
	}
	return -1, errors.New("interface have existed")
}
func (u *UserRepo) CreateTeacher(ctx context.Context, um model.UserModel) (int32, error) {
	_, err := u.data.db.User.Query().Where(user.Phone(um.Phone)).Only(ctx)
	if err != nil {
		_, ok := err.(*ent.NotFoundError)
		if ok {
			fmt.Println("not font")
			us, err := u.data.db.User.Create().SetName(um.Name).SetPassword(um.Password).SetPhone(um.Phone).SetIsTeacher(true).Save(ctx)
			if err != nil {
				return -1, err
			}
			return int32(us.ID), nil
		}
		return -1, err
	}
	return -1, errors.New("interface have existed")
}

func (u *UserRepo) UpdateUser(ctx context.Context, um model.UserModel, id int) error {
	us, err := u.data.db.User.Get(ctx, id)
	if err != nil {
		return err
	}
	up := u.data.db.User.UpdateOne(us)
	if um.Name != "" {
		up = up.SetName(um.Name)
	}
	if um.Password != "" {
		up = up.SetPassword(um.Password)
	}
	if um.Phone != "" {
		_, err := u.data.db.User.Query().Where(user.Phone(um.Phone)).Only(ctx)
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				up = up.SetPhone(um.Phone)
			}
		} else {
			return errors.New("this phone number had been used")
		}
	}
	_, err = up.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, id int) error {
	err := u.data.db.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) VerifyPassword(ctx context.Context, ucheck model.UserCheck) (int, error) {
	us, err := u.data.db.User.Query().Where(user.Phone(ucheck.Phone)).Only(ctx)
	if err != nil {
		return -1, err
	}
	if us.Password == ucheck.Password {
		return us.ID, nil
	}
	return -1, errors.New("password is wrong")
}
