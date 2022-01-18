package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Coupon struct {
	Id         int64
	StartTime  time.Time
	EndTime    time.Time
	WholeStore int8       `json:"whole_store"`
	Category   []Category `json:"category"`
	Type       int        `json:"type"`
	FullMoney  float64    `json:"full_money"`
	Title      string     `json:"title"`
	Minus      float64    `json:"minus"`
	Rate       float64    `json:"rate"`
}

type CouponRepo interface {
	CreateUserCoupon(ctx context.Context, couponId, userId int64) (err error)
	ListUserCoupon(ctx context.Context, userId int64) (list []Coupon, err error)
	CreateCoupon(ctx context.Context) (err error)
	ListCoupon(ctx context.Context) (list []Coupon, err error)
}
type CouponUsecase struct {
	repo CouponRepo
	log  *log.Helper
}

func NewCouponUsecase(repo CouponRepo, logger log.Logger) *CouponUsecase {
	return &CouponUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *CouponUsecase) CreateUserCoupon(ctx context.Context, couponId, userId int64) (err error) {
	panic("implement me")
}

func (uc *CouponUsecase) ListUserCoupon(ctx context.Context, userId int64) (list []Coupon, err error) {
	panic("implement me")
}

func (uc *CouponUsecase) CreateCoupon(ctx context.Context) (err error) {
	panic("implement me")
}

func (uc *CouponUsecase) ListCoupon(ctx context.Context) (list []Coupon, err error) {
	panic("implement me")
}
