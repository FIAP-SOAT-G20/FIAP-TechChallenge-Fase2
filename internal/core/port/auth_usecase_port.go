package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

type AuthUseCase interface {
	Authenticate(ctx context.Context, input dto.AuthenticateInput) (*dto.AuthenticateOutput, error)
}
