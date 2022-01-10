package service

import (
	"context"
	"mall-go/api/app"
	"mall-go/app/app/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type AppService struct {
	app.UnimplementedAppServer
	bu *biz.BannerUsecase
	tu *biz.ThemeUsecase

	log *log.Helper
}

func NewAppService(bu *biz.BannerUsecase, tu *biz.ThemeUsecase,
	logger log.Logger) *AppService {

	return &AppService{
		bu: bu,
		tu: tu,

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
	return
}
func (s *AppService) GetThemeByNames(ctx context.Context, in *app.ThemeByNamesRequest) (out *app.Themes, err error) {
	return
}

func (s *AppService) GetThemeWithSpu(ctx context.Context, in *app.ThemeWithSpuRequest) (out *app.ThemeSpu, err error) {

	return
}

// func (s *AppService) GetActivityByName(ctx context.Context, in *mall.ActivityByNameRequest) (out *mall.Activity, err error) {

// 	rv, err := s.au.GetActivityByName(ctx, in.Name)
// 	if err != nil {
// 		return
// 	}

// 	return &mall.Activity{
// 		Id:          rv.Id,
// 		Title:       rv.Title,
// 		EntranceImg: rv.EntranceImg,
// 		Online:      int32(rv.Online),
// 		Remark:      rv.Remark,
// 		StartTime:   rv.StartTime,
// 		EndTime:     rv.EndTime,
// 	}, nil
// }
func (s *AppService) GetActivityWithCoupon(ctx context.Context, in *app.ActivityWithCouponRequest) (out *app.ActivityCoupon, err error) {
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
