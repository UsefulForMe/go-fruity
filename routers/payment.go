package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func PaymentRoute(route *gin.RouterGroup) {
	h := handlers.NewPaymentHandler(services.NewPaymentService(repository.NewPaymentRepository(config.DB)))
	route.GET("", h.MyPayments()).POST("", h.CreatePayment())
}
