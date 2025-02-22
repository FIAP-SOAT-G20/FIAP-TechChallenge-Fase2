package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
)

func TestOrderProductsUseCase_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderProductGateway(ctrl)
	useCase := usecase.NewOrderProductUseCase(mockGateway)
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
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			orderProducts, total, err := useCase.List(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, orderProducts)
				assert.Equal(t, int64(0), total)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, orderProducts)
				assert.Equal(t, len(mockOrderProducts), len(orderProducts))
				assert.Equal(t, int64(2), total)
			}
		})
	}
}

func TestOrderProductUseCase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderProductGateway(ctrl)
	useCase := usecase.NewOrderProductUseCase(mockGateway)
	ctx := context.Background()

	tests := []struct {
		name        string
		input       dto.CreateOrderProductInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should create order-product successfully",
			input: dto.CreateOrderProductInput{
				OrderID:   1,
				ProductID: 1,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)
			},
			expectError: false,
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreateOrderProductInput{
				OrderID:   1,
				ProductID: 1,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					Create(ctx, gomock.Any()).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			orderProduct, err := useCase.Create(ctx, tt.input)

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
			}
		})
	}
}

func TestOrderProductUseCase_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderProductGateway(ctrl)
	useCase := usecase.NewOrderProductUseCase(mockGateway)
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

			orderProduct, err := useCase.Get(ctx, dto.GetOrderProductInput{
				OrderID:   1,
				ProductID: 1,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, orderProduct)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, mockOrderProduct, orderProduct)
				assert.Equal(t, uint64(1), orderProduct.OrderID)
				assert.Equal(t, uint64(1), orderProduct.ProductID)
			}
		})
	}
}

func TestOrderProductUseCase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderProductGateway(ctrl)
	useCase := usecase.NewOrderProductUseCase(mockGateway)
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

			orderProduct, err := useCase.Update(ctx, tt.input)

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

func TestOrderProductUseCase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderProductGateway(ctrl)
	useCase := usecase.NewOrderProductUseCase(mockGateway)
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
					Return(&entity.OrderProduct{OrderID: 1, ProductID: 1}, nil)

				mockGateway.EXPECT().
					Delete(ctx, uint64(1), uint64(1)).
					Return(nil)
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

			orderProduct, err := useCase.Delete(ctx, dto.DeleteOrderProductInput{
				OrderID:   tt.orderID,
				ProductID: tt.productID,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, orderProduct)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, orderProduct)
				assert.Equal(t, tt.orderID, orderProduct.OrderID)
				assert.Equal(t, tt.productID, orderProduct.ProductID)
			}
		})
	}
}
