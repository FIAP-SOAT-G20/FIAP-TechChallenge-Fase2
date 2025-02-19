package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
)

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

func ErrorHandler(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Execute all the handlers

		// If there are errors, handle the last one
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			handleError(c, err, logger)
		}
	}
}

func handleError(c *gin.Context, err error, logger *slog.Logger) {
	switch e := err.(type) {
	case *domain.ValidationError:
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

	case *domain.NotFoundError:
		c.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: e.Error(),
		})

	case *domain.InvalidInputError:
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

	case *domain.InternalError:
		logger.Error(domain.ErrInternalError,
			"error", e.Error(),
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: domain.ErrInternalError,
		})

	default:
		logger.Error(domain.ErrUnknownError,
			"error", err.Error(),
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: domain.ErrInternalError,
		})
	}
}
