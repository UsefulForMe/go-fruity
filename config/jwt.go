package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/golang-jwt/jwt"
)

func NewJWTToken(issuer string, data map[string]string) (*jwt.Token, *string, *errs.AppError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  issuer,
		"exp":  time.Now().Add(30 * 24 * time.Hour).Unix(),
		"data": data,
	})

	tokenString, err := token.SignedString(Cfg.HmacSecret)
	if err != nil {
		logger.Error("Error when signed string token " + err.Error())
		return nil, nil, errs.NewUnexpectedError("Unexpected error when signed string token " + err.Error())
	}
	return token, &tokenString, nil
}

func VerifyJWTToken(tokenString string) (jwt.MapClaims, *errs.AppError) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Error(fmt.Sprintf("Unexpected signing method: %v", t.Header["alg"]))
			return nil, errors.New("Unexpected signing method")
		}
		return Cfg.HmacSecret, nil
	})

	if err != nil {
		v, _ := err.(*jwt.ValidationError)

		if v.Errors == jwt.ValidationErrorExpired {
			return nil, errs.NewUnauthenticatedError(v.Error())
		}

		return nil, errs.NewUnexpectedError("Unexpected error when parse token: " + err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		logger.Error("Error when verify token")
		return nil, errs.NewUnauthenticatedError("Invalid token")
	}
	return claims, nil
}
