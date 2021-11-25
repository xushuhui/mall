package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)
type Theme struct {
	
}
type ThemeRepo interface {
	GetBannerById(ctx context.Context, id int64) (b Theme, err error)
}

type ThemeUsecase struct {
	repo ThemeRepo
	log  *log.Helper
}
func NewThemeUsecase(){
	
}