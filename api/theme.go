package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall_go/internal/services"
	"mall_go/pkg/core"
)

func ThemeByNames(c *gin.Context) {
	names := c.Param("names")

	data, err := services.ThemeByNames(names)
	if err != nil {
		return
	}
	core.SetData(c, data)
	return

}

func ThemeNameWithSpu(c *gin.Context) {
	name := c.Param("name")
	fmt.Println("n", name)
	data, err := services.ThemeNameWithSpu(name)
	if err != nil {
		return
	}
	core.SetData(c, data)
	return
}
