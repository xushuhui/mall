package biz

import "github.com/go-kratos/kratos/v2/log"

type SpuRepo interface {
}
type SpuUsecase struct {
	repo SpuRepo
	log  *log.Helper
}

func NewSpuUsecase(repo SpuRepo, logger log.Logger) *SpuUsecase {
	return &SpuUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
type SkuRepo interface {

}
type SkuUsecase struct {
	repo SkuRepo
	log  *log.Helper
}
func NewSkuUsecase(repo SkuRepo, logger log.Logger) *SkuUsecase {
	return &SkuUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}