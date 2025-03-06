package handler_test

import (
	"context"
	"io"
	"log/slog"
	"testing"

	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/logger"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/middleware"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type OrderHandlerSuiteTest struct {
	suite.Suite
	handler        *handler.OrderHandler
	router         *gin.Engine
	mockController *mockport.MockOrderController
	ctx            context.Context
	requests       map[string]string
	responses      map[string]string
}

func (s *OrderHandlerSuiteTest) SetupTest() {
	// Create a new router
	s.router = gin.New()
	gin.SetMode(gin.TestMode)
	s.router.Use(middleware.ErrorHandler(&logger.Logger{Logger: slog.New(slog.NewJSONHandler(io.Discard, nil))})) // Remove log output
	// s.router.Use(middleware.ErrorHandler(&Logger{Logger: slog.New(slog.DiscardHandler)})) // TODO: Replace above line with this line, when updating go to 1.24 or higher

	// Create a new handler
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockController = mockport.NewMockOrderController(ctrl)
	s.handler = handler.NewOrderHandler(s.mockController)
	s.ctx = context.Background()

	// Register routes
	s.router.GET("/orders", s.handler.List)
	s.router.POST("/orders", s.handler.Create)
	s.router.PUT("/orders/:id", s.handler.Update)
	s.router.GET("/orders/:id", s.handler.Get)
	s.router.DELETE("/orders/:id", s.handler.Delete)

	// Mock requests
	var err error
	s.requests, err = util.ReadFixtureFiles("order",
		"create_success", "create_invalid_parameter",
		"update_success",
	)
	assert.NoError(s.T(), err)

	// Mock responses
	s.responses, err = util.ReadGoldenFiles("order",
		"error_invalid_parameter", "error_internal_error", "error_not_found",
		"list_success", "list_success_with_query",
		"create_success",
		"update_success",
		"get_success",
		"delete_success",
	)
	assert.NoError(s.T(), err)
}

// func (s *OrderHandlerSuiteTest) BeforeTest(_, _ string) {}

func TestOrderHandlerSuiteTest(t *testing.T) {
	suite.Run(t, new(OrderHandlerSuiteTest))
}
