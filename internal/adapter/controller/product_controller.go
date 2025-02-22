package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type ProductController struct {
	useCase port.ProductUseCase
}

func NewProductController(
	useCase port.ProductUseCase,
) *ProductController {
	return &ProductController{useCase}
}

func (c *ProductController) List(ctx context.Context, p port.Presenter, i dto.ListProductsInput) error {
	products, total, err := c.useCase.List(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{
		Total:  total,
		Page:   i.Page,
		Limit:  i.Limit,
		Result: products,
	})

	return nil
}

func (c *ProductController) Create(ctx context.Context, p port.Presenter, i dto.CreateProductInput) error {
	product, err := c.useCase.Create(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: product})

	return nil
}

func (c *ProductController) Get(ctx context.Context, p port.Presenter, i dto.GetProductInput) error {
	product, err := c.useCase.Get(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: product})

	return nil
}

func (c *ProductController) Update(ctx context.Context, p port.Presenter, i dto.UpdateProductInput) error {
	product, err := c.useCase.Update(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: product})

	return nil
}

func (c *ProductController) Delete(ctx context.Context, p port.Presenter, i dto.DeleteProductInput) error {
	product, err := c.useCase.Delete(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: product})

	return nil
}
