package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type paymentJsonPresenter struct{}

// PaymentJsonResponse represents the response of a payment
func NewPaymentJsonPresenter() port.PaymentPresenter {
	return &paymentJsonPresenter{}
}

// ToCustomerJsonResponse convert entity.Payment to PaymentJsonResponse
func ToPaymentJsonResponse(customer *entity.Payment) PaymentJsonResponse {
	return PaymentJsonResponse{}
}

// Present write the response to the client
func (p *paymentJsonPresenter) Present(pp dto.PaymentPresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.Payment:
		output := ToPaymentJsonResponse(v)
		pp.Writer.JSON(http.StatusOK, output)
	default:
		err := pp.Writer.Error(domain.NewInternalError(errors.New(domain.ErrInternalError)))
		if err != nil {
			pp.Writer.JSON(http.StatusInternalServerError, err)
		}
	}
}
