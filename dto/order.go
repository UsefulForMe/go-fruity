package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type CreateOrderRequest struct {
	UserID     uuid.UUID          `json:"user_id"`
	SellerID   uuid.UUID          `json:"seller_id"`
	PaymentID  uuid.UUID          `json:"payment_id"`
	OrderItems []models.OrderItem `json:"order_items"`
}

type CreateOrderResponse struct {
	Order models.Order `json:"order"`
}

type MyOrdersRequest struct {
	UserID uuid.UUID `json:"user_id"`
}
type MyOrdersResponse struct {
	Orders []models.Order `json:"orders"`
}
