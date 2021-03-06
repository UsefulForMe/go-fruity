package services

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/golang-jwt/jwt"
)

//go:generate mockgen -destination=../mocks/services/mock_user_service.go -package=services  github.com/UsefulForMe/go-ecommerce/services UserService
type UserService interface {
	Create(user dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError)
	List() (*dto.GetAllUserResponse, *errs.AppError)

	Login(user dto.LoginUserRequest) (*dto.LoginUserResponse, *errs.AppError)

	UpdateFCMToken(user dto.UpdateFCMTokenRequest) (*dto.UpdateFCMTokenResponse, *errs.AppError)
	UpdateInfor(r dto.UpdateUserInforRequest) (*dto.UpdateUserInforResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) DefaultUserService {
	return DefaultUserService{
		repo: repo,
	}
}

func (s DefaultUserService) Create(r dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError) {

	user := models.User{
		PhoneNumber: r.PhoneNumber,
		FullName:    r.PhoneNumber,
		Email:       "",
	}

	newUser, err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}
	res := dto.CreateUserResponse{
		User: *newUser,
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
		_user := models.User{
			PhoneNumber: r.PhoneNumber,
			FullName:    r.PhoneNumber,
		}

		newUser, err := s.repo.Save(_user)
		if err != nil {
			return nil, err
		}
		user = newUser
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

func (s DefaultUserService) UpdateFCMToken(r dto.UpdateFCMTokenRequest) (*dto.UpdateFCMTokenResponse, *errs.AppError) {

	user, err := s.repo.FindById(r.UserID)
	if err != nil {
		return nil, err
	}

	user.FCMToken = r.Token
	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}

	res := dto.UpdateFCMTokenResponse{
		Success: true,
	}
	return &res, nil
}

func (s DefaultUserService) UpdateInfor(r dto.UpdateUserInforRequest) (*dto.UpdateUserInforResponse, *errs.AppError) {

	user, err := s.repo.FindById(r.UserID)
	if err != nil {
		return nil, err
	}

	user.FullName = r.FullName
	user.Avatar = r.Avatar
	user.Email = r.Email

	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}

	res := dto.UpdateUserInforResponse{
		User: *user,
	}
	return &res, nil
}
