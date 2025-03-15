package port

import "time"

type JWTClaims struct {
	CustomerID uint64 `json:"customer_id"`
	CPF        string `json:"cpf"`
	Name       string `json:"name"`
}

type JWTService interface {
	GenerateToken(claims JWTClaims, expiresIn time.Duration) (string, error)
	ValidateToken(token string) (*JWTClaims, error)
}
