package mocks

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/stretchr/testify/mock"
)

type SignInUsecaseMock struct {
	mock.Mock
}

func NewSignInUsecaseMock() *SignInUsecaseMock {
	return &SignInUsecaseMock{}
}

func (m *SignInUsecaseMock) GetByCPF(cpf string) (*entity.Customer, error) {
	args := m.Called(cpf)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Customer), args.Error(1)
}
