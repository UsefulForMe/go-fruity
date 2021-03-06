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

		req.UserID = uuid.MustParse("0c9a05cb-407c-43c1-b855-88a034ad8f01")

		res, err := h.orderService.ChangeOrderStatus(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {

			// title := getTitle(res.Order.Status)
			// body := getBody(res.Order)
			// go h.firebaseFCM.SendNotification(dto.SendNotificationRequest{
			// 	Title: title,
			// 	Body:  body,
			// 	Token: user.FCMToken,
			// 	Data: map[string]string{
			// 		"id":     res.Order.ID.String(),
			// 		"action": "/view_order",
			// 	},
			// })
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func getTitle(status string) string {
	switch status {
	case dto.OrderStatusShipping:
		return "???? B???n c?? ????n h??ng ??ang ch??? tr??n ???????ng giao"
	case dto.OrderStatusDelivered:
		return "???? ????n h??ng ???? ???????c giao"
	case dto.OrderStatusCancelled:
		return "???? ????n h??ng ???? b??? h???y"
	default:
		return "???? T???o ????n h??ng th??nh c??ng"
	}
}

func getBody(order models.Order) string {
	switch order.Status {
	case dto.OrderStatusShipping:
		return fmt.Sprintf("????n h??ng %s ??ang trong qu?? tr??nh v???n chuy???n v?? d??? ki???n ???????c giao v??o %s  ", order.TextID, order.ReceivedAt.Format("02/01/2006 15:04:05"))
	case dto.OrderStatusDelivered:
		return fmt.Sprintf("????n h??ng %s ???? ???????c giao v??o %s ", order.TextID, time.Now().Format("02/01/2006 15:04:05"))
	case dto.OrderStatusCancelled:
		return fmt.Sprintf("????n h??ng %s ???? b??? h???y ", order.TextID)
	default:
		return fmt.Sprintf("????n h??ng %s c???a b???n ???? ???????c t???o th??nh c??ng !", order.TextID)
	}
}
