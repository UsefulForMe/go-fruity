package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	middleware "github.com/UsefulForMe/go-ecommerce/middlewares"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup) {
	r := models.NewUserRepository(config.DB)
	h := handlers.NewUserHandler(services.NewUserService(r))

	route.POST("login", h.Login())

	userRouer := route.Group("/users")

	userRouer.Use(middleware.VerifyJWT(r))
	{
		userRouer.GET("", h.GetAll())
		userRouer.POST("", h.Create())
	}

}
