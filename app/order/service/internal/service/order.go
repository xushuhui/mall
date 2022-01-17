package service

import (
	"context"
	"mall-go/app/order/service/internal/biz"

	"mall-go/api/order/service"
)

type OrderService struct {
	service.UnimplementedOrderServer
	ou *biz.OrderUsecase
}

func NewUserService(ou *biz.OrderUsecase) *OrderService {
	return &OrderService{
		ou: ou,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *service.CreateOrderRequest) (*service.CreateOrderReply, error) {
	return &service.CreateOrderReply{}, nil
}
