package controller

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

// TODO: Add more test cenarios
func TestProductController_ListProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductsUseCase := mockport.NewMockProductUseCase(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	productController := NewProductController(mockProductsUseCase)
	productController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.ListProductsInput{
		Name:       "Test",
		CategoryID: 1,
		Page:       1,
		Limit:      10,
	}

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

	mockProductsUseCase.EXPECT().
		List(ctx, input).
		Return(mockProducts, int64(2), nil)

	mockPresenter.EXPECT().
		Present(dto.ProductPresenterInput{
			Result: mockProducts,
			Total:  int64(2),
			Page:   1,
			Limit:  10,
		})

	//gomock.Any()//

	err := productController.ListProducts(ctx, input)
	assert.NoError(t, err)
}

func TestProductController_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := mockport.NewMockProductUseCase(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	productController := NewProductController(mockProductUseCase)
	productController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.CreateProductInput{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
	}

	mockProduct := &entity.Product{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
	}

	mockProductUseCase.EXPECT().
		Create(ctx, input).
		Return(mockProduct, nil)

	mockPresenter.EXPECT().
		Present(dto.ProductPresenterInput{
			Result: mockProduct,
		})

	err := productController.CreateProduct(ctx, input)
	assert.NoError(t, err)
}

func TestProductController_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := mockport.NewMockProductUseCase(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	productController := NewProductController(mockProductUseCase)
	productController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.GetProductInput{
		ID: uint64(1),
	}

	mockProduct := &entity.Product{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
	}

	mockProductUseCase.EXPECT().
		Get(ctx, input).
		Return(mockProduct, nil)

	mockPresenter.EXPECT().
		Present(dto.ProductPresenterInput{
			Result: &entity.Product{
				ID:          1,
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
		})

	err := productController.GetProduct(ctx, input)
	assert.NoError(t, err)
}

func TestProductController_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := mockport.NewMockProductUseCase(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	productController := NewProductController(mockProductUseCase)
	productController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.UpdateProductInput{
		ID:          uint64(1),
		Name:        "Product",
		Description: "Description",
		Price:       99.99,
		CategoryID:  2,
	}

	mockProduct := &entity.Product{
		ID:          1,
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       199.99,
		CategoryID:  2,
	}

	mockProductUseCase.EXPECT().
		Update(ctx, input).
		Return(mockProduct, nil)

	mockPresenter.EXPECT().
		Present(dto.ProductPresenterInput{
			Result: &entity.Product{
				ID:          1,
				Name:        "Updated Product",
				Description: "Updated Description",
				Price:       199.99,
				CategoryID:  2,
			},
		})

	err := productController.UpdateProduct(ctx, input)
	assert.NoError(t, err)
}

func TestProductController_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := mockport.NewMockProductUseCase(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	productController := NewProductController(mockProductUseCase)
	productController.Presenter = mockPresenter

	ctx := context.Background()
	input := dto.DeleteProductInput{
		ID: uint64(1),
	}

	mockProduct := &entity.Product{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
	}

	mockProductUseCase.EXPECT().
		Delete(ctx, input).
		Return(mockProduct, nil)

	mockPresenter.EXPECT().
		Present(dto.ProductPresenterInput{
			Result: mockProduct,
		})

	err := productController.DeleteProduct(ctx, input)
	assert.NoError(t, err)
}
