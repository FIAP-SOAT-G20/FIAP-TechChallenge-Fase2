package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
)

type PaymentGateway interface {
	Create(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
	GetPaymentByOrderIDAndStatus(ctx context.Context, status valueobject.PaymentStatus, orderID uint64) (*entity.Payment, error)
}
