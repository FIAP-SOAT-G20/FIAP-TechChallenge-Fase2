package usecase

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type SignInUsecase struct {
	customerGateway port.CustomerGateway
}

func NewSignInUsecase(customerGateway port.CustomerGateway) port.SignInUsecase {
	return &SignInUsecase{
		customerGateway: customerGateway,
	}
}

func (u *SignInUsecase) GetByCPF(cpf string) (*entity.Customer, error) {
	return u.customerGateway.FindByCPF(context.Background(), cpf)
}
