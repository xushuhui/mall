package model

import (
	"gorm.io/gorm"

	"time"
)

type BaseModel struct {
	ID         uint      `gorm:"primaryKey"`
	CreateTime time.Time `gorm:"-",json:"create_time"`
	UpdateTime time.Time `gorm:"-"`
	DeleteTime gorm.DeletedAt
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
type Theme struct {
	BaseModel
	Title          string `json:"title"`
	Description    string `json:"description"`
	Name           string `json:"name"`
	Extend         string `json:"extend"`
	EntranceImg    string `json:"entrance_img"`
	InternalTopImg string `json:"internal_top_img"`
	Online         bool   `json:"online"`
	TitleImg       string `json:"title_img"`
	TplName        string `json:"tpl_name"`
	SpuList        []Spu  `gorm:"many2many:theme_spu"`
}

func GetBannerById(id uint) (model Banner, e error) {

	return
}
func GetBannerByName(name string) (model Banner, e error) {

	return
}
func GetThemeWithSpu(name string) (model Theme, e error) {

	return
}
func GetThemeByNames(names []string) (model []Theme, e error) {

	return
}
