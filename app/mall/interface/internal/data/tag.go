package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	app "mall-go/api/app/service"
	"mall-go/app/mall/interface/internal/biz"
)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *tagRepo) GetTagByType(ctx context.Context, kind int32) (t []biz.Tag, err error) {
	po, err := r.data.ac.GetTagByType(ctx, &app.TypeRequest{Type: kind})
	if err != nil {
		return
	}
	for _, v := range po.Tag {
		t = append(t, biz.Tag{
			Id:          v.Id,
			Title:       v.Title,
			Highlight:   v.Highlight,
			Description: v.Description,
			Type:        v.Type,
		})

	}
	return t, nil
}
