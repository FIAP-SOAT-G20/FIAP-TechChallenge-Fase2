package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
)

type jwtService struct {
	secretKey []byte
}

type customClaims struct {
	CustomerID uint64 `json:"customer_id"`
	CPF        string `json:"cpf"`
	Name       string `json:"name"`
	jwt.RegisteredClaims
}

func NewJWTService(cfg *config.Config) port.JWTService {
	return &jwtService{
		secretKey: []byte(cfg.JWTSecret),
	}
}

func (s *jwtService) GenerateToken(customerID uint64, cpf string, name string, expiresIn time.Duration) (string, error) {
	// Create claims directly with primitive parameters
	tokenClaims := customClaims{
		CustomerID: customerID,
		CPF:        cpf,
		Name:       name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	signedToken, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signature method")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func (s *jwtService) ExtractClaims(tokenString string) (uint64, string, string, error) {
	claims := &customClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signature method")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return 0, "", "", err
	}

	if !token.Valid {
		return 0, "", "", errors.New("invalid token")
	}

	return claims.CustomerID, claims.CPF, claims.Name, nil
}
