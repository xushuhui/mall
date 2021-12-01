package biz

import "github.com/go-kratos/kratos/v2/log"

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
