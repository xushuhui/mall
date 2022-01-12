package service

import (
	"context"
	"mall-go/api/mall"
	"mall-go/app/mall/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Interface struct {
	mall.UnimplementedInterfaceServer
	bu *biz.BannerUsecase

	log *log.Helper
}

func NewInterface(bu *biz.BannerUsecase,
	logger log.Logger) *Interface {

	return &Interface{
		bu: bu,

		log: log.NewHelper(logger),
	}
}

func (s *Interface) GetBannerById(ctx context.Context, in *mall.BannerByIdRequest) (out *mall.Banner, err error) {
	b, err := s.bu.GetBannerById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &mall.Banner{
		Id:          b.Id,
		Name:        b.Name,
		Title:       b.Title,
		Img:         b.Img,
		Description: b.Description,
	}, nil

}
func (s *Interface) GetBannerByName(ctx context.Context, in *mall.BannerByNameRequest) (out *mall.Banner, err error) {
	return
}
func (s *Interface) GetThemeByNames(ctx context.Context, in *mall.ThemeByNamesRequest) (out *mall.Themes, err error) {

	return
}

func (s *Interface) GetThemeWithSpu(ctx context.Context, in *mall.ThemeWithSpuRequest) (out *mall.ThemeSpu, err error) {

	return
}
func (s *Interface) GetActivityByName(ctx context.Context, in *mall.ActivityByNameRequest) (out *mall.Activity, err error) {
	return

}
func (s *Interface) GetActivityWithCoupon(ctx context.Context, in *mall.ActivityWithCouponRequest) (out *mall.ActivityCoupon, err error) {
	return
}
func (s *Interface) GetAllCategory(ctx context.Context, in *emptypb.Empty) (out *mall.AllCategory, err error) {
	return
}
func (s *Interface) GetGridCategory(ctx context.Context, in *emptypb.Empty) (out *mall.GridCategories, err error) {
	return
}
func (s *Interface) GetTagByType(ctx context.Context, in *mall.TagByTypeRequest) (out *mall.Tags, err error) {
	return
}
