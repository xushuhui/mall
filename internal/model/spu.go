package model

import (
	"mall_go/global"
	"time"
)

type Spu struct {
	BaseModel
	Title          string `json:"title"`
	Subtitle       string `json:"subtitle"`
	CategoryID     int    `json:"category_id"`
	RootCategoryID int    `json:"root_category_id"`
	Price          string `json:"price"`
	Img            string `json:"img"`
	ForThemeImg    string `json:"for_theme_img"`
	Description    string `json:"description"`
	DiscountPrice  string `json:"discount_price"`
	Tags           string `json:"tags"`
	IsTest         bool   `json:"is_test"`
	Online         bool   `json:"online"`
	//ThemeList []Theme `gorm:"many2many:theme_spu"`
}

type SaleExplain struct {
	BaseModel
	Fixed     uint8  `json:"fixed"`
	Text      string `json:"text"`
	SpuID     int32  `json:"spu_id"`
	Index     uint   `json:"index"`
	ReplaceID uint   `json:"replace_id"`
}

type SpuDetailImg struct {
	ID         uint      `json:"id"`
	Img        string    `json:"img"`
	SpuID      uint      `json:"spu_id"`
	Index      uint      `json:"index"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	DeleteTime time.Time `json:"delete_time"`
}
type SpuImg struct {
	ID         uint      `json:"id"`
	Img        string    `json:"img"`
	SpuID      uint      `json:"spu_id"`
	DeleteTime time.Time `json:"delete_time"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
}
type SpuKey struct {
	ID        uint `json:"id"`
	SpuID     uint `json:"spu_id"`
	SpecKeyID uint `json:"spec_key_id"`
}
type SpuTag struct {
	ID    uint `json:"id"`
	SpuID uint `json:"spu_id"`
	TagID uint `json:"tag_id"`
}
type Tag struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UpdateTime  time.Time `json:"update_time"`
	DeleteTime  time.Time `json:"delete_time"`
	CreateTime  time.Time `json:"create_time"`
	Highlight   uint8     `json:"highlight"`
	Type        uint8     `json:"type"`
}

func GetSpu(id uint) (mode Spu, err error) {
	err = global.DBEngine.First(&mode, id).Error
	return
}
func GetSpuByCategory(id uint) (model []Spu, err error) {
	err = global.DBEngine.Find(&model, "category_id = ?", id).Error
	return
}
