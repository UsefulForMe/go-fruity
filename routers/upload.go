package router

import (
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func UploadRouter(route *gin.RouterGroup) {

	h := handlers.NewUploadHandler(services.NewS3Service())
	route.POST("/file", h.Upload())
	route.POST("/files", h.UploadMany())
	route.PUT("/file/delete", h.Delete())
	route.PUT("/files/delete", h.DeleteMany())

}
