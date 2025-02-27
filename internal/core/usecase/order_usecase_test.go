package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
)

func TestOrdersUseCase_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderGateway(ctrl)
	historyUseCaseMock := mockport.NewMockOrderHistoryUseCase(ctrl)
	useCase := usecase.NewOrderUseCase(mockGateway, historyUseCaseMock)
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
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), nil, 1, 10, "").
					Return(mockOrders, int64(2), nil)
			},
			expectError: false,
		},
		{
			name: "should return error when repository fails",
			input: dto.ListOrdersInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), nil, 1, 10, "").
					Return(nil, int64(0), assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should filter by status",
			input: dto.ListOrdersInput{
				Status: []valueobject.OrderStatus{"PENDING"},
				Page:   1,
				Limit:  10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(0), []valueobject.OrderStatus{"PENDING"}, 1, 10, "").
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
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, uint64(1), nil, 1, 10, "").
					Return(mockOrders, int64(2), nil)
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			orders, total, err := useCase.List(ctx, tt.input)

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

func TestOrderUseCase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderGateway(ctrl)
	historyUseCaseMock := mockport.NewMockOrderHistoryUseCase(ctrl)
	useCase := usecase.NewOrderUseCase(mockGateway, historyUseCaseMock)
	ctx := context.Background()

	tests := []struct {
		name        string
		input       dto.CreateOrderInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should create order successfully",
			input: dto.CreateOrderInput{
				CustomerID: 1,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)
				historyUseCaseMock.EXPECT().
					Create(ctx, gomock.Any()).
					Return(&entity.OrderHistory{OrderID: 1, Status: valueobject.OPEN}, nil)
			},
			expectError: false,
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreateOrderInput{
				CustomerID: 1,
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

			order, err := useCase.Create(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, order)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, order)
				assert.Equal(t, tt.input.CustomerID, order.CustomerID)
			}
		})
	}
}

func TestOrderUseCase_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderGateway(ctrl)
	historyUseCaseMock := mockport.NewMockOrderHistoryUseCase(ctrl)
	useCase := usecase.NewOrderUseCase(mockGateway, historyUseCaseMock)
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

			order, err := useCase.Get(ctx, dto.GetOrderInput{
				ID: tt.id,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, order)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, order)
				assert.Equal(t, mockOrder.ID, order.ID)
				assert.Equal(t, mockOrder.CustomerID, order.CustomerID)
				assert.Equal(t, mockOrder.Status, order.Status)
			}
		})
	}
}

func TestOrderUseCase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderGateway(ctrl)
	historyUseCaseMock := mockport.NewMockOrderHistoryUseCase(ctrl)
	useCase := usecase.NewOrderUseCase(mockGateway, historyUseCaseMock)
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

			order, err := useCase.Update(ctx, tt.input)

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

func TestOrderUseCase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockOrderGateway(ctrl)
	historyUseCaseMock := mockport.NewMockOrderHistoryUseCase(ctrl)
	useCase := usecase.NewOrderUseCase(mockGateway, historyUseCaseMock)
	ctx := context.Background()

	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should delete order successfully",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(&entity.Order{ID: 1}, nil)

				mockGateway.EXPECT().
					Delete(ctx, uint64(1)).
					Return(nil)
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
			name: "should return error when gateway fails on find",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should return error when gateway fails on delete",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(&entity.Order{}, nil)

				mockGateway.EXPECT().
					Delete(ctx, uint64(1)).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			order, err := useCase.Delete(ctx, dto.DeleteOrderInput{ID: tt.id})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, order)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, order)
				assert.Equal(t, tt.id, order.ID)
			}
		})
	}
}
