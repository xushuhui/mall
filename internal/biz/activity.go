package biz

import "github.com/go-kratos/kratos/v2/log"

type ActivityRepo interface {
}
type ActivityUsecase struct {
	repo ActivityRepo
	log  *log.Helper
}
func NewActivityUsecase(repo ActivityRepo, logger log.Logger) *ActivityUsecase {
	return &ActivityUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}