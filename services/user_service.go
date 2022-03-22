package services

import (
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
)

type UserService interface {
	Create(user dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError)
	List() (*dto.GetAllUserResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo models.UserRepository
}

func (s DefaultUserService) Create(r dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError) {
	user := models.User{
		PhoneNumber: r.PhoneNumber,
		FullName:    r.PhoneNumber,
		Email:       "",
	}

	userId, err := s.repo.Save(&user)
	if err != nil {
		return nil, err
	}
	res := dto.CreateUserResponse{
		ID: *userId,
	}
	return &res, nil

}

func (s DefaultUserService) List() (*dto.GetAllUserResponse, *errs.AppError) {

	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	res := dto.GetAllUserResponse{
		Users: users,
	}
	return &res, nil

}

func NewUserService(repo models.UserRepository) DefaultUserService {
	return DefaultUserService{
		repo: repo,
	}
}
