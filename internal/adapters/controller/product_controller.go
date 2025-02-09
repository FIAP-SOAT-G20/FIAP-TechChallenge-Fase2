package controller

import (
	"context"

	"tech-challenge-2-app-example/internal/core/dto"
	"tech-challenge-2-app-example/internal/core/port"
)

type ProductController struct {
	listProductsUseCase  port.ListProductsUseCase
	createProductUseCase port.CreateProductUseCase
	getProductUseCase    port.GetProductUseCase
	updateProductUseCase port.UpdateProductUseCase
	deleteProductUseCase port.DeleteProductUseCase
}

func NewProductController(
	listUC port.ListProductsUseCase,
	createUC port.CreateProductUseCase,
	getUC port.GetProductUseCase,
	updateUC port.UpdateProductUseCase,
	deleteUC port.DeleteProductUseCase,
) *ProductController {
	return &ProductController{
		listProductsUseCase:  listUC,
		createProductUseCase: createUC,
		getProductUseCase:    getUC,
		updateProductUseCase: updateUC,
		deleteProductUseCase: deleteUC,
	}
}

func (c *ProductController) ListProducts(ctx context.Context, req dto.ProductListRequest) (*dto.PaginatedResponse, error) {
	return c.listProductsUseCase.Execute(ctx, req)
}

func (c *ProductController) CreateProduct(ctx context.Context, req dto.ProductRequest) (*dto.ProductResponse, error) {
	return c.createProductUseCase.Execute(ctx, req)
}

func (c *ProductController) GetProduct(ctx context.Context, id uint64) (*dto.ProductResponse, error) {
	return c.getProductUseCase.Execute(ctx, id)
}

func (c *ProductController) UpdateProduct(ctx context.Context, id uint64, req dto.ProductRequest) (*dto.ProductResponse, error) {
	return c.updateProductUseCase.Execute(ctx, id, req)
}

func (c *ProductController) DeleteProduct(ctx context.Context, id uint64) error {
	return c.deleteProductUseCase.Execute(ctx, id)
}
