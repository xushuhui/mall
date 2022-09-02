package biz

import (
	"context"
	"errors"
	"mall-go/app/mall/interface/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt"
	"github.com/medivhzhan/weapp/v2"
)

type User struct {
	Id           int64
	Nickname     string
	IdentityType string
	Identifier   string
}

type UserRepo interface {
	CreateUser(ctx context.Context, in *User) (u int64, err error)
	GetOpenidByCode(ctx context.Context, code string) (resp *weapp.LoginResponse, err error)
	GetUserIdentity(ctx context.Context, identityType, identifier, credential string) (User, error)
	CreateUserIdentity(ctx context.Context, identityType, identifier, credential string) (User, error)
}
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
	conf *conf.App
}

func NewUserUsecase(repo UserRepo, logger log.Logger, c *conf.App) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		conf: c,
	}
}

func (u *UserUsecase) MinappLogin(ctx context.Context, code string) (out string, err error) {
	resp, err := u.repo.GetOpenidByCode(ctx, code)
	if err != nil {
		return
	}
	user, err := u.repo.GetUserIdentity(ctx, "weapp", resp.OpenID, "")
	if err != nil {
		return
	}

	var id int64
	if user.Id == 0 {
		id, err = u.repo.CreateUser(ctx, &User{Nickname: "小程序用户", IdentityType: "weapp", Identifier: resp.OpenID})
		if err != nil {
			return
		}
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
	})
	signedString, err := claims.SignedString([]byte(""))
	if err != nil {
		return
	}
	return signedString, nil
}

func (u *UserUsecase) VerifyToken(ctx context.Context, tokenString string) (isValid bool, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if err.(*jwt.ValidationError).Errors&jwt.ValidationErrorMalformed != 0 {
		err = errors.New("that's not even a token")
		return
	}
	if err.(*jwt.ValidationError).Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
		err = errors.New("token expire time")
		return
	}
	if err != nil {
		err = errors.New("token error")
		return
	}

	if token.Valid {
		isValid = true
	}
	return
}
