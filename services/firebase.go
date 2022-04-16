package services

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"google.golang.org/api/option"
)

//go:generate mockgen -destination=../mocks/services/mock_firebase_service.go -package=services  github.com/UsefulForMe/go-ecommerce/services FirebaseService

type FirebaseService interface {
	VerifyIDToken(idToken string) (*auth.Token, *errs.AppError)
}
type DefaultFirebaseService struct {
	firebase *firebase.App
}

func (s DefaultFirebaseService) VerifyIDToken(idToken string) (*auth.Token, *errs.AppError) {
	client, err := s.firebase.Auth(context.Background())
	if err != nil {
		return nil, errs.NewUnexpectedError("Error when getting firebase auth client: " + err.Error())
	}
	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, errs.NewUnauthenticatedError("Error when verifying id token: " + err.Error())
	}
	return token, nil
}

func NewFirebaseService() DefaultFirebaseService {
	wd, err := os.Getwd()
	if err != nil {
		logger.Error("Error when getting working directory: " + err.Error())
		panic(err)
	}

	opt := option.WithCredentialsFile(wd + "/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Error("Error initializing firebase app: " + err.Error())
		panic(err)
	}
	return DefaultFirebaseService{
		firebase: app,
	}
}
