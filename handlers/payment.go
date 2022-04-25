package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) PaymentHandler {
	return PaymentHandler{
		paymentService: paymentService,
	}
}

func (h PaymentHandler) CreatePayment() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req dto.CreatePaymentRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		user := c.MustGet("user").(models.User)
		req.UserID = user.ID

		res, err := h.paymentService.CreatePayment(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusCreated, res)
		}

	}
}

func (h PaymentHandler) MyPayments() gin.HandlerFunc {

	return func(c *gin.Context) {
		var req dto.MyPaymentsRequest
		user := c.MustGet("user").(models.User)
		req.UserID = user.ID
		res, err := h.paymentService.MyPayments(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}

	}
}
