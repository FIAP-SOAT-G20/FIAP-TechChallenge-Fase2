package orderproduct_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	orderproduct "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/order_product"
)

func TestListOrderProductsUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderProductGateway(ctrl)
	mockPresenter := mockport.NewMockOrderProductPresenter(ctrl)
	useCase := orderproduct.NewListOrderProductsUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	currentTime := time.Now()
	mockOrderProducts := []*entity.OrderProduct{
		{
			OrderID:   1,
			ProductID: 1,
			Quantity:  1,
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
		{
			OrderID:   2,
			ProductID: 2,
			Quantity:  2,
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}

	tests := []struct {
		name        string
		input       dto.ListOrderProductsInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should list orderProducts successfully",
			input: dto.ListOrderProductsInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), uint64(0), 1, 10).
					Return(mockOrderProducts, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.OrderProductPresenterInput{
						Result: mockOrderProducts,
						Total:  int64(2),
						Page:   1,
						Limit:  10,
					})
			},
			expectError: false,
		},
		{
			name: "should return error when repository fails",
			input: dto.ListOrderProductsInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), uint64(0), 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should filter by order id",
			input: dto.ListOrderProductsInput{
				OrderID: 1,
				Page:    1,
				Limit:   10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(1), uint64(0), 1, 10).
					Return(mockOrderProducts, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.OrderProductPresenterInput{
						Total:  int64(2),
						Page:   1,
						Limit:  10,
						Result: mockOrderProducts,
					})
			},
			expectError: false,
		},
		{
			name: "should filter by product id",
			input: dto.ListOrderProductsInput{
				ProductID: 1,
				Page:      1,
				Limit:     10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), uint64(1), 1, 10).
					Return(mockOrderProducts, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.OrderProductPresenterInput{
						Total:  int64(2),
						Page:   1,
						Limit:  10,
						Result: mockOrderProducts,
					})
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := useCase.Execute(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
