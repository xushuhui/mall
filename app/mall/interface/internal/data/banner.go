package data

import (
	"context"
	"mall-go/api/mall"
	"mall-go/app/mall/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type bannerRepo struct {
	data *Data
	log  *log.Helper
}

func NewBannerRepo(data *Data, logger log.Logger) biz.BannerRepo {
	return &bannerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *bannerRepo) GetBannerById(ctx context.Context, id int64) (b *mall.Banner, err error) {
	b, err = r.data.ac.GetBannerById(ctx, &mall.BannerByIdRequest{Id: id})

	return

}
func (r *bannerRepo) GetBannerByName(ctx context.Context, name string) (b biz.Banner, err error) {
	return
}
