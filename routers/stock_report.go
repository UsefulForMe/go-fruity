package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func StockReportRouteCMS(route *gin.RouterGroup) {

	h := handlers.NewStockReportHandler(services.NewStockReportService(repository.NewStockReportRepository(config.DB)))

	route.GET("", h.GetAllStockReport())
	route.POST("", h.SaveStockReport())
	route.PUT("/:stock_report_id", h.UpdateStockReport())
	route.GET("/:stock_report_id", h.GetStockReportById())
	// route.GET("/products/:product_id", h.GetStockByProductId())
}
