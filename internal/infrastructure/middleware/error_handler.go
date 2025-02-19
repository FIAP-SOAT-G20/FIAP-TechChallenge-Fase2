package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
)

type ErrorJsonResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

type ErrorXmlResponse struct {
	Code    int    `xml:"code" example:"400"`
	Message string `xml:"message" example:"Bad Request"`
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
		setResponse(c, http.StatusBadRequest, e.Error())
		logger.Warn(domain.ErrValidationError,
			"error", e.Error(),
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)

	case *domain.NotFoundError:
		setResponse(c, http.StatusNotFound, e.Error())
		logger.Warn(domain.ErrNotFound,
			"error", e.Error(),
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)

	case *domain.InvalidInputError:
		setResponse(c, http.StatusBadRequest, e.Error())
		logger.Warn(domain.ErrInvalidInput,
			"error", e.Error(),
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)

	case *domain.InternalError:
		setResponse(c, http.StatusInternalServerError, domain.ErrInternalError)
		logger.Error(domain.ErrInternalError,
			"error", e.Error(),
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)

	default:
		setResponse(c, http.StatusInternalServerError, domain.ErrInternalError)
		logger.Error(domain.ErrUnknownError,
			"error", err.Error(),
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)
	}
}

func setResponse(c *gin.Context, status int, message string) {
	if c.GetHeader("Accept") == "text/xml" {
		c.XML(status, ErrorXmlResponse{
			Code:    status,
			Message: message,
		})
		return
	}

	c.JSON(status, ErrorJsonResponse{
		Code:    status,
		Message: message,
	})
}
