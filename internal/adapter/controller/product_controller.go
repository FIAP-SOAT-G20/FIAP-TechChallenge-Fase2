package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
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

func (c *ProductController) ListProducts(ctx context.Context, input dto.ListProductsInput) error {
	err := c.listProductsUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductController) CreateProduct(ctx context.Context, input dto.CreateProductInput) error {
	err := c.createProductUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductController) GetProduct(ctx context.Context, input dto.GetProductInput) error {
	err := c.getProductUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductController) UpdateProduct(ctx context.Context, input dto.UpdateProductInput) error {
	err := c.updateProductUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductController) DeleteProduct(ctx context.Context, input dto.DeleteProductInput) error {
	return c.deleteProductUseCase.Execute(ctx, input)
}
