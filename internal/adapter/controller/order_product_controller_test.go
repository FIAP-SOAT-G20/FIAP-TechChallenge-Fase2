package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

func TestOrderProductController_ListOrderProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListOrderProductsUseCase := mockport.NewMockListOrderProductsUseCase(ctrl)
	productController := NewOrderProductController(mockListOrderProductsUseCase, nil, nil, nil, nil)

	ctx := context.Background()
	input := dto.ListOrderProductsInput{
		OrderID:   1,
		ProductID: 1,
		Page:      1,
		Limit:     10,
	}

	mockListOrderProductsUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.ListOrderProducts(ctx, input)
	assert.NoError(t, err)
}

func TestOrderProductController_CreateOrderProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateOrderProductUseCase := mockport.NewMockCreateOrderProductUseCase(ctrl)
	productController := NewOrderProductController(nil, mockCreateOrderProductUseCase, nil, nil, nil)

	ctx := context.Background()
	input := dto.CreateOrderProductInput{
		OrderID:   1,
		ProductID: 1,
		Quantity:  1,
	}

	mockCreateOrderProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.CreateOrderProduct(ctx, input)
	assert.NoError(t, err)
}

func TestOrderProductController_GetOrderProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetOrderProductUseCase := mockport.NewMockGetOrderProductUseCase(ctrl)
	productController := NewOrderProductController(nil, nil, mockGetOrderProductUseCase, nil, nil)

	ctx := context.Background()
	input := dto.GetOrderProductInput{
		OrderID:   1,
		ProductID: 1,
	}

	mockGetOrderProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.GetOrderProduct(ctx, input)
	assert.NoError(t, err)
}

func TestOrderProductController_UpdateOrderProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUpdateOrderProductUseCase := mockport.NewMockUpdateOrderProductUseCase(ctrl)
	productController := NewOrderProductController(nil, nil, nil, mockUpdateOrderProductUseCase, nil)

	ctx := context.Background()
	input := dto.UpdateOrderProductInput{
		OrderID:   1,
		ProductID: 1,
		Quantity:  1,
	}

	mockUpdateOrderProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.UpdateOrderProduct(ctx, input)
	assert.NoError(t, err)
}

func TestOrderProductController_DeleteOrderProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeleteOrderProductUseCase := mockport.NewMockDeleteOrderProductUseCase(ctrl)
	productController := NewOrderProductController(nil, nil, nil, nil, mockDeleteOrderProductUseCase)

	ctx := context.Background()
	input := dto.DeleteOrderProductInput{
		OrderID:   1,
		ProductID: 1,
	}

	mockDeleteOrderProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.DeleteOrderProduct(ctx, input)
	assert.NoError(t, err)
}
