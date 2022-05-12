package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type GetAllStockRequest struct{}

type GetAllStockResponse struct {
	Stocks []models.Stock `json:"stocks"`
}

type SaveStockRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type SaveStockResponse struct {
	Stock models.Stock `json:"stock"`
}

type UpdateStockRequest struct {
	StockID  uuid.UUID `json:"stock_id"`
	Quantity int       `json:"quantity"`
}

type UpdateStockResponse struct {
	Stock models.Stock `json:"stock"`
}

type FindStockByIdRequest struct {
	StockID uuid.UUID `json:"stock_id"`
}

type FindStockByIdResponse struct {
	Stock models.Stock `json:"stock"`
}

type FindStockByProductIdRequest struct {
	ProductID uuid.UUID `json:"product_id"`
}

type FindStockByProductIdResponse struct {
	Stock models.Stock `json:"stock"`
}
