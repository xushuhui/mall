package biz

import "mall-go/module/mall/interface/internal/service"

type SkuOrder struct {
	ActualPrice float64
	Count       int32
	CategoryId  int64
}

func NewSkuOrder(sku Sku, skuInfoDTO *service.SkuInfo) SkuOrder {
	return SkuOrder{
		ActualPrice: sku.GetActualPrice(),
		Count:       skuInfoDTO.Count,
		CategoryId:  sku.CategoryID,
	}
}

func (s SkuOrder) GetTotalPrice() float64 {
	return s.ActualPrice * float64(s.Count)
}
