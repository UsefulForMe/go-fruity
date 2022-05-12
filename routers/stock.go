package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func StockRoute(route *gin.RouterGroup) {
	h := handlers.NewStockHandler(services.NewStockService(repository.NewStockRepository(config.DB)))
	route.GET("/products/:product_id", h.GetStockByProductId())
}

func StockRouteCMS(route *gin.RouterGroup) {

	h := handlers.NewStockHandler(services.NewStockService(repository.NewStockRepository(config.DB)))

	route.GET("", h.GetAllStock())
	route.POST("", h.SaveStock())
	route.PUT("/:stock_id", h.UpdateStock())
	route.GET("/:stock_id", h.GetStockById())
	route.GET("/products/:product_id", h.GetStockByProductId())
}
