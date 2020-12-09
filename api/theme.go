package api

import (
	"github.com/gin-gonic/gin"
	"mall_go/internal/logic"
	"mall_go/pkg/core"
	"strings"
)

func ThemeByNames(c *gin.Context) {
	names := c.Query("names")

	nameSlice := strings.Split(names, ",")
	data, err := logic.ThemeByNames(nameSlice)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return

}

func ThemeNameWithSpu(c *gin.Context) {
	name := c.Param("name")

	data, err := logic.ThemeNameWithSpu(name)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}
