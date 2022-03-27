package handlers

import (
	"net/http"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	s3 services.S3Service
}

func (h *UploadHandler) Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.UploadFileRequest
		err := c.ShouldBind(&form)
		if err != nil {
			WriteResponseError(c, errs.NewBadRequestError("Error when bind file "+err.Error()))
			return
		}
		user := c.MustGet("user").(models.User)

		result, appError := h.s3.UploadFile(form, user.ID)

		if appError != nil {
			WriteResponseError(c, appError)
		} else {
			WriteResponse(c, http.StatusCreated, result)
		}

	}
}

func (h *UploadHandler) UploadMany() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.UploadFilesRequest
		err := c.ShouldBind(&form)
		if err != nil {
			WriteResponseError(c, errs.NewBadRequestError("Error when bind file "+err.Error()))
			return
		}
		user := c.MustGet("user").(models.User)
		result, appError := h.s3.UploadFiles(form, user.ID)
		if appError != nil {
			WriteResponseError(c, appError)
		} else {
			WriteResponse(c, http.StatusCreated, result)
		}

	}
}

func (h *UploadHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.DeleteFileRequest
		err := c.ShouldBind(&req)
		if err != nil {
			WriteResponseError(c, errs.NewBadRequestError("Error when bind body "+err.Error()))
			return
		}
		res, appErr := h.s3.DeleteFile(req)
		if appErr != nil {
			WriteResponseError(c, appErr)
			return
		}
		WriteResponse(c, http.StatusOK, res)
	}
}

func (h *UploadHandler) DeleteMany() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.DeleteFilesRequest
		err := c.ShouldBind(&req)
		if err != nil {
			WriteResponseError(c, errs.NewBadRequestError("Error when bind body "+err.Error()))
			return
		}
		res, appErr := h.s3.DeleteFiles(req)
		if appErr != nil {
			WriteResponseError(c, appErr)
			return
		}
		WriteResponse(c, http.StatusOK, res)
	}
}

func NewUploadHandler(s3 services.S3Service) UploadHandler {
	return UploadHandler{
		s3: s3,
	}
}
