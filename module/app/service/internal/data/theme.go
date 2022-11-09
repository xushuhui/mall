package data

import (
	"context"

	app "mall-go/api/app/service"
	"mall-go/module/app/service/internal/biz"
	"mall-go/module/app/service/internal/data/model/theme"

	"github.com/go-kratos/kratos/v2/log"
)

type themeRepo struct {
	data *Data
	log  *log.Helper
}

func NewThemeRepo(data *Data, logger log.Logger) biz.ThemeRepo {
	return &themeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *themeRepo) GetThemeByName(ctx context.Context, name string) (t biz.ThemeSpu, err error) {
	po, err := r.data.db.Theme.Query().Where(theme.Name(name)).WithThemeSpu().First(ctx)
	if err != nil {
		return
	}
	theme := biz.Theme{
		Id:             po.ID,
		Name:           po.Name,
		Title:          po.Title,
		Description:    po.Description,
		EntranceImg:    po.EntranceImg,
		InternalTopImg: po.InternalTopImg,
		TitleImg:       po.TitleImg,
		TplName:        po.TplName,
		Online:         int32(po.Online),
	}
	var spuIds []int64
	for _, spu := range po.Edges.ThemeSpu {
		spuIds = append(spuIds, spu.SpuID)
	}
	return biz.ThemeSpu{
		Theme: theme, SpuIds: spuIds,
	}, nil
}

func (r *themeRepo) GetThemeByNames(ctx context.Context, names []string) (themes []biz.Theme, err error) {
	pos, err := r.data.db.Theme.Query().Where(theme.NameIn(names...)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, po := range pos {
		themes = append(themes, biz.Theme{
			Id:          po.ID,
			Name:        po.Name,
			Title:       po.Title,
			Description: po.Description,
			EntranceImg: po.EntranceImg,

			InternalTopImg: po.InternalTopImg,
			TitleImg:       po.TitleImg,
			TplName:        po.TplName,
			Online:         int32(po.Online),
		})
	}
	return themes, nil
}

func (r *themeRepo) CreateTheme(ctx context.Context, req app.Theme) (err error) {
	return
}

func (r *themeRepo) ListTheme(ctx context.Context) (t []biz.Theme, err error) {
	return
}

func (r *themeRepo) UpdateTheme(ctx context.Context, req app.Theme) (err error) {
	return
}
