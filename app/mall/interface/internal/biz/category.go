package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AllCategory struct {
	Roots []Category
	Subs  []Category
}

type Category struct {
	Id       int64
	Name     string
	IsRoot   bool
	Img      string
	ParentId int64
	Index    int32
}

type GridCategory struct {
	Id             int64
	Name           string
	Title          string
	Img            string
	CategoryId     int64
	RootCategoryId int64
}

type CategoryRepo interface {
	GetAllCategory(ctx context.Context) (c AllCategory, err error)
	GetGridCategory(ctx context.Context) (c []GridCategory, err error)
}

type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (cu *CategoryUsecase) GetAllCategory(ctx context.Context) (c AllCategory, err error) {
	c, err = cu.repo.GetAllCategory(ctx)
	return
}

func (cu *CategoryUsecase) GetGridCategory(ctx context.Context) (c []GridCategory, err error) {
	c, err = cu.repo.GetGridCategory(ctx)
	return
}
