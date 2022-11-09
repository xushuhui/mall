package service

import (
	"context"

	"mall-go/module/coupon/service/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CouponService struct {
	UnimplementedCouponServer
	ou *biz.CouponUsecase
}

func NewCouponService(ou *biz.CouponUsecase) *CouponService {
	return &CouponService{
		ou: ou,
	}
}

func (s *CouponService) GetCouponByCategory(ctx context.Context, in *IdRequest) (*Coupons, error) {
	return &Coupons{}, nil
}

func (s *CouponService) CreateUserCoupon(ctx context.Context, in *CreateUserCouponRequest) (out *emptypb.Empty, err error) {
	panic("implement me")
}

func (s *CouponService) GetUserCouponByStatusWithCategory(ctx context.Context, in *StatusRequest) (out *Coupons, err error) {
	panic("implement me")
}

func (s *CouponService) GetUserCouponByStatus(ctx context.Context, in *StatusRequest) (out *Coupons, err error) {
	panic("implement me")
}

func (s *CouponService) GetCouponByType(ctx context.Context, in *TypeRequest) (out *Coupons, err error) {
	panic("implement me")
}
