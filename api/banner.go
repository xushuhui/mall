package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall_go/internal/services"
	"mall_go/pkg/core"
	"mall_go/pkg/utils"
)

func BannerID(c *gin.Context) {
	id, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	data, err := services.BannerById(id)
	if err != nil {
		return
	}
	core.SetData(c, data)
	return

}
func BannerName(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(name)
	data, err := services.BannerByName(name)
	if err != nil {
		return
	}
	core.SetData(c, data)
	return

}
