package middleware

import (
	"strings"

	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/handlers"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func VerifyJWT(r repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authentiationHeader := c.Request.Header.Get("Authorization")

		if authentiationHeader == "" {
			handlers.WriteResponseError(c, errs.NewUnauthenticatedError("Unauthorized"))
			return
		}
		arr := strings.Split(authentiationHeader, " ")
		if len(arr) <= 1 {
			handlers.WriteResponseError(c, errs.NewUnauthenticatedError("Invalid token"))
			return
		}
		token := arr[1]

		claims, err := config.VerifyJWTToken(token)
		if err != nil {
			handlers.WriteResponseError(c, err)
			return
		}

		userId := claims["iss"].(string)

		userUUID := uuid.Must(uuid.Parse(userId))

		user, err := r.FindById(userUUID)
		if err != nil {
			handlers.WriteResponseError(c, err)
			return
		}
		c.Set("user", *user)
		c.Next()
	}
}
