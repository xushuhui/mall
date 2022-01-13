package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"mall-go/app/mall/interface/internal/biz"
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

func (r *categoryRepo) GetAllCategory(ctx context.Context) (c biz.AllCategory, err error) {
	po, err := r.data.ac.ListCategory(ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	var roots []biz.Category
	var subs []biz.Category
	for _, v := range po.Roots {
		roots = append(roots, biz.Category{
			Id:       v.Id,
			Name:     v.Name,
			IsRoot:   v.IsRoot,
			Img:      v.Img,
			ParentId: v.ParentId,
			Index:    v.Index,
		})
	}
	for _, v := range po.Subs {
		subs = append(subs, biz.Category{
			Id:       v.Id,
			Name:     v.Name,
			IsRoot:   v.IsRoot,
			Img:      v.Img,
			ParentId: v.ParentId,
			Index:    v.Index,
		})
	}
	return biz.AllCategory{
		Roots: roots,
		Subs:  subs,
	}, nil
}

func (r *categoryRepo) GetGridCategory(ctx context.Context) (c []biz.GridCategory, err error) {
	po, err := r.data.ac.ListGridCategory(ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	for _, v := range po.Category {
		c = append(c, biz.GridCategory{
			Id:             v.Id,
			Name:           v.Name,
			Title:          v.Title,
			Img:            v.Img,
			CategoryId:     v.CategoryId,
			RootCategoryId: v.RootCategoryId,
		})
	}
	return c, nil
}
