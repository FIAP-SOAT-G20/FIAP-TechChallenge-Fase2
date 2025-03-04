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

type StaffUsecaseSuiteTest struct {
	suite.Suite
	mockGateway *mockport.MockStaffGateway
	useCase     port.StaffUseCase
	ctx         context.Context
}

func (s *StaffUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockGateway = mockport.NewMockStaffGateway(ctrl)
	s.useCase = usecase.NewStaffUseCase(s.mockGateway)
	s.ctx = context.Background()
}

func TestStaffUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(StaffUsecaseSuiteTest))
}
