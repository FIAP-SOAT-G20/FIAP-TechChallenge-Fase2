package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type ProductController struct {
	useCase   port.ProductUseCase
	Presenter port.Presenter
}

func NewProductController(
	useCase port.ProductUseCase,
) *ProductController {
	return &ProductController{useCase, nil}
}

func (c *ProductController) List(ctx context.Context, input dto.ListProductsInput) error {
	products, total, err := c.useCase.List(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: products,
	})

	return nil
}

func (c *ProductController) Create(ctx context.Context, input dto.CreateProductInput) error {
	product, err := c.useCase.Create(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: product,
	})

	return nil
}

func (c *ProductController) Get(ctx context.Context, input dto.GetProductInput) error {
	product, err := c.useCase.Get(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: product,
	})

	return nil
}

func (c *ProductController) Update(ctx context.Context, input dto.UpdateProductInput) error {
	product, err := c.useCase.Update(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: product,
	})

	return nil
}

func (c *ProductController) Delete(ctx context.Context, input dto.DeleteProductInput) error {
	product, err := c.useCase.Delete(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: product,
	})

	return nil
}
