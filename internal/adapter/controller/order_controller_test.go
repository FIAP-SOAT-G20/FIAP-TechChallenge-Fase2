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

	mokOrdercUseCase := mockport.NewMockOrderUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	orderController := NewOrderController(mokOrdercUseCase)
	orderController.Presenter = mockPresenter

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

	mokOrdercUseCase.EXPECT().
		List(ctx, input).
		Return(mockOrders, int64(2), nil)

	mockPresenter.EXPECT().
		Present(dto.OrderPresenterInput{
			Result: mockOrders,
			Total:  int64(2),
			Page:   1,
			Limit:  10,
		})

	err := orderController.ListOrders(ctx, input)
	assert.NoError(t, err)
}

func TestOrderController_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mokOrdercUseCase := mockport.NewMockOrderUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	orderController := NewOrderController(mokOrdercUseCase)
	orderController.Presenter = mockPresenter

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

	mokOrdercUseCase.EXPECT().
		Create(ctx, input).
		Return(mockOrder, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderPresenterInput{
			Result: mockOrder,
		})

	err := orderController.CreateOrder(ctx, input)
	assert.NoError(t, err)
}

func TestOrderController_GetOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mokOrdercUseCase := mockport.NewMockOrderUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	orderController := NewOrderController(mokOrdercUseCase)
	orderController.Presenter = mockPresenter

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

	mokOrdercUseCase.EXPECT().
		Get(ctx, input).
		Return(mockOrder, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderPresenterInput{
			Result: mockOrder,
		})

	err := orderController.GetOrder(ctx, input)
	assert.NoError(t, err)
}

func TestOrderController_UpdateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mokOrdercUseCase := mockport.NewMockOrderUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	orderController := NewOrderController(mokOrdercUseCase)
	orderController.Presenter = mockPresenter

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

	mokOrdercUseCase.EXPECT().
		Update(ctx, input).
		Return(mockOrder, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderPresenterInput{
			Result: mockOrder,
		})

	err := orderController.UpdateOrder(ctx, input)
	assert.NoError(t, err)
}

func TestOrderController_DeleteOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mokOrdercUseCase := mockport.NewMockOrderUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	orderController := NewOrderController(mokOrdercUseCase)
	orderController.Presenter = mockPresenter

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

	mokOrdercUseCase.EXPECT().
		Delete(ctx, input).
		Return(mockOrder, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderPresenterInput{
			Result: mockOrder,
		})

	err := orderController.DeleteOrder(ctx, input)
	assert.NoError(t, err)
}
