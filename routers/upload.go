package router

import (
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func UploadRouter(route *gin.RouterGroup) {

	h := handlers.NewUploadHandler(services.NewS3Service())

	uploadRouter := route.Group("/upload")
	
	uploadRouter.POST("/file", h.Upload())
	uploadRouter.POST("/files", h.UploadMany())
	uploadRouter.PUT("/file/delete", h.Delete())
	uploadRouter.PUT("/files/delete", h.DeleteMany())

}
