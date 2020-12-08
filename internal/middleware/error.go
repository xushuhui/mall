package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mall_go/pkg/core"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		e := c.Errors.Last()
		if e == nil {
			return
		}
		errors := e.Err

		switch errors.(type) {
		case error:
			if errors == gorm.ErrRecordNotFound {
				core.NotFoundResp(c, e.Err.Error())
			}
		case core.Error:

			codeErr := errors.(core.Error)

			core.FailResp(c, codeErr.Code)

		case *json.UnmarshalTypeError:
			unmarshalTypeError := errors.(*json.UnmarshalTypeError)
			errStr := fmt.Errorf("%s 类型错误，期望类型 %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String()).Error()
			core.InvalidParamsResp(c, errStr)
		default:

			errStr := errors.Error()

			core.InvalidParamsResp(c, errStr)
		}

	}

}
