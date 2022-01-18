package data

import (
	"context"
	"mall-go/app/spu/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type skuRepo struct {
	data *Data
	log  *log.Helper
}

func NewSkuRepo(data *Data, logger log.Logger) biz.SkuRepo {
	return &skuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *skuRepo) GetSkuById(ctx context.Context, id int64) (o biz.Sku, err error) {

	return biz.Sku{
		Id: 0,
	}, nil
}
