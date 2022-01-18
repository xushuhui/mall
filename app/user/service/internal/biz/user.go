package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall-go/api/user/service"
)

type User struct {
	Id int64
}

type UserRepo interface {
	CreateUser(ctx context.Context, in *service.CreateUserRequest) (userId int64, err error)

	GetUserIdentiy(ctx context.Context, identityType, identifier, credential string) (User, error)
	CreateUserIdentiy(ctx context.Context, userId int64, identityType, identifier, credential string) (err error)
}
type UserUsecase struct {
	repo   UserRepo
	jwtKey string
	log    *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
