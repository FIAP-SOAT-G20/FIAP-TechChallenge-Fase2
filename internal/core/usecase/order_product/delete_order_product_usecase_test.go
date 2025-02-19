package orderproduct_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	orderproduct "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/order_product"
)

func TestDeleteOrderProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderProductGateway(ctrl)
	mockPresenter := mockport.NewMockOrderProductPresenter(ctrl)
	useCase := orderproduct.NewDeleteOrderProductUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	tests := []struct {
		name        string
		orderID     uint64
		productID   uint64
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name:      "should delete orderProduct successfully",
			orderID:   1,
			productID: 1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1), uint64(1)).
					Return(&entity.OrderProduct{}, nil)

				mockGateway.EXPECT().
					Delete(ctx, uint64(1), uint64(1)).
					Return(nil)

				mockPresenter.EXPECT().
					Present(gomock.Any())
			},
			expectError: false,
		},
		{
			name:      "should return not found error when orderProduct doesn't exist",
			orderID:   1,
			productID: 1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1), uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name:      "should return error when gateway fails on find",
			orderID:   1,
			productID: 1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1), uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name:      "should return error when gateway fails on delete",
			orderID:   1,
			productID: 1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1), uint64(1)).
					Return(&entity.OrderProduct{}, nil)

				mockGateway.EXPECT().
					Delete(ctx, uint64(1), uint64(1)).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := useCase.Execute(ctx, dto.DeleteOrderProductInput{
				OrderID:   tt.orderID,
				ProductID: tt.productID,
			})

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
