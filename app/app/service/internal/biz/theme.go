package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Theme struct {
	Id             int64  ` json:"id,omitempty"`
	Title          string ` json:"title,omitempty"`
	Description    string ` json:"description,omitempty"`
	Name           string ` json:"name,omitempty"`
	EntranceImg    string ` json:"entrance_img,omitempty"`
	
	InternalTopImg string ` json:"internal_top_img,omitempty"`
	TitleImg       string `json:"title_img,omitempty"`
	TplName        string ` json:"tpl_name,omitempty"`
	Online         int32   `json:"online,omitempty"`
}
type Spu struct {
	
}
type ThemeSpu struct {
	Theme
	SpuList []*Spu ` json:"spu_list,omitempty"`
}
type ThemeRepo interface {
	GetThemeByNames(ctx context.Context, names []string) (t []Theme, err error)
	GetThemeWithSpu(ctx context.Context, name string) (t ThemeSpu, err error)
}

type ThemeUsecase struct {
	repo ThemeRepo
	log  *log.Helper
}

func NewThemeUsecase(repo ThemeRepo, logger log.Logger) *ThemeUsecase {
	return &ThemeUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
func (uc *ThemeUsecase) GetThemeByNames(ctx context.Context, names []string) (t []Theme, err error) {
	t, err = uc.repo.GetThemeByNames(ctx, names)
	return
}
func (uc *ThemeUsecase) GetThemeWithSpu(ctx context.Context, name string) (t ThemeSpu, err error) {
	t, err = uc.repo.GetThemeWithSpu(ctx, name)
	return
}
