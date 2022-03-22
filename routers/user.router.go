package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserGroup(route *gin.RouterGroup) {
	route.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{
				"message": "hello from user",
			},
		)
	})
}
