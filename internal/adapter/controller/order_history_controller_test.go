package controller

import (
	"context"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

// TODO: Add more test cenarios
func TestOrderHistoryController_ListOrderHistories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderHistoriesUseCase := mockport.NewMockOrderHistoryUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewOrderHistoryController(mockOrderHistoriesUseCase)

	ctx := context.Background()
	input := dto.ListOrderHistoriesInput{
		OrderID: 1,
		Status:  "OPEN",
		Page:    1,
		Limit:   10,
	}

	currentTime := time.Now()
	mockOrderHistories := []*entity.OrderHistory{
		{
			ID:        1,
			OrderID:   1,
			Status:    valueobject.OPEN,
			CreatedAt: currentTime,
		},
		{
			ID:        2,
			OrderID:   1,
			Status:    valueobject.PENDING,
			CreatedAt: currentTime,
		},
		{
			ID:        3,
			OrderID:   2,
			Status:    valueobject.OPEN,
			CreatedAt: currentTime,
		},
	}

	mockOrderHistoriesUseCase.EXPECT().
		List(ctx, input).
		Return(mockOrderHistories, int64(3), nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{
			Result: mockOrderHistories,
			Total:  int64(3),
			Page:   1,
			Limit:  10,
		})

	err := controller.List(ctx, mockPresenter, input)
	assert.NoError(t, err)
}

func TestOrderHistoryController_CreateOrderHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderHistoryUseCase := mockport.NewMockOrderHistoryUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewOrderHistoryController(mockOrderHistoryUseCase)

	ctx := context.Background()
	input := dto.CreateOrderHistoryInput{
		OrderID: 1,
		Status:  "READY",
	}

	mockOrderHistory := &entity.OrderHistory{
		ID:      1,
		OrderID: 1,
		Status:  valueobject.READY,
	}

	mockOrderHistoryUseCase.EXPECT().
		Create(ctx, input).
		Return(mockOrderHistory, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{
			Result: mockOrderHistory,
		})

	err := controller.Create(ctx, mockPresenter, input)
	assert.NoError(t, err)
}

func TestOrderHistoryController_GetOrderHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderHistoryUseCase := mockport.NewMockOrderHistoryUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewOrderHistoryController(mockOrderHistoryUseCase)

	ctx := context.Background()
	input := dto.GetOrderHistoryInput{
		ID: uint64(1),
	}

	mockOrderHistory := &entity.OrderHistory{
		ID:      1,
		OrderID: 1,
		Status:  valueobject.OPEN,
	}

	mockOrderHistoryUseCase.EXPECT().
		Get(ctx, input).
		Return(mockOrderHistory, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{
			Result: &entity.OrderHistory{
				ID:      1,
				OrderID: 1,
				Status:  valueobject.OPEN,
			},
		})

	err := controller.Get(ctx, mockPresenter, input)
	assert.NoError(t, err)
}

func TestOrderHistoryController_DeleteOrderHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderHistoryUseCase := mockport.NewMockOrderHistoryUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewOrderHistoryController(mockOrderHistoryUseCase)

	ctx := context.Background()
	input := dto.DeleteOrderHistoryInput{
		ID: uint64(1),
	}

	mockOrderHistory := &entity.OrderHistory{
		ID:      1,
		OrderID: 1,
		Status:  valueobject.OPEN,
	}

	mockOrderHistoryUseCase.EXPECT().
		Delete(ctx, input).
		Return(mockOrderHistory, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{
			Result: mockOrderHistory,
		})

	err := controller.Delete(ctx, mockPresenter, input)
	assert.NoError(t, err)
}
