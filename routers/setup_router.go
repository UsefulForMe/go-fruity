package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	middleware "github.com/UsefulForMe/go-ecommerce/middlewares"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/gin-gonic/gin"
)

func SetupRoute(app *gin.Engine) {
	jwtMiddleware := middleware.NewJWTMiddleware(repository.NewUserRepository(config.DB))

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := app.Group("/api/v1")

	AuthRouter(v1)

	CategoryRoute(v1.Group("/categories"))
	ProductRoute(v1.Group("/products"))
	StockRoute(v1.Group("/stocks"))
	SellerRoute(v1.Group("/sellers"))

	v1.Use(jwtMiddleware.Verify())
	{
		UserRouter(v1.Group("/users"))
		UploadRouter(v1.Group("/upload"))
		PaymentRoute(v1.Group("/payments"))
		OrderRoute(v1.Group("/orders"))
		UserAddressRoute(v1.Group("/user-addresses"))
	}

	// setup for cms
	cmsV1 := app.Group("/cms/v1")
	CategoryRouteCMS(cmsV1.Group("/categories"))
	ProductRouteCMS(cmsV1.Group("/products"))
	OrderRouteCMS(cmsV1.Group("/orders"))
	SellerRouteCMS(cmsV1.Group("/sellers"))
	StockRouteCMS(cmsV1.Group("/stocks"))
	StockReportRouteCMS(cmsV1.Group("/stocks-reports"))
	UserRouterCMS(cmsV1.Group("/users"))
	UploadRouter(cmsV1.Group("/upload"))

}
