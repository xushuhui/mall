package data

import (
	"context"

	user "mall-go/api/user/service"
	"mall-go/app/mall/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/medivhzhan/weapp/v2"
)

type userRepo struct {
	log  *log.Helper
	data *Data
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) GetOpenidByCode(ctx context.Context, code string) (resp *weapp.LoginResponse, err error) {
	resp, err = weapp.Login(r.data.wc.Appid, r.data.wc.Secret, code)
	if err != nil {
		return
	}
	err = resp.CommonError.GetResponseError()
	return
}

func (r *userRepo) GetUserIdentiy(ctx context.Context, identityType, identifier, credential string) (u biz.User, err error) {
	po, err := r.data.uc.GetUserIdentiy(ctx, &user.UserIdentiyRequest{Identifier: identifier, Credential: credential, IdentityType: identityType})
	if err != nil {
		return
	}

	return biz.User{
		Id: po.Id,
	}, nil
}

func (r *userRepo) CreateUserIdentiy(ctx context.Context, identityType, identifier, credential string) (u biz.User, err error) {
	po, err := r.data.uc.CreateUserIdentiy(ctx, &user.UserIdentiyRequest{Identifier: identifier, Credential: credential, IdentityType: identityType})
	if err != nil {
		return
	}
	return biz.User{
		Id: po.Id,
	}, nil
}

func (r *userRepo) CreateUser(ctx context.Context, in *biz.User) (u int64, err error) {

	po, err := r.data.uc.CreateUser(ctx, &user.CreateUserRequest{
		Nickname:     "",
		Identifier:   "",
		Credential:   "",
		IdentityType: "",
	})
	if err != nil {
		return
	}

	return po.Id, nil
}
