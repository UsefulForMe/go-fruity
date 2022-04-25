package dto

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
)

type LoginUserRequest struct {
	PhoneNumber string `json:"phone_number"  binding:"required"`
	IdToken     string `json:"id_token" binding:"required"`
}

func (r LoginUserRequest) Validate() *errs.AppError {
	if r.PhoneNumber == "" {
		return errs.NewBadRequestError("phone_number is required")
	}
	if r.IdToken == "" {
		return errs.NewBadRequestError("id_token is required")
	}
	return nil
}

type LoginUserResponse struct {
	User     models.User `json:"user" binding:"required"`
	Token    string      `json:"token" binding:"required"`
	ExpireAt int64       `json:"expire_at" binding:"required"`
}
