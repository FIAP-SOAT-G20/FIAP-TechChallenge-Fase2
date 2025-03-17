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

type PaymentUsecaseSuiteTest struct {
	suite.Suite
	mockPayments     []*entity.Payment
	mockGateway      *mockport.MockPaymentGateway
	mockOrderUseCase *mockport.MockOrderUseCase
	useCase          port.PaymentUseCase
	ctx              context.Context
}

func (s *PaymentUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockGateway = mockport.NewMockPaymentGateway(ctrl)
	s.mockOrderUseCase = mockport.NewMockOrderUseCase(ctrl)
	s.useCase = usecase.NewPaymentUseCase(s.mockGateway, s.mockOrderUseCase)
	s.ctx = context.Background()
	currentTime := time.Now()
	s.mockPayments = []*entity.Payment{
		{
			ID:                1,
			Status:            valueobject.PROCESSING,
			ExternalPaymentID: "123456789",
			QrData:            "123456789",
			OrderID:           1,
			CreatedAt:         currentTime,
			UpdatedAt:         currentTime,
		},
		{
			ID:                2,
			Status:            valueobject.CONFIRMED,
			ExternalPaymentID: "123456789",
			QrData:            "123456789",
			OrderID:           2,
			CreatedAt:         currentTime,
			UpdatedAt:         currentTime,
		},
	}
}

func TestPaymentUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(PaymentUsecaseSuiteTest))
}
