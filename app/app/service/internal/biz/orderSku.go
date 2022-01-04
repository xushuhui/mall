package biz

import "mall-go/api/mall"

type OrderSku struct {
	id          int64
	spuId       int64
	finalPrice  float64
	singlePrice float64
	specValues  []string
	count       int32
	img         string
	title       string
}

func NewOrderSku(sku Sku, skuInfoDTO *mall.SkuInfo) OrderSku {
	return OrderSku{
		id:          sku.Id,
		spuId:       sku.SpuID,
		singlePrice: sku.GetActualPrice(),
		finalPrice:  sku.GetActualPrice() * float64(skuInfoDTO.Count),
		count:       skuInfoDTO.Count,
		img:         sku.Img,
		title:       sku.Title,
		specValues:  sku.getSpecValueList(),
	}
}
