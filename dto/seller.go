package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type CreateSellerRequest struct {
	Name          string               `json:"name"`
	Logo          string               `json:"logo"`
	Banner        string               `json:"banner"`
	Type          string               `json:"type"`
	PhoneNumber   string               `json:"phone_number"`
	Description   string               `json:"description"`
	HeadQuarter   string               `json:"head_quarter"`
	Rating        float32              `json:"rating"`
	AvailableTime models.AvailableTime `json:"available_time"`
	Note          string               `json:"note"`
	Email         string               `json:"email"`
	TotalVote     int                  `json:"total_vote"`
}

type CreateSellerResponse struct {
	Seller models.Seller `json:"seller"`
}

type GetAllSellerRequest struct {
	IDs []string `json:"ids"`
}

type GetAllSellerResponse struct {
	Sellers []models.Seller `json:"sellers"`
}

type GetSellerByIDRequest struct {
	ID uuid.UUID `json:"id"`
}

type GetSellerByIDResponse struct {
	Seller models.Seller `json:"seller"`
}

type GetProductsBySellerIDRequest struct {
	SellerID uuid.UUID `json:"seller_id"`
}

type GetProductsBySellerIDResponse struct {
	Products []models.Product `json:"products"`
}
