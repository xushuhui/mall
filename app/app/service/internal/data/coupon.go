package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall-go/app/app/service/internal/biz"
)

type couponRepo struct {
	data *Data
	log  *log.Helper
}

func NewCouponRepo(data *Data, logger log.Logger) biz.CouponRepo {
	return &couponRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *couponRepo) CreateUserCoupon(ctx context.Context, couponId, userId int64) (err error) {
	//err = r.data.db.UserCoupon.Create().SetCouponID(couponId).SetUserID(userId).Exec(ctx)
	return
}

func (r *couponRepo) ListUserCoupon(ctx context.Context, userId int64) (list []biz.Coupon, err error) {
	//pos, err := r.data.db.UserCoupon.Query().Where(usercoupon.UserID(userId)).WithCoupon().All(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//
	//for _, v := range pos {
	//	c := v.Edges.Coupon
	//	list = append(list, biz.Coupon{
	//		Id:        v.CouponID,
	//		Title:     c.Title,
	//		StartTime: c.StartTime,
	//		EndTime:   c.EndTime,
	//	})
	//}
	return
}

func (r *couponRepo) CreateCoupon(ctx context.Context) (err error) {
	panic("implement me")
}

func (r *couponRepo) ListCoupon(ctx context.Context) (list []biz.Coupon, err error) {
	panic("implement me")
}
