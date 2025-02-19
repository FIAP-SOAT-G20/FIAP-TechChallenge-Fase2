package orderproduct

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type listOrderProductsUseCase struct {
	gateway port.OrderProductGateway
}

// NewListOrderProductsUseCase creates a new ListOrderProductsUseCase
func NewListOrderProductsUseCase(gateway port.OrderProductGateway) port.ListOrderProductsUseCase {
	return &listOrderProductsUseCase{gateway}
}

// Execute lists all orderProducts
func (uc *listOrderProductsUseCase) Execute(ctx context.Context, input dto.ListOrderProductsInput) ([]*entity.OrderProduct, int64, error) {
	orderProducts, total, err := uc.gateway.FindAll(ctx, input.OrderID, input.ProductID, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}
	return orderProducts, total, nil
}
