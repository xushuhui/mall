package data

import (
	"context"
	"mall-go/api/mall"
	"mall-go/internal/data/model"
	"mall-go/internal/data/model/sku"
)

type Sku struct {
	*model.Sku
}

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
	ActualPrice float64
	Count       int32
	CategoryId  int64
}

func NewSkuOrder(sku Sku, skuInfoDTO *mall.SkuInfo) SkuOrder {
	return SkuOrder{
		ActualPrice: sku.GetActualPrice(),
		Count:       skuInfoDTO.Count,
		CategoryId:  sku.CategoryID,
	}
}

func (s SkuOrder) GetTotalPrice() float64 {
	return s.ActualPrice * float64(s.Count)
}
func GetSkuListByIds(ctx context.Context, ids []int64) (skus []*model.Sku, err error) {
	skus, err = GetDB().Sku.Query().Where(sku.IDIn(ids...)).All(ctx)
	return
}
