package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)
type Activity struct {
	
}
type ActivityCoupon struct {
	
}
type ActivityRepo interface {
	GetActivityByName(ctx context.Context,name string) (a Activity, err error)
	GetActivityWithCoupon(ctx context.Context,name string) (a ActivityCoupon, err error)
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


