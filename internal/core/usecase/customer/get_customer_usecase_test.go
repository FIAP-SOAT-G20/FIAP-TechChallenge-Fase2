package customer_test

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
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/customer"
)

func TestGetCustomerUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockCustomerGateway(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)
	useCase := customer.NewGetCustomerUseCase(mockGateway)
	ctx := context.Background()

	currentTime := time.Now()
	mockCustomer := &entity.Customer{
		ID:        1,
		Name:      "Test Customer",
		Email:     "test.customer@email.com",
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
			name: "should get customer successfully",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(mockCustomer, nil)
			},
			expectError: false,
		},
		{
			name: "should return not found error when customer doesn't exist",
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

			customer, err := useCase.Execute(ctx, dto.GetCustomerInput{
				ID:     tt.id,
				Writer: mockWriter,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, customer)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, customer)
				assert.Equal(t, mockCustomer.ID, customer.ID)
				assert.Equal(t, mockCustomer.Name, customer.Name)
				assert.Equal(t, mockCustomer.Email, customer.Email)
				assert.Equal(t, mockCustomer.CreatedAt, customer.CreatedAt)
				assert.Equal(t, mockCustomer.UpdatedAt, customer.UpdatedAt)
			}
		})
	}
}
