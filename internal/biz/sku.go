package biz

import "github.com/go-kratos/kratos/v2/log"

type Spec struct {
	Value string `json:"value"`
}
type Sku struct {
	Id int64 
	
	DiscountPrice float64

	Price float64

	Specs []*Spec
	Online int8 `json:"online,omitempty"`

	Img string `json:"img,omitempty"`

	Title string `json:"title,omitempty"`
	
	SpuID int64 `json:"spu_id,omitempty"`
	
	Code string `json:"code,omitempty"`

	Stock int `json:"stock,omitempty"`

	CategoryID int64 `json:"category_id,omitempty"`
	
	RootCategoryID int64 `json:"root_category_id,omitempty"`
}
type SkuRepo interface {
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
