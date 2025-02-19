package order_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/order"
)

func TestListOrdersUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderGateway(ctrl)
	useCase := order.NewListOrdersUseCase(mockGateway)
	ctx := context.Background()

	currentTime := time.Now()
	mockOrders := []*entity.Order{
		{
			ID:         1,
			CustomerID: uint64(1),
			TotalBill:  99.99,
			Status:     valueobject.PENDING,
			CreatedAt:  currentTime,
			UpdatedAt:  currentTime,
		},
		{
			ID:         2,
			CustomerID: uint64(2),
			TotalBill:  199.99,
			Status:     valueobject.PENDING,
			CreatedAt:  currentTime,
			UpdatedAt:  currentTime,
		},
	}

	tests := []struct {
		name        string
		input       dto.ListOrdersInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should list orders successfully",
			input: dto.ListOrdersInput{
				Page:   1,
				Limit:  10,
				Writer: nil,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), valueobject.OrderStatus(""), 1, 10).
					Return(mockOrders, int64(2), nil)
			},
			expectError: false,
		},
		{
			name: "should return error when repository fails",
			input: dto.ListOrdersInput{
				Page:   1,
				Limit:  10,
				Writer: nil,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), valueobject.OrderStatus(""), 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should filter by status",
			input: dto.ListOrdersInput{
				Status: "PENDING",
				Page:   1,
				Limit:  10,
				Writer: nil,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), valueobject.OrderStatus("PENDING"), 1, 10).
					Return(mockOrders, int64(2), nil)
			},
			expectError: false,
		},
		{
			name: "should filter by customer",
			input: dto.ListOrdersInput{
				CustomerID: 1,
				Page:       1,
				Limit:      10,
				Writer:     nil,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(1), valueobject.OrderStatus(""), 1, 10).
					Return(mockOrders, int64(2), nil)
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			orders, total, err := useCase.Execute(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, orders)
				assert.Zero(t, total)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, orders)
				assert.Len(t, orders, 2)
				assert.Equal(t, int64(2), total)
			}
		})
	}
}
