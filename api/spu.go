package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall_go/internal/logic"
	"mall_go/pkg/core"
	"mall_go/pkg/utils"
)

func SaleExplainFixed(c *gin.Context) {

}
func Search(c *gin.Context) {
	q := c.Query("q")
	fmt.Println(q)
}
func TagType(c *gin.Context) {
	types := c.Param("type")
	fmt.Println(types)
}
func Detail(c *gin.Context) {

	id, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	data, err := logic.SpuById(id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}
func Latest(c *gin.Context) {

}
func SpuByCategory(c *gin.Context) {
	id, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	data, err := logic.SpuByCategory(id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return

}
