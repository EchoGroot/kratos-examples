package service

import (
	"context"

	adminv1 "github.com/EchoGroot/kratos-examples/user-manage/api/user-manage/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LivezService struct {
	adminv1.UnimplementedLivezServiceServer
}

func NewLivezService() *LivezService {
	return &LivezService{}
}

func (s *LivezService) Livez(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
