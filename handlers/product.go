package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) ProductHandler {
	return ProductHandler{
		productService: productService,
	}
}

func (h ProductHandler) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.CreateProductRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}
		fmt.Print(request)
		product, err := h.productService.CreateProduct(&request)
		if err != nil {
			WriteResponseError(c, err)
			return
		}
		WriteResponse(c, http.StatusCreated, product)
	}
}

func (h ProductHandler) GetProductAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.GetProductsRequest
		products, err := h.productService.GetProducts(request)
		if err != nil {
			WriteResponseError(c, err)
			return
		}
		WriteResponse(c, http.StatusOK, products)
	}
}

func (h ProductHandler) GetProductByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.GetProductRequest
		productId := c.Param("id")
		request.ID = uuid.Must(uuid.Parse(productId))
		product, err := h.productService.GetProduct(request)
		if err != nil {
			WriteResponseError(c, err)
			return
		}
		WriteResponse(c, http.StatusOK, product)
	}
}
func (h ProductHandler) GetTopSaleProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.GetTopSaleProductsRequest
		limit, errLim := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if errLim != nil {
			WriteResponseError(c, errs.NewBadRequestError(errLim.Error()))
			return
		}
		request.Limit = limit
		products, err := h.productService.GetTopSaleProducts(request)
		if err != nil {
			WriteResponseError(c, err)
			return
		}
		WriteResponse(c, http.StatusOK, products)
	}
}
func (h ProductHandler) GetProductsSaleOff() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.GetProductsSaleOffRequest
		limit, errLim := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if errLim != nil {
			WriteResponseError(c, errs.NewBadRequestError(errLim.Error()))
			return
		}
		req.Limit = limit
		products, err := h.productService.GetSaleOffProducts(req)
		if err != nil {
			WriteResponseError(c, err)
			return
		}
		WriteResponse(c, http.StatusOK, products)
	}
}
