package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xushuhui/goal/core"
	"mall-go/internal/biz"
)

func ActivityByName(c *fiber.Ctx) error {

	return nil
}

func ActivityNameWithSpu(c *fiber.Ctx) error {
	name := c.Params("name")

	data, err := biz.BannerByName(c.Context(), name)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}
