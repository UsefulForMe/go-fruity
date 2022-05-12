package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StockReportHandler struct {
	stockReportService services.StockReportService
}

func NewStockReportHandler(stockReportService services.StockReportService) StockReportHandler {
	return StockReportHandler{
		stockReportService: stockReportService,
	}
}

func (h StockReportHandler) GetAllStockReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.GetAllStockReportRequest{}
		res, err := h.stockReportService.GetAllStockReport(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func (h StockReportHandler) SaveStockReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.SaveStockReportRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}
		res, err := h.stockReportService.SaveStockReport(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusCreated, res)
		}
	}
}

func (h StockReportHandler) UpdateStockReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.UpdateStockReportRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		stockReportId := c.Param("stock_report_id")
		req.StockReportID = uuid.MustParse(stockReportId)

		res, err := h.stockReportService.UpdateStockReport(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

// func (h StockHandler) GetStockById() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		req := dto.FindStockByIdRequest{}
// 		stockId := c.Param("stock_id")
// 		req.StockID = uuid.MustParse(stockId)
// 		res, err := h.stockService.GetStockById(req)
// 		if err != nil {
// 			WriteResponseError(c, err)
// 		} else {
// 			WriteResponse(c, http.StatusOK, res)
// 		}

// 	}
// }
