package middleware

import (
	"github.com/gin-gonic/gin"
	"mall_go/pkg/core"
	"mall_go/pkg/errcode"
	"mall_go/pkg/lib"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		Authorization := c.Request.Header.Get("Authorization")

		if Authorization == "" {
			core.InvalidParamsResp(c, "empty Authorization")
			return
		}

		claims, e := lib.ParseToken(Authorization)
		if e != nil {
			core.FailResp(c, errcode.ErrorAuthToken)
			return

		}
		if time.Now().Unix() > claims.ExpiresAt {
			core.FailResp(c, errcode.TimeoutAuthToken)
			return
		}
		c.Set("uid", claims.Uid)
		c.Next()
	}

}
