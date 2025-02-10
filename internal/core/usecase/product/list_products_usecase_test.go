package product

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/domain/errors"
	mockport "tech-challenge-2-app-example/internal/core/port/mocks"
	"tech-challenge-2-app-example/internal/core/usecase"
)

func TestListProductsUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	useCase := NewListProductsUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	currentTime := time.Now()
	mockProducts := []*entity.Product{
		{
			ID:          1,
			Name:        "Test Product 1",
			Description: "Description 1",
			Price:       99.99,
			CategoryID:  1,
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
		},
		{
			ID:          2,
			Name:        "Test Product 2",
			Description: "Description 2",
			Price:       199.99,
			CategoryID:  1,
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
		},
	}

	mockOutput := &usecase.ListProductPaginatedOutput{
		PaginatedOutput: usecase.PaginatedOutput{
			Total: 2,
			Page:  1,
			Limit: 10,
		},
		Products: []usecase.ProductOutput{
			{
				ID:          1,
				Name:        "Test Product 1",
				Description: "Description 1",
				Price:       99.99,
				CategoryID:  1,
				CreatedAt:   currentTime.Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt:   currentTime.Format("2006-01-02T15:04:05Z07:00"),
			},
			{
				ID:          2,
				Name:        "Test Product 2",
				Description: "Description 2",
				Price:       199.99,
				CategoryID:  1,
				CreatedAt:   currentTime.Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt:   currentTime.Format("2006-01-02T15:04:05Z07:00"),
			},
		},
	}

	tests := []struct {
		name        string
		input       usecase.ListProductsInput
		setupMocks  func()
		expectError bool
		errorType   error
		checkOutput func(t *testing.T, output *usecase.ListProductPaginatedOutput)
	}{
		{
			name: "should list products successfully",
			input: usecase.ListProductsInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "", uint64(0), 1, 10).
					Return(mockProducts, int64(2), nil)

				mockPresenter.EXPECT().
					ToPaginatedOutput(mockProducts, int64(2), 1, 10).
					Return(mockOutput)
			},
			expectError: false,
			checkOutput: func(t *testing.T, output *usecase.ListProductPaginatedOutput) {
				assert.Equal(t, int64(2), output.Total)
				assert.Equal(t, 1, output.Page)
				assert.Equal(t, 10, output.Limit)
				assert.Len(t, output.Products, 2)

				assert.Equal(t, uint64(1), output.Products[0].ID)
				assert.Equal(t, "Test Product 1", output.Products[0].Name)
				assert.Equal(t, 99.99, output.Products[0].Price)

				assert.Equal(t, uint64(2), output.Products[1].ID)
				assert.Equal(t, "Test Product 2", output.Products[1].Name)
				assert.Equal(t, 199.99, output.Products[1].Price)
			},
		},
		{
			name: "should return error when page is invalid",
			input: usecase.ListProductsInput{
				Page:  0,
				Limit: 10,
			},
			setupMocks:  func() {},
			expectError: true,
			errorType:   &errors.InvalidInputError{},
		},
		{
			name: "should return error when limit is too high",
			input: usecase.ListProductsInput{
				Page:  1,
				Limit: 101,
			},
			setupMocks:  func() {},
			expectError: true,
			errorType:   &errors.InvalidInputError{},
		},
		{
			name: "should return error when repository fails",
			input: usecase.ListProductsInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "", uint64(0), 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			expectError: true,
			errorType:   &errors.InternalError{},
		},
		{
			name: "should filter by name",
			input: usecase.ListProductsInput{
				Name:  "Test",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "Test", uint64(0), 1, 10).
					Return(mockProducts, int64(2), nil)

				mockPresenter.EXPECT().
					ToPaginatedOutput(mockProducts, int64(2), 1, 10).
					Return(mockOutput)
			},
			expectError: false,
			checkOutput: func(t *testing.T, output *usecase.ListProductPaginatedOutput) {
				assert.Equal(t, int64(2), output.Total)
				assert.Len(t, output.Products, 2)
			},
		},
		{
			name: "should filter by category",
			input: usecase.ListProductsInput{
				CategoryID: 1,
				Page:       1,
				Limit:      10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "", uint64(1), 1, 10).
					Return(mockProducts, int64(2), nil)

				mockPresenter.EXPECT().
					ToPaginatedOutput(mockProducts, int64(2), 1, 10).
					Return(mockOutput)
			},
			expectError: false,
			checkOutput: func(t *testing.T, output *usecase.ListProductPaginatedOutput) {
				assert.Equal(t, int64(2), output.Total)
				assert.Len(t, output.Products, 2)
			},
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
				tt.checkOutput(t, output)
			}
		})
	}
}
