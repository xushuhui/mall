package data

import (
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
