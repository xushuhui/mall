package service

import (
	"context"
	"mall-go/api/mall"
	"mall-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type CouponService struct {
	mall.UnimplementedCouponServer
	cu  *biz.CouponUsecase
	log *log.Helper
}

func NewCouponService(cu *biz.CouponUsecase, logger log.Logger) *CouponService {
	return &CouponService{cu: cu, log: log.NewHelper(logger)}
}
func (s *CouponService) CollectCoupon(ctx context.Context, in *mall.CollectCouponRequest) (out *emptypb.Empty,err error) {
	
	return
}
func (s *CouponService) GetCouponByCategory(ctx context.Context, in *mall.CouponByCategoryRequest) (out *mall.Coupons,err error) {
	return
}
func (s *CouponService) GetMyAvailableCoupon(ctx context.Context, in *emptypb.Empty) (out *mall.Coupons,err error) {
	return
}
func (s *CouponService) GetMyCouponByStatus(ctx context.Context, in *mall.MyCouponByStatusRequest) (out *mall.Coupons,err error) {
	return
}
func (s *CouponService) GetWholeCoupon(ctx context.Context, in *emptypb.Empty) (out *mall.Coupons,err error) {
	return
}
