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

func (s *UserService) CreateUser(ctx context.Context, in *CreateUserRequest) (*UserVO, error) {
	return &UserVO{}, nil
}

func (s *UserService) GetUser(ctx context.Context, in *IdRequest) (*UserVO, error) {
	return &UserVO{}, nil
}

func (s *UserService) GetUserIdentity(ctx context.Context, in *UserIdentityRequest) (*UserVO, error) {
	return &UserVO{}, nil
}

func (s *UserService) CreateUserIdentity(ctx context.Context, in *UserIdentityRequest) (*UserVO, error) {
	return &UserVO{}, nil
}

func (s *UserService) ListUser(ctx context.Context, in *IdsRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
