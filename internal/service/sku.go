package service

import (
	"mall-go/api/mall"
	"mall-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type SkuService struct {
	mall.UnimplementedSkuServer
	ku  *biz.SkuUsecase
	pu  *biz.SpuUsecase
	log *log.Helper
}

func NewSkuService(ku *biz.SkuUsecase, pu *biz.SpuUsecase, logger log.Logger) *SkuService {
	return &SkuService{ku: ku, pu: pu, log: log.NewHelper(logger)}
}
