package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) CategoryHandler {
	return CategoryHandler{
		categoryService: categoryService,
	}
}

func (h CategoryHandler) GetAllCategories() gin.HandlerFunc {
	return func(c *gin.Context) {

		var res *dto.ListCategoryResponse
		var err *errs.AppError
		parentId := c.Query("parent_id")

		if parentId != "" {
			if id, _err := uuid.Parse(parentId); _err != nil {
				WriteResponseError(c, errs.NewBadRequestError(_err.Error()))
				return
			} else {
				res, err = h.categoryService.GetAllChildCategories(dto.ListCategoryRequest{
					ParentID: &id,
				})
			}
		} else {
			res, err = h.categoryService.GetAllParentCategories()
		}

		if err != nil {
			WriteResponseError(c, err)
		} else {

			WriteResponse(c, http.StatusOK, res)
		}

	}
}

func (h CategoryHandler) CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateCategoryRequest
		if err := c.BindJSON(&req); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		res, err := h.categoryService.CreateCategory(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusCreated, res)
		}

	}
}

func (h CategoryHandler) GetProductsByCategoryID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := uuid.Parse(c.Param("id")); err != nil {
			WriteResponseError(c, errs.NewBadRequestError(err.Error()))
			return
		}

		req := dto.GetProductsByCategoryRequest{
			CategoryID: uuid.Must(uuid.Parse(c.Param("id"))),
		}

		res, err := h.categoryService.GetProductsByCategory(req)
		if err != nil {
			WriteResponseError(c, err)
		} else {
			WriteResponse(c, http.StatusOK, res)
		}

	}
}
