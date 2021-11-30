package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Banner struct {
	Id          int64
	Name        string
	Title       string
	Img         string
	Description string
	Items       []BannerItem
}
type BannerItem struct {
	ID       int64  `json:"id"`
	Img      string `json:"img"`
	Keyword  string `json:"keyword"`
	Type     int    `json:"type"`
	Name     string `json:"name"`
	BannerID int64  `json:"banner_id"`
}
type BannerRepo interface {
	GetBannerById(ctx context.Context, id int64) (b Banner, err error)
	GetBannerByName(ctx context.Context, name string) (b Banner, err error)
}

type BannerUsecase struct {
	repo BannerRepo
	log  *log.Helper
}

func NewBannerUsecase(repo BannerRepo, logger log.Logger) *BannerUsecase {
	return &BannerUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
func (uc *BannerUsecase) GetBannerById(ctx context.Context, id int64) (b Banner, err error) {
	b, err = uc.repo.GetBannerById(ctx, id)
	return
}
func (uc *BannerUsecase) GetBannerByName(ctx context.Context, name string) (b Banner, err error) {
	b, err = uc.repo.GetBannerByName(ctx, name)
	return
}