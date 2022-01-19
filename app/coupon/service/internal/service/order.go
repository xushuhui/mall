package service

import (
	"context"
	"mall-go/app/coupon/service/internal/biz"

	"mall-go/api/coupon/service"
)

type CouponService struct {
	service.UnimplementedCouponServer
	ou *biz.CouponUsecase
}

func NewCouponService(ou *biz.CouponUsecase) *CouponService {
	return &CouponService{
		ou: ou,
	}
}

func (s *CouponService) GetCouponByCategory(ctx context.Context, req *service.IdRequest) (*service.Coupons, error) {
	return &service.Coupons{}, nil
}
