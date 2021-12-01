package data

import (
	"mall-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type activityRepo struct {
	data *Data
	log  *log.Helper
}
func NewActivityRepo(data *Data, logger log.Logger) biz.ActivityRepo {
	return &activityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}