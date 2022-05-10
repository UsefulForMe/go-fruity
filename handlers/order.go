package handlers

import (
	"fmt"
	"net/http"
	"time"

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

func (h OrderHandler) GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.GetAllOrdersRequest

		res, err := h.orderService.GetAllOrders(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
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
			title := getTitle(order.Order.Status)
			body := getBody(order.Order)
			go h.firebaseFCM.SendNotification(dto.SendNotificationRequest{
				Title: title,
				Body:  body,
				Token: user.FCMToken,
				Data: map[string]string{
					"id":     order.Order.ID.String(),
					"action": "/view_order",
				},
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
			Status: c.DefaultQuery("status", ""),
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

func (h OrderHandler) ChangeOrderStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.ChangeOrderStatusRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}
		if err := req.IsValidStatus(); err != nil {
			WriteResponseError(c, err)
			return
		}

		orderId := uuid.MustParse(c.Param("order_id"))
		req.OrderID = orderId

		user := c.MustGet("user").(models.User)
		req.UserID = user.ID

		res, err := h.orderService.ChangeOrderStatus(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {

			title := getTitle(res.Order.Status)
			body := getBody(res.Order)
			go h.firebaseFCM.SendNotification(dto.SendNotificationRequest{
				Title: title,
				Body:  body,
				Token: user.FCMToken,
				Data: map[string]string{
					"id":     res.Order.ID.String(),
					"action": "/view_order",
				},
			})
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func getTitle(status string) string {
	switch status {
	case dto.OrderStatusShipping:
		return "🛒 Bạn có đơn hàng đang chờ trên đường giao"
	case dto.OrderStatusDelivered:
		return "🛒 Đơn hàng đã được giao"
	case dto.OrderStatusCancelled:
		return "🛒 Đơn hàng đã bị hủy"
	default:
		return "🛒 Tạo đơn hàng thành công"
	}
}

func getBody(order models.Order) string {
	switch order.Status {
	case dto.OrderStatusShipping:
		return fmt.Sprintf("Đơn hàng %s đang trong quá trình vận chuyển và dự kiến được giao vào %s  ", order.TextID, order.ReceivedAt.Format("02/01/2006 15:04:05"))
	case dto.OrderStatusDelivered:
		return fmt.Sprintf("Đơn hàng %s đã được giao vào %s ", order.TextID, time.Now().Format("02/01/2006 15:04:05"))
	case dto.OrderStatusCancelled:
		return fmt.Sprintf("Đơn hàng %s đã bị hủy ", order.TextID)
	default:
		return fmt.Sprintf("Đơn hàng %s của bạn đã được tạo thành công !", order.TextID)
	}
}
