package service

import (
	"mall-go/api/mall"
	"mall-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
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
