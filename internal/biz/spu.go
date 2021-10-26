package biz

import (
	"mall-go/internal/data"
)

func SpuById(id int) (spu data.Spu, err error) {
	spu, err = data.GetSpuById(id)
	return
}
func SpuByCategory(id int) (spu []data.Spu, err error) {
	spu, err = data.GetSpuByCategory(id)
	return
}
