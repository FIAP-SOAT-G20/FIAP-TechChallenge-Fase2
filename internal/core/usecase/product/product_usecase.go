package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type productUseCase struct {
	gateway port.ProductGateway
}

// NewProductUseCase creates a new StaffUseCase
func NewProductUseCase(gateway port.ProductGateway) port.ProductUseCase {
	return &productUseCase{
		gateway: gateway,
	}
}

// Execute lists all products
func (uc *productUseCase) List(ctx context.Context, input dto.ListProductsInput) ([]*entity.Product, int64, error) {
	products, total, err := uc.gateway.FindAll(ctx, input.Name, input.CategoryID, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return products, total, nil
}

func (uc productUseCase) Create(ctx context.Context, input dto.CreateProductInput) (*entity.Product, error) {
	product := entity.NewProduct(input.Name, input.Description, input.Price, input.CategoryID)

	if err := uc.gateway.Create(ctx, product); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return product, nil
}

func (uc productUseCase) Get(ctx context.Context, input dto.GetProductInput) (*entity.Product, error) {
	staff, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return staff, nil
}

func (uc productUseCase) Update(ctx context.Context, input dto.UpdateProductInput) (*entity.Product, error) {
	product, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if product == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	product.Update(input.Name, input.Description, input.Price, input.CategoryID)

	if err := uc.gateway.Update(ctx, product); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return product, nil
}

func (uc productUseCase) Delete(ctx context.Context, input dto.DeleteProductInput) (*entity.Product, error) {
	product, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if product == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return product, nil
}
