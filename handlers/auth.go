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

type AuthHandler struct {
	userService         services.UserService
	firebaseAuthService services.FirebaseAuthService
}

func (h AuthHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginUserRequest

		if err := c.BindJSON(&req); err != nil {
			logger.Error(err.Error())
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		req.PhoneNumber = utils.InternationPhoneToNational(req.PhoneNumber)

		if config.Cfg.IsProduction() {
			token, err := h.firebaseAuthService.VerifyIDToken(req.IdToken)
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

func NewAuthHandler(userService services.UserService, firebaseAuthService services.FirebaseAuthService) AuthHandler {
	return AuthHandler{
		userService:         userService,
		firebaseAuthService: firebaseAuthService,
	}
}
