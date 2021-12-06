package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Theme struct {
	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title          string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description    string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Name           string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	EntranceImg    string `protobuf:"bytes,5,opt,name=entrance_img,json=entranceImg,proto3" json:"entrance_img,omitempty"`
	Extend         string `protobuf:"bytes,6,opt,name=extend,proto3" json:"extend,omitempty"`
	InternalTopImg string `protobuf:"bytes,7,opt,name=internal_top_img,json=internalTopImg,proto3" json:"internal_top_img,omitempty"`
	TitleImg       string `protobuf:"bytes,8,opt,name=title_img,json=titleImg,proto3" json:"title_img,omitempty"`
	TplName        string `protobuf:"bytes,9,opt,name=tpl_name,json=tplName,proto3" json:"tpl_name,omitempty"`
	Online         int8   `protobuf:"varint,10,opt,name=online,proto3" json:"online,omitempty"`
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
