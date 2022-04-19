package config

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func newFirebaseApp() *firebase.App {
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
	return app

}

func InitFirebase() {
	FirebaseApp = newFirebaseApp()
}
