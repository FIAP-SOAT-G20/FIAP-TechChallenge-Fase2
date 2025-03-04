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

type OrderHistoryUsecaseSuiteTest struct {
	suite.Suite
	mockGateway *mockport.MockOrderHistoryGateway
	useCase     port.OrderHistoryUseCase
	ctx         context.Context
}

func (s *OrderHistoryUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockGateway = mockport.NewMockOrderHistoryGateway(ctrl)
	s.useCase = usecase.NewOrderHistoryUseCase(s.mockGateway)
	s.ctx = context.Background()
}

func TestOrderHistoryUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(OrderHistoryUsecaseSuiteTest))
}
