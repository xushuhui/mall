package api

import (
	"github.com/gofiber/fiber/v2"
	"mall_go/internal/biz"
	"mall_go/pkg/core"
)

func ActivityByName(c *fiber.Ctx) error {

	name := c.Params("name")

	data, err := biz.ActivityByName(name)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}

func ActivityNameWithSpu(c *fiber.Ctx) error {
	name := c.Params("name")

	data, err := biz.BannerByName(name)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}
