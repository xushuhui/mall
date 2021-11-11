package biz

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/xushuhui/goal/utils"
	"mall-go/internal/data/model"
)

func LocalUser(c *fiber.Ctx) (user model.User) {
	local := c.Locals("user")
	if local == nil {
		return
	}
	jwtToken := local.(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	bytes, _ := utils.JSONEncode(claims["user"])
	utils.JSONDecode(bytes, &user)
	return
}
