package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type orderJsonPresenter struct{}

// OrderJsonResponse represents the response of a order
func NewOrderJsonPresenter() port.OrderPresenter {
	return &orderJsonPresenter{}
}

// toOrderJsonResponse convert entity.Order to OrderJsonResponse
func toOrderJsonResponse(order *entity.Order) OrderJsonResponse {
	return OrderJsonResponse{
		ID:        order.ID,
		CustomerID: order.CustomerID,
		TotalBill: order.TotalBill,
		Status: 	order.Status.ToString(),
		CreatedAt: order.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: order.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// Present write the response to the client
func (p *orderJsonPresenter) Present(pp dto.OrderPresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.Order:
		output := toOrderJsonResponse(v)
		pp.Writer.JSON(http.StatusOK, output)
	case []*entity.Order:
		orderOutputs := make([]OrderJsonResponse, len(v))
		for i, order := range v {
			orderOutputs[i] = toOrderJsonResponse(order)
		}

		output := &OrderJsonPaginatedResponse{
			JsonPagination: JsonPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			Orders: orderOutputs,
		}
		pp.Writer.JSON(http.StatusOK, output)
	default:
		err := pp.Writer.Error(domain.NewInternalError(errors.New(domain.ErrInternalError)))
		if err != nil {
			pp.Writer.JSON(http.StatusInternalServerError, err)
		}
	}
}
