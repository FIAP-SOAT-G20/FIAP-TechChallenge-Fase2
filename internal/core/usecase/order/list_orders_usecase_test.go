package order

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
)

func TestListOrdersUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderGateway(ctrl)
	mockPresenter := mockport.NewMockOrderPresenter(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)
	useCase := NewListOrdersUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	currentTime := time.Now()
	mockOrders := []*entity.Order{
		{
			ID:         1,
			CustomerID: uint64(1),
			TotalBill:  99.99,
			Status:     "PENDING",
			CreatedAt:  currentTime,
			UpdatedAt:  currentTime,
		},
		{
			ID:         2,
			CustomerID: uint64(2),
			TotalBill:  199.99,
			Status:     "DELIVERED",
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
				Writer: mockWriter,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), "", 1, 10).
					Return(mockOrders, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.OrderPresenterInput{
						Result: mockOrders,
						Total:  int64(2),
						Page:   1,
						Limit:  10,
						Writer: mockWriter,
					})
			},
			expectError: false,
		},
		{
			name: "should return error when repository fails",
			input: dto.ListOrdersInput{
				Page:   1,
				Limit:  10,
				Writer: mockWriter,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), "", 1, 10).
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
				Writer: mockWriter,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), "PENDING", 1, 10).
					Return(mockOrders, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.OrderPresenterInput{
						Writer: mockWriter,
						Total:  int64(2),
						Page:   1,
						Limit:  10,
						Result: mockOrders,
					})
			},
			expectError: false,
		},
		{
			name: "should filter by customer",
			input: dto.ListOrdersInput{
				CustomerID: 1,
				Page:       1,
				Limit:      10,
				Writer:     mockWriter,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(1), "", 1, 10).
					Return(mockOrders, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.OrderPresenterInput{
						Writer: mockWriter,
						Total:  int64(2),
						Page:   1,
						Limit:  10,
						Result: mockOrders,
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
