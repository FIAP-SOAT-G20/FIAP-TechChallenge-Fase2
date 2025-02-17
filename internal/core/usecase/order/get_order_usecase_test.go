package order_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	mockdto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/order"
)

func TestGetOrderUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderGateway(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)
	useCase := order.NewGetOrderUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	currentTime := time.Now()
	mockOrder := &entity.Order{
		ID:         1,
		CustomerID: 1,
		Status:     "PENDING",
		TotalBill:  100.0,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}

	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should get order successfully",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(mockOrder, nil)

				mockPresenter.EXPECT().
					Present(dto.OrderPresenterInput{
						Writer: mockWriter,
						Result: mockOrder,
					})
			},
			expectError: false,
		},
		{
			name: "should return not found error when order doesn't exist",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return internal error when gateway fails",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := useCase.Execute(ctx, dto.GetOrderInput{
				ID:     tt.id,
				Writer: mockWriter,
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
