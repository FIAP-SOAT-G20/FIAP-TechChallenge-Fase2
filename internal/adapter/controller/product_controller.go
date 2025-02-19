package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type ProductController struct {
	listUC    port.ListProductsUseCase
	createUC  port.CreateProductUseCase
	getUC     port.GetProductUseCase
	updateUC  port.UpdateProductUseCase
	deleteUC  port.DeleteProductUseCase
	Presenter port.ProductPresenter
}

func NewProductController(
	listUC port.ListProductsUseCase,
	createUC port.CreateProductUseCase,
	getUC port.GetProductUseCase,
	updateUC port.UpdateProductUseCase,
	deleteUC port.DeleteProductUseCase,
) *ProductController {
	return &ProductController{
		listUC,
		createUC,
		getUC,
		updateUC,
		deleteUC,
		nil,
	}
}

func (c *ProductController) ListProducts(ctx context.Context, input dto.ListProductsInput) error {
	products, total, err := c.listUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.ProductPresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: products,
	})

	return nil
}

func (c *ProductController) CreateProduct(ctx context.Context, input dto.CreateProductInput) error {
	product, err := c.createUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.ProductPresenterInput{
		Result: product,
	})

	return nil
}

func (c *ProductController) GetProduct(ctx context.Context, input dto.GetProductInput) error {
	product, err := c.getUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.ProductPresenterInput{
		Result: product,
	})

	return nil
}

func (c *ProductController) UpdateProduct(ctx context.Context, input dto.UpdateProductInput) error {
	product, err := c.updateUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.ProductPresenterInput{
		Result: product,
	})

	return nil
}

func (c *ProductController) DeleteProduct(ctx context.Context, input dto.DeleteProductInput) error {
	product, err := c.deleteUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.ProductPresenterInput{
		Result: product,
	})

	return nil
}
