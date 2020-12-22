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
	activity := router.Group("/activity")
	{
		activity.GET("/name/:name", api.ActivityByName)
		activity.GET("/name/:name/with_spu", api.ActivityNameWithSpu)
	}
	coupon := router.Group("/coupon")
	{
		coupon.GET("/by/category/:cid", api.CouponByCategory)
		coupon.GET("/whole_store", api.CouponWholeStore)
		coupon.GET("/myself/by/status/:status", api.CouponMyselfByStatus)
		coupon.GET("/myself/available", api.CouponMyselfAvailable)
		coupon.POST("/collect/:id", api.CouponCollect)

	}
	spu := router.Group("/spu")
	{
		spu.GET("/id/:id/detail", api.Detail)
		spu.GET("/latest", api.Latest)
		spu.GET("/by/category/:id", api.SpuByCategory)
	}
	router.GET("/search", api.Search)
	router.GET("/sale_explain/fixed", api.SaleExplainFixed)
	router.GET("/category/grid/all", api.CategoryGrid)
	router.GET("/category/all", api.CategoryAll)
	router.GET("/tag/type/:type", api.TagType)

	//router.Use(middleware.Auth())
	//需要登录的接口
	return router
}
