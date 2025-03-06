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

type OrderProductUsecaseSuiteTest struct {
	suite.Suite
	mockGateway *mockport.MockOrderProductGateway
	useCase     port.OrderProductUseCase
	ctx         context.Context
}

func (s *OrderProductUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockGateway = mockport.NewMockOrderProductGateway(ctrl)
	s.useCase = usecase.NewOrderProductUseCase(s.mockGateway)
	s.ctx = context.Background()
}

func TestOrderProductUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(OrderProductUsecaseSuiteTest))
}
