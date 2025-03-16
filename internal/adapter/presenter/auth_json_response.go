package presenter

import "encoding/json"

type AuthenticationResponse struct {
	Customer    CustomerJsonResponse `json:"customer"`
	AccessToken string               `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	TokenType   string               `json:"token_type" example:"Bearer"`
	ExpiresIn   int64                `json:"expires_in" example:"86400"`
}

func (r AuthenticationResponse) String() string {
	o, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(o)
}
