package server

import (
	"github.com/gofiber/fiber/v2"
	"mall-go/api"
)

func InitRoute(app *fiber.App) {

	app.Post("/login", api.Login)

	banner := app.Group("/banner")
	{
		banner.Get("/id/:id", api.BannerID)
		banner.Get("/name/:name", api.BannerName)
	}
	theme := app.Group("/theme")
	{
		theme.Get("/theme/by/names", api.ThemeByNames)
		theme.Get("/theme/name/:name/with_spu", api.ThemeNameWithSpu)
	}
	activity := app.Group("/activity")
	{
		activity.Get("/name/:name", api.ActivityByName)
		activity.Get("/name/:name/with_spu", api.ActivityNameWithSpu)
	}
	coupon := app.Group("/coupon")
	{
		coupon.Get("/by/category/:cid", api.CouponByCategory)
		coupon.Get("/whole_store", api.CouponWholeStore)
		coupon.Get("/myself/by/status/:status", api.CouponMyselfByStatus)
		coupon.Get("/myself/available", api.CouponMyselfAvailable)
		coupon.Post("/collect/:id", api.CouponCollect)

	}
	spu := app.Group("/spu")
	{
		spu.Get("/id/:id/detail", api.Detail)
		spu.Get("/latest", api.Latest)
		spu.Get("/by/category/:id", api.SpuByCategory)
	}
	app.Get("/search", api.Search)
	app.Get("/sale_explain/fixed", api.SaleExplainFixed)
	app.Get("/category/grid/all", api.CategoryGrid)
	app.Get("/category/all", api.CategoryAll)
	app.Get("/tag/type/:type", api.TagType)

	//app.Use(middleware.Auth())
	//需要登录的接口

}
