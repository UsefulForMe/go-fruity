package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func UserAddressRoute(router *gin.RouterGroup) {
	h := handlers.NewUserAddressHandler(services.NewUserAddressService(repository.NewUserAddressRepository(config.DB)))
	router.GET("", h.MyAddresses()).POST("", h.CreateUserAddress())

}
