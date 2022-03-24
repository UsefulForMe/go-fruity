package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
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
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		res, err := h.userService.Create(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusCreated, res)
		}

	}
}

func (h UserHandler) GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {
		res, err := h.userService.List()
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}

	}
}

func (h UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginUserRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}
		res, err := h.userService.Login(req)

		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func NewUserHandler(service services.UserService) UserHandler {
	return UserHandler{
		userService: service,
	}
}
