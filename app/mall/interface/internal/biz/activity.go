package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Activity struct {
	Id          int64
	Title       string
	EntranceImg string
	StartTime   time.Time
	EndTime     time.Time
	Remark      string
	Online      int32
}

type ActivityCoupon struct {
	Activity
	Coupons []Coupon
}

type ActivityRepo interface {
	GetActivityByName(ctx context.Context, name string) (c Activity, err error)
	GetActivityWithCoupon(ctx context.Context, name string) (c ActivityCoupon, err error)
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

func (ac *ActivityUsecase) GetActivityByName(ctx context.Context, name string) (c Activity, err error) {
	c, err = ac.repo.GetActivityByName(ctx, name)
	return
}

func (ac *ActivityUsecase) GetActivityWithCoupon(ctx context.Context, name string) (c ActivityCoupon, err error) {
	c, err = ac.repo.GetActivityWithCoupon(ctx, name)
	return
}
