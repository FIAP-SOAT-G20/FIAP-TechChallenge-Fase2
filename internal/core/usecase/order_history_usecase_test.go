package usecase_test

import (
	"testing"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *OrderHistoryUsecaseSuiteTest) TestOrderHistoriesUseCase_List() {
	currentTime := time.Now()
	mockOrderHistories := []*entity.OrderHistory{
		{
			ID:        1,
			OrderID:   1,
			Status:    "OPEN",
			CreatedAt: currentTime,
		},
		{
			ID:        2,
			OrderID:   1,
			Status:    "PENDING",
			CreatedAt: currentTime,
		},
	}

	tests := []struct {
		name        string
		input       dto.ListOrderHistoriesInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should list staffs successfully",
			input: dto.ListOrderHistoriesInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				var status valueobject.OrderStatus
				s.mockGateway.EXPECT().
					FindAll(s.ctx, uint64(0), status, 1, 10).
					Return(mockOrderHistories, int64(2), nil)
			},
			expectError: false,
		},
		{
			name: "should return error when repository fails",
			input: dto.ListOrderHistoriesInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				var status valueobject.OrderStatus
				s.mockGateway.EXPECT().
					FindAll(s.ctx, uint64(0), status, 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should filter by orderID",
			input: dto.ListOrderHistoriesInput{
				OrderID: 1,
				Page:    1,
				Limit:   10,
			},
			setupMocks: func() {
				var status valueobject.OrderStatus
				s.mockGateway.EXPECT().
					FindAll(s.ctx, uint64(1), status, 1, 10).
					Return(mockOrderHistories, int64(2), nil)
			},
			expectError: false,
		},
		{
			name: "should filter by status",
			input: dto.ListOrderHistoriesInput{
				Status: "OPEN",
				Page:   1,
				Limit:  10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, uint64(0), valueobject.OPEN, 1, 10).
					Return(mockOrderHistories, int64(1), nil)

			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			_, _, err := s.useCase.List(s.ctx, tt.input)

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

func (s *OrderHistoryUsecaseSuiteTest) TestOrderHistoryUseCase_Create() {
	tests := []struct {
		name        string
		input       dto.CreateOrderHistoryInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should create staff successfully",
			input: dto.CreateOrderHistoryInput{
				OrderID: 1,
				Status:  "OPEN",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					Create(s.ctx, gomock.Any()).
					Return(nil)
			},
			expectError: false,
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreateOrderHistoryInput{
				OrderID: 1,
				Status:  "OPEN",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					Create(s.ctx, gomock.Any()).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			_, err := s.useCase.Create(s.ctx, tt.input)

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

func (s *OrderHistoryUsecaseSuiteTest) TestOrderHistoryUseCase_Get() {
	currentTime := time.Now()
	mockOrderHistory := &entity.OrderHistory{
		ID:        1,
		OrderID:   1,
		Status:    "OPEN",
		CreatedAt: currentTime,
	}

	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should get order history successfully",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(mockOrderHistory, nil)
			},
			expectError: false,
		},
		{
			name: "should return not found error when order history doesn't exist",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return internal error when gateway fails",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			_, err := s.useCase.Get(s.ctx, dto.GetOrderHistoryInput{
				ID: tt.id,
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

func (s *OrderHistoryUsecaseSuiteTest) TestOrderHistoryUseCase_Delete() {
	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should delete order history successfully",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(&entity.OrderHistory{}, nil)

				s.mockGateway.EXPECT().
					Delete(s.ctx, uint64(1)).
					Return(nil)
			},
			expectError: false,
		},
		{
			name: "should return not found error when order history doesn't exist",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return error when gateway fails on find",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should return error when gateway fails on delete",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(&entity.OrderHistory{}, nil)

				s.mockGateway.EXPECT().
					Delete(s.ctx, uint64(1)).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			_, err := s.useCase.Delete(s.ctx, dto.DeleteOrderHistoryInput{ID: tt.id})

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
