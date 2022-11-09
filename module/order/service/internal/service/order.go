package service

import (
	"context"

	"mall-go/module/order/service/internal/biz"
)

type OrderService struct {
	UnimplementedOrderServer
	ou *biz.OrderUsecase
}

func NewUserService(ou *biz.OrderUsecase) *OrderService {
	return &OrderService{
		ou: ou,
	}
}

func (s *OrderService) GetOrderById(ctx context.Context, req *IdRequest) (*OrderVO, error) {
	return &OrderVO{}, nil
}
