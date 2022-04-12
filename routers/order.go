package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func OrderRoute(route *gin.RouterGroup) {
	h := handlers.NewOrderHandler(services.NewOrderService(repository.NewOrderRepository(config.DB)))
	route.GET("", h.MyOrders()).POST("", h.CreateOrder())
}
