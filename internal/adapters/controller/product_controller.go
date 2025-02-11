package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
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

func (c *ProductController) ListProducts(ctx context.Context, rw dto.ResponseWriter, req dto.ProductListRequest) error {
	input := dto.ListProductsInput{
		Name:       req.Name,
		CategoryID: req.CategoryID,
		Page:       req.Page,
		Limit:      req.Limit,
		Writer:     rw,
	}

	err := c.listProductsUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductController) CreateProduct(ctx context.Context, rw dto.ResponseWriter, req dto.ProductRequest) error {
	input := dto.CreateProductInput{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
		Writer:      rw,
	}

	err := c.createProductUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductController) GetProduct(ctx context.Context, rw dto.ResponseWriter, id uint64) error {
	input := dto.GetProductInput{
		ID:     id,
		Writer: rw,
	}
	err := c.getProductUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductController) UpdateProduct(ctx context.Context, rw dto.ResponseWriter, id uint64, req dto.ProductRequest) error {
	input := dto.UpdateProductInput{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
		Writer:      rw,
	}

	err := c.updateProductUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductController) DeleteProduct(ctx context.Context, rw dto.ResponseWriter, id uint64) error {
	input := dto.DeleteProductInput{
		ID:     id,
		Writer: rw,
	}
	return c.deleteProductUseCase.Execute(ctx, input)
}
