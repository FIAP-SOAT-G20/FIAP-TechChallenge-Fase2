package orderproduct

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type getOrderProductUseCase struct {
	gateway port.OrderProductGateway
}

// NewGetOrderProductUseCase creates a new GetOrderProductUseCase
func NewGetOrderProductUseCase(gateway port.OrderProductGateway) port.GetOrderProductUseCase {
	return &getOrderProductUseCase{gateway}
}

// Execute gets a orderProduct
func (uc *getOrderProductUseCase) Execute(ctx context.Context, input dto.GetOrderProductInput) (*entity.OrderProduct, error) {
	orderProduct, err := uc.gateway.FindByID(ctx, input.OrderID, input.ProductID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if orderProduct == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return orderProduct, nil
}
