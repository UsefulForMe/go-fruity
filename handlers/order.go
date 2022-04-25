package handlers

import (
	"fmt"
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderHandler struct {
	orderService services.OrderService
	firebaseFCM  services.FirebaseMessageService
}

func NewOrderHandler(orderService services.OrderService, firebaseFCM services.FirebaseMessageService) OrderHandler {
	return OrderHandler{
		orderService: orderService,
		firebaseFCM:  firebaseFCM,
	}
}

func (h OrderHandler) CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateOrderRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}
		user := c.MustGet("user").(models.User)
		req.UserID = user.ID

		order, err := h.orderService.CreateOrder(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			title := "🛒 Tạo đơn hàng thành công"
			body := fmt.Sprintf("Đơn hàng %s của bạn đã được tạo thành công !", order.Order.TextID)
			go h.firebaseFCM.SendNotification(dto.SendNotificationRequest{
				Title: title,
				Body:  body,
				Token: user.FCMToken,
			})

			WriteResponse(c, http.StatusCreated, order)
		}

	}
}

func (h OrderHandler) MyOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)
		req := dto.MyOrdersRequest{
			UserID: user.ID,
			Status: c.DefaultQuery("status", "processing"),
		}

		res, err := h.orderService.MyOrders(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func (h OrderHandler) GetOrderByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.GetOrderByIDRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}
		user := c.MustGet("user").(models.User)

		if req.UserID != user.ID {
			WriteResponseError(c, errs.NewForbiddenError("UserID không hợp lệ"))
			return
		}

		orderId := uuid.MustParse(c.Param("order_id"))
		req.OrderID = orderId

		res, err := h.orderService.GetOrderByID(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}
