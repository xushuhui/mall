package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
}


type UserRepo interface {
	GenerateToken(ctx context.Context, account string) (token string, err error)
	VerifyToken(ctx context.Context, token string) (isValid bool, err error)
}
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (u *UserUsecase) GenerateToken(ctx context.Context, account string) (token string, err error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserUsecase) VerifyToken(ctx context.Context, token string) (isValid bool, err error) {
	panic("not implemented") // TODO: Implement
}
