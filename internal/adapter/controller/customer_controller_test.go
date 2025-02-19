package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

func TestCustomerController_ListCustomers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListCustomersUseCase := mockport.NewMockListCustomersUseCase(ctrl)
	mockPresenter := mockport.NewMockCustomerPresenter(ctrl)
	productController := NewCustomerController(mockListCustomersUseCase, nil, nil, nil, nil)
	productController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.ListCustomersInput{
		Name:  "Test",
		Page:  1,
		Limit: 10,
	}

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

	mockListCustomersUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockCustomers, int64(2), nil)

	mockPresenter.EXPECT().
		Present(dto.CustomerPresenterInput{
			Total:  int64(2),
			Page:   1,
			Limit:  10,
			Result: mockCustomers,
		})

	err := productController.ListCustomers(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_CreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateCustomerUseCase := mockport.NewMockCreateCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockCustomerPresenter(ctrl)
	productController := NewCustomerController(nil, mockCreateCustomerUseCase, nil, nil, nil)
	productController.Presenter = mockPresenter

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

	mockCreateCustomerUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.CustomerPresenterInput{
			Result: mockCustomer,
		})

	err := productController.CreateCustomer(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_GetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetCustomerUseCase := mockport.NewMockGetCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockCustomerPresenter(ctrl)
	productController := NewCustomerController(nil, nil, mockGetCustomerUseCase, nil, nil)
	productController.Presenter = mockPresenter

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

	mockGetCustomerUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.CustomerPresenterInput{
			Result: mockCustomer,
		})

	err := productController.GetCustomer(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_UpdateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUpdateCustomerUseCase := mockport.NewMockUpdateCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockCustomerPresenter(ctrl)
	productController := NewCustomerController(nil, nil, nil, mockUpdateCustomerUseCase, nil)
	productController.Presenter = mockPresenter

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

	mockUpdateCustomerUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.CustomerPresenterInput{
			Result: mockCustomer,
		})

	err := productController.UpdateCustomer(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_DeleteCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeleteCustomerUseCase := mockport.NewMockDeleteCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockCustomerPresenter(ctrl)
	productController := NewCustomerController(nil, nil, nil, nil, mockDeleteCustomerUseCase)
	productController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.DeleteCustomerInput{
		ID: uint64(1),
	}

	mockCustomer := &entity.Customer{
		ID:    1,
		Name:  "Test Customer",
		Email: "test.customer@email.com",
	}

	mockDeleteCustomerUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.CustomerPresenterInput{
			Result: mockCustomer,
		})

	err := productController.DeleteCustomer(ctx, input)
	assert.NoError(t, err)
}
