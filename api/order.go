package api

import (
	"mall-go/internal/biz"
	"mall-go/internal/request"

	"github.com/xushuhui/goal/core"

	"github.com/gofiber/fiber/v2"
)

func PlaceOrder(c *fiber.Ctx) (err error) {
	var req request.PlaceOrder

	if err = core.ParseRequest(c, &req); err != nil {
		return
	}

	user := biz.LocalUser(c)
	orderChecker, err := biz.OrderIsOk(c.Context(),user.ID, req)
	if err != nil {
		return err
	}
	orderId, err := biz.PlaceOrder(user.ID, req, orderChecker)
	if err != nil {
		return err
	}
	m := make(map[string]interface{})
	m["id"] = orderId
	return core.SetData(c, m)
}
