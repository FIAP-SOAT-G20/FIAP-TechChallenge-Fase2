package orderproduct

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type createOrderProductUseCase struct {
	gateway port.OrderProductGateway
}

// NewCreateOrderProductUseCase creates a new CreateOrderProductUseCase
func NewCreateOrderProductUseCase(gateway port.OrderProductGateway) port.CreateOrderProductUseCase {
	return &createOrderProductUseCase{gateway}
}

// Execute creates a new OrderProduct
func (uc *createOrderProductUseCase) Execute(ctx context.Context, input dto.CreateOrderProductInput) (*entity.OrderProduct, error) {
	orderProduct := entity.NewOrderProduct(input.OrderID, input.ProductID, input.Quantity)

	if err := uc.gateway.Create(ctx, orderProduct); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return orderProduct, nil
}
