package service

import (
	"context"
	"mall-go/api/mall/admin"
	"mall-go/app/mall/admin/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type MallAdmin struct {
	admin.UnimplementedAdminServer
	bu  *biz.BannerUsecase
	log *log.Helper
}

func NewAdmin(bu *biz.BannerUsecase,
	logger log.Logger) *MallAdmin {

	return &MallAdmin{
		bu: bu,

		log: log.NewHelper(logger),
	}
}

func (s *MallAdmin) GetBannerById(ctx context.Context, in *admin.IdRequest) (out *admin.Banner, err error) {
	b, err := s.bu.GetBannerById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	var items []*admin.BannerItem
	for _, v := range b.Items {
		items = append(items, &admin.BannerItem{
			Id:       v.ID,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			Name:     v.Name,
			BannerId: v.BannerId,
		})

	}
	return &admin.Banner{
		Id:          b.Id,
		Name:        b.Name,
		Title:       b.Title,
		Img:         b.Img,
		Description: b.Description,
		Items:       items,
	}, nil

}
