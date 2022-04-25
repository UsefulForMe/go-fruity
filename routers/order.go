package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func OrderRoute(route *gin.RouterGroup) {
	firebaseFCM := services.NewFirebaseMessageService(config.FirebaseApp)
	h := handlers.NewOrderHandler(services.NewOrderService(repository.NewOrderRepository(config.DB)), firebaseFCM)
	route.GET("", h.MyOrders()).POST("", h.CreateOrder())

	route.GET("/:order_id", h.GetOrderByID())
}
