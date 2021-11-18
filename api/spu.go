package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func SaleExplainFixed(c *fiber.Ctx) error {
	return nil
}
func Search(c *fiber.Ctx) error {
	q := c.Query("q")
	fmt.Println(q)
	return nil
}
func TagType(c *fiber.Ctx) error {
	types := c.Params("type")
	fmt.Println(types)
	return nil
}
func Detail(c *fiber.Ctx) error {

	return nil

}
func Latest(c *fiber.Ctx) error {
	return nil
}
func SpuByCategory(c *fiber.Ctx) (err error) {

	return nil
}
