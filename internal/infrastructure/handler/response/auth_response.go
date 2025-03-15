package response

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"

type AuthenticationResponse struct {
	Customer    presenter.CustomerJsonResponse `json:"customer"`
	AccessToken string                         `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	TokenType   string                         `json:"token_type" example:"Bearer"`
	ExpiresIn   int64                          `json:"expires_in" example:"86400"`
}
