package handlers

import (
	"errors"

	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/gin-gonic/gin"
)

func WriteResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func WriteResponseError(c *gin.Context, err *errs.AppError) {
	c.Error(errors.New(err.Message))
	c.JSON(err.Code, err.Error())
}
