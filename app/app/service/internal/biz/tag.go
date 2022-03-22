package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Tag struct {
	Id          int64
	Title       string
	Highlight   int
	Description string
	Type        int
}

type TagRepo interface {
	GetTagByType(ctx context.Context, kind int) (t []Tag, err error)
}

type TagUsecase struct {
	repo TagRepo
	log  *log.Helper
}

func NewTagUsecase(repo TagRepo, logger log.Logger) *TagUsecase {
	return &TagUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
func (uc *TagUsecase) GetTagByType(ctx context.Context, kind int) (t []Tag, err error) {
	t, err = uc.repo.GetTagByType(ctx, kind)
	return
}
