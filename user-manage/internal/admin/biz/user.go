package biz

import (
	"context"

	adminv1 "github.com/EchoGroot/kratos-examples/user-manage/api/user-manage/v1"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/data"
)

type UserUsecase struct {
	repo *data.UserRepo
}

func NewUserUsecase(
	userRepo *data.UserRepo,
) *UserUsecase {
	return &UserUsecase{
		repo: userRepo,
	}
}

func (u *UserUsecase) createCheck(ctx context.Context, user *data.User) error {
	old, err := u.repo.SelectOne(ctx, &data.User{
		Username: user.Username,
	})
	if err != nil {
		return err
	}
	if old != nil {
		return adminv1.ErrorUserNameRepeat("用户名 %s 重复", user.Username)
	}

	return nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *data.User) (*data.User, error) {
	if err := u.createCheck(ctx, user); err != nil {
		return nil, err
	}
	err := u.repo.Insert(ctx, user)
	return user, err
}
