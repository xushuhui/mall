package data

import (
	"context"
	"mall-go/internal/data/model"
	"mall-go/internal/data/model/theme"
)

func GetThemeWithSpu(ctx context.Context, name string) (t *model.Theme, err error) {
	t, err = GetDB().Theme.Query().Where(theme.Name(name)).WithSpu().First(ctx)
	return
}
func GetThemeByNames(ctx context.Context, names []string) (themes []*model.Theme, err error) {
	themes, err = GetDB().Theme.Query().Where(theme.NameIn(names...)).All(ctx)
	return
}
