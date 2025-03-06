package controller_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/util"
)

func (s *CustomerControllerSuiteTest) TestCustomerController_ListCustomers() {
	tests := []struct {
		name        string
		input       dto.ListCustomersInput
		setupMocks  func()
		checkResult func(*testing.T, []byte, error)
	}{
		{
			name: "List customers success",
			input: dto.ListCustomersInput{
				Name:  "Test",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockCustomers := []*entity.Customer{
					{
						ID:    1,
						Name:  "Test Customer 1",
						Email: "test.customer.1@email.com",
						CPF:   "12345678901",
					},
					{
						ID:    2,
						Name:  "Test Customer 2",
						Email: "test.customer.2@email.com",
						CPF:   "12345678902",
					},
				}
				s.mockUseCase.EXPECT().
					List(s.ctx, dto.ListCustomersInput{
						Name:  "Test",
						Page:  1,
						Limit: 10,
					}).
					Return(mockCustomers, int64(2), nil)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				want, _ := util.ReadGoldenFile("customer/list_success")
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, want, util.RemoveAllSpaces(string(output)))
			},
		},
		{
			name: "List customers use case error",
			input: dto.ListCustomersInput{
				Name:  "Test",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockUseCase.EXPECT().
					List(s.ctx, dto.ListCustomersInput{
						Name:  "Test",
						Page:  1,
						Limit: 10,
					}).
					Return(nil, int64(0), assert.AnError)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				assert.Error(t, err)
				assert.Nil(t, output)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			output, err := s.controller.List(s.ctx, s.mockPresenter, tt.input)

			// Assert
			tt.checkResult(t, output, err)
		})
	}
}

func TestCustomerController_CreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockport.NewMockCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := controller.NewCustomerController(mockUseCase)

	ctx := context.Background()
	input := dto.CreateCustomerInput{
		Name:  "Test Customer",
		Email: "test.customer.1@email.com",
		CPF:   "123.456.789-00",
	}

	mockCustomer := &entity.Customer{
		ID:    1,
		Name:  "Test Customer",
		Email: "test.customer@email.com",
		CPF:   "123.456.789-00",
	}

	mockUseCase.EXPECT().
		Create(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockCustomer}).
		Return([]byte{}, nil)

	output, err := controller.Create(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestCustomerController_GetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockport.NewMockCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := controller.NewCustomerController(mockUseCase)

	ctx := context.Background()
	input := dto.GetCustomerInput{
		ID: uint64(1),
	}

	mockCustomer := &entity.Customer{
		ID:    1,
		Name:  "Test Customer",
		Email: "test.customer@email.com",
		CPF:   "12345678901",
	}

	mockUseCase.EXPECT().
		Get(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockCustomer}).
		Return([]byte{}, nil)

	output, err := controller.Get(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestCustomerController_UpdateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockport.NewMockCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := controller.NewCustomerController(mockUseCase)

	ctx := context.Background()
	input := dto.UpdateCustomerInput{
		ID:    uint64(1),
		Name:  "Test Customer",
		Email: "test.customer@email.com",
	}

	mockCustomer := &entity.Customer{
		ID:    1,
		Name:  "Updated Customer",
		Email: "updated.customer@email.com",
	}

	mockUseCase.EXPECT().
		Update(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockCustomer}).
		Return([]byte{}, nil)

	output, err := controller.Update(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestCustomerController_DeleteCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockport.NewMockCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := controller.NewCustomerController(mockUseCase)

	ctx := context.Background()
	input := dto.DeleteCustomerInput{
		ID: uint64(1),
	}

	mockCustomer := &entity.Customer{
		ID:    1,
		Name:  "Test Customer",
		Email: "test.customer@email.com",
	}

	mockUseCase.EXPECT().
		Delete(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockCustomer}).
		Return([]byte{}, nil)

	output, err := controller.Delete(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}
