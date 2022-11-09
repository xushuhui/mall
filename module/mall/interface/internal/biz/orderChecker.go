package biz

import (
	"errors"
	"mall-go/module/mall/interface/internal/service"
)

type (
	OrderChecker struct {
		orderDTO      *service.PlaceOrderRequest
		serverSkuList []Sku
		couponChecker *CouponChecker
		maxSkuLimit   int
		orderSkuList  []OrderSku
	}
)

func NewOrderChecker(req *service.PlaceOrderRequest, serverSkuList []Sku, checker *CouponChecker, maxSkuLimit int) *OrderChecker {
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
	if o.couponChecker != nil {
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
		err = errors.New("invalid params")
		return
	}
	return
}

func calculateSkuOrderPrice(skuInfoDTO *service.SkuInfo, sku Sku) (price float64, err error) {
	if skuInfoDTO.Count <= 0 {
		err = errors.New("invalid params")
		return
	}
	price = sku.GetActualPrice() * float64(skuInfoDTO.Count)
	return
}

func skuNotOnSale(count1, count2 int) (err error) {
	if count1 != count2 {
		err = errors.New("invalid params")
	}
	return
}

func containsSoldOutSku(sku Sku) (err error) {
	if sku.Stock == 0 {
		err = errors.New("invalid params")
	}
	return
}

func beyondSkuStock(sku Sku, skuInfoDTO *service.SkuInfo) (err error) {
	if sku.Stock < int(skuInfoDTO.Count) {
		err = errors.New("invalid params")
	}
	return
}

func (o *OrderChecker) beyondMaxSkuLimit(skuInfoDTO *service.SkuInfo) (err error) {
	if int(skuInfoDTO.Count) > o.maxSkuLimit {
		err = errors.New("invalid params")
	}
	return
}
