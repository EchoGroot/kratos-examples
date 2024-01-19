package data

import (
	"github.com/EchoGroot/kratos-examples/pkg/gormx/base"
	"github.com/EchoGroot/kratos-examples/pkg/gormx/repo"
	adminv1 "github.com/EchoGroot/kratos-examples/user-manage/api/user-manage/v1"
	"gorm.io/plugin/soft_delete"

	"gorm.io/gorm"
)

type User struct {
	base.Model[string]
	Deleted  soft_delete.DeletedAt
	Username string             // 用户名
	NickName string             // 昵称
	Password string             // 密码
	Status   adminv1.UserStatus // 状态
	Tel      string             // 联系电话
	Email    string             // 邮箱
}

func (User) TableName() string {
	return "user"
}

type UserRepo struct {
	repo.BaseRepo[User]
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		BaseRepo: repo.NewBaseRepo[User](db),
	}
}
