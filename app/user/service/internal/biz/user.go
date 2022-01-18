package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id int64
}

type UserRepo interface {
	CreateUser(ctx context.Context) (err error)
	GetUserByAccount(ctx context.Context, account string)
	GetUserIdentiy(ctx context.Context, identifier, credential, identityType string) (User, error)
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
