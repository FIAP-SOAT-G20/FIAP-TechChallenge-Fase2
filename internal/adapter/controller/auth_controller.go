package controller

import (
	"context"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type authController struct {
	customerUseCase port.CustomerUseCase
	jwtService      port.JWTService
}

func NewAuthController(customerUseCase port.CustomerUseCase, jwtService port.JWTService) port.AuthController {
	return &authController{
		customerUseCase: customerUseCase,
		jwtService:      jwtService,
	}
}

func (c *authController) Authenticate(ctx context.Context, presenter port.Presenter, input dto.AuthenticateInput) ([]byte, error) {
	customer, err := c.customerUseCase.FindByCPF(ctx, dto.FindCustomerByCPFInput(input))
	if err != nil {
		return nil, err
	}

	token, err := c.jwtService.GenerateToken(
		customer.ID,
		customer.CPF,
		customer.Name,
		24*time.Hour,
	)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	return presenter.Present(dto.PresenterInput{
		Result: map[string]interface{}{
			"customer":     customer,
			"access_token": token,
			"token_type":   "Bearer",
			"expires_in":   int64((24 * time.Hour).Seconds()),
		},
	})
}
