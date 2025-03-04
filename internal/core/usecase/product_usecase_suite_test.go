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

type ProductUsecaseSuiteTest struct {
	suite.Suite
	mockGateway *mockport.MockProductGateway
	useCase     port.ProductUseCase
	ctx         context.Context
}

func (s *ProductUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockGateway = mockport.NewMockProductGateway(ctrl)
	s.useCase = usecase.NewProductUseCase(s.mockGateway)
	s.ctx = context.Background()
}

func TestProductUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(ProductUsecaseSuiteTest))
}
