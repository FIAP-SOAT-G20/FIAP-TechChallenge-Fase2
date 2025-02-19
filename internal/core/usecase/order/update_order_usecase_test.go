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
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/order"
)

func TestUpdateOrderUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderGateway(ctrl)
	useCase := order.NewUpdateOrderUseCase(mockGateway)
	ctx := context.Background()

	currentTime := time.Now()
	existingOrder := &entity.Order{
		ID:         1,
		CustomerID: 1,
		Status:     "PENDING",
		TotalBill:  100.0,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}

	tests := []struct {
		name        string
		input       dto.UpdateOrderInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should update order successfully",
			input: dto.UpdateOrderInput{
				ID:         1,
				CustomerID: 1,
				Status:     "RECEIVED",
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(existingOrder, nil)

				mockGateway.EXPECT().
					Update(ctx, gomock.Any()).
					DoAndReturn(func(_ context.Context, p *entity.Order) error {
						assert.Equal(t, uint64(1), p.ID)
						return nil
					})
			},
			expectError: false,
		},
		{
			name: "should return error when order not found",
			input: dto.UpdateOrderInput{
				ID:         1,
				CustomerID: 1,
				Status:     "RECEIVED",
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return error when gateway update fails",
			input: dto.UpdateOrderInput{
				ID:         1,
				CustomerID: 1,
				Status:     "RECEIVED",
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(existingOrder, nil)

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

			order, err := useCase.Execute(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, order)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, order)
				assert.Equal(t, tt.input.Status, order.Status)
			}
		})
	}
}
