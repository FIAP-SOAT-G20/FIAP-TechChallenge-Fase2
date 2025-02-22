package presenter

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type paymentJsonPresenter struct {
	writer ResponseWriter
}

// PaymentJsonResponse represents the response of a payment
func NewPaymentJsonPresenter(writer ResponseWriter) port.Presenter {
	return &paymentJsonPresenter{}
}

// ToCustomerJsonResponse convert entity.Payment to PaymentJsonResponse
func ToPaymentJsonResponse(customer *entity.Payment) PaymentJsonResponse {
	return PaymentJsonResponse{}
}

// Present write the response to the client
func (p *paymentJsonPresenter) Present(pp dto.PresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.Payment:
		output := ToPaymentJsonResponse(v)
		p.writer.JSON(http.StatusOK, output)
	default:
		fmt.Println("paymentJsonPresenter Unknown type")
		p.writer.JSON(
			http.StatusInternalServerError,
			domain.NewInternalError(errors.New(domain.ErrInternalError)),
		)
	}
}
