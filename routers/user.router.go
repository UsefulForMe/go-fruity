package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup) {

	h := handlers.NewUserHandler(services.NewUserService(models.NewUserRepository(config.DB)))

	route.GET("", h.GetAll())

	route.POST("", h.Create())
}
