package services

import (
	"golang.org/x/crypto/bcrypt"
	"mall_go/internal/model"
	"mall_go/internal/request"
	"mall_go/pkg/core"
	"mall_go/pkg/errcode"
	"mall_go/pkg/lib"
)

func Login(req request.Login) (data map[string]interface{}, e error) {

	userModel, e := model.GetAccountUserOne("phone=?", req.Phone)
	if e != nil {
		return
	}
	// 正确密码验证
	e = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password))
	if e != nil {
		e = core.NewError(errcode.ErrorPassWord)
		return
	}
	token, e := lib.GenerateToken(userModel.ID)
	if e != nil {
		return
	}
	data = make(map[string]interface{})
	data["access_token"] = token
	return
}
