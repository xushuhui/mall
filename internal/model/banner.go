package model

import (
	"fmt"
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
	ID             uint        `gorm:"primaryKey"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	Name           string      `json:"name"`
	Extend         interface{} `json:"extend"`
	EntranceImg    string      `json:"entrance_img"`
	InternalTopImg string      `json:"internal_top_img"`
	Online         bool        `json:"online"`
	TitleImg       string      `json:"title_img"`
	TplName        string      `json:"tpl_name"`
	SpuList        []Spu       ` gorm:"many2many:theme_spu;"`
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

type Car struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	CityList []City `gorm:"many2many:car_city"`
}
type City struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}

func GetBannerById(id uint) (model Banner, e error) {
	e = global.DBEngine.Preload("Items").First(&model, id).Error
	return
}
func GetBannerByName(name string) (model Banner, e error) {
	e = global.DBEngine.Preload("Items").Where("name = ?", name).First(&model).Error
	return
}
func GetThemeWithSpu(name string) (e error) {
	//var s SysAuthority
	//e = global.DBEngine.Preload("Menus").First(&s, "id = ?", 1).Error
	//
	//fmt.Println("e----",e,s.Menus)
	var c Car
	e = global.DBEngine.First(&c, "id=?", 1).Error
	fmt.Println("e---", e, c.CityList)
	var t Theme
	e = global.DBEngine.First(&t, "id=?", 1).Error

	fmt.Println("e---", e, t.Title)
	//fmt.Println(spuList)
	return
}
func GetThemeByNames(names string) (model []Theme, e error) {

	return
}
