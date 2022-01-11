package biz

import (
	"context"
	"mall-go/api/mall"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt"
)

type User struct {
	Id int64
}

type UserRepo interface {
	CreateUser(ctx context.Context)
	GetUserByAccount(ctx context.Context, account string)
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

func (u *UserUsecase) Login(ctx context.Context, account string) (out *mall.LoginReply, err error) {
	// verify
	// generate token

	var userId int64
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
	})
	signedString, err := claims.SignedString([]byte(u.jwtKey))
	if err != nil {
		return nil, mall.ErrorLoginfail("generate token failed: %s", err.Error())
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
