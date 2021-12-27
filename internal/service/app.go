package service

import (
	"context"
	"mall-go/api/mall"
	"mall-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type AppService struct {
	mall.UnimplementedAppServer
	bu *biz.BannerUsecase
	tu *biz.ThemeUsecase
	au *biz.ActivityUsecase

	log *log.Helper
}

func NewAppService(bu *biz.BannerUsecase, tu *biz.ThemeUsecase, au *biz.ActivityUsecase,
	logger log.Logger) *AppService {

	return &AppService{
		bu: bu,
		tu: tu,
		au: au,

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
func (s *AppService) GetBannerByName(ctx context.Context, in *mall.BannerByNameRequest) (out *mall.Banner, err error) {
	rv, err := s.bu.GetBannerByName(ctx, in.Name)
	if err != nil {
		return
	}
	var items []*mall.BannerItem
	for _, v := range rv.Items {

		item := &mall.BannerItem{
			Id:       v.ID,
			Name:     v.Name,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			
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
func (s *AppService) GetThemeByNames(ctx context.Context, in *mall.ThemeByNamesRequest) (out *mall.Themes, err error) {
	rv, err := s.tu.GetThemeByNames(ctx, in.Names)
	if err != nil {
		return
	}
	var themes []*mall.Theme
	for _, v := range rv {
		theme := &mall.Theme{
			Id:          v.Id,
			Name:        v.Name,
			EntranceImg: v.EntranceImg,
			Title:       v.Title,
			Description: v.Description,
		}
		themes = append(themes, theme)
	}
	out = &mall.Themes{
		Theme: themes,
	}
	return
}

func (s *AppService) GetThemeWithSpu(ctx context.Context, in *mall.ThemeWithSpuRequest) (out *mall.ThemeSpu, err error) {
	rv, err := s.tu.GetThemeWithSpu(ctx, in.Name)
	if err != nil {
		return
	}
	out = &mall.ThemeSpu{
		Id:          rv.Id,
		Name:        rv.Name,
		EntranceImg: rv.EntranceImg,
		Title:       rv.Title,
		Description: rv.Description,
	}
	return
}
func (s *AppService) GetActivityByName(ctx context.Context, in *mall.ActivityByNameRequest) (out *mall.Activity, err error) {
	return
}
func (s *AppService) GetActivityWithCoupon(ctx context.Context, in *mall.ActivityWithCouponRequest) (out *mall.ActivityCoupon, err error) {
	return
}
func (s *AppService) GetAllCategory(ctx context.Context, in *emptypb.Empty) (out *mall.AllCategory, err error) {
	return
}
func (s *AppService) GetGridCategory(ctx context.Context, in *emptypb.Empty) (out *mall.GridCategories, err error) {
	return
}
func (s *AppService) GetTagByType(ctx context.Context, in *mall.TagByTypeRequest) (out *mall.Tags, err error) {
	return
}
