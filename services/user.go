package services

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Create(user dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError)
	List() (*dto.GetAllUserResponse, *errs.AppError)

	Login(user dto.LoginUserRequest) (*dto.LoginUserResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo repository.UserRepository
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

func (s DefaultUserService) Login(r dto.LoginUserRequest) (*dto.LoginUserResponse, *errs.AppError) {

	user, _ := s.repo.FindByPhoneNumber(r.PhoneNumber)

	if user == nil {
		newUser := models.User{
			PhoneNumber: r.PhoneNumber,
			FullName:    r.PhoneNumber,
		}

		id, err := s.repo.Save(&newUser)
		if err != nil {
			return nil, err
		}
		newUser.ID = *id
		user = &newUser
	}

	jwtToken, token, err := config.NewJWTToken(user.ID.String(), map[string]string{})
	if err != nil {
		return nil, err
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	expiredAt := claims["exp"].(int64)

	res := dto.LoginUserResponse{
		User:     *user,
		Token:    *token,
		ExpireAt: expiredAt,
	}

	return &res, nil
}

func NewUserService(repo repository.UserRepository) DefaultUserService {
	return DefaultUserService{
		repo: repo,
	}
}
