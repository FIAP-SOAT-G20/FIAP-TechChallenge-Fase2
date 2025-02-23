package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type PaymentExternalDatasource interface {
	Create(context context.Context, payment *entity.CreatePaymentExternalInput) (*entity.CreatePaymentExternalOutput, error)
}
