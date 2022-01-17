package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Sku struct {
	Id int64
}
type SkuRepo interface {
	GetSkuById(ctx context.Context, id int64) (o Sku, err error)
}
type SkuUsecase struct {
	repo SkuRepo
	log  *log.Helper
}

func NewSkuUsecase(repo SkuRepo, logger log.Logger) *SkuUsecase {
	return &SkuUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
func (uc *SkuUsecase) GetOrderById(ctx context.Context, id int64) (o Sku, err error) {
	o, err = uc.repo.GetSkuById(ctx, id)
	return
}
