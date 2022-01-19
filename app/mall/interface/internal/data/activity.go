package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	app "mall-go/api/app/service"
	"mall-go/app/mall/interface/internal/biz"
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

func (r *activityRepo) GetActivityByName(ctx context.Context, name string) (c biz.Activity, err error) {
	po, err := r.data.ac.GetActivityByName(ctx, &app.NameRequest{Name: name})
	if err != nil {
		return
	}
	return biz.Activity{
		Id:          po.Id,
		Title:       po.Title,
		StartTime:   po.StartTime.AsTime(),
		EndTime:     po.EndTime.AsTime(),
		Remark:      po.Remark,
		Online:      po.Online,
		EntranceImg: po.EntranceImg,
	}, nil
}

func (r *activityRepo) GetActivityWithCoupon(ctx context.Context, name string) (b biz.ActivityCoupon, err error) {
	po, err := r.data.ac.GetActivityWithCoupon(ctx, &app.NameRequest{Name: name})
	if err != nil {
		return
	}
	activity := biz.Activity{
		Id:          po.Id,
		Title:       po.Title,
		EntranceImg: po.EntranceImg,
		StartTime:   po.StartTime.AsTime(),
		EndTime:     po.EndTime.AsTime(),
		Remark:      po.Remark,
		Online:      po.Online,
	}
	var coupons []biz.Coupon
	for _, v := range po.Coupon {
		coupons = append(coupons, biz.Coupon{
			Id:          v.Id,
			Title:       v.Title,
			StartTime:   v.StartTime.AsTime(),
			EndTime:     v.EndTime.AsTime(),
			Description: v.Description,
			FullMoney:   v.FullMoney,
			Rate:        v.Rate,
			Minus:       v.Minus,
			Type:        int(v.Type),
			Remark:      v.Remark,
			WholeStore:  int(v.WholeStore),
		})
	}
	return biz.ActivityCoupon{
		Activity: activity,
		Coupons:  coupons,
	}, nil
}
