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
	v1 := app.Group("/v1")

	AuthRouter(v1)

	CategoryRoute(v1.Group("/categories"))
	ProductRoute(v1.Group("/products"))

	v1.Use(jwtMiddleware.Verify())
	{
		UserRouter(v1.Group("/users"))
		UploadRouter(v1.Group("/upload"))
	}

}
