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

func (s *OrderService) GetOrderById(ctx context.Context, req *service.IdRequest) (*service.OrderVO, error) {
	return &service.OrderVO{}, nil
}
