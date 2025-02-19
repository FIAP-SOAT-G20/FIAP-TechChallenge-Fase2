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
func TestOrderController_ListOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListOrdersUseCase := mockport.NewMockListOrdersUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	productController := NewOrderController(mockListOrdersUseCase, nil, nil, nil, nil, mockPresenter)

	ctx := context.Background()
	input := dto.ListOrdersInput{
		CustomerID: 1,
		Status:     "PENDING",
		Page:       1,
		Limit:      10,
	}

	mockOrders := []*entity.Order{
		{
			ID:         1,
			CustomerID: 1,
			Status:     "PENDING",
			TotalBill:  100.0,
		},
		{
			ID:         2,
			CustomerID: 1,
			Status:     "PENDING",
			TotalBill:  200.0,
		},
	}

	mockListOrdersUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockOrders, int64(2), nil)

	mockPresenter.EXPECT().
		Present(dto.OrderPresenterInput{
			Result: mockOrders,
			Total:  int64(2),
			Page:   1,
			Limit:  10,
		})

	err := productController.ListOrders(ctx, input)
	assert.NoError(t, err)
}

func TestOrderController_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateOrderUseCase := mockport.NewMockCreateOrderUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	productController := NewOrderController(nil, mockCreateOrderUseCase, nil, nil, nil, mockPresenter)

	ctx := context.Background()
	input := dto.CreateOrderInput{
		CustomerID: 1,
	}

	mockOrder := &entity.Order{
		ID:         1,
		CustomerID: 1,
		Status:     "OPEN",
		TotalBill:  0.0,
	}

	mockCreateOrderUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockOrder, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderPresenterInput{
			Result: mockOrder,
		})

	err := productController.CreateOrder(ctx, input)
	assert.NoError(t, err)
}

func TestOrderController_GetOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetOrderUseCase := mockport.NewMockGetOrderUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	productController := NewOrderController(nil, nil, mockGetOrderUseCase, nil, nil, mockPresenter)

	ctx := context.Background()
	input := dto.GetOrderInput{
		ID: uint64(1),
	}

	mockOrder := &entity.Order{
		ID:         1,
		CustomerID: 1,
		Status:     "PENDING",
		TotalBill:  100.0,
	}

	mockGetOrderUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockOrder, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderPresenterInput{
			Result: mockOrder,
		})

	err := productController.GetOrder(ctx, input)
	assert.NoError(t, err)
}

func TestOrderController_UpdateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUpdateOrderUseCase := mockport.NewMockUpdateOrderUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	productController := NewOrderController(nil, nil, nil, mockUpdateOrderUseCase, nil, mockPresenter)

	ctx := context.Background()
	input := dto.UpdateOrderInput{
		ID:         uint64(1),
		CustomerID: 1,
		Status:     "OPEN",
	}

	mockOrder := &entity.Order{
		ID:         1,
		CustomerID: 1,
		Status:     "PENDING",
		TotalBill:  100.0,
	}

	mockUpdateOrderUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockOrder, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderPresenterInput{
			Result: mockOrder,
		})

	err := productController.UpdateOrder(ctx, input)
	assert.NoError(t, err)
}

func TestOrderController_DeleteOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeleteOrderUseCase := mockport.NewMockDeleteOrderUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	productController := NewOrderController(nil, nil, nil, nil, mockDeleteOrderUseCase, mockPresenter)

	ctx := context.Background()
	input := dto.DeleteOrderInput{
		ID: uint64(1),
	}

	mockOrder := &entity.Order{
		ID:         1,
		CustomerID: 1,
		Status:     "PENDING",
		TotalBill:  100.0,
	}

	mockDeleteOrderUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockOrder, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderPresenterInput{
			Result: mockOrder,
		})

	err := productController.DeleteOrder(ctx, input)
	assert.NoError(t, err)
}
