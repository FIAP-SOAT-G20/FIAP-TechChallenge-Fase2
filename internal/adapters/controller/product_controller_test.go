package controller

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"tech-challenge-2-app-example/internal/adapters/dto"
	"tech-challenge-2-app-example/internal/core/domain/errors"
	mockport "tech-challenge-2-app-example/internal/core/port/mocks"
	"tech-challenge-2-app-example/internal/core/usecase"
)

func TestProductController_ListProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockport.NewMockListProductsUseCase(ctrl)
	controller := NewProductController(mockUseCase, nil, nil, nil, nil)
	ctx := context.Background()

	currentTime := time.Now().Format("2006-01-02T15:04:05Z07:00")
	mockOutput := &usecase.ListProductPaginatedOutput{
		PaginatedOutput: usecase.PaginatedOutput{
			Total: 1,
			Page:  1,
			Limit: 10,
		},
		Products: []usecase.ProductOutput{
			{
				ID:          1,
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
				CreatedAt:   currentTime,
				UpdatedAt:   currentTime,
			},
		},
	}

	tests := []struct {
		name        string
		request     dto.ProductListRequest
		setupMocks  func()
		expectError bool
	}{
		{
			name: "should list products successfully",
			request: dto.ProductListRequest{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, usecase.ListProductsInput{Page: 1, Limit: 10}).
					Return(mockOutput, nil)
			},
			expectError: false,
		},
		{
			name: "should return error when use case fails",
			request: dto.ProductListRequest{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, usecase.ListProductsInput{Page: 1, Limit: 10}).
					Return(nil, errors.NewInternalError(assert.AnError))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			response, err := controller.ListProducts(ctx, tt.request)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, mockOutput.Page, response.Page)
				assert.Equal(t, mockOutput.Limit, response.Limit)
				assert.Equal(t, mockOutput.Total, response.Total)
				assert.Len(t, response.Products, len(mockOutput.Products))
				assert.Equal(t, mockOutput.Products[0].Name, response.Products[0].Name)
			}
		})
	}
}

func TestProductController_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockport.NewMockCreateProductUseCase(ctrl)
	controller := NewProductController(nil, mockUseCase, nil, nil, nil)
	ctx := context.Background()

	currentTime := time.Now().Format("2006-01-02T15:04:05Z07:00")
	mockOutput := &usecase.ProductOutput{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	tests := []struct {
		name        string
		request     dto.ProductRequest
		setupMocks  func()
		expectError bool
	}{
		{
			name: "should create product successfully",
			request: dto.ProductRequest{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, usecase.CreateProductInput{
						Name:        "Test Product",
						Description: "Test Description",
						Price:       99.99,
						CategoryID:  1,
					}).
					Return(mockOutput, nil)
			},
			expectError: false,
		},
		{
			name: "should return error when use case fails",
			request: dto.ProductRequest{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, gomock.Any()).
					Return(nil, errors.NewInternalError(assert.AnError))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			response, err := controller.CreateProduct(ctx, tt.request)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, mockOutput.ID, response.ID)
				assert.Equal(t, mockOutput.Name, response.Name)
				assert.Equal(t, mockOutput.Description, response.Description)
				assert.Equal(t, mockOutput.Price, response.Price)
				assert.Equal(t, mockOutput.CategoryID, response.CategoryID)
				assert.Equal(t, mockOutput.CreatedAt, response.CreatedAt)
				assert.Equal(t, mockOutput.UpdatedAt, response.UpdatedAt)
			}
		})
	}
}

func TestProductController_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockport.NewMockGetProductUseCase(ctrl)
	controller := NewProductController(nil, nil, mockUseCase, nil, nil)
	ctx := context.Background()

	currentTime := time.Now().Format("2006-01-02T15:04:05Z07:00")
	mockOutput := &usecase.ProductOutput{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
	}{
		{
			name: "should get product successfully",
			id:   1,
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, uint64(1)).
					Return(mockOutput, nil)
			},
			expectError: false,
		},
		{
			name: "should return error when product not found",
			id:   1,
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, uint64(1)).
					Return(nil, errors.NewNotFoundError("produto não encontrado"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			response, err := controller.GetProduct(ctx, tt.id)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, mockOutput.ID, response.ID)
				assert.Equal(t, mockOutput.Name, response.Name)
				assert.Equal(t, mockOutput.Description, response.Description)
				assert.Equal(t, mockOutput.Price, response.Price)
				assert.Equal(t, mockOutput.CategoryID, response.CategoryID)
				assert.Equal(t, mockOutput.CreatedAt, response.CreatedAt)
				assert.Equal(t, mockOutput.UpdatedAt, response.UpdatedAt)
			}
		})
	}
}

func TestProductController_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockport.NewMockUpdateProductUseCase(ctrl)
	controller := NewProductController(nil, nil, nil, mockUseCase, nil)
	ctx := context.Background()

	currentTime := time.Now().Format("2006-01-02T15:04:05Z07:00")
	mockOutput := &usecase.ProductOutput{
		ID:          1,
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       199.99,
		CategoryID:  2,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	tests := []struct {
		name        string
		id          uint64
		request     dto.ProductRequest
		setupMocks  func()
		expectError bool
	}{
		{
			name: "should update product successfully",
			id:   1,
			request: dto.ProductRequest{
				Name:        "Updated Product",
				Description: "Updated Description",
				Price:       199.99,
				CategoryID:  2,
			},
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, uint64(1), usecase.UpdateProductInput{
						Name:        "Updated Product",
						Description: "Updated Description",
						Price:       199.99,
						CategoryID:  2,
					}).
					Return(mockOutput, nil)
			},
			expectError: false,
		},
		{
			name: "should return error when use case fails",
			id:   1,
			request: dto.ProductRequest{
				Name:        "Updated Product",
				Description: "Updated Description",
				Price:       199.99,
				CategoryID:  2,
			},
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, uint64(1), gomock.Any()).
					Return(nil, errors.NewNotFoundError("produto não encontrado"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			response, err := controller.UpdateProduct(ctx, tt.id, tt.request)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, mockOutput.ID, response.ID)
				assert.Equal(t, mockOutput.Name, response.Name)
				assert.Equal(t, mockOutput.Description, response.Description)
				assert.Equal(t, mockOutput.Price, response.Price)
				assert.Equal(t, mockOutput.CategoryID, response.CategoryID)
				assert.Equal(t, mockOutput.CreatedAt, response.CreatedAt)
				assert.Equal(t, mockOutput.UpdatedAt, response.UpdatedAt)
			}
		})
	}
}

func TestProductController_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockport.NewMockDeleteProductUseCase(ctrl)
	controller := NewProductController(nil, nil, nil, nil, mockUseCase)
	ctx := context.Background()

	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
	}{
		{
			name: "should delete product successfully",
			id:   1,
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, uint64(1)).
					Return(nil)
			},
			expectError: false,
		},
		{
			name: "should return error when product not found",
			id:   1,
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, uint64(1)).
					Return(errors.NewNotFoundError("produto não encontrado"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := controller.DeleteProduct(ctx, tt.id)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
