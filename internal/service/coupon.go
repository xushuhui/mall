package service

import (
	"mall-go/api/mall"
	"mall-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type CouponService struct {
	mall.UnimplementedCouponServer
	cu  *biz.CouponUsecase
	log *log.Helper
}
func NewCouponService(cu *biz.CouponUsecase, logger log.Logger) *CouponService {
	return &CouponService{cu: cu, log: log.NewHelper(logger)}
}