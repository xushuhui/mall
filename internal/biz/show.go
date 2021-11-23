package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)
type Banner struct {
	Id int64
	Name string
	Title string
	Img string
	Description string
	Items []BannerItem
}
type BannerItem struct {
	ID       int64         `json:"id"`
	Img      string      `json:"img"`
	Keyword  string      `json:"keyword"`
	Type     int         `json:"type"`
	Name     string `json:"name"`
	BannerID int64         `json:"banner_id"`
}
type BannerRepo interface {
	GetBannerById(ctx context.Context, id int64) (b Banner,err error)
	
}


type ShowUsecase struct {
	bannerRepo BannerRepo
	log  *log.Helper
}
func NewShowUsecase(repo BannerRepo, logger log.Logger)*ShowUsecase{
	return &ShowUsecase{
		bannerRepo: repo,
		log: log.NewHelper(logger),
	}
}
func (uc *ShowUsecase)GetBannerById (ctx context.Context, id int64) (b Banner,err error){
	b,err = uc.bannerRepo.GetBannerById(ctx,id)
	return
}