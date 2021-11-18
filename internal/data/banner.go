package data

import (
	"context"
	"mall-go/internal/data/model"
)

type Banner struct {
}
type Theme struct {
}

func GetBannerById(ctx context.Context, id int) (b *model.Banner, err error) {

	return
}
func GetBannerByName(ctx context.Context, name string) (b *model.Banner, err error) {
	return
}

func GetAccountUser(phone string) (user *model.User, err error) {
	return
}
