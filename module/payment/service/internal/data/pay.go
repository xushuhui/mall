package data

import (
	"mall-go/module/payment/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type payRepo struct {
	data *Data
	log  *log.Helper
}

func NewPayRepo(data *Data, logger log.Logger) biz.PayRepo {
	return &payRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
