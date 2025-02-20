package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type ProductController struct {
	useCase   port.ProductUseCase
	Presenter port.ProductPresenter
}

func NewProductController(
	useCase port.ProductUseCase,
) *ProductController {
	return &ProductController{useCase, nil}
}

func (c *ProductController) ListProducts(ctx context.Context, input dto.ListProductsInput) error {
	products, total, err := c.useCase.List(ctx, input)
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
	product, err := c.useCase.Create(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.ProductPresenterInput{
		Result: product,
	})

	return nil
}

func (c *ProductController) GetProduct(ctx context.Context, input dto.GetProductInput) error {
	product, err := c.useCase.Get(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.ProductPresenterInput{
		Result: product,
	})

	return nil
}

func (c *ProductController) UpdateProduct(ctx context.Context, input dto.UpdateProductInput) error {
	product, err := c.useCase.Update(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.ProductPresenterInput{
		Result: product,
	})

	return nil
}

func (c *ProductController) DeleteProduct(ctx context.Context, input dto.DeleteProductInput) error {
	product, err := c.useCase.Delete(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.ProductPresenterInput{
		Result: product,
	})

	return nil
}
