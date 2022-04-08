package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func ProductRoute(route *gin.RouterGroup) {
	h := handlers.NewProductHandler(services.NewProductService(repository.NewProductRepository(config.DB)))
	route.GET("", h.GetProductAll()).POST("", h.CreateProduct())
	route.GET("/get-top-sale", h.GetTopSaleProduct())
	route.GET("/:id", h.GetProductByID())
}
