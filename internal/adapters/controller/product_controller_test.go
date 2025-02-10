package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/dto"
	mockport "tech-challenge-2-app-example/internal/core/port/mocks"
)

func TestProductController_ListProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockport.NewMockListProductsUseCase(ctrl)
	controller := NewProductController(mockUseCase, nil, nil, nil, nil)
	ctx := context.Background()

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
					Execute(ctx, dto.ProductListRequest{Page: 1, Limit: 10}).
					Return(&dto.PaginatedResponse{
						Total: 1,
						Page:  1,
						Limit: 10,
						Products: []dto.ProductResponse{
							{ID: 1, Name: "Test Product"},
						},
					}, nil)
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
					Execute(ctx, dto.ProductListRequest{Page: 1, Limit: 10}).
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
				assert.Equal(t, 1, response.Page)
				assert.Equal(t, 10, response.Limit)
				assert.Len(t, response.Products, 1)
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
					Execute(ctx, gomock.Any()).
					Return(&dto.ProductResponse{
						ID:          1,
						Name:        "Test Product",
						Description: "Test Description",
						Price:       99.99,
						CategoryID:  1,
					}, nil)
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
				assert.Equal(t, tt.request.Name, response.Name)
				assert.Equal(t, tt.request.Description, response.Description)
				assert.Equal(t, tt.request.Price, response.Price)
				assert.Equal(t, tt.request.CategoryID, response.CategoryID)
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
					Return(&dto.ProductResponse{
						ID:          1,
						Name:        "Test Product",
						Description: "Test Description",
						Price:       99.99,
						CategoryID:  1,
					}, nil)
			},
			expectError: false,
		},
		{
			name: "should return error when product not found",
			id:   1,
			setupMocks: func() {
				mockUseCase.EXPECT().
					Execute(ctx, uint64(1)).
					Return(nil, errors.NewNotFoundError("Produto não encontrado"))
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
				assert.Equal(t, tt.id, response.ID)
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
					Execute(ctx, uint64(1), gomock.Any()).
					Return(&dto.ProductResponse{
						ID:          1,
						Name:        "Updated Product",
						Description: "Updated Description",
						Price:       199.99,
						CategoryID:  2,
					}, nil)
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
					Return(nil, errors.NewNotFoundError("Produto não encontrado"))
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
				assert.Equal(t, tt.request.Name, response.Name)
				assert.Equal(t, tt.request.Description, response.Description)
				assert.Equal(t, tt.request.Price, response.Price)
				assert.Equal(t, tt.request.CategoryID, response.CategoryID)
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
					Return(errors.NewNotFoundError("Produto não encontrado"))
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
