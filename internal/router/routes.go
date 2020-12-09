package router

import (
	"github.com/gin-gonic/gin"
	"mall_go/api"
	"mall_go/internal/middleware"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	router.StaticFS("/upload", http.Dir("storage/upload"))

	router.Use(middleware.Cors(), middleware.AccessLog(), middleware.ErrorHandle())
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gsssssn")
	})

	router.POST("/login", api.Login)

	banner := router.Group("/banner")
	{
		banner.GET("/id/:id", api.BannerID)
		banner.GET("/name/:name", api.BannerName)
	}
	theme := router.Group("/theme")
	{
		theme.GET("/theme/by/names", api.ThemeByNames)
		theme.GET("/theme/name/:name/with_spu", api.ThemeNameWithSpu)
	}
	spu := router.Group("/spu")
	{
		spu.GET("/id/:id/detail", api.Detail)
		spu.GET("/latest", api.Latest)
		spu.GET("/by/category/:id", api.SpuByCategory)
	}
	//router.Use(middleware.Auth())
	//需要登录的接口
	return router
}
