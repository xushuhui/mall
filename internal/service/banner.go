package service

import (
	"context"
	"mall-go/api/mall"
	"mall-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AppService struct {
	mall.UnimplementedAppServer
	bu  *biz.BannerUsecase
	log *log.Helper
}

func NewAppService(uc *biz.BannerUsecase, logger log.Logger) *AppService {
	return &AppService{
		bu:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *AppService) GetBannerById(ctx context.Context, in *mall.BannerByIdRequest) (*mall.Banner, error) {
	rv, err := s.bu.GetBannerById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var items []*mall.BannerItem
	for _, v := range rv.Items {

		item := &mall.BannerItem{
			Id:       v.ID,
			Name:     v.Name,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			BannerId: v.BannerID,
		}
		items = append(items, item)
	}
	return &mall.Banner{
		Id:          rv.Id,
		Name:        rv.Name,
		Img:         rv.Img,
		Title:       rv.Title,
		Description: rv.Description,
		Items:       items,
	}, nil
}
