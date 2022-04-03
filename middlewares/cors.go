package middleware

import (
	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: config.Cfg.Debug == "true",
	})
}
