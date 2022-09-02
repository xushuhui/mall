package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id           int64
	Nickname     string
	Identifier   string
	IdentityType string
	Credential   string
}

type UserRepo interface {
	CreateUser(ctx context.Context, in *User) (userId int64, err error)

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
