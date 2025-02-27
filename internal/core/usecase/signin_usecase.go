package usecase

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type SignInUsecase struct {
	customerDatasource port.CustomerDatasourcePort
}

func NewSignInUsecase(customerDatasource port.CustomerDatasourcePort) port.SignInUsecasePort {
	return &SignInUsecase{
		customerDatasource: customerDatasource,
	}
}

func (u *SignInUsecase) GetByCPF(cpf string) (*entity.Customer, error) {
	return u.customerDatasource.GetByCPF(cpf)
} 