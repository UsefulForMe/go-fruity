package services

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
)

//go:generate mockgen -destination=../mocks/services/mock_firebase_service.go -package=services  github.com/UsefulForMe/go-ecommerce/services FirebaseService

type FirebaseAuthService interface {
	VerifyIDToken(idToken string) (*auth.Token, *errs.AppError)
}
type DefaultFirebaseAuthService struct {
	client *auth.Client
}

func (s DefaultFirebaseAuthService) VerifyIDToken(idToken string) (*auth.Token, *errs.AppError) {

	token, err := s.client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, errs.NewUnauthenticatedError("Error when verifying id token: " + err.Error())
	}
	return token, nil
}

func NewFirebaseService(app *firebase.App) DefaultFirebaseAuthService {

	authClient, err := app.Auth(context.Background())
	if err != nil {
		logger.Error("Error initializing firebase auth client: " + err.Error())
		panic(err)
	}

	return DefaultFirebaseAuthService{
		client: authClient,
	}
}
