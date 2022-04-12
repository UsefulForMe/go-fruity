package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type CreatePaymentRequest struct {
	Name      string    `json:"name"`
	Provider  string    `json:"provider"`
	AccountNo string    `json:"account_no"`
	UserID    uuid.UUID `json:"user_id"`
	Logo      string    `json:"logo"`
}

type CreatePaymentResponse struct {
	Payment models.Payment `json:"payment"`
}

type MyPaymentsRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type MyPaymentsResponse struct {
	Payments []models.Payment `json:"payments"`
}
