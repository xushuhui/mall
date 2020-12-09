package model

type Sku struct {
	BaseModel
	Price         float32 `json:"price"`
	DiscountPrice float32 `json:"discount_price"`
	Online        int     `json:"online"`
	Img           string  `json:"img"`
	Title         string  `json:"title"`
	SpuID         uint    `json:"spu_id"`

	Specs          string `json:"specs"`
	Code           string `json:"code"`
	Stock          uint   `json:"stock"`
	CategoryID     uint   `json:"category_id"`
	RootCategoryID uint   `json:"root_category_id"`
}
type SkuSpec struct {
	ID      uint `json:"id"`
	SpuID   uint `json:"spu_id"`
	SkuID   uint `json:"sku_id"`
	KeyID   uint `json:"key_id"`
	ValueID uint `json:"value_id"`
}
type SpecKey struct {
	BaseModel

	Name     string `json:"name"`
	Unit     string `json:"unit"`
	Standard uint8  `json:"standard"`

	Description string `json:"description"`
}
type SpecValue struct {
	BaseModel
	Value  string `json:"value"`
	SpecID uint   `json:"spec_id"`

	Extend string `json:"extend"`
}
