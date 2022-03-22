package middleware

import (
	"time"

	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		path := c.Request.URL.Path

		query := c.Request.URL.RawQuery

		c.Next()

		cost := time.Since(start)

		logger.Info(path,

			zap.Int("status", c.Writer.Status()),

			zap.String("method", c.Request.Method),

			zap.String("path", path),

			zap.String("query", query),

			zap.String("ip", c.ClientIP()),

			zap.String("user-agent", c.Request.UserAgent()),

			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),

			zap.Duration("latency", cost),
		)
	}
}
