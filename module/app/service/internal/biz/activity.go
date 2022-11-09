package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Activity struct {
	Id             int64
	Title          string
	Description    string
	StartTime      time.Time
	EndTime        time.Time
	Remark         string
	Online         int
	EntranceImg    string
	InternalTopImg string
}

type ActivityCoupon struct {
	Activity
	Coupons []Coupons
}

type Coupons struct {
	Id          int64
	Title       string
	StartTime   time.Time
	EndTime     time.Time
	Description string
	FullMoney   float64
	Rate        float64
	Minus       float64
	Type        int
	Remark      string
	WholeStore  int
}

type ActivityRepo interface {
	GetActivityByName(ctx context.Context, name string) (a Activity, err error)
	// GetActivityWithCoupon(ctx context.Context, name string) (a ActivityCoupon, err error)
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

func (uc *ActivityUsecase) GetActivityWithCoupon(ctx context.Context, name string) (a ActivityCoupon, err error) {
	// a, err = uc.repo.GetActivityWithCoupon(ctx, name)
	return
}
