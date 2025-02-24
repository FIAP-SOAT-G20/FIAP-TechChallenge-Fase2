package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
)

type PaymentDataSource interface {
	Create(context context.Context, payment *entity.Payment) (*entity.Payment, error)
	GetByOrderID(context context.Context, orderID uint64) (*entity.Payment, error)
	UpdateStatus(status valueobject.PaymentStatus, externalPaymentID string) error
	GetByExternalPaymentID(externalPaymentID string) (*entity.Payment, error)
}
