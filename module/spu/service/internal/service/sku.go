package service

import (
	"context"

	"mall-go/module/spu/service/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SpuService struct {
	UnimplementedSpuServer
	su *biz.SkuUsecase
	pu *biz.SpuUsecase
}

func NewSpuService(su *biz.SkuUsecase, pu *biz.SpuUsecase) *SpuService {
	return &SpuService{
		su: su,
		pu: pu,
	}
}

func (s *SpuService) GetSkuById(ctx context.Context, req *IdRequest) (*SkuVO, error) {
	return &SkuVO{}, nil
}

func (s *SpuService) GetSaleExplain(ctx context.Context, empty *emptypb.Empty) (out *SaleExplains, err error) {
	panic("implement me")
}

func (s *SpuService) GetSpuByCategory(ctx context.Context, in *IdRequest) (out *SpuPage, err error) {
	panic("implement me")
}

func (s *SpuService) GetSpuById(ctx context.Context, in *IdRequest) (out *SpuDetail, err error) {
	panic("implement me")
}

func (s *SpuService) ListSpuByIds(ctx context.Context, in *IdsRequest) (out *Spus, err error) {
	pos, err := s.pu.ListSpuByIds(ctx, in.Ids)
	if err != nil {
		return
	}
	var items []*SpuVO
	for _, po := range pos {
		items = append(items, &SpuVO{
			Id:       po.Id,
			Title:    po.Title,
			Subtitle: po.Subtitle,
		})
	}

	return &Spus{
		Spus: items,
	}, nil
}

func (s *SpuService) GetSpuLatest(ctx context.Context, empty *emptypb.Empty) (out *SpuPage, err error) {
	panic("implement me")
}
