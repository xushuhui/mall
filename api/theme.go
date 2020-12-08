package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall_go/internal/services"
	"mall_go/pkg/core"
	"strings"
)

func ThemeByNames(c *gin.Context) {
	names := c.Query("names")

	nameSlice := strings.Split(names, ",")
	data, err := services.ThemeByNames(nameSlice)
	fmt.Println("err", err)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return

}

func ThemeNameWithSpu(c *gin.Context) {
	name := c.Param("name")

	data, err := services.ThemeNameWithSpu(name)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}
