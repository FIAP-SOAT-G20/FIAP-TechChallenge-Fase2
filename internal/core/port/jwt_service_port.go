package port

import "time"

// JWTService provides token generation and validation methods
type JWTService interface {
	// GenerateToken creates a new JWT token with the given claims
	GenerateToken(customerID uint64, cpf string, name string, expiresIn time.Duration) (token string, err error)

	// ValidateToken verifies if a token is valid without extracting data
	ValidateToken(token string) error

	// ExtractClaims extracts customer data from a valid token
	ExtractClaims(token string) (customerID uint64, cpf string, name string, err error)
}
