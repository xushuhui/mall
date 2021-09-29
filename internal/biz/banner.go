package biz

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
func ThemeNameWithSpu(name string) (theme model.Theme, e error) {
	theme, e = model.GetThemeWithSpu(name)

	return
}
func ThemeByNames(names []string) (theme []model.Theme, e error) {
	theme, e = model.GetThemeByNames(names)
	return
}
