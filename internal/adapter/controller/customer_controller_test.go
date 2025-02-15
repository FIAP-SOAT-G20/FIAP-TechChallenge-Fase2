package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	mockdto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto/mocks"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

func TestCustomerController_ListCustomers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListCustomersUseCase := mockport.NewMockListCustomersUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewCustomerController(mockListCustomersUseCase, nil, nil, nil, nil)

	ctx := context.Background()
	input := dto.ListCustomersInput{
		Name:   "Test",
		Page:   1,
		Limit:  10,
		Writer: mockResponseWriter,
	}

	mockListCustomersUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.ListCustomers(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_CreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateCustomerUseCase := mockport.NewMockCreateCustomerUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewCustomerController(nil, mockCreateCustomerUseCase, nil, nil, nil)

	ctx := context.Background()
	input := dto.CreateCustomerInput{
		Name:   "Test Customer",
		Email:  "test.customer.1@email.com",
		CPF:    "123.456.789-00",
		Writer: mockResponseWriter,
	}

	mockCreateCustomerUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.CreateCustomer(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_GetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetCustomerUseCase := mockport.NewMockGetCustomerUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewCustomerController(nil, nil, mockGetCustomerUseCase, nil, nil)

	ctx := context.Background()
	input := dto.GetCustomerInput{
		ID:     uint64(1),
		Writer: mockResponseWriter,
	}

	mockGetCustomerUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.GetCustomer(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_UpdateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUpdateCustomerUseCase := mockport.NewMockUpdateCustomerUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewCustomerController(nil, nil, nil, mockUpdateCustomerUseCase, nil)

	ctx := context.Background()
	input := dto.UpdateCustomerInput{
		ID:     uint64(1),
		Name:   "Updated Customer",
		Email:  "updated.customer@email.com",
		Writer: mockResponseWriter,
	}

	mockUpdateCustomerUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.UpdateCustomer(ctx, input)
	assert.NoError(t, err)
}

func TestCustomerController_DeleteCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeleteCustomerUseCase := mockport.NewMockDeleteCustomerUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewCustomerController(nil, nil, nil, nil, mockDeleteCustomerUseCase)

	ctx := context.Background()
	input := dto.DeleteCustomerInput{
		ID:     uint64(1),
		Writer: mockResponseWriter,
	}

	mockDeleteCustomerUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.DeleteCustomer(ctx, input)
	assert.NoError(t, err)
}
