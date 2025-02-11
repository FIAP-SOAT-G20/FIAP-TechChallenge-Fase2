package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	mockdto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto/mocks"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

func TestProductController_ListProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListProductsUseCase := mockport.NewMockListProductsUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewProductController(mockListProductsUseCase, nil, nil, nil, nil)

	ctx := context.Background()
	req := dto.ProductListRequest{
		Name:       "Test",
		CategoryID: 1,
		Page:       1,
		Limit:      10,
	}

	mockListProductsUseCase.EXPECT().
		Execute(ctx, dto.ListProductsInput{
			Name:       req.Name,
			CategoryID: req.CategoryID,
			Page:       req.Page,
			Limit:      req.Limit,
			Writer:     mockResponseWriter,
		}).
		Return(nil)

	err := productController.ListProducts(ctx, mockResponseWriter, req)
	assert.NoError(t, err)
}

func TestProductController_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateProductUseCase := mockport.NewMockCreateProductUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewProductController(nil, mockCreateProductUseCase, nil, nil, nil)

	ctx := context.Background()
	req := dto.ProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
	}

	mockCreateProductUseCase.EXPECT().
		Execute(ctx, dto.CreateProductInput{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			CategoryID:  req.CategoryID,
			Writer:      mockResponseWriter,
		}).
		Return(nil)

	err := productController.CreateProduct(ctx, mockResponseWriter, req)
	assert.NoError(t, err)
}

func TestProductController_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetProductUseCase := mockport.NewMockGetProductUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewProductController(nil, nil, mockGetProductUseCase, nil, nil)

	ctx := context.Background()
	id := uint64(1)

	mockGetProductUseCase.EXPECT().
		Execute(ctx, dto.GetProductInput{
			ID:     id,
			Writer: mockResponseWriter,
		}).
		Return(nil)

	err := productController.GetProduct(ctx, mockResponseWriter, id)
	assert.NoError(t, err)
}

func TestProductController_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUpdateProductUseCase := mockport.NewMockUpdateProductUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewProductController(nil, nil, nil, mockUpdateProductUseCase, nil)

	ctx := context.Background()
	id := uint64(1)
	req := dto.ProductRequest{
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       199.99,
		CategoryID:  2,
	}

	mockUpdateProductUseCase.EXPECT().
		Execute(ctx, dto.UpdateProductInput{
			ID:          id,
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			CategoryID:  req.CategoryID,
			Writer:      mockResponseWriter,
		}).
		Return(nil)

	err := productController.UpdateProduct(ctx, mockResponseWriter, id, req)
	assert.NoError(t, err)
}

func TestProductController_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeleteProductUseCase := mockport.NewMockDeleteProductUseCase(ctrl)
	mockResponseWriter := mockdto.NewMockResponseWriter(ctrl)
	productController := NewProductController(nil, nil, nil, nil, mockDeleteProductUseCase)

	ctx := context.Background()
	id := uint64(1)

	mockDeleteProductUseCase.EXPECT().
		Execute(ctx, dto.DeleteProductInput{
			ID:     id,
			Writer: mockResponseWriter,
		}).
		Return(nil)

	err := productController.DeleteProduct(ctx, mockResponseWriter, id)
	assert.NoError(t, err)
}
