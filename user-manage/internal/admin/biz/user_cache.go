package biz

import (
	"context"

	adminv1 "github.com/EchoGroot/kratos-examples/user-manage/api/user-manage/v1"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/data"
)

// UserCacheUsecase cacheRepo 的使用示例，基本和不带 cache 的 baseRepo 一致
type UserCacheUsecase struct {
	repo *data.UserCacheRepo
}

func NewUserCacheUsecase(
	userRepo *data.UserCacheRepo,
) *UserCacheUsecase {
	return &UserCacheUsecase{
		repo: userRepo,
	}
}

func (u *UserCacheUsecase) createCheck(ctx context.Context, user *data.User) error {
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
func (u *UserCacheUsecase) Create(ctx context.Context, user *data.User) (*data.User, error) {
	if err := u.createCheck(ctx, user); err != nil {
		return nil, err
	}
	err := u.repo.Insert(ctx, user)
	return user, err
}

func (u *UserCacheUsecase) Delete(ctx context.Context, id string) error {
	row, err := u.repo.DeleteByPK(ctx, id)
	if err != nil {
		return err
	}
	if row == 0 {
		return adminv1.ErrorUserNotFound("用户 %s 不存在", id)
	}
	return nil
}

func (u *UserCacheUsecase) Get(ctx context.Context, id string) (*data.User, error) {
	user, err := u.repo.SelectOneByPK(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, adminv1.ErrorUserNotFound("用户 %s 不存在", id)
	}
	return user, nil
}

func (u *UserCacheUsecase) Update(ctx context.Context, id string, field map[string]interface{}) (*data.User, error) {
	row, err := u.repo.UpdateByPKWithMap(ctx, id, field)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, adminv1.ErrorUserNotFound("用户 %s 不存在", id)
	}
	return u.repo.SelectOneByPK(ctx, id)
}
