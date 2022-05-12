package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type GetAllStockReportRequest struct{}

type GetAllStockReportResponse struct {
	StockReports []models.StockReport `json:"stockreports"`
}

type SaveStockReportRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Content   string    `json:"content"`
	Quantity  int       `json:"quantity"`
}

type SaveStockReportResponse struct {
	StockReport models.StockReport `json:"stockreport"`
}

type UpdateStockReportRequest struct {
	StockReportID uuid.UUID `json:"stock_report_id"`
	Quantity      int       `json:"quantity"`
	Content       string    `json:"content"`
}

type UpdateStockReportResponse struct {
	StockReport models.StockReport `json:"stockreport"`
}
