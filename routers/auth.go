package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func AuthRouter(route *gin.RouterGroup) {

	h := handlers.NewAuthHandler(services.NewUserService(repository.NewUserRepository(config.DB)), services.NewFirebaseService(config.FirebaseApp))

	route.POST("login", h.Login())

}
