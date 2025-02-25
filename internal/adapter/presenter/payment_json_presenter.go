package presenter

import (
	"errors"
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
	return &paymentJsonPresenter{writer}
}

// Present write the response to the client
func (p *paymentJsonPresenter) Present(pp dto.PresenterInput) ([]byte, error) {
	switch v := pp.Result.(type) {
	case *entity.Payment:
		output := ToPaymentJsonResponse(v)
		p.writer.JSON(http.StatusOK, output)
		return nil, nil
	default:
		p.writer.JSON(
			http.StatusInternalServerError,
			domain.NewInternalError(errors.New(domain.ErrInternalError)),
		)
		return nil, nil
	}
}

// ToPaymentJsonResponse convert entity.Payment to PaymentJsonResponse
func ToPaymentJsonResponse(p *entity.Payment) PaymentJsonResponse {
	return PaymentJsonResponse{
		ID:                p.ID,
		Status:            p.Status,
		OrderID:           p.OrderID,
		ExternalPaymentID: p.ExternalPaymentID,
		QrData:            p.QrData,
	}
}
