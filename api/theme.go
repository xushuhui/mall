package api

import (
	"github.com/gofiber/fiber/v2"
	"mall-go/internal/biz"

	"mall-go/pkg/core"
	"strings"
)

func ThemeByNames(c *fiber.Ctx) (err error) {
	names := c.Query("names")

	nameSlice := strings.Split(names, ",")
	data, err := biz.ThemeByNames(nameSlice)
	if err != nil {

		return
	}
	return core.SetData(c, data)

}

func ThemeNameWithSpu(c *fiber.Ctx) error {
	name := c.Params("name")

	data, err := biz.ThemeNameWithSpu(name)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}
