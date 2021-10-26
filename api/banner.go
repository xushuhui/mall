package api

import (
	"github.com/gofiber/fiber/v2"
	"mall-go/internal/biz"
	"mall-go/pkg/core"
	"mall-go/pkg/utils"
)

func BannerID(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {

		return err
	}
	data, err := biz.BannerById(id)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}
func BannerName(c *fiber.Ctx) error {
	name := c.Params("name")

	data, err := biz.BannerByName(name)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}
