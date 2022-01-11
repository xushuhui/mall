package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"mall-go/api/app"
	"mall-go/app/app/service/internal/biz"
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

func (s *AppService) GetBannerById(ctx context.Context, in *app.IdRequest) (out *app.Banner, err error) {
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
func (s *AppService) GetBannerByName(ctx context.Context, in *app.NameRequest) (out *app.Banner, err error) {
	rv, err := s.bu.GetBannerByName(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	var items []*app.BannerItem
	for _, v := range rv.Items {
		items = append(items, &app.BannerItem{
			Id:      v.ID,
			Name:    v.Name,
			Img:     v.Img,
			Keyword: v.Keyword,
			Type:    int32(v.Type),
		})
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

func (s *AppService) GetThemeWithSpu(ctx context.Context, in *app.NameRequest) (out *app.ThemeSpu, err error) {

	return
}

func (s *AppService) GetActivityByName(ctx context.Context, in *app.NameRequest) (out *app.Activity, err error) {

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
		StartTime:   rv.StartTime.Unix(),
		EndTime:     rv.EndTime.Unix(),
	}, nil
}

func (s *AppService) GetActivityWithCoupon(ctx context.Context, in *app.NameRequest) (out *app.ActivityCoupon, err error) {
	rv, err := s.au.GetActivityWithCoupon(ctx, in.Name)
	if err != nil {
		return
	}

	var coupons []*app.CouponBo
	for _, v := range rv.Coupons {

		coupons = append(coupons, &app.CouponBo{
			Id:          v.Id,
			Title:       v.Title,
			StartTime:   v.StartTime.Unix(),
			EndTime:     v.EndTime.Unix(),
			Description: rv.Description,
			FullMoney:   v.FullMoney,
			Minus:       v.Minus,
			Rate:        v.Rate,
			Type:        int32(v.Type),
			Remark:      v.Remark,
			WholeStore:  true,
		})
	}

	out = &app.ActivityCoupon{
		Id:          rv.Activity.Id,
		Title:       rv.Activity.Title,
		EntranceImg: rv.Activity.EntranceImg,
		Online:      int32(rv.Activity.Online),
		Remark:      rv.Activity.Remark,
		StartTime:   rv.Activity.StartTime.Unix(),
		EndTime:     rv.Activity.EndTime.Unix(),
		Coupon:      coupons,
	}
	return

}
func (s *AppService) GetAllCategory(ctx context.Context, in *emptypb.Empty) (out *app.Categories, err error) {
	return
}
func (s *AppService) GetGridCategory(ctx context.Context, in *emptypb.Empty) (out *app.GridCategories, err error) {
	return
}
func (s *AppService) GetTagByType(ctx context.Context, in *app.TypeRequest) (out *app.Tags, err error) {
	return
}
func (s *AppService) CreateUserCoupon(ctx context.Context, in *app.CreateUserCouponRequest) (out *emptypb.Empty, err error) {
	panic("implement me")
}

func (s *AppService) GetCouponByCategory(ctx context.Context, in *app.IdRequest) (out *app.Coupons, err error) {
	panic("implement me")
}

func (s *AppService) GetUserCouponByStatusWithCategory(ctx context.Context, in *app.StatusRequest) (out *app.Coupons, err error) {
	panic("implement me")
}

func (s *AppService) GetUserCouponByStatus(ctx context.Context, in *app.StatusRequest) (out *app.Coupons, err error) {
	panic("implement me")
}

func (s *AppService) GetSaleExplain(ctx context.Context, empty *emptypb.Empty) (out *app.SaleExplains, err error) {
	panic("implement me")
}

func (s *AppService) GetSpuByCategory(ctx context.Context, in *app.IdRequest) (out *app.SpuPage, err error) {
	panic("implement me")
}

func (s *AppService) GetSpuById(ctx context.Context, in *app.IdRequest) (out *app.SpuDetail, err error) {
	panic("implement me")
}

func (s *AppService) GetSpuLatest(ctx context.Context, empty *emptypb.Empty) (out *app.SpuPage, err error) {
	panic("implement me")
}

func (s *AppService) GetCouponByType(ctx context.Context, in *app.TypeRequest) (out *app.Coupons, err error) {
	panic("implement me")
}
