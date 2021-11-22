package biz

import (
	"context"
	"mall-go/internal/conf"
	"mall-go/internal/data"
	"mall-go/internal/logic"
	"mall-go/internal/request"

	"github.com/xushuhui/goal/core"
)

func PlaceOrder(userId int, req request.PlaceOrder, oc logic.OrderChecker) (orderId int, err error) {
	return
}

func OrderIsOk(ctx context.Context, userId int, orderDTO request.PlaceOrder) (oc logic.OrderChecker, err error) {
	if orderDTO.FinalTotalPrice <= 0 {
		err = core.ParamsError(core.InvalidParams)
		return
	}
	//
	var skuIds []int
	for _, v := range orderDTO.SkuInfoList {
		skuIds = append(skuIds, v.Id)
	}
	skuList, err := data.GetSkuListByIds(ctx, skuIds)
	if err != nil {
		return
	}
	couponId := orderDTO.CouponId
	var couponChecker logic.CouponChecker
	if couponId > 0 {
		coupon, err := data.GetCouponById(ctx, couponId)
		if err != nil {
			return oc, err
		}
		_, err = data.GetUserCoupon(ctx, userId, couponId)
		if err != nil {
			return oc, err
		}
		couponChecker = *logic.NewCouponChecker(coupon)

	}
	maxSkuLimit := conf.GlobalConfig.App.MaxSkuLimit
	oc = *logic.NewOrderChecker(orderDTO, skuList, couponChecker, maxSkuLimit)
	err = oc.IsOk()

	return
}
