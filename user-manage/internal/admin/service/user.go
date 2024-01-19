package service

import (
	"context"

	"github.com/EchoGroot/kratos-examples/pkg/lang/structx/copier"
	adminv1 "github.com/EchoGroot/kratos-examples/user-manage/api/user-manage/v1"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/biz"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/data"
)

type UserService struct {
	adminv1.UnimplementedUserServiceServer
	userUC *biz.UserUsecase
}

func NewUserService(
	userUC *biz.UserUsecase,
) *UserService {
	return &UserService{
		userUC: userUC,
	}
}

func (r *UserService) CreateUser(ctx context.Context, req *adminv1.CreateUserRequest) (*adminv1.User, error) {
	user := &data.User{}
	if err := copier.Copy(user, req.User).Error; err != nil {
		return nil, err
	}

	createdUser, err := r.userUC.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser.Password = ""
	resp := &adminv1.User{}
	return resp, copier.Copy(resp, createdUser).Time2Timestamppb().Error
}
