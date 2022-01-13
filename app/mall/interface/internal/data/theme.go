package data

import (
	"context"
	"mall-go/api/app"
	"mall-go/app/mall/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type themeRepo struct {
	log  *log.Helper
	data *Data
}

func NewThemeRepo(data *Data, logger log.Logger) biz.ThemeRepo {
	return &themeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *themeRepo) GetThemeWithSpu(ctx context.Context, name string) (t biz.ThemeSpu, err error) {
	po, err := r.data.ac.GetThemeWithSpu(ctx, &app.NameRequest{Name: name})
	var themeSups []*biz.Spu
	for _, v := range po.SpuList {
		themeSups = append(themeSups, &biz.Spu{
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
	theme := biz.Theme{
		Id:             po.Id,
		Title:          po.Title,
		Description:    po.Description,
		Name:           po.Name,
		EntranceImg:    po.EntranceImg,
		InternalTopImg: po.InternalTopImg,
		TitleImg:       po.TitleImg,
		TplName:        po.TplName,
		Online:         po.Online,
	}
	return biz.ThemeSpu{
		Theme:   theme,
		SpuList: themeSups,
	}, nil
}
func (r *themeRepo) GetThemeByNames(ctx context.Context, names string) (themes []biz.Theme, err error) {
	po, err := r.data.ac.GetThemeByNames(ctx, &app.ThemeByNamesRequest{Names: names})
	if err != nil {
		return
	}
	for _, v := range po.Theme {
		themes = append(themes, biz.Theme{
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
	return themes, nil
}
