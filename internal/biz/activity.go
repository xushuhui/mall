package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Activity struct {
	Id             uint32
	Title          string
	Description    string
	StartTime      string
	EndTime        string
	Remark         string
	Online         bool
	EntranceImg    string
	InternalTopImg string
}
type ActivityCoupon struct {
}
type ActivityRepo interface {
	GetActivityByName(ctx context.Context, name string) (a Activity, err error)
	GetActivityWithCoupon(ctx context.Context, name string) (a ActivityCoupon, err error)
}
type ActivityUsecase struct {
	repo ActivityRepo
	log  *log.Helper
}

func NewActivityUsecase(repo ActivityRepo, logger log.Logger) *ActivityUsecase {
	return &ActivityUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ActivityUsecase) GetActivityByName(ctx context.Context, name string) (a Activity, err error) {
	a, err = uc.repo.GetActivityByName(ctx, name)
	return
}
