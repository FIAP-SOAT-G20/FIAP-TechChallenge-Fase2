package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

type PaymentUseCase interface {
	Create(ctx context.Context, input dto.CreatePaymentInput) (*entity.Payment, error)
	Update(ctx context.Context, payment dto.UpdatePaymentInput) (*entity.Payment, error)
}
