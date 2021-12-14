package biz

import "github.com/go-kratos/kratos/v2/log"
type User struct {
	
}
type UserRepo interface {
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
