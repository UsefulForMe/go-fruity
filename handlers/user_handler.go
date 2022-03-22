package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func (h UserHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateUserRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		res, err := h.userService.Create(req)
		if err != nil {
			c.JSON(err.Code, err.Error())
		} else {
			c.JSON(http.StatusOK, res)
		}

	}
}

func (h UserHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := h.userService.List()
		if err != nil {
			c.JSON(err.Code, err.Error())
		} else {
			c.JSON(http.StatusOK, res)
		}

	}
}

func NewUserHandler(service services.UserService) UserHandler {
	return UserHandler{
		userService: service,
	}
}
