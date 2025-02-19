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

func TestUpdateOrderProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderProductGateway(ctrl)
	useCase := orderproduct.NewUpdateOrderProductUseCase(mockGateway)
	ctx := context.Background()

	currentTime := time.Now()
	existingOrderProduct := &entity.OrderProduct{
		OrderID:   1,
		ProductID: 1,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	tests := []struct {
		name        string
		input       dto.UpdateOrderProductInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should update orderProduct successfully",
			input: dto.UpdateOrderProductInput{
				OrderID:   1,
				ProductID: 1,
				Quantity:  1,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1), uint64(1)).
					Return(existingOrderProduct, nil)

				mockGateway.EXPECT().
					Update(ctx, gomock.Any()).
					DoAndReturn(func(_ context.Context, p *entity.OrderProduct) error {
						assert.Equal(t, uint64(1), p.OrderID)
						assert.Equal(t, uint64(1), p.ProductID)
						assert.Equal(t, uint32(1), p.Quantity)
						return nil
					})
			},
			expectError: false,
		},
		{
			name: "should return error when orderProduct not found",
			input: dto.UpdateOrderProductInput{
				OrderID:   1,
				ProductID: 1,
				Quantity:  1,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1), uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return error when gateway update fails",
			input: dto.UpdateOrderProductInput{
				OrderID:   1,
				ProductID: 1,
				Quantity:  1,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1), uint64(1)).
					Return(existingOrderProduct, nil)

				mockGateway.EXPECT().
					Update(ctx, gomock.Any()).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			orderProduct, err := useCase.Execute(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, orderProduct)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, orderProduct)
				assert.Equal(t, tt.input.OrderID, orderProduct.OrderID)
				assert.Equal(t, tt.input.ProductID, orderProduct.ProductID)
				assert.Equal(t, tt.input.Quantity, orderProduct.Quantity)
			}
		})
	}
}
