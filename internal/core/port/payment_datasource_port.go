package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type PaymentDataSource interface {
	Create(context context.Context, payment *entity.Payment) (*entity.Payment, error)
	GetPaymentByOrderIDAndStatus(context context.Context, status entity.PaymentStatus, orderID uint64) (*entity.Payment, error)
	UpdateStatus(status entity.PaymentStatus, externalPaymentID string) error
	GetByExternalPaymentID(externalPaymentID string) (*entity.Payment, error)
}
