package data

import (
	"context"
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
func (r *activityRepo) GetActivityByName(ctx context.Context, name string) (a biz.Activity, err error) {
	panic("not implemented") // TODO: Implement
}

func (r *activityRepo) GetActivityWithCoupon(ctx context.Context, name string) (a biz.ActivityCoupon, err error) {
	panic("not implemented") // TODO: Implement
}
