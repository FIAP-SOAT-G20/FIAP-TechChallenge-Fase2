package usecase_test

import (
	"context"
	"testing"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type OrderUsecaseSuiteTest struct {
	suite.Suite
	historyUseCaseMock *mockport.MockOrderHistoryUseCase
	mockGateway        *mockport.MockOrderGateway
	useCase            port.OrderUseCase
	ctx                context.Context
}

func (s *OrderUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.historyUseCaseMock = mockport.NewMockOrderHistoryUseCase(ctrl)
	s.mockGateway = mockport.NewMockOrderGateway(ctrl)
	s.useCase = usecase.NewOrderUseCase(s.mockGateway, s.historyUseCaseMock)
	s.ctx = context.Background()
}

func TestOrderUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(OrderUsecaseSuiteTest))
}
