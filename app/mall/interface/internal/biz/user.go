package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt"
	"github.com/medivhzhan/weapp/v2"
	"mall-go/api/mall"
	"mall-go/api/user/service"
)

type User struct {
	Id int64
}

type UserRepo interface {
	CreateUser(ctx context.Context, in *service.CreateUserRequest) (u User, err error)
	GetOpenidByCode(ctx context.Context, code string) (resp *weapp.LoginResponse, err error)
	GetUserIdentiy(ctx context.Context, identityType, identifier, credential string) (User, error)
	CreateUserIdentiy(ctx context.Context, identityType, identifier, credential string) (User, error)
}
type UserUsecase struct {
	repo   UserRepo
	jwtKey string
	log    *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (u *UserUsecase) MinappLogin(ctx context.Context, code string) (out *mall.LoginReply, err error) {
	resp, err := u.repo.GetOpenidByCode(ctx, code)
	if err != nil {
		return
	}
	user, err := u.repo.GetUserIdentiy(ctx, "weapp", resp.OpenID, "")
	if err != nil {
		return nil, mall.ErrorLoginFail("GetUserIdentiy failed: %s", err.Error())
	}
	//TODO notfound error
	if user.Id == 0 {
		user, err = u.repo.CreateUser(ctx, &service.CreateUserRequest{Nickname: "小程序用户", IdentityType: "weapp", Identifier: resp.OpenID})
		if err != nil {
			return nil, mall.ErrorLoginFail("CreateUserIdentiy failed: %s", err.Error())
		}
	}
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
	})
	signedString, err := claims.SignedString([]byte(u.jwtKey))
	if err != nil {
		return nil, mall.ErrorLoginFail("generate token failed: %s", err.Error())
	}
	return &mall.LoginReply{
		Token: signedString,
	}, nil

}

func (u *UserUsecase) VerifyToken(ctx context.Context, tokenString string) (isValid bool, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.jwtKey), nil
	})
	if err.(*jwt.ValidationError).Errors&jwt.ValidationErrorMalformed != 0 {
		err = mall.ErrorToken("That's not even a token")
		return
	}
	if err.(*jwt.ValidationError).Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
		err = mall.ErrorToken("token expire time")
		return
	}
	if err != nil {
		err = mall.ErrorToken("token error")
		return
	}

	if token.Valid {
		isValid = true
	}
	return
}
