package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Type       ErrorType
	Message    string
	Err        error
	StatusCode int
}

type ErrorType string

const (
	NotFound     ErrorType = "NOT_FOUND"
	Validation   ErrorType = "VALIDATION"
	Internal     ErrorType = "INTERNAL"
	InvalidInput ErrorType = "INVALID_INPUT"
	Unauthorized ErrorType = "UNAUTHORIZED"
)

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:       NotFound,
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

func NewValidationError(err error) *AppError {
	return &AppError{
		Type:       Validation,
		Message:    "Erro de validação",
		Err:        err,
		StatusCode: http.StatusBadRequest,
	}
}

func NewInternalError(err error) *AppError {
	return &AppError{
		Type:       Internal,
		Message:    "Erro interno do servidor",
		Err:        err,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewInvalidInputError(message string) *AppError {
	return &AppError{
		Type:       InvalidInput,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Type:       Unauthorized,
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}
