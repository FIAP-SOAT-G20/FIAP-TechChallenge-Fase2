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

func (s *jwtService) GenerateToken(claims port.JWTClaims, expiresIn time.Duration) (string, error) {
	tokenClaims := customClaims{
		CustomerID: claims.CustomerID,
		CPF:        claims.CPF,
		Name:       claims.Name,
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

// ValidateToken valida um token JWT e retorna os claims
func (s *jwtService) ValidateToken(tokenString string) (*port.JWTClaims, error) {
	claims := &customClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return &port.JWTClaims{
		CustomerID: claims.CustomerID,
		CPF:        claims.CPF,
		Name:       claims.Name,
	}, nil
}
