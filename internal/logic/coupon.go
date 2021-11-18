package logic

import (
	"mall-go/internal/data"
	"mall-go/pkg/enum"
	"mall-go/pkg/utils"
	"time"

	"github.com/xushuhui/goal/core"
)

type CouponChecker struct {
	Coupon   data.Coupon
	discount float64
}

func NewCouponChecker(coupon data.Coupon, discount float64) *CouponChecker {
	return &CouponChecker{
		Coupon:   coupon,
		discount: discount,
	}
}

func (c *CouponChecker) IsOk() (err error) {
	in := utils.IsInTime(time.Now(), c.Coupon.StartTime, c.Coupon.EndTime)
	if !in {
		err = core.ParamsError(core.InvalidParams)
		return
	}
	return
}

func (c *CouponChecker) CanBeUsed(skuOrderList []data.SkuOrder, serverTotalPrice float64) (err error) {
	var orderCategoryPrice float64
	var cids []int

	if c.Coupon.WholeStore == 1 {
		orderCategoryPrice = serverTotalPrice
	} else {
		for _, v := range c.Coupon.Edges.Category {
			cids = append(cids, v.ID)
		}
		orderCategoryPrice = c.getSumByCategoryList(skuOrderList, cids)
	}
	err = c.couponCanBeUsed(orderCategoryPrice)

	return
}

func (c *CouponChecker) couponCanBeUsed(orderCategoryPrice float64) (err error) {
	switch c.Coupon.Type {
	case enum.FULL_MINUS:
	case enum.NO_THRESHOLD_MINUS:

	}
	return
}

func (c *CouponChecker) getSumByCategoryList(skuOrderList []data.SkuOrder, cids []int) (sum float64) {
	for _, cid := range cids {
		sum = sum + c.getSumByCategory(skuOrderList, cid)
	}
	return
}

func (c *CouponChecker) getSumByCategory(skuOrderList []data.SkuOrder, cid int) (sum float64) {
	for _, v := range skuOrderList {
		if v.CategoryId == cid {
			sum = sum + v.GetTotalPrice()
		}
	}
	return
}

//private void couponCanBeUsed(BigDecimal orderCategoryPrice) {
//switch (CouponType.toType(this.coupon.getType())){
//case FULL_OFF:
//case FULL_MINUS:
//int compare = this.coupon.getFullMoney().compareTo(orderCategoryPrice);
//if(compare > 0){
//throw new ParameterException(40008);
//}
//break;
//case NO_THRESHOLD_MINUS:
//break;
//default:
//throw new ParameterException(40009);
//}
//}
//

//

func (c *CouponChecker) FinalTotalPriceIsOk(orderFinalTotalPrice float64, serverTotalPrice float64) (err error) {
	return
}
