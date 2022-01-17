package service

import (
	"context"
	"mall-go/app/sku/service/internal/biz"

	"mall-go/api/sku/service"
)

type SkuService struct {
	service.UnimplementedSkuServer
	su *biz.SkuUsecase
}

func NewSkuService(su *biz.SkuUsecase) *SkuService {
	return &SkuService{
		su: su,
	}
}

func (s *SkuService) GetSkuById(ctx context.Context, req *service.IdRequest) (*service.SkuVO, error) {
	return &service.SkuVO{}, nil
}
