package presenter

import (
	"encoding/json"
	"errors"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type authPresenter struct{}

func NewAuthPresenter() port.Presenter {
	return &authPresenter{}
}

func (p *authPresenter) Present(input dto.PresenterInput) ([]byte, error) {
	data, ok := input.Result.(map[string]interface{})
	if !ok {
		return nil, errors.New("unexpected input data type")
	}

	customer, ok := data["customer"].(*entity.Customer)
	if !ok {
		return nil, errors.New("unexpected customer data type")
	}

	accessToken, ok := data["access_token"].(string)
	if !ok {
		return nil, errors.New("access_token must be a string")
	}

	tokenType, ok := data["token_type"].(string)
	if !ok {
		return nil, errors.New("token_type must be a string")
	}

	expiresIn, ok := data["expires_in"].(int64)
	if !ok {
		return nil, errors.New("expires_in must be an int64")
	}

	response := AuthenticationResponse{
		Customer:    ToCustomerJsonResponse(customer),
		AccessToken: accessToken,
		TokenType:   tokenType,
		ExpiresIn:   expiresIn,
	}

	return json.Marshal(response)
}
