package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"tech-challenge-2-app-example/internal/core/domain/errors"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func ErrorHandler(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Execute os handlers

		// Verifica se há erros
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			handleError(c, err, logger)
		}
	}
}

func handleError(c *gin.Context, err error, logger *slog.Logger) {
	switch e := err.(type) {
	case *errors.AppError:
		response := ErrorResponse{
			Code:    e.StatusCode,
			Type:    string(e.Type),
			Message: e.Message,
		}
		if e.Err != nil {
			response.Details = e.Err.Error()
		}
		c.JSON(e.StatusCode, response)

	case validator.ValidationErrors:
		validationErrors := make([]map[string]string, 0)
		for _, err := range e {
			validationErrors = append(validationErrors, map[string]string{
				"field": err.Field(),
				"tag":   err.Tag(),
				"value": err.Param(),
			})
		}
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Type:    string(errors.Validation),
			Message: "Erro de validação",
			Details: validationErrors,
		})

	default:
		logger.Error("internal server error", "error", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Type:    string(errors.Internal),
			Message: "Erro interno do servidor",
		})
	}
}
