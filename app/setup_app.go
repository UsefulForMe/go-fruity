package app

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	middleware "github.com/UsefulForMe/go-ecommerce/middlewares"
	router "github.com/UsefulForMe/go-ecommerce/routers"
	"github.com/gin-gonic/gin"
)

func SetupApp() *gin.Engine {

	config.InitDatabase()

	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(middleware.Cors(), middleware.RequestLogger(), gin.Recovery())
	router.SetupRoute(app)

	return app
}
