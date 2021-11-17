package api

import (
	"github.com/gofiber/fiber/v2"
	"mall-go/internal/biz"

	"github.com/xushuhui/goal/core"
	"strings"
)

func ThemeByNames(c *fiber.Ctx) (err error) {
	names := c.Query("names")

	data, err := biz.ThemeByNames(c.Context(), strings.Split(names, ","))
	if err != nil {

		return
	}
	return core.SetData(c, data)

}

func ThemeNameWithSpu(c *fiber.Ctx) error {
	name := c.Params("name")

	data, err := biz.ThemeNameWithSpu(c.Context(), name)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}
