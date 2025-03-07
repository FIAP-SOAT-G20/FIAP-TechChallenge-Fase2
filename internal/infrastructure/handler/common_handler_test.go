package handler_test

import (
	"io"
	"log/slog"
	"maps"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/logger"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/middleware"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/server"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/util"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.Use(middleware.ErrorHandler(&logger.Logger{Logger: slog.New(slog.NewJSONHandler(io.Discard, nil))})) // Remove log output
	// router.Use(middleware.ErrorHandler(&Logger{Logger: slog.New(slog.DiscardHandler)})) // TODO: Replace above line with this line, when updating go to 1.24 or higher
	server.RegisterCustomValidation()
	return router
}

func addCommonResponses(r *map[string]string) {
	commonResponses, err := util.ReadGoldenFiles("common",
		"error_invalid_parameter", "error_internal_error", "error_not_found",
	)
	if err != nil {
		panic(err)
	}
	maps.Copy(*r, commonResponses)
}
