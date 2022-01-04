package biz

import "github.com/go-kratos/kratos/v2/log"

type Order struct {
	Id int64 
}
type OrderRepo interface {
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
