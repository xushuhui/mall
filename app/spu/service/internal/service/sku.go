package service

import (
	"context"
	"mall-go/app/spu/service/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"

	"mall-go/api/spu/service"
)

type SkuService struct {
	service.UnimplementedSkuServer
	su *biz.SkuUsecase
	pu *biz.SpuUsecase
}

func NewSkuService(su *biz.SkuUsecase, pu *biz.SpuUsecase) *SkuService {
	return &SkuService{
		su: su,
		pu: pu,
	}
}

func (s *SkuService) GetSkuById(ctx context.Context, req *service.IdRequest) (*service.SkuVO, error) {
	return &service.SkuVO{}, nil
}
func (s *SkuService) GetSpuByTheme(ctx context.Context, req *service.IdRequest) (*service.SpuByThemeReply, error) {
	return &service.SpuByThemeReply{}, nil
}

func (s *SkuService) GetSaleExplain(ctx context.Context, empty *emptypb.Empty) (out *service.SaleExplains, err error) {
	panic("implement me")
}

func (s *SkuService) GetSpuByCategory(ctx context.Context, in *service.IdRequest) (out *service.SpuPage, err error) {
	panic("implement me")
}

func (s *SkuService) GetSpuById(ctx context.Context, in *service.IdRequest) (out *service.SpuDetail, err error) {
	panic("implement me")
}

func (s *SkuService) GetSpuLatest(ctx context.Context, empty *emptypb.Empty) (out *service.SpuPage, err error) {
	panic("implement me")
}
