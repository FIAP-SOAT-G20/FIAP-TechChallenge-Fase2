package usecase_test

import (
	"context"
	"testing"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_paymentUseCase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentGateway := mockport.NewMockPaymentGateway(ctrl)
	mockOrderUseCase := mockport.NewMockOrderUseCase(ctrl)
	useCase := usecase.NewPaymentUseCase(mockPaymentGateway, mockOrderUseCase)
	ctx := context.Background()

	tests := []struct {
		name        string
		input       dto.CreatePaymentInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should create payment successfully",
			input: dto.CreatePaymentInput{
				OrderID: uint64(1),
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().FindByOrderIDAndStatusProcessing(ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				mockOrderUseCase.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Order{ID: uint64(1), OrderProducts: []entity.OrderProduct{{OrderID: 1, ProductID: 1}}}, nil)
				mockPaymentGateway.EXPECT().CreateExternal(ctx, gomock.Any()).Return(&entity.CreatePaymentExternalOutput{}, nil)
				mockPaymentGateway.EXPECT().Create(ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				mockOrderUseCase.EXPECT().Update(ctx, gomock.Any()).Return(&entity.Order{ID: 1}, nil)
			},
			expectError: false,
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreatePaymentInput{
				OrderID: uint64(1),
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().FindByOrderIDAndStatusProcessing(ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				mockOrderUseCase.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Order{ID: uint64(1), OrderProducts: []entity.OrderProduct{{OrderID: 1, ProductID: 1}}}, nil)
				mockPaymentGateway.EXPECT().CreateExternal(ctx, gomock.Any()).Return(&entity.CreatePaymentExternalOutput{}, nil)
				mockPaymentGateway.EXPECT().Create(ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				mockOrderUseCase.EXPECT().Update(ctx, gomock.Any()).Return(nil, &domain.InternalError{})
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreatePaymentInput{
				OrderID: uint64(1),
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().FindByOrderIDAndStatusProcessing(ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				mockOrderUseCase.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Order{ID: uint64(1), OrderProducts: []entity.OrderProduct{{OrderID: 1, ProductID: 1}}}, nil)
				mockPaymentGateway.EXPECT().CreateExternal(ctx, gomock.Any()).Return(&entity.CreatePaymentExternalOutput{}, nil)
				mockPaymentGateway.EXPECT().Create(ctx, gomock.Any()).Return(&entity.Payment{}, &domain.InternalError{})
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreatePaymentInput{
				OrderID: uint64(1),
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().FindByOrderIDAndStatusProcessing(ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				mockOrderUseCase.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Order{ID: uint64(1), OrderProducts: []entity.OrderProduct{{OrderID: 1, ProductID: 1}}}, nil)
				mockPaymentGateway.EXPECT().CreateExternal(ctx, gomock.Any()).Return(&entity.CreatePaymentExternalOutput{}, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should return error when dont have order product",
			input: dto.CreatePaymentInput{
				OrderID: uint64(1),
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().FindByOrderIDAndStatusProcessing(ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				mockOrderUseCase.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Order{ID: 1}, nil)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreatePaymentInput{
				OrderID: uint64(1),
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().FindByOrderIDAndStatusProcessing(ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				mockOrderUseCase.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Order{}, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreatePaymentInput{
				OrderID: uint64(1),
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().FindByOrderIDAndStatusProcessing(ctx, gomock.Any()).Return(&entity.Payment{}, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should return the existing payment",
			input: dto.CreatePaymentInput{
				OrderID: uint64(1),
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().FindByOrderIDAndStatusProcessing(ctx, gomock.Any()).Return(&entity.Payment{ID: 1}, nil)
			},
			expectError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			payment, err := useCase.Create(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, payment)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, payment)
			}
		})
	}
}

func Test_paymentUseCase_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentGateway := mockport.NewMockPaymentGateway(ctrl)
	mockOrderUseCase := mockport.NewMockOrderUseCase(ctrl)
	useCase := usecase.NewPaymentUseCase(mockPaymentGateway, mockOrderUseCase)
	ctx := context.Background()

	tests := []struct {
		name        string
		input       dto.GetPaymentInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should return the payment",
			input: dto.GetPaymentInput{
				OrderID: uint64(1),
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().FindByOrderID(ctx, gomock.Any()).Return(&entity.Payment{ID: 1}, nil)
			},
			expectError: false,
		},
		{
			name: "should return error when gateway fails",
			input: dto.GetPaymentInput{
				OrderID: uint64(1),
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().FindByOrderID(ctx, gomock.Any()).Return(&entity.Payment{}, assert.AnError)
			},
			expectError: true,
			errorType:   assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			payment, err := useCase.Get(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, payment)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, payment)
			}
		})
	}
}

func Test_paymentUseCase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentGateway := mockport.NewMockPaymentGateway(ctrl)
	mockOrderUseCase := mockport.NewMockOrderUseCase(ctrl)
	useCase := usecase.NewPaymentUseCase(mockPaymentGateway, mockOrderUseCase)
	ctx := context.Background()

	tests := []struct {
		name        string
		input       dto.UpdatePaymentInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should update the payment",
			input: dto.UpdatePaymentInput{
				Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
				Topic:    "payment",
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().Update(ctx, gomock.Any(), gomock.Any()).Return(nil)
				mockPaymentGateway.EXPECT().FindByExternalPaymentID(ctx, gomock.Any()).Return(&entity.Payment{ID: 1}, nil)
				mockOrderUseCase.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Order{ID: 1}, nil)
				mockOrderUseCase.EXPECT().Update(ctx, gomock.Any()).Return(&entity.Order{ID: 1}, nil)
			},
			expectError: false,
		},
		{
			name: "should return error when gateway fails",
			input: dto.UpdatePaymentInput{
				Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
				Topic:    "payment",
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().Update(ctx, gomock.Any(), gomock.Any()).Return(nil)
				mockPaymentGateway.EXPECT().FindByExternalPaymentID(ctx, gomock.Any()).Return(&entity.Payment{ID: 1}, nil)
				mockOrderUseCase.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Order{ID: 1}, nil)
				mockOrderUseCase.EXPECT().Update(ctx, gomock.Any()).Return(nil, &domain.InternalError{})
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should return error when gateway fails",
			input: dto.UpdatePaymentInput{
				Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
				Topic:    "payment",
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().Update(ctx, gomock.Any(), gomock.Any()).Return(nil)
				mockPaymentGateway.EXPECT().FindByExternalPaymentID(ctx, gomock.Any()).Return(&entity.Payment{ID: 1}, nil)
				mockOrderUseCase.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Order{}, &domain.InternalError{})
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should return error when gateway fails",
			input: dto.UpdatePaymentInput{
				Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
				Topic:    "payment",
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().Update(ctx, gomock.Any(), gomock.Any()).Return(nil)
				mockPaymentGateway.EXPECT().FindByExternalPaymentID(ctx, gomock.Any()).Return(&entity.Payment{}, assert.AnError)
			},
			expectError: true,
			errorType:   assert.AnError,
		},
		{
			name: "should return error when gateway fails",
			input: dto.UpdatePaymentInput{
				Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
				Topic:    "payment",
			},
			setupMocks: func() {
				mockPaymentGateway.EXPECT().Update(ctx, gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
			expectError: true,
			errorType:   assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			payment, err := useCase.Update(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, payment)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, payment)
			}
		})
	}
}
