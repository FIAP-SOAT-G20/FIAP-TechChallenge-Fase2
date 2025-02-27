package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"

type SignInUsecasePort interface {
	GetByCPF(cpf string) (*entity.Customer, error)
}
