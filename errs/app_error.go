package errs

import "net/http"

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewUnauthenticatedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}
