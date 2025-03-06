package controller_test

import (
	"context"
	"testing"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type CustomerControllerSuiteTest struct {
	suite.Suite
	mockUseCase   *mockport.MockCustomerUseCase
	mockPresenter port.Presenter
	controller    port.CustomerController
	ctx           context.Context
}

func (s *CustomerControllerSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockUseCase = mockport.NewMockCustomerUseCase(ctrl)
	s.mockPresenter = presenter.NewCustomerJsonPresenter()
	s.controller = controller.NewCustomerController(s.mockUseCase)
	s.ctx = context.Background()
}

func TestCustomerControllerSuiteTest(t *testing.T) {
	suite.Run(t, new(CustomerControllerSuiteTest))
}
