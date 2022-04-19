package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func UserRouter(userRouer *gin.RouterGroup) {
	r := repository.NewUserRepository(config.DB)
	h := handlers.NewUserHandler(services.NewUserService(r))

	userRouer.GET("", h.GetAll())
	userRouer.POST("", h.Create())

	userRouer.PUT("/update-my-fcm-token", h.UpdateFCMToken())

}
