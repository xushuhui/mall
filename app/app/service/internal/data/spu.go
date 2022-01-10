package data

import (
	"context"
	"mall-go/app/app/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type spuRepo struct {
	data *Data
	log  *log.Helper
}

func NewSpuRepo(data *Data, logger log.Logger) biz.SpuRepo {
	return &spuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *spuRepo) GetSpuById(ctx context.Context, id int64) (Spu biz.Spu, err error) {
	return
}
func (r *spuRepo) GetSpuByCategory(ctx context.Context, id int64) (Spus []biz.Spu, err error) {
	return
}
