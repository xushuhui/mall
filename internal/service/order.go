package service

import (
	"mall-go/api/mall"
	"mall-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type OrderService struct {
	mall.UnimplementedOrderServer
	ou  *biz.OrderUsecase
	log *log.Helper
}
func NewOrderService(ou *biz.OrderUsecase,logger log.Logger) *OrderService {
	return &OrderService{
		ou:ou,
		log: log.NewHelper(logger),
	}
}