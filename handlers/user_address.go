package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserAddressHandler struct {
	userAddressService services.UserAddressService
}

func NewUserAddressHandler(userAddressService services.UserAddressService) UserAddressHandler {
	return UserAddressHandler{
		userAddressService: userAddressService,
	}
}

func (h UserAddressHandler) CreateUserAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateUserAddressRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		user := c.MustGet("user").(models.User)
		req.UserID = user.ID

		resp, err := h.userAddressService.CreateUserAddress(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusCreated, resp)
		}
	}
}

func (h UserAddressHandler) MyAddresses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.MyAddressesRequest
		user := c.MustGet("user").(models.User)
		req.UserID = user.ID
		resp, err := h.userAddressService.MyAddresses(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, resp)
		}
	}
}

func (h UserAddressHandler) GetUserAddressByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.GetUserAddressByIDRequest
		userAddressId := uuid.MustParse(c.Param("id"))
		req.UserAddressID = userAddressId
		resp, err := h.userAddressService.GetUserAddressByID(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, resp)
		}
	}
}

func (h UserAddressHandler) UpdateUserAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.UpdateUserAddressRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		userAddressId := uuid.MustParse(c.Param("id"))
		req.UserAddressID = userAddressId

		user := c.MustGet("user").(models.User)

		if user.ID != req.UserID {
			WriteResponseError(c, errs.NewForbiddenError("user id not match"))
			return
		}

		resp, err := h.userAddressService.UpdateUserAddress(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, resp)
		}
	}
}

func (h UserAddressHandler) DeleteUserAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.DeleteUserAddressRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		user := c.MustGet("user").(models.User)

		if user.ID != req.UserID {
			WriteResponseError(c, errs.NewForbiddenError("user id not match"))
			return
		}

		userAddressId := uuid.MustParse(c.Param("id"))
		req.UserAddressID = userAddressId

		err := h.userAddressService.DeleteUserAddress(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, nil)
		}
	}
}
