package biz

import (
	"errors"
	"time"

	"mall-go/pkg/enum"
	"mall-go/pkg/utils"
)

type CouponChecker struct {
	Coupon
}

func NewCouponChecker(coupon Coupon) *CouponChecker {
	return &CouponChecker{
		Coupon: coupon,
	}
}

func (c *CouponChecker) IsOk() (err error) {
	in := utils.IsInTime(time.Now(), c.Coupon.StartTime, c.Coupon.EndTime)
	if !in {
		err = errors.New("invalid params")
		return
	}
	return
}

func (c *CouponChecker) CanBeUsed(skuOrderList []SkuOrder, serverTotalPrice float64) (err error) {
	var orderCategoryPrice float64
	var cids []int64

	if c.Coupon.WholeStore == 1 {
		orderCategoryPrice = serverTotalPrice
	} else {
		for _, v := range c.Coupon.Category {
			cids = append(cids, v.Id)
		}
		orderCategoryPrice = getSumByCategoryList(skuOrderList, cids)
	}
	err = c.couponCanBeUsed(orderCategoryPrice)

	return
}

func (c *CouponChecker) couponCanBeUsed(orderCategoryPrice float64) (err error) {
	switch c.Coupon.Type {
	case enum.FULL_MINUS:
	case enum.FULL_OFF:
		if c.Coupon.FullMoney > orderCategoryPrice {
			err = errors.New("invalid params")

			return
		}
	case enum.NO_THRESHOLD_MINUS:
		return
	default:
		err = errors.New("invalid params")

		return
	}
	return
}

func getSumByCategoryList(skuOrderList []SkuOrder, cids []int64) (sum float64) {
	for _, cid := range cids {
		sum = sum + getSumByCategory(skuOrderList, cid)
	}
	return
}

func getSumByCategory(skuOrderList []SkuOrder, cid int64) (sum float64) {
	for _, v := range skuOrderList {
		if v.CategoryId == cid {
			sum = sum + v.GetTotalPrice()
		}
	}
	return
}

func (c *CouponChecker) FinalTotalPriceIsOk(orderFinalTotalPrice float64, serverTotalPrice float64) (err error) {
	var serverFinalTotalPrice float64
	switch c.Coupon.Type {
	case enum.FULL_MINUS:
	case enum.NO_THRESHOLD_MINUS:
		serverFinalTotalPrice = serverTotalPrice - c.Coupon.Minus
		if serverFinalTotalPrice <= 0 {
			err = errors.New("invalid params")
			return
		}

	case enum.FULL_OFF:
		// todo
		serverFinalTotalPrice = serverTotalPrice * c.Coupon.Rate

	default:
		err = errors.New("invalid params")

		return
	}
	if serverFinalTotalPrice != orderFinalTotalPrice {
		err = errors.New("invalid params")

		return
	}
	return
}
