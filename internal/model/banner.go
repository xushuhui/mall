package model

import (
	"gorm.io/gorm"
	"mall_go/global"
	"time"
)

type BaseModel struct {
	ID         uint `gorm:"primaryKey"`
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime gorm.DeletedAt `gorm:"index"`
}
type Banner struct {
	BaseModel
	Name        string `json:"name"`
	Title       string `json:"title"`
	Img         string `json:"img"`
	Description string `json:"description"`

	Items []BannerItem `json:"items"`
}
type BannerItem struct {
	ID       uint   `json:"id"`
	Img      string `json:"img"`
	Keyword  string `json:"keyword"`
	Name     string `json:"name"`
	BannerID uint   `json:"banner_id"`
	Type     int    `json:"type"`
}

func GetBannerById(id uint) (model Banner, e error) {
	e = global.DBEngine.Preload("Items").First(&model, id).Error
	return
}
func GetBannerByName(name string) (model Banner, e error) {
	e = global.DBEngine.Preload("Items").Where("name = ?", name).First(&model).Error
	return
}
