package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type OrderUsecaseSuiteTest struct {
	suite.Suite
	mockOrders         []*entity.Order
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
	currentTime := time.Now()
	s.mockOrders = []*entity.Order{
		{
			ID:         1,
			CustomerID: uint64(1),
			TotalBill:  99.99,
			Status:     valueobject.PENDING,
			CreatedAt:  currentTime,
			UpdatedAt:  currentTime,
		},
		{
			ID:         2,
			CustomerID: uint64(2),
			TotalBill:  199.99,
			Status:     valueobject.PENDING,
			CreatedAt:  currentTime,
			UpdatedAt:  currentTime,
		},
	}
}

func TestOrderUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(OrderUsecaseSuiteTest))
}
