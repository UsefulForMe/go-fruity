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
	FindUserAddressByID(userAddressID uuid.UUID) (*models.UserAddress, *errs.AppError)
	Update(userAddress models.UserAddress) (*models.UserAddress, *errs.AppError)
	Delete(userAddressID uuid.UUID) *errs.AppError
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

	tx := r.db.Begin()

	if userAddress.IsDefault {
		if err := r.db.Model(&models.UserAddress{}).Where("user_id = ?", userAddress.UserID).Update("is_default", false).Error; err != nil {
			logger.Error("Error while updating user addresses " + err.Error())
			tx.Rollback()
			return nil, errs.NewUnexpectedError("Error while updating user addresses " + err.Error())
		}
	}
	if err := r.db.Create(&userAddress).Error; err != nil {
		logger.Error("Error while saving user address " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Error while saving user address " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		logger.Error("Error while saving user address " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Error while saving user address " + err.Error())
	}

	return &userAddress, nil
}

func (r DefaultUserAddressRepository) FindUserAddressesByUserID(userID uuid.UUID) ([]models.UserAddress, *errs.AppError) {
	var userAddresses []models.UserAddress

	if err := r.db.Where("user_id = ? and status = ?", userID, "active").Find(&userAddresses).Error; err != nil {
		logger.Error("Error while finding user addresses " + err.Error())
		return nil, errs.NewUnexpectedError("Error while finding user addresses " + err.Error())
	}
	return userAddresses, nil
}

func (r DefaultUserAddressRepository) FindUserAddressByID(userAddressID uuid.UUID) (*models.UserAddress, *errs.AppError) {
	var userAddress models.UserAddress

	if err := r.db.Where("id = ?", userAddressID).First(&userAddress).Error; err != nil {
		logger.Error("Error while finding user address " + err.Error())
		return nil, errs.NewUnexpectedError("Error while finding user address " + err.Error())
	}
	return &userAddress, nil
}

func (r DefaultUserAddressRepository) Update(userAddress models.UserAddress) (*models.UserAddress, *errs.AppError) {

	tx := r.db.Begin()

	if userAddress.IsDefault {
		if err := r.db.Model(&models.UserAddress{}).Where("user_id = ?", userAddress.UserID).Update("is_default", false).Error; err != nil {
			logger.Error("Error while updating user addresses " + err.Error())
			tx.Rollback()
			return nil, errs.NewUnexpectedError("Error while updating user addresses " + err.Error())
		}
	}
	if err := r.db.Save(&userAddress).Error; err != nil {
		logger.Error("Error while updating user address " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Error while updating user address " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		logger.Error("Error while updating user address " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Error while updating user address " + err.Error())
	}

	return &userAddress, nil
}

func (r DefaultUserAddressRepository) Delete(userAddressID uuid.UUID) *errs.AppError {
	tx := r.db.Begin()

	var userAddress models.UserAddress

	if err := r.db.Where("id = ?", userAddressID).First(&userAddress).Error; err != nil {
		logger.Error("Error while finding user address " + err.Error())
		tx.Rollback()
		return errs.NewUnexpectedError("Error while finding user address " + err.Error())
	}

	if err := r.db.Model(&models.UserAddress{}).Where("id = ?", userAddressID).Update("status", "inactive").Error; err != nil {
		logger.Error("Error while deleting user address " + err.Error())
		tx.Rollback()
		return errs.NewUnexpectedError("Error while deleting user address " + err.Error())
	}

	// update user first address to default if user has no address
	if userAddress.IsDefault {
		if err := r.db.Model(&models.UserAddress{}).Where("user_id = ?", userAddress.UserID).Update("is_default", true).Error; err != nil {
			logger.Error("Error while updating user addresses " + err.Error())
			tx.Rollback()
			return errs.NewUnexpectedError("Error while updating user addresses " + err.Error())
		}
	}

	if err := tx.Commit().Error; err != nil {
		logger.Error("Error while deleting user address " + err.Error())
		tx.Rollback()
		return errs.NewUnexpectedError("Error while deleting user address " + err.Error())
	}

	return nil
}
