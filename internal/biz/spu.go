package biz

import "mall_go/internal/model"

func SpuById(id uint) (spu model.Spu, err error) {
	spu, err = model.GetSpu(id)
	return
}
func SpuByCategory(id uint) (spu []model.Spu, err error) {
	spu, err = model.GetSpuByCategory(id)
	return
}
