package data

import (
	"mall-go/internal/data/model"
	"mall-go/internal/request"
)

type Sku struct {
	*model.Sku
}
type OrderSku struct {
	id          int
	spuId       int
	finalPrice  float64
	singlePrice float64
	specValues  []string
	count       int
	img         string
	title       string
}

func NewOrderSku(sku Sku, skuInfoDTO request.SkuInfo) OrderSku {
	return OrderSku{
		id:          sku.ID,
		spuId:       sku.SpuID,
		singlePrice: sku.GetActualPrice(),
		finalPrice:  sku.GetActualPrice() * float64(skuInfoDTO.Count),
		count:       skuInfoDTO.Count,
		img:         sku.Img,
		title:       sku.Title,
		specValues:  sku.getSpecValueList(),
	}
}
func (sku Sku) GetActualPrice() float64 {
	if sku.DiscountPrice != 0 {
		return sku.DiscountPrice
	}
	return sku.Price
}
func (sku Sku) getSpecValueList() (val []string) {
	for _, v := range sku.Specs {
		val = append(val, v.Value)
	}
	return
}

type SkuOrder struct {
}
