package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"mall-go/api/app"
	"mall-go/app/app/service/internal/biz"
	"mall-go/pkg/utils"
	"strings"
)

type AppService struct {
	app.UnimplementedAppServer
	bu *biz.BannerUsecase
	tu *biz.ThemeUsecase
	au *biz.ActivityUsecase
	cu *biz.CateGoryUsecase

	log *log.Helper
}

func NewAppService(bu *biz.BannerUsecase, tu *biz.ThemeUsecase, au *biz.ActivityUsecase, cu *biz.CateGoryUsecase,
	logger log.Logger) *AppService {

	return &AppService{
		bu:  bu,
		tu:  tu,
		au:  au,
		cu:  cu,
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
			Id:       v.ID,
			Name:     v.Name,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			BannerId: v.BannerId,
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
			Id:       v.ID,
			Name:     v.Name,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			BannerId: v.BannerId,
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
	rv, err := s.tu.GetThemeByNames(ctx, strings.Split(in.Names, `,`))
	if err != nil {
		return nil, err
	}
	var items []*app.Theme
	for _, v := range rv {
		items = append(items, &app.Theme{
			Id:     v.Id,
			Online: v.Online,
			Name:   v.Name,
			Title:  v.Title,
		})
	}

	return &app.Themes{
		Theme: items,
	}, nil
}

func (s *AppService) GetThemeWithSpu(ctx context.Context, in *app.NameRequest) (out *app.ThemeSpu, err error) {
	rv, err := s.tu.GetThemeWithSpu(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	var items []*app.Spu
	for _, v := range rv.SpuList {
		items = append(items, &app.Spu{
			Id:       v.Id,
			Title:    v.Title,
			Subtitle: v.Subtitle,
			Online:   v.Online,

			Img:            v.Img,
			ForThemeImg:    v.ForThemeImg,
			RootCategoryId: v.RootCategoryId,
			CategoryId:     v.CategoryId,
		})
	}
	return &app.ThemeSpu{
		Id:          rv.Id,
		Title:       rv.Title,
		Name:        rv.Name,
		EntranceImg: rv.EntranceImg,
		Online:      rv.Online,
		TitleImg:    rv.TitleImg,
		SpuList:     items,
	}, nil

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
func (s *AppService) ListCategory(ctx context.Context, in *emptypb.Empty) (out *app.Categories, err error) {
	rv, err := s.cu.ListCategory(ctx)
	if err != nil {
		return
	}

	var roots []*app.Category
	var subs []*app.Category
	for _, v := range rv.Roots {
		roots = append(roots, &app.Category{
			Id:       v.Id,
			Name:     v.Name,
			IsRoot:   utils.Int2Bool(v.IsRoot),
			Img:      v.Img,
			ParentId: v.ParentId,
			Index:    int32(v.Index),
		})
	}

	for _, v := range rv.Subs {
		subs = append(subs, &app.Category{
			Id:       v.Id,
			Name:     v.Name,
			IsRoot:   utils.Int2Bool(v.IsRoot),
			Img:      v.Img,
			ParentId: v.ParentId,
			Index:    int32(v.Index),
		})
	}
	out = &app.Categories{
		Roots: roots,
		Subs:  subs,
	}
	return out, nil
}
func (s *AppService) ListGridCategory(ctx context.Context, in *emptypb.Empty) (out *app.GridCategories, err error) {
	rv, err := s.cu.ListGridCategory(ctx)
	if err != nil {
		return
	}
	var category []*app.GridCategories_GridCategory
	for _, v := range rv {
		category = append(category, &app.GridCategories_GridCategory{
			Id:             v.Id,
			Name:           v.Name,
			Title:          v.Title,
			Img:            v.Img,
			CategoryId:     int64(v.CategoryId),
			RootCategoryId: int64(v.RootCategoryId),
		})
	}
	out = &app.GridCategories{Category: category}
	return out, nil
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
