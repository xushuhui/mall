package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xushuhui/goal/core"
	"github.com/xushuhui/goal/utils"
	"mall-go/internal/biz"
)

func BannerID(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {

		return err
	}
	data, err := biz.BannerById(c.Context(), id)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}
func BannerName(c *fiber.Ctx) error {
	name := c.Params("name")

	data, err := biz.BannerByName(c.Context(), name)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}
