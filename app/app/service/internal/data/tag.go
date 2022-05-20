package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall-go/api/mall"
	"mall-go/app/app/service/internal/biz"
	"mall-go/app/app/service/internal/data/model"
	"mall-go/app/app/service/internal/data/model/tag"
)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepoRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *tagRepo) GetTagByType(ctx context.Context, kind int) (t []biz.Tag, err error) {
	po, err := r.data.db.Tag.Query().Where(tag.Type(kind)).All(ctx)
	if err != nil {
		return
	}
	if model.IsNotFound(err) {
		err = mall.ErrorNotFound("tag")
		return
	}
	for _, v := range po {
		t = append(t, biz.Tag{
			Id:          v.ID,
			Title:       v.Title,
			Highlight:   v.Highlight,
			Description: v.Description,
			Type:        v.Type,
		})
	}
	return t, nil
}
