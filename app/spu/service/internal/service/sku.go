package service

import (
	"context"
	"mall-go/app/spu/service/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"

	"mall-go/api/spu/service"
)

type SpuService struct {
	service.UnimplementedSpuServer
	su *biz.SkuUsecase
	pu *biz.SpuUsecase
}

func NewSpuService(su *biz.SkuUsecase, pu *biz.SpuUsecase) *SpuService {
	return &SpuService{
		su: su,
		pu: pu,
	}
}

func (s *SpuService) GetSkuById(ctx context.Context, req *service.IdRequest) (*service.SkuVO, error) {
	return &service.SkuVO{}, nil
}

func (s *SpuService) GetSaleExplain(ctx context.Context, empty *emptypb.Empty) (out *service.SaleExplains, err error) {
	panic("implement me")
}

func (s *SpuService) GetSpuByCategory(ctx context.Context, in *service.IdRequest) (out *service.SpuPage, err error) {
	panic("implement me")
}

func (s *SpuService) GetSpuById(ctx context.Context, in *service.IdRequest) (out *service.SpuDetail, err error) {
	panic("implement me")
}
func (s *SpuService) ListSpuByIds(ctx context.Context, in *service.IdsRequest) (out *service.Spus, err error) {
	pos, err := s.pu.ListSpuByIds(ctx, in.Ids)
	if err != nil {
		return
	}
	var items []*service.SpuVO
	for _, po := range pos {
		items = append(items, &service.SpuVO{
			Id:       po.Id,
			Title:    po.Title,
			Subtitle: po.Subtitle,
		})
	}

	return &service.Spus{
		Spus: items,
	}, nil
}
func (s *SpuService) GetSpuLatest(ctx context.Context, empty *emptypb.Empty) (out *service.SpuPage, err error) {
	panic("implement me")
}
