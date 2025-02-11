package product

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	mockdto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

func TestCreateProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)
	useCase := NewCreateProductUseCase(mockGateway, mockPresenter)
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

				mockPresenter.EXPECT().
					Present(gomock.Any())
			},
			expectError: false,
		},
		{
			name: "should return error when validation fails",
			input: dto.CreateProductInput{
				Name:        "", // Nome vazio deve falhar na validação
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
				Writer:      mockWriter,
			},
			setupMocks:  func() {},
			expectError: true,
			errorType:   &domain.ValidationError{},
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
