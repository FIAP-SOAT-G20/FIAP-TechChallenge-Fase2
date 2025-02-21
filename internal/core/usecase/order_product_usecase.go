package usecase

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type orderProductUseCase struct {
	gateway port.OrderProductGateway
}

// NewOrderProductUseCase creates a new ListOrderProductsUseCase
func NewOrderProductUseCase(gateway port.OrderProductGateway) port.OrderProductUseCase {
	return &orderProductUseCase{gateway}
}

// List lists all orderProducts
func (uc *orderProductUseCase) List(ctx context.Context, input dto.ListOrderProductsInput) ([]*entity.OrderProduct, int64, error) {
	orderProducts, total, err := uc.gateway.FindAll(ctx, input.OrderID, input.ProductID, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}
	return orderProducts, total, nil
}

// Create creates a new orderProduct
func (uc *orderProductUseCase) Create(ctx context.Context, input dto.CreateOrderProductInput) (*entity.OrderProduct, error) {
	orderProduct := entity.NewOrderProduct(input.OrderID, input.ProductID, input.Quantity)

	if err := uc.gateway.Create(ctx, orderProduct); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return orderProduct, nil
}

// Get returns a orderProduct by ID
func (uc *orderProductUseCase) Get(ctx context.Context, input dto.GetOrderProductInput) (*entity.OrderProduct, error) {
	orderProduct, err := uc.gateway.FindByID(ctx, input.OrderID, input.ProductID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if orderProduct == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return orderProduct, nil
}

func (uc *orderProductUseCase) Update(ctx context.Context, input dto.UpdateOrderProductInput) (*entity.OrderProduct, error) {
	orderProduct, err := uc.gateway.FindByID(ctx, input.OrderID, input.ProductID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if orderProduct == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	orderProduct.Update(input.Quantity)

	if err := uc.gateway.Update(ctx, orderProduct); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return orderProduct, nil
}

func (uc *orderProductUseCase) Delete(ctx context.Context, input dto.DeleteOrderProductInput) (*entity.OrderProduct, error) {
	order, err := uc.gateway.FindByID(ctx, input.OrderID, input.ProductID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if order == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.OrderID, input.ProductID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return order, nil
}
