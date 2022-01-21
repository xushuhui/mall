package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
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
func (s *CouponService) CreateUserCoupon(ctx context.Context, in *service.CreateUserCouponRequest) (out *emptypb.Empty, err error) {
	panic("implement me")
}

func (s *CouponService) GetUserCouponByStatusWithCategory(ctx context.Context, in *service.StatusRequest) (out *service.Coupons, err error) {
	panic("implement me")
}

func (s *CouponService) GetUserCouponByStatus(ctx context.Context, in *service.StatusRequest) (out *service.Coupons, err error) {
	panic("implement me")
}

func (s *CouponService) GetCouponByType(ctx context.Context, in *service.TypeRequest) (out *service.Coupons, err error) {
	panic("implement me")
}
