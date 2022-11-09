package data

import (
	"context"

	"mall-go/module/order/service/internal/biz"
	"mall-go/module/order/service/internal/data/model"
	"mall-go/module/order/service/internal/data/model/order"

	"github.com/go-kratos/kratos/v2/log"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *orderRepo) GetOrderById(ctx context.Context, id int64) (o biz.Order, err error) {
	po, err := r.data.db.Order.Query().Where(order.ID(id)).First(ctx)
	if model.MaskNotFound(err) != nil {
		return
	}
	return biz.Order{
		Id: po.ID,
	}, nil
}

func (r *orderRepo) GetOrderByOrderNo(ctx context.Context, orderNo string) (o biz.Order, err error) {
	po, err := r.data.db.Order.Query().Where(order.OrderNo(orderNo)).First(ctx)
	if model.MaskNotFound(err) != nil {
		return
	}

	return biz.Order{
		Id: po.ID,
	}, nil
}

func (r *orderRepo) ListUserOrder(ctx context.Context, userid int64) (list []biz.Order, err error) {
	return
}
