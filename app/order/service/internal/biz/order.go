package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	Id int64
}
type OrderRepo interface {
	GetOrderById(ctx context.Context, id int64) (o Order, err error)
	GetOrderByOrderNo(ctx context.Context, orderNo string) (o Order, err error)
	ListUserOrder(ctx context.Context, userid int64) (list []Order, err error)
}
type OrderUsecase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *OrderUsecase) GetOrderById(ctx context.Context, id int64) (o Order, err error) {
	o, err = uc.repo.GetOrderById(ctx, id)
	return
}

func (uc *OrderUsecase) GetOrderByOrderNo(ctx context.Context, orderNo string) (o Order, err error) {
	o, err = uc.repo.GetOrderByOrderNo(ctx, orderNo)
	return
}
