package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	app "mall-go/api/app/service"
	"mall-go/app/mall/admin/internal/biz"
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
func (r *bannerRepo) GetBannerById(ctx context.Context, id int64) (b biz.Banner, err error) {
	po, err := r.data.ac.GetBannerById(ctx, &app.IdRequest{Id: id})
	if err != nil {
		return
	}
	var items []biz.BannerItem
	for _, v := range po.Items {

		items = append(items, biz.BannerItem{
			ID:       v.Id,
			Name:     v.Name,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     int(v.Type),
			BannerId: v.BannerId,
		})
	}
	return biz.Banner{
		Id:          po.Id,
		Name:        po.Name,
		Title:       po.Title,
		Description: po.Description,
		Img:         po.Description,
		Items:       items,
	}, nil
}
