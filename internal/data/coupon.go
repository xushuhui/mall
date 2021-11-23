package data

import (
	"context"
	"mall-go/internal/data/model"
)

type Coupon struct {
	*model.Coupon
}

func GetCouponById(ctx context.Context, id int) (c *model.Coupon, err error) {

	return
}
func GetUserCoupon(ctx context.Context, userId, couponId int) (c *model.UserCoupon, err error) {

	return
}
