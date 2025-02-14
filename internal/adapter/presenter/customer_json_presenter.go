package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type customerJsonPresenter struct{}

// CustomerJsonResponse represents the response of a customer
func NewCustomerJsonPresenter() port.CustomerPresenter {
	return &customerJsonPresenter{}
}

// toCustomerJsonResponse convert entity.Customer to CustomerJsonResponse
func toCustomerJsonResponse(customer *entity.Customer) CustomerJsonResponse {
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
func (p *customerJsonPresenter) Present(pp dto.CustomerPresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.Customer:
		output := toCustomerJsonResponse(v)
		pp.Writer.JSON(http.StatusOK, output)
	case []*entity.Customer:
		customerOutputs := make([]CustomerJsonResponse, len(v))
		for i, customer := range v {
			customerOutputs[i] = toCustomerJsonResponse(customer)
		}

		output := &CustomerJsonPaginatedResponse{
			JsonPagination: JsonPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			Customers: customerOutputs,
		}
		pp.Writer.JSON(http.StatusOK, output)
	default:
		err := pp.Writer.Error(domain.NewInternalError(errors.New(domain.ErrInternalError)))
		if err != nil {
			pp.Writer.JSON(http.StatusInternalServerError, err)
		}
	}
}
