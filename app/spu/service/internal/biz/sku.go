package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Spec struct {
	Value string `json:"value"`
}
type Sku struct {
	Id             int64
	DiscountPrice  float64 `json:"discount_price"`
	Price          float64 `json:"price"`
	Specs          []*Spec `json:"specs"`
	Online         int8    `json:"online"`
	Img            string  `json:"img"`
	Title          string  `json:"title"`
	SpuID          int64   `json:"spu_id"`
	Code           string  `json:"code"`
	Stock          int     `json:"stock"`
	CategoryID     int64   `json:"category_id"`
	RootCategoryID int64   `json:"root_category_id"`
}

func (sku *Sku) GetActualPrice() float64 {
	if sku.DiscountPrice != 0 {
		return sku.DiscountPrice
	}
	return sku.Price
}

func (sku *Sku) getSpecValueList() (val []string) {
	for _, v := range sku.Specs {
		val = append(val, v.Value)
	}
	return
}

type SkuRepo interface {
	GetSkuById(ctx context.Context, id int64) (o Sku, err error)
}
type SkuUsecase struct {
	repo SkuRepo
	log  *log.Helper
}

func NewSkuUsecase(repo SkuRepo, logger log.Logger) *SkuUsecase {
	return &SkuUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
func (uc *SkuUsecase) GetOrderById(ctx context.Context, id int64) (o Sku, err error) {
	o, err = uc.repo.GetSkuById(ctx, id)
	return
}
