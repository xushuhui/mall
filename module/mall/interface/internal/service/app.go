package service

import (
	"context"

	"mall-go/module/mall/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type MallInterface struct {
	UnimplementedInterfaceServer
	bu  *biz.BannerUsecase
	tu  *biz.ThemeUsecase
	au  *biz.ActivityUsecase
	cu  *biz.CategoryUsecase
	gu  *biz.TagUsecase
	uu  *biz.UserUsecase
	log *log.Helper
}

func NewMallInterface(bu *biz.BannerUsecase, tu *biz.ThemeUsecase, au *biz.ActivityUsecase,
	cu *biz.CategoryUsecase, gu *biz.TagUsecase, uu *biz.UserUsecase, logger log.Logger) *MallInterface {
	return &MallInterface{
		bu:  bu,
		tu:  tu,
		au:  au,
		cu:  cu,
		gu:  gu,
		uu:  uu,
		log: log.NewHelper(logger),
	}
}

func (s *MallInterface) GetBannerById(ctx context.Context, in *IdRequest) (out *Banner, err error) {
	b, err := s.bu.GetBannerById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	var items []*BannerItem
	for _, v := range b.Items {
		items = append(items, &BannerItem{
			Id:       v.ID,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			Name:     v.Name,
			BannerId: v.BannerId,
		})
	}
	return &Banner{
		Id:          b.Id,
		Name:        b.Name,
		Title:       b.Title,
		Img:         b.Img,
		Description: b.Description,
		Items:       items,
	}, nil
}

func (s *MallInterface) GetBannerByName(ctx context.Context, in *BannerByNameRequest) (out *Banner, err error) {
	b, err := s.bu.GetBannerByName(ctx, in.Name)
	if err != nil {
		return
	}
	var items []*BannerItem
	for _, v := range b.Items {
		items = append(items, &BannerItem{
			Id:       v.ID,
			Img:      v.Img,
			Name:     v.Name,
			Keyword:  v.Keyword,
			Type:     int32(v.Type),
			BannerId: v.BannerId,
		})
	}
	return &Banner{
		Id:          b.Id,
		Name:        b.Name,
		Title:       b.Title,
		Img:         b.Img,
		Description: b.Description,
		Items:       items,
	}, nil
}

func (s *MallInterface) GetThemeByNames(ctx context.Context, in *ThemeByNamesRequest) (out *Themes, err error) {
	t, err := s.tu.GetThemeByNames(ctx, in.Names)
	if err != nil {
		return
	}
	var thems []*Theme
	for _, v := range t {
		thems = append(thems, &Theme{
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
	return &Themes{
		Theme: thems,
	}, nil
}

func (s *MallInterface) GetThemeWithSpu(ctx context.Context, in *ThemeWithSpuRequest) (out *ThemeSpu, err error) {
	t, err := s.tu.GetThemeWithSpu(ctx, in.Name)
	if err != nil {
		return
	}
	var spuList []*Spu
	for _, v := range t.SpuList {
		spuList = append(spuList, &Spu{
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

	return &ThemeSpu{
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

func (s *MallInterface) GetActivityByName(ctx context.Context, in *ActivityByNameRequest) (out *Activity, err error) {
	c, err := s.au.GetActivityByName(ctx, in.Name)
	if err != nil {
		return
	}
	out = &Activity{
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

func (s *MallInterface) GetActivityWithCoupon(ctx context.Context, in *ActivityWithCouponRequest) (out *ActivityCoupon, err error) {
	c, err := s.au.GetActivityWithCoupon(ctx, in.Name)
	if err != nil {
		return
	}
	var coupons []*CouponBo
	for _, v := range c.Coupons {
		coupons = append(coupons, &CouponBo{
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
	return &ActivityCoupon{
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

func (s *MallInterface) GetAllCategory(ctx context.Context, in *emptypb.Empty) (out *AllCategory, err error) {
	c, err := s.cu.GetAllCategory(ctx)
	if err != nil {
		return
	}
	var roots []*Category
	var subs []*Category
	for _, v := range c.Roots {
		roots = append(roots, &Category{
			Id:       v.Id,
			Name:     v.Name,
			IsRoot:   v.IsRoot,
			Img:      v.Img,
			ParentId: v.ParentId,
			Index:    v.Index,
		})
	}
	for _, v := range c.Subs {
		subs = append(subs, &Category{
			Id:       v.Id,
			Name:     v.Name,
			IsRoot:   v.IsRoot,
			Img:      v.Img,
			ParentId: v.ParentId,
			Index:    v.Index,
		})
	}
	return &AllCategory{
		Roots: roots,
		Subs:  subs,
	}, nil
}

func (s *MallInterface) GetGridCategory(ctx context.Context, in *emptypb.Empty) (out *GridCategories, err error) {
	c, err := s.cu.GetGridCategory(ctx)
	if err != nil {
		return
	}
	var category []*GridCategories_GridCategory
	for _, v := range c {
		category = append(category, &GridCategories_GridCategory{
			Id:             v.Id,
			Name:           v.Name,
			Title:          v.Title,
			Img:            v.Img,
			CategoryId:     v.CategoryId,
			RootCategoryId: v.RootCategoryId,
		})
	}
	return &GridCategories{
		Category: category,
	}, nil
}

func (s *MallInterface) GetTagByType(ctx context.Context, in *TagByTypeRequest) (out *Tags, err error) {
	t, err := s.gu.GetTagByType(ctx, in.Type)
	if err != nil {
		return
	}
	var tags []*Tags_Tag
	for _, v := range t {
		tags = append(tags, &Tags_Tag{
			Id:          v.Id,
			Title:       v.Title,
			Highlight:   v.Highlight,
			Description: v.Description,
			Type:        v.Type,
		})
	}
	return &Tags{Tag: tags}, nil
}

func (s *MallInterface) MiniappLogin(ctx context.Context, in *MiniappLoginRequest) (out *LoginReply, err error) {
	return
}
