package api

import (
	"github.com/gofiber/fiber/v2"
	"mall-go/internal/biz"
	"mall-go/internal/request"
	"mall-go/pkg/core"
)

func Login(c *fiber.Ctx) (err error) {
	var req request.Login

	if err = core.ParseRequest(c, &req); err != nil {
		return
	}

	data, err := biz.Login(req)
	if err != nil {
		return
	}
	return core.SetData(c, data)

}
