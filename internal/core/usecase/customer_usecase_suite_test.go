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

type CustomerUsecaseSuiteTest struct {
	suite.Suite
	// conn    *sql.DB
	// mock    sqlmock.Sqlmock
	// handler handler
	mockGateway *mockport.MockCustomerGateway
	useCase     port.CustomerUseCase
	ctx         context.Context
}

func (s *CustomerUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockGateway = mockport.NewMockCustomerGateway(ctrl)
	s.useCase = usecase.NewCustomerUseCase(s.mockGateway)
	s.ctx = context.Background()
}

// func (ts *CustomerUsecaseSuiteTest) AfterTest(_, _ string) {
// 	// assert.NoError(ts.T(), ts.mock.ExpectationsWereMet())
// }

func TestCustomerUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(CustomerUsecaseSuiteTest))
}
