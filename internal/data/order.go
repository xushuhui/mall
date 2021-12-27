package data

import (
	"context"
	"mall-go/api/mall"
	"mall-go/internal/biz"
	"mall-go/internal/data/model"
	"mall-go/internal/data/model/order"

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
	if model.IsNotFound(err) {
		err = mall.ErrorNotFound("banner")
		return
	}
	if err != nil {
		return
	}

	return biz.Order{
		Id: po.ID,
	}, nil
}
func (r *orderRepo) GetOrderByOrderNo(ctx context.Context, orderNo string) (o biz.Order, err error) {
	return
}
