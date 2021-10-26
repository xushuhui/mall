package biz

import (
	"mall-go/internal/data"
)

func BannerById(id int) (banner data.Banner, err error) {
	banner, err = data.GetBannerById(id)
	return
}
func BannerByName(name string) (banner data.Banner, err error) {
	banner, err = data.GetBannerByName(name)
	return
}
func ThemeNameWithSpu(name string) (theme data.Theme, err error) {
	theme, err = data.GetThemeWithSpu(name)

	return
}
func ThemeByNames(names []string) (theme []data.Theme, err error) {
	theme, err = data.GetThemeByNames(names)
	return
}
