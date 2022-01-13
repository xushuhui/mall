package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Activity struct {
	Id          int64
	Title       string
	EntranceImg string
	StartTime   int64
	EndTime     int64
	Remark      string
	Online      int32
}

type ActivityCoupon struct {
	Activity
	Coupons []Coupon
}

type Coupon struct {
	Id          int64
	Title       string
	StartTime   int64
	EndTime     int64
	Description string
	FullMoney   float64
	Rate        float64
	Minus       float64
	Type        int32
	Remark      string
	WholeStore  bool
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
