package data

import (
	"context"
	"mall-go/app/spu/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type spuRepo struct {
	data *Data
	log  *log.Helper
}

func NewSpuRepo(data *Data, logger log.Logger) biz.SpuRepo {
	return &spuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *spuRepo) GetSpuById(ctx context.Context, id int64) (Spu biz.Spu, err error) {

	return
}
func (r *spuRepo) GetSpuByCategory(ctx context.Context, id int64) (Spus []biz.Spu, err error) {
	return
}
func (r *spuRepo) ListSpuByIds(ctx context.Context, ids []int64) (spuList []biz.Spu, err error) {
	//pos, err := r.data.db.Spu.Query().Where(spu.IDIn(ids...)).All(ctx)
	//if err != nil {
	//	return
	//}
	//for _, po := range pos {
	//	spuList = append(spuList, biz.Spu{
	//		Id:       po.ID,
	//		Title:    po.Title,
	//		Subtitle: po.Subtitle,
	//		Img:      po.Img,
	//	})
	//}

	return
}
func (r *spuRepo) CreateSpu(ctx context.Context) (err error) {
	return
}
