package services

import (
	"mall_go/internal/model"
)

func BannerById(id uint) (banner model.Banner, e error) {
	banner, e = model.GetBannerById(id)
	return
}
func BannerByName(name string) (banner model.Banner, e error) {
	banner, e = model.GetBannerByName(name)
	return
}
