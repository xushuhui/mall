package data

import (
	"context"
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
	return
}
func (r *themeRepo) GetThemeByNames(ctx context.Context, names []string) (themes []biz.Theme, err error) {

	return
}
