package service

import (
	"context"
	"mall-go/api/mall"
	"mall-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	mall.UnimplementedUserServer
	uu  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uu *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uu:  uu,
		log: log.NewHelper(logger),
	}
}
func (s *UserService) GenerateToken(ctx context.Context, in *mall.GenerateTokenRequest) (out *mall.GenerateTokenReply, err error) {
	return
}
func (s *UserService) VerifyToken(ctx context.Context, in *mall.VerifyTokenRequest) (out *mall.VerifyTokenReply, err error) {
	return
}
func (s *UserService) UpdateInfo(ctx context.Context, in *mall.UpdateInfoRequest) (out *emptypb.Empty, err error) {
	return
}
