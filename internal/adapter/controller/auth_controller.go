package controller

import (
	"context"
	"encoding/json"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type authController struct {
	useCase port.AuthUseCase
}

func NewAuthController(useCase port.AuthUseCase) port.AuthController {
	return &authController{useCase}
}

func (c *authController) Authenticate(ctx context.Context, p port.Presenter, i dto.AuthenticateInput) ([]byte, error) {
	authOutput, err := c.useCase.Authenticate(ctx, i)
	if err != nil {
		return nil, err
	}

	customerBytes, err := p.Present(dto.PresenterInput{Result: authOutput.Customer})
	if err != nil {
		return nil, err
	}

	var customerMap map[string]interface{}
	if err := json.Unmarshal(customerBytes, &customerMap); err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"customer":     customerMap,
		"access_token": authOutput.AccessToken,
		"token_type":   authOutput.TokenType,
		"expires_in":   authOutput.ExpiresIn,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
