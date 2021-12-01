package data

import (
	"mall-go/internal/biz"
	"mall-go/internal/data/model"

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
func GetSpuById(id int) (Spu *model.Spu, err error) {
	return
}
func GetSpuByCategory(id int) (Spus []*model.Spu, err error) {
	return
}
