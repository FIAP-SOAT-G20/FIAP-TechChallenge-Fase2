package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

// AuthUseCase defines the authentication use case interface
type AuthUseCase interface {
	// Authenticate authenticates a customer by CPF and return the token
	Authenticate(ctx context.Context, input dto.AuthenticateInput) (string, error)
}
