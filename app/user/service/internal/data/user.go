package data

import (
	"context"
	"mall-go/app/user/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	Data *Data
	Log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		Data: data,
		Log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context) (err error) {
	panic("not implemented") // TODO: Implement
}

func (r *userRepo) GetUserByAccount(ctx context.Context, account string) {
	panic("not implemented") // TODO: Implement
}
