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
	ID             uint   `gorm:"primaryKey"`
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

	//CityList []City `gorm:"many2many:car_city"`
}

type Spu struct {
	ID             uint   `gorm:"primaryKey"`
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

//type Car struct {
//	ID       uint   `gorm:"primaryKey"`
//	Name     string `json:"name"`
//	CityList []City `gorm:"many2many:car_city"`
//}
//type City struct {
//	ID   uint   `gorm:"primaryKey"`
//	Name string `json:"name"`
//}

func GetBannerById(id uint) (model Banner, e error) {
	e = global.DBEngine.Preload("Items").First(&model, id).Error
	return
}
func GetBannerByName(name string) (model Banner, e error) {
	e = global.DBEngine.Preload("Items").Where("name = ?", name).First(&model).Error
	return
}
func GetThemeWithSpu(name string) (model Theme, e error) {
	e = global.DBEngine.Preload("SpuList").First(&model, "name=?", name).Error

	return
}
func GetThemeByNames(names []string) (model []Theme, e error) {
	e = global.DBEngine.Find(&model, "name in ?", names).Error
	return
}
