package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
)

type CreatePaymentUseCase interface {
	Execute(ctx context.Context, OrderID uint64, writer dto.ResponseWriter) error
}
