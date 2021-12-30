package biz

import (
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

	Minus float64 `json:"minus"`

	Rate float64 `json:"rate"`
}
type Category struct {
	Id int64
}
type CouponRepo interface {
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
