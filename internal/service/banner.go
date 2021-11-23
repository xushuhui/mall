package service

import (
	"context"
	"mall-go/api/mall"
	"mall-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ShowService struct {
	mall.UnimplementedShowServer
	uc  *biz.BannerUsecase
	log *log.Helper
}

func NewShowService(uc *biz.BannerUsecase, logger log.Logger) *ShowService {
	return &ShowService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *ShowService) GetBannerById(ctx context.Context, in *mall.BannerByIdRequest) (*mall.BannerByIdReply, error) {
	rv, err := s.uc.GetBannerById(ctx, in.Id)
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
	return &mall.BannerByIdReply{
		Id:          rv.Id,
		Name:        rv.Name,
		Img:         rv.Img,
		Title:       rv.Title,
		Description: rv.Description,
		Items:       items,
	}, nil
}
