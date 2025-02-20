package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"

type CustomerPresenter interface {
	Present(pp dto.CustomerPresenterInput)
}
