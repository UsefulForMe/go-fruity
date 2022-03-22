package services

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
)

type UserService interface {
	Create(user *models.User) *errs.AppError
	GetByID(id uint) (*models.User, *errs.AppError)
	GetAll() ([]models.User, *errs.AppError)
	Update(user *models.User) *errs.AppError
	Delete(id uint) *errs.AppError
}

type DefaultUserService struct {
	UserRepository models.UserRepository
}


