package data

import (
	"context"

	"mall-go/module/app/service/internal/biz"
	"mall-go/module/app/service/internal/data/model/activity"

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
	po, err := r.data.db.Activity.Query().Where(activity.Name(name)).First(ctx)

	if err != nil {
		return
	}

	return biz.Activity{
		Id:             po.ID,
		Title:          po.Title,
		Description:    po.Description,
		StartTime:      po.StartTime,
		EndTime:        po.EndTime,
		Remark:         po.Remark,
		Online:         po.Online,
		EntranceImg:    po.EntranceImg,
		InternalTopImg: po.InternalTopImg,
	}, nil
}

//func (r *activityRepo) GetActivityWithCoupon(ctx context.Context, name string) (a biz.ActivityCoupon, err error) {
//	po, err := r.data.db.Activity.Query().Where(activity.Name(name)).WithCoupon(func(query *model.CouponQuery) {
//		query.Where(coupon.EndTimeGT(time.Now()))
//	}).First(ctx)
//
//	if model.IsNotFound(err) {
//		err = mall.ErrorNotFound("activity")
//		return
//	}
//
//	if err != nil {
//		return
//	}
//
//	var coupons []biz.Coupons
//
//	for _, v := range po.Edges.Coupon {
//		coupons = append(coupons, biz.Coupons{
//			Id:          v.ID,
//			Title:       v.Title,
//			StartTime:   v.StartTime,
//			EndTime:     v.EndTime,
//			Description: v.Description,
//			FullMoney:   v.FullMoney,
//			Rate:        v.Rate,
//			Type:        v.Type,
//			Remark:      v.Remark,
//			WholeStore:  v.WholeStore,
//			Minus:       v.Minus,
//		})
//	}
//
//	act := biz.Activity{
//		Id:          po.ID,
//		Title:       po.Title,
//		EntranceImg: po.EntranceImg,
//		Online:      po.Online,
//		Remark:      po.Remark,
//		StartTime:   po.StartTime,
//		EndTime:     po.EndTime,
//	}
//
//	return biz.ActivityCoupon{
//		Activity: act,
//		Coupons:  coupons,
//	}, nil
//
//}
