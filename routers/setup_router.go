package router

import "github.com/gin-gonic/gin"

func SetupRoute(app *gin.Engine) {
	route := app.Group("/v1")

	UserRouter(route)

}
