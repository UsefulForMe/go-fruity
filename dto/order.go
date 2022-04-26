package dto

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type CreateOrderRequest struct {
	UserID        uuid.UUID          `json:"user_id"`
	SellerID      uuid.UUID          `json:"seller_id"`
	PaymentID     uuid.UUID          `json:"payment_id"`
	OrderItems    []models.OrderItem `json:"order_items"`
	ReceivedAt    LocalTime          `json:"received_at" `
	UserAddressID uuid.UUID          `json:"user_address_id"`
	Note          string             `json:"note"`
}

type CreateOrderResponse struct {
	Order models.Order `json:"order"`
}

type MyOrdersRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Status string    `json:"status"`
}
type MyOrdersResponse struct {
	Orders []models.Order `json:"orders"`
}

type GetOrderByIDRequest struct {
	OrderID uuid.UUID `json:"order_id"`
}

type GetOrderByIDResponse struct {
	Order models.Order `json:"order"`
}

type ChangeOrderStatusRequest struct {
	OrderID uuid.UUID `json:"order_id"`
	Status  string    `json:"status"`
	Note    string    `json:"note"`
	UserID  uuid.UUID `json:"user_id"`
}

const (
	OrderStatusProccessing = "processing"
	OrderStatusShipping    = "shipping"
	OrderStatusDelivered   = "delivered"
	OrderStatusCancelled   = "cancelled"
)

func (r ChangeOrderStatusRequest) IsValidStatus() *errs.AppError {
	isValid := r.Status == OrderStatusProccessing ||
		r.Status == OrderStatusShipping ||
		r.Status == OrderStatusDelivered ||
		r.Status == OrderStatusCancelled
	if !isValid {
		return errs.NewBadRequestError("invalid order status")
	}
	return nil

}

type ChangeOrderStatusResponse struct {
	Order models.Order `json:"order"`
}
