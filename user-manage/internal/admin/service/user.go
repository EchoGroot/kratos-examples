package service

import (
	"context"

	"github.com/EchoGroot/kratos-examples/pkg/lang/structx"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/EchoGroot/kratos-examples/pkg/lang/structx/copier"
	adminv1 "github.com/EchoGroot/kratos-examples/user-manage/api/user-manage/v1"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/biz"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/data"
)

type UserService struct {
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

func (r *UserService) DeleteUser(ctx context.Context, req *adminv1.DeleteUserRequest) (*emptypb.Empty, error) {
	err := r.userUC.DeleteUser(ctx, req.Id)
	return &emptypb.Empty{}, err
}

func (r *UserService) GetUser(ctx context.Context, req *adminv1.GetUserRequest) (*adminv1.User, error) {
	user, err := r.userUC.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	resp := &adminv1.User{}
	return resp, copier.Copy(resp, user).Time2Timestamppb().Error
}

func (r *UserService) UpdateUser(ctx context.Context, req *adminv1.UpdateUserRequest) (*adminv1.User, error) {
	updateField, err := structx.StructToMapByFieldMask(req.UpdateMask, req.User)
	if err != nil {
		return nil, err
	}
	user, err := r.userUC.UpdateUser(ctx, req.User.Id, updateField)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	var vo adminv1.User
	return &vo, copier.Copy(&vo, user).Time2Timestamppb().Error
}

func (r *UserService) ListUsers(ctx context.Context, req *adminv1.ListUsersRequest) (*adminv1.ListUsersResponse, error) {
	var params data.ListUsersParams
	if err := copier.Copy(&params, req).Error; err != nil {
		return nil, err
	}

	users, totalSize, err := r.userUC.ListUsers(ctx, &params)
	if err != nil {
		return nil, err
	}

	var respUsers []*adminv1.User
	if err := copier.Copy(&respUsers, users).Time2Timestamppb().Filter(req.SelectMask.GetPaths()).Error; err != nil {
		return nil, err
	}

	return &adminv1.ListUsersResponse{
		Users: respUsers,
		Page: &adminv1.PageResponse{
			TotalSize: totalSize,
		},
	}, nil
}
