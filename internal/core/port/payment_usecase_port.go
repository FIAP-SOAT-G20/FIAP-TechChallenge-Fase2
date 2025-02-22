package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type CreatePaymentUseCase interface {
	Execute(ctx context.Context, OrderID uint64) (*entity.Payment, error)
}
