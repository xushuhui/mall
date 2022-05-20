package service

import (
	"context"
	mall"mall-go/api/mall/interface"
	"mall-go/app/mall/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type MallInterface struct {
	mall.UnimplementedInterfaceServer
	bu *biz.BannerUsecase
	tu *biz.ThemeUsecase
	au *biz.ActivityUsecase
	cu *biz.CategoryUsecase

	log *log.Helper
}

func NewInterface(bu *biz.BannerUsecase, tu *biz.ThemeUsecase, au *biz.ActivityUsecase, cu *biz.CategoryUsecase,
	logger log.Logger) *MallInterface {

	return &MallInterface{
		bu:  bu,
		tu:  tu,
		au:  au,
		cu:  cu,
		log: log.NewHelper(logger),
	}
}

func (s *MallInterface) GetBannerById(ctx context.Context, in *mall.IdRequest) (out *mall.Banner, err error) {
	b, err := s.bu.GetBannerById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	var items []*mall.BannerItem
	for _, v := range b.Items {
		items = append(items, &mall.BannerItem{
			Id:       v.ID,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			Name:     v.Name,
			BannerId: v.BannerId,
		})

	}
	return &mall.Banner{
		Id:          b.Id,
		Name:        b.Name,
		Title:       b.Title,
		Img:         b.Img,
		Description: b.Description,
		Items:       items,
	}, nil

}
func (s *MallInterface) GetBannerByName(ctx context.Context, in *mall.BannerByNameRequest) (out *mall.Banner, err error) {
	b, err := s.bu.GetBannerByName(ctx, in.Name)
	if err != nil {
		return
	}
	var items []*mall.BannerItem
	for _, v := range b.Items {
		items = append(items, &mall.BannerItem{
			Id:       v.ID,
			Img:      v.Img,
			Name:     v.Name,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			BannerId: v.BannerId,
		})
	}
	return &mall.Banner{
		Id:          b.Id,
		Name:        b.Name,
		Title:       b.Title,
		Img:         b.Img,
		Description: b.Description,
		Items:       items,
	}, nil
}
func (s *MallInterface) GetThemeByNames(ctx context.Context, in *mall.ThemeByNamesRequest) (out *mall.Themes, err error) {
	t, err := s.tu.GetThemeByNames(ctx, in.Names)
	if err != nil {
		return
	}
	var thems []*mall.Theme
	for _, v := range t {
		thems = append(thems, &mall.Theme{
			Id:             v.Id,
			Title:          v.Title,
			Description:    v.Description,
			Name:           v.Name,
			EntranceImg:    v.EntranceImg,
			InternalTopImg: v.InternalTopImg,
			TitleImg:       v.TitleImg,
			TplName:        v.TplName,
			Online:         v.Online,
		})
	}
	return &mall.Themes{
		Theme: thems,
	}, nil
}

func (s *MallInterface) GetThemeWithSpu(ctx context.Context, in *mall.ThemeWithSpuRequest) (out *mall.ThemeSpu, err error) {
	t, err := s.tu.GetThemeWithSpu(ctx, in.Name)
	if err != nil {
		return
	}
	var spuList []*mall.Spu
	for _, v := range t.SpuList {
		spuList = append(spuList, &mall.Spu{
			Id:             v.Id,
			Title:          v.Title,
			Subtitle:       v.Subtitle,
			CategoryId:     v.CategoryId,
			RootCategoryId: v.RootCategoryId,
			Price:          v.Price,
			Img:            v.Img,
			ForThemeImg:    v.ForThemeImg,
			Description:    v.Description,
			DiscountPrice:  v.DiscountPrice,
			Tags:           v.Tags,
			Online:         v.Online,
		})
	}

	return &mall.ThemeSpu{
		Id:             t.Id,
		Title:          t.Title,
		Description:    t.Description,
		Name:           t.Name,
		EntranceImg:    t.EntranceImg,
		InternalTopImg: t.InternalTopImg,
		TitleImg:       t.TitleImg,
		TplName:        t.TplName,
		Online:         t.Online,
		SpuList:        spuList,
	}, nil
}
func (s *MallInterface) GetActivityByName(ctx context.Context, in *mall.ActivityByNameRequest) (out *mall.Activity, err error) {
	c, err := s.au.GetActivityByName(ctx, in.Name)
	if err != nil {
		return
	}
	out = &mall.Activity{
		Id:          c.Id,
		Title:       c.Title,
		EntranceImg: c.EntranceImg,
		Online:      c.Online,
		Remark:      c.Remark,
		StartTime:   c.StartTime.Unix(),
		EndTime:     c.EndTime.Unix(),
	}
	return out, nil
}
func (s *MallInterface) GetActivityWithCoupon(ctx context.Context, in *mall.ActivityWithCouponRequest) (out *mall.ActivityCoupon, err error) {
	c, err := s.au.GetActivityWithCoupon(ctx, in.Name)
	if err != nil {
		return
	}
	var coupons []*mall.CouponBo
	for _, v := range c.Coupons {
		coupons = append(coupons, &mall.CouponBo{
			Id:          v.Id,
			Title:       v.Title,
			StartTime:   v.StartTime.Unix(),
			EndTime:     v.EndTime.Unix(),
			Description: v.Description,
			FullMoney:   v.FullMoney,
			Minus:       v.Minus,
			Rate:        v.Rate,
			Type:        int32(v.Type),
			Remark:      v.Remark,
			WholeStore:  int32(v.WholeStore),
		})
	}
	return &mall.ActivityCoupon{
		Id:          c.Id,
		Title:       c.Title,
		EntranceImg: c.EntranceImg,
		Online:      c.Online,
		Remark:      c.Remark,
		StartTime:   c.StartTime.Unix(),
		EndTime:     c.EndTime.Unix(),
		Coupon:      coupons,
	}, nil
}
func (s *MallInterface) GetAllCategory(ctx context.Context, in *emptypb.Empty) (out *mall.AllCategory, err error) {
	c, err := s.cu.GetAllCategory(ctx)
	if err != nil {
		return
	}
	var roots []*mall.Category
	var subs []*mall.Category
	for _, v := range c.Roots {
		roots = append(roots, &mall.Category{
			Id:       v.Id,
			Name:     v.Name,
			IsRoot:   v.IsRoot,
			Img:      v.Img,
			ParentId: v.ParentId,
			Index:    v.Index,
		})
	}
	for _, v := range c.Subs {
		subs = append(subs, &mall.Category{
			Id:       v.Id,
			Name:     v.Name,
			IsRoot:   v.IsRoot,
			Img:      v.Img,
			ParentId: v.ParentId,
			Index:    v.Index,
		})
	}
	return &mall.AllCategory{
		Roots: roots,
		Subs:  subs,
	}, nil
}
func (s *MallInterface) GetGridCategory(ctx context.Context, in *emptypb.Empty) (out *mall.GridCategories, err error) {
	c, err := s.cu.GetGridCategory(ctx)
	if err != nil {
		return
	}
	var category []*mall.GridCategories_GridCategory
	for _, v := range c {
		category = append(category, &mall.GridCategories_GridCategory{
			Id:             v.Id,
			Name:           v.Name,
			Title:          v.Title,
			Img:            v.Img,
			CategoryId:     v.CategoryId,
			RootCategoryId: v.RootCategoryId,
		})
	}
	return &mall.GridCategories{
		Category: category,
	}, nil
}
func (s *MallInterface) GetTagByType(ctx context.Context, in *mall.TagByTypeRequest) (out *mall.Tags, err error) {
	return
}
func (s *MallInterface) MiniappLogin(ctx context.Context, in *mall.MiniappLoginRequest) (out *mall.LoginReply, err error) {

	return
}
