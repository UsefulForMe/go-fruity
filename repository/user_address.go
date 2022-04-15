package repository

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAddressRepository interface {
	Save(userAddress models.UserAddress) (*models.UserAddress, *errs.AppError)

	FindUserAddressesByUserID(userID uuid.UUID) ([]models.UserAddress, *errs.AppError)
}

type DefaultUserAddressRepository struct {
	db *gorm.DB
}

func NewUserAddressRepository(db *gorm.DB) DefaultUserAddressRepository {
	return DefaultUserAddressRepository{
		db: db,
	}
}

func (r DefaultUserAddressRepository) Save(userAddress models.UserAddress) (*models.UserAddress, *errs.AppError) {
	if err := r.db.Create(&userAddress).Error; err != nil {
		logger.Error("Error while saving user address " + err.Error())
		return nil, errs.NewUnexpectedError("Error while saving user address " + err.Error())
	}
	return &userAddress, nil
}

func (r DefaultUserAddressRepository) FindUserAddressesByUserID(userID uuid.UUID) ([]models.UserAddress, *errs.AppError) {
	var userAddresses []models.UserAddress
	if err := r.db.Where("user_id = ?", userID).Find(&userAddresses).Error; err != nil {
		logger.Error("Error while finding user addresses " + err.Error())
		return nil, errs.NewUnexpectedError("Error while finding user addresses " + err.Error())
	}
	return userAddresses, nil
}
