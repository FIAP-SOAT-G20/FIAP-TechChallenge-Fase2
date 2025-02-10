package product

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"tech-challenge-2-app-example/internal/core/domain/errors"
	mockport "tech-challenge-2-app-example/internal/core/port/mocks"
	"tech-challenge-2-app-example/internal/core/usecase"
)

func TestCreateProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	useCase := NewCreateProductUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	currentTime := time.Now()
	mockOutput := &usecase.ProductOutput{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
		CreatedAt:   currentTime.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   currentTime.Format("2006-01-02T15:04:05Z07:00"),
	}

	tests := []struct {
		name        string
		input       usecase.CreateProductInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should create product successfully",
			input: usecase.CreateProductInput{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)

				mockPresenter.EXPECT().
					ToOutput(gomock.Any()).
					Return(mockOutput)
			},
			expectError: false,
		},
		{
			name: "should return error when validation fails",
			input: usecase.CreateProductInput{
				Name:        "", // Nome vazio deve falhar na validação
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
			setupMocks:  func() {},
			expectError: true,
			errorType:   &errors.ValidationError{},
		},
		{
			name: "should return error when gateway fails",
			input: usecase.CreateProductInput{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					Create(ctx, gomock.Any()).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &errors.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			output, err := useCase.Execute(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, mockOutput, output)
			}
		})
	}
}
