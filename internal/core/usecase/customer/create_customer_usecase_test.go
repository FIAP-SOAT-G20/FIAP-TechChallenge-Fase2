package customer_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	mockdto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/customer"
)

func TestCreateCustomerUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockCustomerGateway(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)
	useCase := customer.NewCreateCustomerUseCase(mockGateway)
	ctx := context.Background()

	tests := []struct {
		name        string
		input       dto.CreateCustomerInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should create customer successfully",
			input: dto.CreateCustomerInput{
				Name:   "Test Customer",
				Email:  "test.customer.1@email.com",
				CPF:    "123.456.789-00",
				Writer: mockWriter,
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
			input: dto.CreateCustomerInput{
				Name:   "Test Customer",
				Email:  "test.customer.2@email.com",
				CPF:    "123.456.789-01",
				Writer: mockWriter,
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

			customer, err := useCase.Execute(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, customer)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, customer)
				assert.Equal(t, tt.input.Name, customer.Name)
				assert.Equal(t, tt.input.Email, customer.Email)
				assert.Equal(t, tt.input.CPF, customer.CPF)
			}
		})
	}
}
