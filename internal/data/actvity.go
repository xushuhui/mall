package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall-go/api/mall"
	"mall-go/internal/biz"
	"mall-go/internal/data/model"
	"mall-go/internal/data/model/activity"
	"mall-go/pkg/utils"
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
	po, err := r.data.db.Activity.Query().Where(activity.Name(name)).First(ctx)
	r.log.Info(po, 111)
	if model.IsNotFound(err) {
		err = mall.ErrorNotfound("activity")
		return
	}

	if err != nil {
		return
	}

	return biz.Activity{
		Id:             po.ID,
		Title:          po.Title,
		Description:    po.Description,
		StartTime:      utils.TimeBecomeString(po.StartTime),
		EndTime:        utils.TimeBecomeString(po.EndTime),
		Remark:         po.Remark,
		Online:         po.Online,
		EntranceImg:    po.EntranceImg,
		InternalTopImg: po.InternalTopImg,
	}, nil
}

func (r *activityRepo) GetActivityWithCoupon(ctx context.Context, name string) (a biz.ActivityCoupon, err error) {
	panic("not implemented") // TODO: Implement
}
