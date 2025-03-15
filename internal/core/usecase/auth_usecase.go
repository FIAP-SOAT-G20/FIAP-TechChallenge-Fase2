package usecase

import (
	"context"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type authUseCase struct {
	customerGateway port.CustomerGateway
	jwtService      port.JWTService
	jwtExpiration   time.Duration
}

func NewAuthUseCase(customerGateway port.CustomerGateway, jwtService port.JWTService, jwtExpiration time.Duration) port.AuthUseCase {
	return &authUseCase{
		customerGateway: customerGateway,
		jwtService:      jwtService,
		jwtExpiration:   jwtExpiration,
	}
}

func (uc *authUseCase) Authenticate(ctx context.Context, input dto.AuthenticateInput) (*dto.AuthenticateOutput, error) {
	customer, err := uc.customerGateway.FindByCPF(ctx, input.CPF)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if customer == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	claims := port.JWTClaims{
		CustomerID: customer.ID,
		CPF:        customer.CPF,
		Name:       customer.Name,
	}

	token, err := uc.jwtService.GenerateToken(claims, uc.jwtExpiration)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	return &dto.AuthenticateOutput{
		Customer:    customer,
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int64(uc.jwtExpiration.Seconds()),
	}, nil
}
