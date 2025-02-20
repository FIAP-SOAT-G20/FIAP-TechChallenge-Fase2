package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type customerJsonPresenter struct {
	writer ResponseWriter
}

// CustomerJsonResponse represents the response of a customer
func NewCustomerJsonPresenter(writer ResponseWriter) port.Presenter {
	return &customerJsonPresenter{writer}
}

// ToCustomerJsonResponse convert entity.Customer to CustomerJsonResponse
func ToCustomerJsonResponse(customer *entity.Customer) CustomerJsonResponse {
	return CustomerJsonResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		CPF:       customer.CPF,
		CreatedAt: customer.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: customer.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// Present write the response to the client
func (p *customerJsonPresenter) Present(pp dto.PresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.Customer:
		output := ToCustomerJsonResponse(v)
		p.writer.JSON(http.StatusOK, output)
	case []*entity.Customer:
		customerOutputs := make([]CustomerJsonResponse, len(v))
		for i, customer := range v {
			customerOutputs[i] = ToCustomerJsonResponse(customer)
		}

		output := &CustomerJsonPaginatedResponse{
			JsonPagination: JsonPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			Customers: customerOutputs,
		}
		p.writer.JSON(http.StatusOK, output)
	default:
		err := p.writer.Error(domain.NewInternalError(errors.New(domain.ErrInternalError)))
		if err != nil {
			p.writer.JSON(http.StatusInternalServerError, err)
		}
	}
}
