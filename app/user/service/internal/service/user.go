package service

import (
	"context"

	"mall-go/app/user/service/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	UnimplementedUserServer
	uu *biz.UserUsecase
}

func NewUserService(uu *biz.UserUsecase) *UserService {
	return &UserService{
		uu: uu,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserVO, error) {
	return &UserVO{}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *IdRequest) (*UserVO, error) {
	return &UserVO{}, nil
}

func (s *UserService) GetUserIdentiy(ctx context.Context, req *UserIdentiyRequest) (*UserVO, error) {
	return &UserVO{}, nil
}

func (s *UserService) CreateUserIdentiy(ctx context.Context, req *UserIdentiyRequest) (*UserVO, error) {
	return &UserVO{}, nil
}

func (s *UserService) ListUser(ctx context.Context, req *IdsRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
