package api

import (
	"github.com/gofiber/fiber/v2"
	"mall-go/internal/biz"
	"mall-go/pkg/core"
)

func ActivityByName(c *fiber.Ctx) error {

	return nil
}

func ActivityNameWithSpu(c *fiber.Ctx) error {
	name := c.Params("name")

	data, err := biz.BannerByName(name)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}
