package services

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
)

type FirebaseMessageService interface {
	SendNotification(req dto.SendNotificationRequest) (*dto.SendNotificationResponse, *errs.AppError)
}

type DefaultFirebaseMessageService struct {
	client *messaging.Client
}

func NewFirebaseMessageService(firebase *firebase.App) DefaultFirebaseMessageService {
	client, err := firebase.Messaging(context.Background())
	if err != nil {
		logger.Error("Error initializing firebase messaging client: " + err.Error())
		panic(err)
	}

	return DefaultFirebaseMessageService{
		client: client,
	}
}

func (s DefaultFirebaseMessageService) SendNotification(req dto.SendNotificationRequest) (*dto.SendNotificationResponse, *errs.AppError) {

	msg := messaging.Message{
		Notification: &messaging.Notification{
			Title: req.Title,
			Body:  req.Body,
		},
		Token: req.Token,
		Data:  req.Data,
	}

	res, err := s.client.Send(context.Background(), &msg)

	if err != nil {
		logger.Error("Error when sending message: " + err.Error())
		return nil, errs.NewUnexpectedError("Error when sending message: " + err.Error())
	}

	return &dto.SendNotificationResponse{
		MessageID: res,
	}, nil
}
