package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall-go/api/mall"
	"mall-go/app/app/service/internal/biz"
	"mall-go/app/app/service/internal/data/model"
	"mall-go/app/app/service/internal/data/model/category"
	"mall-go/pkg/enum"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *categoryRepo) ListCategory(ctx context.Context) (c biz.AllCateGory, err error) {
	rootList, err := r.data.db.Category.Query().Where(category.IsRoot(enum.ROOT)).All(ctx)
	if model.IsNotFound(err) {
		err = mall.ErrorNotfound("category")
	}
	if err != nil {
		return
	}
	subList, err := r.data.db.Category.Query().Where(category.IsRoot(enum.SUB)).All(ctx)
	if model.IsNotFound(err) {
		err = mall.ErrorNotfound("category")
	}
	if err != nil {
		return
	}
	var roots []biz.CateGory
	var subs []biz.CateGory

	for _, v := range rootList {

		roots = append(roots, biz.CateGory{
			Id:       v.ID,
			Name:     v.Name,
			IsRoot:   v.IsRoot,
			Img:      v.Img,
			ParentId: v.ParentID,
			Index:    v.Index,
		})
	}

	for _, v := range subList {

		subs = append(subs, biz.CateGory{
			Id:       v.ID,
			Name:     v.Name,
			IsRoot:   v.IsRoot,
			Img:      v.Img,
			ParentId: v.ParentID,
		})
	}

	return biz.AllCateGory{Roots: roots, Subs: subs}, nil
}

func (r *categoryRepo) ListGridCategory(ctx context.Context) (c []biz.GridCategory, err error) {
	po, err := r.data.db.GridCategory.Query().All(ctx)
	if model.IsNotFound(err) {
		err = mall.ErrorNotfound("gridCategory")
	}
	if err != nil {
		return
	}

	for _, v := range po {
		c = append(c, biz.GridCategory{
			Id:             v.ID,
			Title:          v.Title,
			Img:            v.Img,
			Name:           v.Name,
			CategoryId:     v.CategoryID,
			RootCategoryId: v.RootCategoryID,
		})
	}
	return
}
