package main

import (
	"fmt"

	"github.com/UsefulForMe/go-ecommerce/app"
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/logger"
)

func main() {
	config.InitConfig()
	config.InitFirebase()
	r := app.SetupApp()

	port := config.Cfg.Port
	logger.Info("App is running on port " + port)
	r.Run(fmt.Sprintf(":%s", port))
}
