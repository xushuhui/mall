package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AllCateGory struct {
	Roots []CateGory
	Subs  []CateGory
}

type CateGory struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	IsRoot   int    `json:"is_root"`
	Img      string `json:"img"`
	ParentId int64  `json:"parent_id"`
	Index    int    `json:"index"`
}

type GridCategory struct {
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	Img            string `json:"img"`
	Name           string `json:"name"`
	CategoryId     int    `json:"category_id"`
	RootCategoryId int    `json:"root_category_id"`
}

type CategoryRepo interface {
	ListCategory(ctx context.Context) (c AllCateGory, err error)
	ListGridCategory(ctx context.Context) (c []GridCategory, err error)
}

type CateGoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCateGoryUsecase(repo CategoryRepo, logger log.Logger) *CateGoryUsecase {
	return &CateGoryUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (cu *CateGoryUsecase) ListCategory(ctx context.Context) (c AllCateGory, err error) {
	c, err = cu.repo.ListCategory(ctx)
	return
}

func (cu *CateGoryUsecase) ListGridCategory(ctx context.Context) (c []GridCategory, err error) {
	c, err = cu.repo.ListGridCategory(ctx)
	return
}
