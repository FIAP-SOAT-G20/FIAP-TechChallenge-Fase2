package orderproduct

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type updateOrderProductUseCase struct {
	gateway port.OrderProductGateway
}

// NewUpdateOrderProductUseCase creates a new UpdateOrderProductUseCase
func NewUpdateOrderProductUseCase(gateway port.OrderProductGateway) port.UpdateOrderProductUseCase {
	return &updateOrderProductUseCase{gateway}
}

// Execute updates a orderProduct
func (uc *updateOrderProductUseCase) Execute(ctx context.Context, input dto.UpdateOrderProductInput) (*entity.OrderProduct, error) {
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
