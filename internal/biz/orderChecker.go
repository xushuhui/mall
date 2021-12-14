package biz

import (
	"mall-go/api/mall"

	"github.com/xushuhui/goal/core"
)

type (
	OrderChecker struct {
		orderDTO      *mall.PlaceOrderRequest
		serverSkuList []Sku
		couponChecker CouponChecker
		maxSkuLimit   int
		orderSkuList  []OrderSku
	}
)

func NewOrderChecker(req *mall.PlaceOrderRequest, serverSkuList []Sku, checker CouponChecker, maxSkuLimit int) *OrderChecker {
	return &OrderChecker{
		orderDTO:      req,
		serverSkuList: serverSkuList,
		couponChecker: checker,
		maxSkuLimit:   maxSkuLimit,
	}
}

func (o *OrderChecker) GetLeaderImg() string {
	return o.serverSkuList[0].Img
}

func (o *OrderChecker) GetLeaderTitle() string {
	return o.serverSkuList[0].Title
}

func (o *OrderChecker) GetTotalCount() (totalCount int32) {
	for _, v := range o.orderDTO.SkuInfoList {
		totalCount = totalCount + v.Count
	}
	return
}

func (o *OrderChecker) IsOk() (err error) {
	var serverTotalPrice float64
	var skuOrderList []SkuOrder
	err = skuNotOnSale(len(o.orderDTO.SkuInfoList), len(o.serverSkuList))
	if err != nil {
		return
	}
	for i := 0; i < len(o.serverSkuList); i++ {
		sku := o.serverSkuList[i]

		skuInfoDTO := o.orderDTO.SkuInfoList[i]
		err = containsSoldOutSku(sku)
		if err != nil {
			return err
		}
		err = beyondSkuStock(sku, skuInfoDTO)
		if err != nil {
			return err
		}
		err = o.beyondMaxSkuLimit(skuInfoDTO)
		if err != nil {
			return err
		}
		price, err := calculateSkuOrderPrice(skuInfoDTO, sku)
		if err != nil {
			return err
		}
		serverTotalPrice = serverTotalPrice + price
		o.orderSkuList = append(o.orderSkuList, NewOrderSku(sku, skuInfoDTO))

		skuOrderList = append(skuOrderList, NewSkuOrder(sku, skuInfoDTO))
	}
	err = totalPriceIsOk(o.orderDTO.TotalPrice, serverTotalPrice)
	if err != nil {
		return err
	}
	if o.couponChecker != (CouponChecker{}) {
		err = o.couponChecker.IsOk()
		if err != nil {
			return err
		}
		err = o.couponChecker.CanBeUsed(skuOrderList, serverTotalPrice)
		if err != nil {
			return err
		}
		err = o.couponChecker.FinalTotalPriceIsOk(o.orderDTO.FinalTotalPrice, serverTotalPrice)
		if err != nil {
			return err
		}
	}
	return
}

func totalPriceIsOk(orderTotalPrice float64, serverTotalPrice float64) (err error) {
	if orderTotalPrice != serverTotalPrice {
		err = core.ParamsError(core.InvalidParams)
		return
	}
	return
}

func calculateSkuOrderPrice(skuInfoDTO *mall.SkuInfo, sku Sku) (price float64, err error) {
	if skuInfoDTO.Count <= 0 {
		err = core.ParamsError(core.InvalidParams)
		return
	}
	price = sku.GetActualPrice() * float64(skuInfoDTO.Count)
	return
}

func skuNotOnSale(count1, count2 int) (err error) {
	if count1 != count2 {
		err = core.ParamsError(core.InvalidParams)
	}
	return
}

func containsSoldOutSku(sku Sku) (err error) {
	if sku.Stock == 0 {
		err = core.ParamsError(core.InvalidParams)
	}
	return
}

func beyondSkuStock(sku Sku, skuInfoDTO *mall.SkuInfo) (err error) {
	if sku.Stock < int(skuInfoDTO.Count) {
		err = core.ParamsError(core.InvalidParams)
	}
	return
}

func (o *OrderChecker) beyondMaxSkuLimit(skuInfoDTO *mall.SkuInfo) (err error) {
	if int(skuInfoDTO.Count) > o.maxSkuLimit {
		err = core.ParamsError(core.InvalidParams)
	}
	return
}
