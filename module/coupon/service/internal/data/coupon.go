package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mall-go/module/coupon/service/internal/biz"
)

type couponRepo struct {
	data *Data
	log  *log.Helper
}

func NewCouponRepo(data *Data, logger log.Logger) biz.CouponRepo {
	return &couponRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
