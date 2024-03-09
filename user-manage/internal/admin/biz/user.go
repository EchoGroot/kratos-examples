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

// Create 命名没有使用 CreateUser，原因是这里可以根据上下文（结构体）辅助命名
func (u *UserUsecase) Create(ctx context.Context, user *data.User) (*data.User, error) {
	if err := u.createCheck(ctx, user); err != nil {
		return nil, err
	}
	err := u.repo.Insert(ctx, user)
	return user, err
}

func (u *UserUsecase) Delete(ctx context.Context, id string) error {
	row, err := u.repo.DeleteByPK(ctx, id)
	if err != nil {
		return err
	}
	if row == 0 {
		return adminv1.ErrorUserNotFound("用户 %s 不存在", id)
	}
	return nil
}

func (u *UserUsecase) Get(ctx context.Context, id string) (*data.User, error) {
	user, err := u.repo.SelectOneByPK(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, adminv1.ErrorUserNotFound("用户 %s 不存在", id)
	}
	return user, nil
}

func (u *UserUsecase) Update(ctx context.Context, id string, field map[string]interface{}) (*data.User, error) {
	row, err := u.repo.UpdateByPKWithMap(ctx, id, field)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, adminv1.ErrorUserNotFound("用户 %s 不存在", id)
	}
	return u.repo.SelectOneByPK(ctx, id)
}

func (u *UserUsecase) List(ctx context.Context, params *data.ListUsersParams) ([]*data.User, int32, error) {
	users, total, err := u.repo.List(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	for _, user := range users {
		user.Password = ""
	}
	return users, total, err
}
