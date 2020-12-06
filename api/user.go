package api

import (
	"mall_go/internal/request"
	"mall_go/internal/services"
	"mall_go/pkg/core"

	"github.com/gin-gonic/gin"
)

/*

登录

URL:
	/login

POST参数：

	"phone":"2" //手机号
	"password": "1" //密码


返回值：

	"code": 0
	"message": "ok"

*/
func Login(c *gin.Context) {
	var req request.Login

	if e := core.ParseRequest(c, &req); e != nil {

		c.Error(e)
		return
	}

	data, e := services.Login(req)
	if e != nil {
		c.Error(e)
		return
	}
	core.SetData(c, data)
	return

}
