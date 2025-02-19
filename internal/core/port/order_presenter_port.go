package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"

type OrderPresenter interface {
	Present(pp dto.OrderPresenterInput)
}
