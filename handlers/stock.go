package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StockHandler struct {
	stockService services.StockService
}

func NewStockHandler(stockService services.StockService) StockHandler {
	return StockHandler{
		stockService: stockService,
	}
}

func (h StockHandler) GetAllStock() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.GetAllStockRequest{}
		res, err := h.stockService.GetAllStock(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func (h StockHandler) SaveStock() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.SaveStockRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}
		res, err := h.stockService.SaveStock(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusCreated, res)
		}
	}
}

func (h StockHandler) UpdateStock() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.UpdateStockRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		stockId := c.Param("stock_id")
		req.StockID = uuid.MustParse(stockId)

		res, err := h.stockService.UpdateStock(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func (h StockHandler) GetStockById() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.FindStockByIdRequest{}
		stockId := c.Param("stock_id")
		req.StockID = uuid.MustParse(stockId)
		res, err := h.stockService.GetStockById(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}

	}
}

func (h StockHandler) GetStockByProductId() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.FindStockByProductIdRequest{}
		stockId := c.Param("product_id")
		req.ProductID = uuid.MustParse(stockId)
		res, err := h.stockService.GetStockByProductId(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}

	}
}
