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
	route.GET("", h.GetProductAll())
	route.GET("/get-top-sale", h.GetTopSaleProduct())
	route.GET("/get-sale-off", h.GetProductsSaleOff())
	route.GET("/get-sale-shock", h.GetProductSaleShock())
	route.GET("/get-by-ids", h.GetProductsByID())
	route.GET("/:id", h.GetProductByID())
}

func ProductRouteCMS(route *gin.RouterGroup) {
	h := handlers.NewProductHandler(services.NewProductService(repository.NewProductRepository(config.DB)))
	route.GET("", h.GetProductAll()).POST("", h.CreateProduct())
	route.GET("/:id", h.GetProductByID())
}
