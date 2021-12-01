package data

import (
	"context"
	"mall-go/internal/biz"
	"mall-go/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
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
func GetCouponById(ctx context.Context, id int) (c *model.Coupon, err error) {

	return
}
func GetUserCoupon(ctx context.Context, userId, couponId int) (c *model.UserCoupon, err error) {

	return
}
