package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"

type PaymentPresenter interface {
	Present(pp dto.PaymentPresenterInput)
}

type PaymentPresenterDTO struct {
	Result any
	Total  int64
	Page   int
	Limit  int
	Writer dto.ResponseWriter
}
