package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SellerHandler struct {
	sellerService services.SellerService
}

func NewSellerHandler(sellerService services.SellerService) SellerHandler {
	return SellerHandler{
		sellerService: sellerService,
	}
}

func (h SellerHandler) CreateSeller() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateSellerRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		res, err := h.sellerService.CreateSeller(req)

		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func (h SellerHandler) GetAllSellers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.GetAllSellerRequest
		ids := c.QueryArray("ids")
		req.IDs = ids

		res, err := h.sellerService.GetAllSeller(req)

		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func (h SellerHandler) GetSellerByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.GetSellerByIDRequest
		id := c.Param("id")
		req.ID = uuid.MustParse(id)

		res, err := h.sellerService.GetSellerByID(req)

		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}

func (h SellerHandler) GetProductsBySellerID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.GetProductsBySellerIDRequest
		id := c.Param("id")
		req.SellerID = uuid.MustParse(id)

		res, err := h.sellerService.GetProductsBySellerID(req)

		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}
	}
}
