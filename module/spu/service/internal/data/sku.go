package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"mall-go/module/spu/service/internal/biz"
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
