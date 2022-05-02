package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func UserRouter(userRouter *gin.RouterGroup) {
	r := repository.NewUserRepository(config.DB)
	h := handlers.NewUserHandler(services.NewUserService(r))

	userRouter.GET("", h.GetAll())
	userRouter.POST("", h.Create())

	userRouter.PUT("/update-my-fcm-token", h.UpdateFCMToken())
	userRouter.PUT("/update-my-profile", h.UpdateInfor())

}
