package app

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	router "github.com/UsefulForMe/go-ecommerce/routers"
	"github.com/gin-gonic/gin"
)

func SetupApp() *gin.Engine {

	config.InitDatabase()

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	// app.Use(middleware.Cors(), middleware.RequestLogger(), gin.Recovery())
	router.SetupRoute(app)

	return app
}
