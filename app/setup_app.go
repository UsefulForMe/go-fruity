package app

import (
	"os"
	"reflect"

	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/dto"
	middleware "github.com/UsefulForMe/go-ecommerce/middlewares"
	router "github.com/UsefulForMe/go-ecommerce/routers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

func ValidateJSONDateType(field reflect.Value) interface{} {
	if field.Type() == reflect.TypeOf(dto.LocalTime{}) {
		timeStr := field.Interface().(dto.LocalTime).String()
		if timeStr == "0001-01-01 00:00:00" {
			return nil
		}
		return timeStr
	}
	return nil
}

func SetupApp() *gin.Engine {

	config.InitDatabase()

	os.Setenv("TZ", config.Cfg.Tz)

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterCustomTypeFunc(ValidateJSONDateType, dto.LocalTime{})
	}

	app.Use(middleware.Cors())
	// app.Use(middleware.Cors(), middleware.RequestLogger(), gin.Recovery())
	router.SetupRoute(app)

	return app
}
