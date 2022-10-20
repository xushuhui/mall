package service

import (
	"context"
	"strings"

	"mall-go/app/app/service/internal/biz"
	"mall-go/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppService struct {
	UnimplementedAppServer
	bu *biz.BannerUsecase
	tu *biz.ThemeUsecase
	au *biz.ActivityUsecase
	cu *biz.CateGoryUsecase

	log *log.Helper
}

func NewAppService(bu *biz.BannerUsecase, tu *biz.ThemeUsecase, au *biz.ActivityUsecase, cu *biz.CateGoryUsecase,
	logger log.Logger,
) *AppService {
	return &AppService{
		bu:  bu,
		tu:  tu,
		au:  au,
		cu:  cu,
		log: log.NewHelper(logger),
	}
}

func (s *AppService) GetBannerById(ctx context.Context, in *IdRequest) (out *Banner, err error) {
	rv, err := s.bu.GetBannerById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var items []*BannerItem
	for _, v := range rv.Items {

		item := &BannerItem{
			Id:       v.ID,
			Name:     v.Name,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			BannerId: v.BannerId,
		}
		items = append(items, item)
	}
	return &Banner{
		Id:          rv.Id,
		Name:        rv.Name,
		Img:         rv.Img,
		Title:       rv.Title,
		Description: rv.Description,
		Items:       items,
	}, nil
}

func (s *AppService) GetBannerByName(ctx context.Context, in *NameRequest) (out *Banner, err error) {
	rv, err := s.bu.GetBannerByName(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	var items []*BannerItem
	for _, v := range rv.Items {
		items = append(items, &BannerItem{
			Id:       v.ID,
			Name:     v.Name,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			BannerId: v.BannerId,
		})
	}
	return &Banner{
		Id:          rv.Id,
		Name:        rv.Name,
		Img:         rv.Img,
		Title:       rv.Title,
		Description: rv.Description,
		Items:       items,
	}, nil
}

func (s *AppService) GetThemeByName(ctx context.Context, in *NameRequest) (out *ThemeSpu, err error) {
	rv, err := s.tu.GetThemeByName(ctx, in.Name)
	if err != nil {
		return nil, err
	}

	return &ThemeSpu{
		Id:     rv.Id,
		Online: rv.Online,
		Name:   rv.Name,
		Title:  rv.Title,
		SpuIds: rv.SpuIds,
	}, nil
}

func (s *AppService) GetThemeByNames(ctx context.Context, in *ThemeByNamesRequest) (out *Themes, err error) {
	rv, err := s.tu.GetThemeByNames(ctx, strings.Split(in.Names, `,`))
	if err != nil {
		return nil, err
	}
	var items []*Theme
	for _, v := range rv {
		items = append(items, &Theme{
			Id:     v.Id,
			Online: v.Online,
			Name:   v.Name,
			Title:  v.Title,
		})
	}

	return &Themes{
		Theme: items,
	}, nil
}

func (s *AppService) GetActivityByName(ctx context.Context, in *NameRequest) (out *Activity, err error) {
	rv, err := s.au.GetActivityByName(ctx, in.Name)
	if err != nil {
		return
	}

	return &Activity{
		Id:          rv.Id,
		Title:       rv.Title,
		EntranceImg: rv.EntranceImg,
		Online:      int32(rv.Online),
		Remark:      rv.Remark,
		StartTime:   timestamppb.New(rv.StartTime),
		EndTime:     timestamppb.New(rv.EndTime),
	}, nil
}

func (s *AppService) GetActivityWithCoupon(ctx context.Context, in *NameRequest) (out *ActivityCoupon, err error) {
	rv, err := s.au.GetActivityWithCoupon(ctx, in.Name)
	if err != nil {
		return
	}

	var coupons []*CouponBo
	for _, v := range rv.Coupons {
		coupons = append(coupons, &CouponBo{
			Id:          v.Id,
			Title:       v.Title,
			StartTime:   timestamppb.New(v.StartTime),
			EndTime:     timestamppb.New(v.EndTime),
			Description: rv.Description,
			FullMoney:   v.FullMoney,
			Minus:       v.Minus,
			Rate:        v.Rate,
			Type:        int32(v.Type),
			Remark:      v.Remark,
			WholeStore:  int32(v.WholeStore),
		})
	}

	out = &ActivityCoupon{
		Id:          rv.Activity.Id,
		Title:       rv.Activity.Title,
		EntranceImg: rv.Activity.EntranceImg,
		Online:      int32(rv.Activity.Online),
		Remark:      rv.Activity.Remark,
		StartTime:   timestamppb.New(rv.StartTime),
		EndTime:     timestamppb.New(rv.EndTime),
		Coupon:      coupons,
	}
	return
}

func (s *AppService) ListCategory(ctx context.Context, in *emptypb.Empty) (out *Categories, err error) {
	rv, err := s.cu.ListCategory(ctx)
	if err != nil {
		return
	}

	var roots []*Category
	var subs []*Category
	for _, v := range rv.Roots {
		roots = append(roots, &Category{
			Id:       v.Id,
			Name:     v.Name,
			IsRoot:   utils.Int2Bool(v.IsRoot),
			Img:      v.Img,
			ParentId: v.ParentId,
			Index:    int32(v.Index),
		})
	}

	for _, v := range rv.Subs {
		subs = append(subs, &Category{
			Id:       v.Id,
			Name:     v.Name,
			IsRoot:   utils.Int2Bool(v.IsRoot),
			Img:      v.Img,
			ParentId: v.ParentId,
			Index:    int32(v.Index),
		})
	}
	out = &Categories{
		Roots: roots,
		Subs:  subs,
	}
	return out, nil
}

func (s *AppService) ListGridCategory(ctx context.Context, in *emptypb.Empty) (out *GridCategories, err error) {
	rv, err := s.cu.ListGridCategory(ctx)
	if err != nil {
		return
	}
	var category []*GridCategories_GridCategory
	for _, v := range rv {
		category = append(category, &GridCategories_GridCategory{
			Id:             v.Id,
			Name:           v.Name,
			Title:          v.Title,
			Img:            v.Img,
			CategoryId:     int64(v.CategoryId),
			RootCategoryId: int64(v.RootCategoryId),
		})
	}
	out = &GridCategories{Category: category}
	return out, nil
}

func (s *AppService) GetTagByType(ctx context.Context, in *TypeRequest) (out *Tags, err error) {
	return
}
