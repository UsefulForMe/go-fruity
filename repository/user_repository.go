package repository

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) (*models.User, *errs.AppError)
	FindById(id uuid.UUID) (*models.User, *errs.AppError)
	FindByPhoneNumber(phoneNumber string) (*models.User, *errs.AppError)
	FindAll() ([]models.User, *errs.AppError)
	Update(user *models.User) *errs.AppError
	Delete(id uuid.UUID) *errs.AppError
}

type DefaultUserRepository struct {
	db *gorm.DB
}

func (d DefaultUserRepository) Save(user models.User) (*models.User, *errs.AppError) {
	if err := d.db.Create(&user).Error; err != nil {
		logger.Error("Error when create user " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when create user " + err.Error())
	}
	return &user, nil
}

func (d DefaultUserRepository) FindById(id uuid.UUID) (*models.User, *errs.AppError) {
	var user models.User

	if err := d.db.First(&user, id).Error; err != nil {
		logger.Error("Error when find user by id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find user by id " + err.Error())
	}
	return &user, nil
}

func (d DefaultUserRepository) FindAll() ([]models.User, *errs.AppError) {
	var users []models.User
	if err := d.db.Model(&users).Preload("Addresses").Find(&users).Error; err != nil {
		logger.Error("Error when find all users " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find all users " + err.Error())
	}
	return users, nil
}

func (d DefaultUserRepository) Update(user *models.User) *errs.AppError {
	if err := d.db.Save(&user).Error; err != nil {
		logger.Error("Error when update user " + err.Error())
		return errs.NewUnexpectedError("Unexpected error when update user " + err.Error())
	}
	return nil
}

func (d DefaultUserRepository) Delete(id uuid.UUID) *errs.AppError {
	if err := d.db.Delete(models.User{}, id).Error; err != nil {
		logger.Error("Error when delete user " + err.Error())
		return errs.NewUnexpectedError("Unexpected error when delete user " + err.Error())
	}
	return nil
}

func (d DefaultUserRepository) FindByPhoneNumber(phoneNumber string) (*models.User, *errs.AppError) {
	var user models.User
	result := d.db.Where("phone_number=?", phoneNumber).First(&user)
	if err := result.Error; err != nil {
		logger.Error("Error when find user by phone number " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find user by phone number " + err.Error())
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) *DefaultUserRepository {
	return &DefaultUserRepository{db}
}
