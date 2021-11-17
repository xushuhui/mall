package data

import (
	"context"
	"mall-go/internal/data/model"
	"mall-go/internal/data/model/banner"
)

type Banner struct {
}
type Theme struct {
}

func GetBannerById(ctx context.Context, id int) (b *model.Banner, err error) {

	b, err = GetDB().Banner.Query().Where(banner.ID(id)).First(ctx)
	return
}
func GetBannerByName(ctx context.Context, name string) (b *model.Banner, err error) {
	b, err = GetDB().Banner.Query().Where(banner.Name(name)).First(ctx)
	return
}

func GetAccountUser(phone string) (user *model.User, err error) {
	return
}
