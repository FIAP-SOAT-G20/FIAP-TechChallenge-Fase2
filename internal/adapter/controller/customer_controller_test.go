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

// TODO: Add more test cenarios
func TestCustomerController_ListCustomers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewCustomerController(mockCustomerUseCase)
	controller.Presenter = mockPresenter

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

	mockCustomerUseCase.EXPECT().
		List(ctx, input).
		Return(mockCustomers, int64(2), nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{
			Total:  int64(2),
			Page:   1,
			Limit:  10,
			Result: mockCustomers,
		})

	err := controller.List(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_CreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewCustomerController(mockCustomerUseCase)
	controller.Presenter = mockPresenter

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

	mockCustomerUseCase.EXPECT().
		Create(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{
			Result: mockCustomer,
		})

	err := controller.Create(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_GetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewCustomerController(mockCustomerUseCase)
	controller.Presenter = mockPresenter

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

	mockCustomerUseCase.EXPECT().
		Get(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{
			Result: mockCustomer,
		})

	err := controller.Get(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_UpdateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewCustomerController(mockCustomerUseCase)
	controller.Presenter = mockPresenter

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

	mockCustomerUseCase.EXPECT().
		Update(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{
			Result: mockCustomer,
		})

	err := controller.Update(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_DeleteCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewCustomerController(mockCustomerUseCase)
	controller.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.DeleteCustomerInput{
		ID: uint64(1),
	}

	mockCustomer := &entity.Customer{
		ID:    1,
		Name:  "Test Customer",
		Email: "test.customer@email.com",
	}

	mockCustomerUseCase.EXPECT().
		Delete(ctx, input).
		Return(mockCustomer, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{
			Result: mockCustomer,
		})

	err := controller.Delete(ctx, input)
	assert.NoError(t, err)
}
