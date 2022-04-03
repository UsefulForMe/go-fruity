package router

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func CategoryRoute(route *gin.RouterGroup) {

	h := handlers.NewCategoryHandler(services.NewCategoryService(repository.NewCategoryRepository(config.DB)))

	route.GET("", h.GetAllCategories()).POST("", h.CreateCategory())

}
