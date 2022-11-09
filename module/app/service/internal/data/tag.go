package data

import (
	"context"

	"mall-go/module/app/service/internal/biz"
	"mall-go/module/app/service/internal/data/model/tag"

	"github.com/go-kratos/kratos/v2/log"
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
