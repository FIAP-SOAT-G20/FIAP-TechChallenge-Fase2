package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

type AuthController interface {
	Authenticate(ctx context.Context, presenter Presenter, input dto.AuthenticateInput) ([]byte, error)
}
