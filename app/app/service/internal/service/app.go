package service

import (
	"context"
	"mall-go/api/app"
	"mall-go/app/app/service/internal/biz"
	"mall-go/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type AppService struct {
	app.UnimplementedAppServer
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

func (s *AppService) GetBannerById(ctx context.Context, in *app.BannerByIdRequest) (out *app.Banner, err error) {
	rv, err := s.bu.GetBannerById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var items []*app.BannerItem
	for _, v := range rv.Items {

		item := &app.BannerItem{
			Id:      v.ID,
			Name:    v.Name,
			Img:     v.Img,
			Keyword: v.Keyword,
			Type:    int32(v.Type),
		}
		items = append(items, item)
	}
	return &app.Banner{
		Id:          rv.Id,
		Name:        rv.Name,
		Img:         rv.Img,
		Title:       rv.Title,
		Description: rv.Description,
		Items:       items,
	}, nil

}
func (s *AppService) GetBannerByName(ctx context.Context, in *app.BannerByNameRequest) (out *app.Banner, err error) {
	rv, err := s.bu.GetBannerByName(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	var items []*app.BannerItem
	for _, v := range rv.Items {

		item := &app.BannerItem{
			Id:      v.ID,
			Name:    v.Name,
			Img:     v.Img,
			Keyword: v.Keyword,
			Type:    int32(v.Type),
		}
		items = append(items, item)
	}
	return &app.Banner{
		Id:          rv.Id,
		Name:        rv.Name,
		Img:         rv.Img,
		Title:       rv.Title,
		Description: rv.Description,
		Items:       items,
	}, nil

}
func (s *AppService) GetThemeByNames(ctx context.Context, in *app.ThemeByNamesRequest) (out *app.Themes, err error) {
	return
}

func (s *AppService) GetThemeWithSpu(ctx context.Context, in *app.ThemeWithSpuRequest) (out *app.ThemeSpu, err error) {

	return
}

func (s *AppService) GetActivityByName(ctx context.Context, in *app.ActivityByNameRequest) (out *app.Activity, err error) {

	rv, err := s.au.GetActivityByName(ctx, in.Name)
	if err != nil {
		return
	}

	return &app.Activity{
		Id:          rv.Id,
		Title:       rv.Title,
		EntranceImg: rv.EntranceImg,
		Online:      int32(rv.Online),
		Remark:      rv.Remark,
		StartTime:   utils.TimeBecomeString(rv.StartTime),
		EndTime:     utils.TimeBecomeString(rv.EndTime),
	}, nil
}

func (s *AppService) GetActivityWithCoupon(ctx context.Context, in *app.ActivityWithCouponRequest) (out *app.ActivityCoupon, err error) {
	rv, err := s.au.GetActivityWithCoupon(ctx, in.Name)
	if err != nil {
		return
	}

	var coupons []*app.CouponBo
	for _, v := range rv.Coupons {

		coupon := &app.CouponBo{
			Id:          uint32(v.Id),
			Title:       v.Title,
			StartTime:   utils.Time2uint64(v.StartTime),
			EndTime:     utils.Time2uint64(v.EndTime),
			Description: rv.Description,
			FullMoney:   v.FullMoney,
			Minus:       v.Minus,
			Rate:        v.Rate,
			Type:        uint32(v.Type),
			Remark:      v.Remark,
			WholeStore:  true,
		}
		coupons = append(coupons, coupon)
	}

	out = &app.ActivityCoupon{
		Id:          rv.Activity.Id,
		Title:       rv.Activity.Title,
		EntranceImg: rv.Activity.EntranceImg,
		Online:      int32(rv.Activity.Online),
		Remark:      rv.Activity.Remark,
		StartTime:   utils.TimeBecomeString(rv.Activity.StartTime),
		EndTime:     utils.TimeBecomeString(rv.Activity.EndTime),
		Coupon:      coupons,
	}
	return

}
func (s *AppService) GetAllCategory(ctx context.Context, in *emptypb.Empty) (out *app.AllCategory, err error) {
	return
}
func (s *AppService) GetGridCategory(ctx context.Context, in *emptypb.Empty) (out *app.GridCategories, err error) {
	return
}
func (s *AppService) GetTagByType(ctx context.Context, in *app.TagByTypeRequest) (out *app.Tags, err error) {
	return
}
