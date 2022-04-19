package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
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

func (h UserHandler) UpdateFCMToken() gin.HandlerFunc {

	return func(c *gin.Context) {

		var req dto.UpdateFCMTokenRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}
		user := c.MustGet("user").(models.User)
		req.UserID = user.ID
		res, err := h.userService.UpdateFCMToken(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}

	}
}

func NewUserHandler(userService services.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}
