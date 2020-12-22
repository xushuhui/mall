package api

import (
	"github.com/gin-gonic/gin"
	"mall_go/internal/logic"
	"mall_go/pkg/core"
)

func ActivityByName(c *gin.Context) {

	name := c.Param("name")

	data, err := logic.BannerByName(name)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return

}

func ActivityNameWithSpu(c *gin.Context) {
	name := c.Param("name")

	data, err := logic.BannerByName(name)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return

}
