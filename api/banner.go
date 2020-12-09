package api

import (
	"github.com/gin-gonic/gin"
	"mall_go/internal/logic"
	"mall_go/pkg/core"
	"mall_go/pkg/utils"
)

func BannerID(c *gin.Context) {
	id, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		return
	}
	data, err := logic.BannerById(id)
	if err != nil {
		return
	}
	core.SetData(c, data)
	return

}
func BannerName(c *gin.Context) {
	name := c.Param("name")

	data, err := logic.BannerByName(name)
	if err != nil {
		return
	}
	core.SetData(c, data)
	return

}
