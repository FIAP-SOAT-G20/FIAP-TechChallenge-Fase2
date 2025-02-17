package orderproduct_test

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
	orderproduct "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/order_product"
)

func TestGetOrderProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderProductGateway(ctrl)
	mockPresenter := mockport.NewMockOrderProductPresenter(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)
	useCase := orderproduct.NewGetOrderProductUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	currentTime := time.Now()
	mockOrderProduct := &entity.OrderProduct{
		OrderID:   1,
		ProductID: 1,
		Quantity:  1,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should get orderProduct successfully",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1), uint64(1)).
					Return(mockOrderProduct, nil)

				mockPresenter.EXPECT().
					Present(dto.OrderProductPresenterInput{
						Writer: mockWriter,
						Result: mockOrderProduct,
					})
			},
			expectError: false,
		},
		{
			name: "should return not found error when orderProduct doesn't exist",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1), uint64(1)).
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
					FindByID(ctx, uint64(1), uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := useCase.Execute(ctx, dto.GetOrderProductInput{
				OrderID:   1,
				ProductID: 1,
				Writer:    mockWriter,
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
