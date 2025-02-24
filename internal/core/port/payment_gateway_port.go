package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type PaymentGateway interface {
	Create(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
	GetByOrderID(ctx context.Context, orderID uint64) (*entity.Payment, error)
}
