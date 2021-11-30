package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)
type Theme struct {
	
}
type ThemeRepo interface {
	GetThemeByNames(ctx context.Context, names []string) (t []Theme, err error)
	GetThemeWithSpu(ctx context.Context, name string) (t Theme, err error)
}

type ThemeUsecase struct {
	repo ThemeRepo
	log  *log.Helper
}
func NewThemeUsecase(repo ThemeRepo, logger log.Logger) *ThemeUsecase {
	return &ThemeUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}