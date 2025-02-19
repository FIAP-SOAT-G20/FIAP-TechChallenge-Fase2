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
func TestOrderProductController_ListOrderProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListOrderProductsUseCase := mockport.NewMockListOrderProductsUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderProductPresenter(ctrl)
	orderProductController := NewOrderProductController(mockListOrderProductsUseCase, nil, nil, nil, nil)
	orderProductController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.ListOrderProductsInput{
		OrderID: 1,
		Page:    1,
		Limit:   10,
	}

	mockOrderProducts := []*entity.OrderProduct{
		{
			OrderID:   1,
			ProductID: 1,
			Quantity:  1,
		},
		{
			OrderID:   1,
			ProductID: 2,
			Quantity:  2,
		},
	}

	mockListOrderProductsUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockOrderProducts, int64(2), nil)

	mockPresenter.EXPECT().
		Present(dto.OrderProductPresenterInput{
			Result: mockOrderProducts,
			Total:  int64(2),
			Page:   1,
			Limit:  10,
		})

	err := orderProductController.ListOrderProducts(ctx, input)
	assert.NoError(t, err)
}

func TestOrderProductController_CreateOrderProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateOrderProductUseCase := mockport.NewMockCreateOrderProductUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderProductPresenter(ctrl)
	orderProductController := NewOrderProductController(nil, mockCreateOrderProductUseCase, nil, nil, nil)
	orderProductController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.CreateOrderProductInput{
		OrderID:   1,
		ProductID: 1,
		Quantity:  1,
	}

	mockOrderProduct := &entity.OrderProduct{
		OrderID:   1,
		ProductID: 1,
		Quantity:  1,
	}

	mockCreateOrderProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockOrderProduct, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderProductPresenterInput{
			Result: mockOrderProduct,
		})

	err := orderProductController.CreateOrderProduct(ctx, input)
	assert.NoError(t, err)
}

func TestOrderProductController_GetOrderProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetOrderProductUseCase := mockport.NewMockGetOrderProductUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderProductPresenter(ctrl)
	orderProductController := NewOrderProductController(nil, nil, mockGetOrderProductUseCase, nil, nil)
	orderProductController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.GetOrderProductInput{
		OrderID:   1,
		ProductID: 1,
	}

	mockOrderProduct := &entity.OrderProduct{
		OrderID:   1,
		ProductID: 1,
		Quantity:  1,
	}

	mockGetOrderProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockOrderProduct, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderProductPresenterInput{
			Result: mockOrderProduct,
		})

	err := orderProductController.GetOrderProduct(ctx, input)
	assert.NoError(t, err)
}

func TestOrderProductController_UpdateOrderProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUpdateOrderProductUseCase := mockport.NewMockUpdateOrderProductUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderProductPresenter(ctrl)
	orderProductController := NewOrderProductController(nil, nil, nil, mockUpdateOrderProductUseCase, nil)
	orderProductController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.UpdateOrderProductInput{
		OrderID:   1,
		ProductID: 1,
		Quantity:  2,
	}

	mockOrderProduct := &entity.OrderProduct{
		OrderID:   1,
		ProductID: 1,
		Quantity:  2,
	}

	mockUpdateOrderProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockOrderProduct, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderProductPresenterInput{
			Result: mockOrderProduct,
		})

	err := orderProductController.UpdateOrderProduct(ctx, input)
	assert.NoError(t, err)
}

func TestOrderProductController_DeleteOrderProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeleteOrderProductUseCase := mockport.NewMockDeleteOrderProductUseCase(ctrl)
	mockPresenter := mockport.NewMockOrderProductPresenter(ctrl)
	orderProductController := NewOrderProductController(nil, nil, nil, nil, mockDeleteOrderProductUseCase)
	orderProductController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.DeleteOrderProductInput{
		OrderID:   1,
		ProductID: 1,
	}

	mockOrderProduct := &entity.OrderProduct{
		OrderID:   1,
		ProductID: 1,
		Quantity:  1,
	}

	mockDeleteOrderProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(mockOrderProduct, nil)

	mockPresenter.EXPECT().
		Present(dto.OrderProductPresenterInput{
			Result: mockOrderProduct,
		})

	err := orderProductController.DeleteOrderProduct(ctx, input)
	assert.NoError(t, err)
}
