package controller

import (
	"context"

	"tech-challenge-2-app-example/internal/adapters/dto"
	"tech-challenge-2-app-example/internal/core/port"
	"tech-challenge-2-app-example/internal/core/usecase"
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
	input := usecase.ListProductsInput{
		Name:       req.Name,
		CategoryID: req.CategoryID,
		Page:       req.Page,
		Limit:      req.Limit,
	}

	output, err := c.listProductsUseCase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	products := make([]dto.ProductResponse, len(output.Products))
	for i, p := range output.Products {
		products[i] = dto.ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			CategoryID:  p.CategoryID,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		}
	}

	return &dto.PaginatedResponse{
		Pagination: dto.Pagination{
			Total: output.Total,
			Page:  output.Page,
			Limit: output.Limit,
		},
		Products: products,
	}, nil
}

func (c *ProductController) CreateProduct(ctx context.Context, req dto.ProductRequest) (*dto.ProductResponse, error) {
	input := usecase.CreateProductInput{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	output, err := c.createProductUseCase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:          output.ID,
		Name:        output.Name,
		Description: output.Description,
		Price:       output.Price,
		CategoryID:  output.CategoryID,
		CreatedAt:   output.CreatedAt,
		UpdatedAt:   output.UpdatedAt,
	}, nil
}

func (c *ProductController) GetProduct(ctx context.Context, id uint64) (*dto.ProductResponse, error) {
	output, err := c.getProductUseCase.Execute(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:          output.ID,
		Name:        output.Name,
		Description: output.Description,
		Price:       output.Price,
		CategoryID:  output.CategoryID,
		CreatedAt:   output.CreatedAt,
		UpdatedAt:   output.UpdatedAt,
	}, nil
}

func (c *ProductController) UpdateProduct(ctx context.Context, id uint64, req dto.ProductRequest) (*dto.ProductResponse, error) {
	input := usecase.UpdateProductInput{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	output, err := c.updateProductUseCase.Execute(ctx, id, input)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:          output.ID,
		Name:        output.Name,
		Description: output.Description,
		Price:       output.Price,
		CategoryID:  output.CategoryID,
		CreatedAt:   output.CreatedAt,
		UpdatedAt:   output.UpdatedAt,
	}, nil
}

func (c *ProductController) DeleteProduct(ctx context.Context, id uint64) error {
	return c.deleteProductUseCase.Execute(ctx, id)
}
