package biz

import (
	"context"
	"mall-go/internal/data"
)

func BannerById(ctx context.Context, id int) (res interface{}, err error) {
	banner, err := data.GetBannerById(ctx, id)
	res = banner
	return
}
func BannerByName(ctx context.Context, name string) (res interface{}, err error) {
	banner, err := data.GetBannerByName(ctx, name)
	res = banner
	return
}
func ThemeNameWithSpu(ctx context.Context, name string) (res interface{}, err error) {
	theme, err := data.GetThemeWithSpu(ctx, name)
	res = theme

	return
}
func ThemeByNames(ctx context.Context, names []string) (res interface{}, err error) {
	theme, err := data.GetThemeByNames(ctx, names)
	res = theme
	return
}
