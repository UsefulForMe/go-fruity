package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func SellerRoute(route *gin.RouterGroup) {

	h := handlers.NewSellerHandler(services.NewSellerService(repository.NewSellerRepository(config.DB)))

	route.GET("", h.GetAllSellers()).POST("", h.CreateSeller())
	route.GET("/:id", h.GetSellerByID())
	route.GET("/:id/products", h.GetProductsBySellerID())
}
