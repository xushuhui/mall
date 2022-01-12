package service

import (
	"context"
	"mall-go/api/user/service"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	service.UnimplementedUserServer
	uu  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uu *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uu:  uu,
		log: log.NewHelper(logger),
	}
}
func (s *UserService) Login(ctx context.Context, in *service.LoginRequest) (out *service.LoginReply, err error) {

	return s.uu.Login(ctx, in.Account)
}
func (s *UserService) VerifyToken(ctx context.Context, in *service.VerifyTokenRequest) (out *service.VerifyTokenReply, err error) {
	return
}
func (s *UserService) UpdateInfo(ctx context.Context, in *service.UpdateInfoRequest) (out *emptypb.Empty, err error) {
	return
}
func (s *UserService) ListAddress(ctx context.Context, in *service.LoginRequest) (out *service.LoginReply, err error) {
	return
}
func (s *UserService) GetAddress(ctx context.Context, in *service.LoginRequest) (out *service.LoginReply, err error) {
	return
}
func (s *UserService) SetDefaultAddr(ctx context.Context, in *service.LoginRequest) (out *service.LoginReply, err error) {
	return
}
