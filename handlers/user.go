package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/UsefulForMe/go-ecommerce/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService     services.UserService
	firebaseService services.FirebaseService
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
			logger.Error(err.Error())
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		req.PhoneNumber = utils.InternationPhoneToNational(req.PhoneNumber)

		if config.Cfg.IsProduction() {
			token, err := h.firebaseService.VerifyIDToken(req.IdToken)
			if err != nil {
				WriteResponseError(c, err)
				return
			}

			verifiedPhoneNumber := utils.InternationPhoneToNational(token.Claims["phone_number"].(string))
			if req.PhoneNumber != verifiedPhoneNumber {
				WriteResponseError(c, errs.NewUnauthenticatedError("phone number does not match"))
				return
			}
		}
		res, err := h.userService.Login(req)

		if err != nil {
			WriteResponseError(c, err)
		} else {
			print(res.Token)
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func NewUserHandler(userService services.UserService, firebaseService services.FirebaseService) UserHandler {
	return UserHandler{
		userService:     userService,
		firebaseService: firebaseService,
	}
}
