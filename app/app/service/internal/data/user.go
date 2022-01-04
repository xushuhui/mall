package data

import (
	"context"
	"mall-go/internal/biz"

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
func (r *userRepo) GenerateToken(ctx context.Context, account string) (token string, err error) {
	panic("not implemented") // TODO: Implement
}

func (r *userRepo) VerifyToken(ctx context.Context, token string) (isValid bool, err error) {
	panic("not implemented") // TODO: Implement
}
