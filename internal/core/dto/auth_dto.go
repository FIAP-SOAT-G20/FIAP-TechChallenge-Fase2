package dto

type AuthenticateInput struct {
	CPF string
}

type AuthenticateOutput struct {
	Customer    interface{} `json:"customer"`
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"token_type"`
	ExpiresIn   int64       `json:"expires_in"` // Time in seconds
}
