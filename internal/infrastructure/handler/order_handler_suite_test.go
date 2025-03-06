package handler_test

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type OrderHandlerSuiteTest struct {
	suite.Suite
	handler        *handler.OrderHandler
	router         *gin.Engine
	mockController *mockport.MockOrderController
	ctx            context.Context
}

func (s *OrderHandlerSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockController = mockport.NewMockOrderController(ctrl)
	s.handler = handler.NewOrderHandler(s.mockController)
	s.ctx = context.Background()
}

func (s *OrderHandlerSuiteTest) BeforeTest(_, _ string) {
	s.router = gin.Default()
	s.router.Use(middleware.ErrorHandler(slog.New(slog.NewJSONHandler(&bytes.Buffer{}, nil))))
	s.router.GET("/orders", s.handler.List)
}

func TestOrderHandlerSuiteTest(t *testing.T) {
	suite.Run(t, new(OrderHandlerSuiteTest))
}
