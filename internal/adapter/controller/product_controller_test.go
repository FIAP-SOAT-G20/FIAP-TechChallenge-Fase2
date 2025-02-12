package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	mockdto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto/mocks"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

func TestProductController_ListProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListProductsUseCase := mockport.NewMockListProductsUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewProductController(mockListProductsUseCase, nil, nil, nil, nil)

	ctx := context.Background()
	input := dto.ListProductsInput{
		Name:       "Test",
		CategoryID: 1,
		Page:       1,
		Limit:      10,
		Writer:     mockResponseWriter,
	}

	mockListProductsUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.ListProducts(ctx, input)
	assert.NoError(t, err)
}

func TestProductController_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateProductUseCase := mockport.NewMockCreateProductUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewProductController(nil, mockCreateProductUseCase, nil, nil, nil)

	ctx := context.Background()
	input := dto.CreateProductInput{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
		Writer:      mockResponseWriter,
	}

	mockCreateProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.CreateProduct(ctx, input)
	assert.NoError(t, err)
}

func TestProductController_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetProductUseCase := mockport.NewMockGetProductUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewProductController(nil, nil, mockGetProductUseCase, nil, nil)

	ctx := context.Background()
	input := dto.GetProductInput{
		ID:     uint64(1),
		Writer: mockResponseWriter,
	}

	mockGetProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.GetProduct(ctx, input)
	assert.NoError(t, err)
}

func TestProductController_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUpdateProductUseCase := mockport.NewMockUpdateProductUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewProductController(nil, nil, nil, mockUpdateProductUseCase, nil)

	ctx := context.Background()
	input := dto.UpdateProductInput{
		ID:          uint64(1),
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       199.99,
		CategoryID:  2,
		Writer:      mockResponseWriter,
	}

	mockUpdateProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.UpdateProduct(ctx, input)
	assert.NoError(t, err)
}

func TestProductController_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeleteProductUseCase := mockport.NewMockDeleteProductUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewProductController(nil, nil, nil, nil, mockDeleteProductUseCase)

	ctx := context.Background()
	input := dto.DeleteProductInput{
		ID:     uint64(1),
		Writer: mockResponseWriter,
	}

	mockDeleteProductUseCase.EXPECT().
		Execute(ctx, input).
		Return(nil)

	err := productController.DeleteProduct(ctx, input)
	assert.NoError(t, err)
}
