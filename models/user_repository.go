package models

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *User) (*uint, *errs.AppError)
	FindById(id uint) (*User, *errs.AppError)
	FindAll() ([]User, *errs.AppError)
	Update(user *User) *errs.AppError
	Delete(id uint) *errs.AppError
}

type DefaultUserRepository struct {
	db *gorm.DB
}

func (d DefaultUserRepository) Save(user User) (*uint, *errs.AppError) {
	if err := d.db.Create(&user).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &user.ID, nil
}

func (d DefaultUserRepository) FindById(id uint) (*User, *errs.AppError) {
	var user User
	if err := d.db.First(user, id).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &user, nil
}

func (d DefaultUserRepository) FindAll() ([]User, *errs.AppError) {
	var users []User
	if err := d.db.Find(&users).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return users, nil
}

func (d DefaultUserRepository) Update(user *User) *errs.AppError {
	if err := d.db.Save(user).Error; err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

func (d DefaultUserRepository) Delete(id uint) *errs.AppError {
	if err := d.db.Delete(User{}, id).Error; err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}
