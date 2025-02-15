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

func TestOrderController_ListOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListOrdersUseCase := mockport.NewMockListOrdersUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	// productController := NewOrderController(mockListOrdersUseCase, nil, nil, nil, nil)
	productController := NewOrderController(mockListOrdersUseCase, nil, nil)

	ctx := context.Background()
	input := dto.ListOrdersInput{
		CustomerID: 1,
		Status: 	 "PENDING",
		Page:       1,
		Limit:      10,
		Writer:     mockResponseWriter,
	}

	mockListOrdersUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.ListOrders(ctx, input)
	assert.NoError(t, err)
}

func TestOrderController_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateOrderUseCase := mockport.NewMockCreateOrderUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	// productController := NewOrderController(nil, mockCreateOrderUseCase, nil, nil, nil)
	productController := NewOrderController(nil, mockCreateOrderUseCase, nil)

	ctx := context.Background()
	input := dto.CreateOrderInput{
		CustomerID: 1,
		Writer:      mockResponseWriter,
	}

	mockCreateOrderUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.CreateOrder(ctx, input)
	assert.NoError(t, err)
}

func TestOrderController_GetOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetOrderUseCase := mockport.NewMockGetOrderUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	// productController := NewOrderController(nil, nil, mockGetOrderUseCase, nil, nil)
	productController := NewOrderController(nil, nil, mockGetOrderUseCase)

	ctx := context.Background()
	input := dto.GetOrderInput{
		ID:     uint64(1),
		Writer: mockResponseWriter,
	}

	mockGetOrderUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.GetOrder(ctx, input)
	assert.NoError(t, err)
}

// func TestOrderController_UpdateOrder(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUpdateOrderUseCase := mockport.NewMockUpdateOrderUseCase(ctrl)
// 	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
// 	productController := NewOrderController(nil, nil, nil, mockUpdateOrderUseCase, nil)

// 	ctx := context.Background()
// 	input := dto.UpdateOrderInput{
// 		ID:          uint64(1),
// 		Name:        "Updated Order",
// 		Description: "Updated Description",
// 		Price:       199.99,
// 		CategoryID:  2,
// 		Writer:      mockResponseWriter,
// 	}

// 	mockUpdateOrderUseCase.EXPECT().
// 		Execute(ctx, input).
// 		Return(nil)

// 	err := productController.UpdateOrder(ctx, input)
// 	assert.NoError(t, err)
// }

// func TestOrderController_DeleteOrder(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockDeleteOrderUseCase := mockport.NewMockDeleteOrderUseCase(ctrl)
// 	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
// 	productController := NewOrderController(nil, nil, nil, nil, mockDeleteOrderUseCase)

// 	ctx := context.Background()
// 	input := dto.DeleteOrderInput{
// 		ID:     uint64(1),
// 		Writer: mockResponseWriter,
// 	}

// 	mockDeleteOrderUseCase.EXPECT().
// 		Execute(ctx, input).
// 		Return(nil)

// 	err := productController.DeleteOrder(ctx, input)
// 	assert.NoError(t, err)
// }
