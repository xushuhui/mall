package data

import (
	"context"
	"mall-go/internal/data/model"
	"mall-go/internal/data/model/coupon"
	"mall-go/internal/data/model/usercoupon"
)

type Coupon struct {
	*model.Coupon
}
func GetCouponById(ctx context.Context,id int)(c *model.Coupon,err error){
	c,err = GetDB().Coupon.Query().Where(coupon.ID(id)).First(ctx)
	return
}
func GetUserCoupon(ctx context.Context,userId,couponId int)(c *model.UserCoupon,err error){
	c,err = GetDB().UserCoupon.Query().Where(usercoupon.CouponID(couponId),usercoupon.UserID(userId)).First(ctx)
	return
}