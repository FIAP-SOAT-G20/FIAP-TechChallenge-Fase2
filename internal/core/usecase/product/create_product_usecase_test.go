package product_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	mockdto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/product"
)

func TestCreateProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)
	useCase := product.NewCreateProductUseCase(mockGateway)
	ctx := context.Background()

	tests := []struct {
		name        string
		input       dto.CreateProductInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should create product successfully",
			input: dto.CreateProductInput{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
				Writer:      mockWriter,
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
			input: dto.CreateProductInput{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
				Writer:      mockWriter,
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

			product, err := useCase.Execute(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, product)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, product)
				assert.Equal(t, tt.input.Name, product.Name)
				assert.Equal(t, tt.input.Description, product.Description)
				assert.Equal(t, tt.input.Price, product.Price)
				assert.Equal(t, tt.input.CategoryID, product.CategoryID)
			}
		})
	}
}
