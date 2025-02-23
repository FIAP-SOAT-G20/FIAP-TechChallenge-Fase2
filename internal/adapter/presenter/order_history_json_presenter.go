package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type orderHistoryJsonPresenter struct {
	writer ResponseWriter
}

// OrderHistoryJsonResponse represents the response of a orderHistory
func NewOrderHistoryJsonPresenter(writer ResponseWriter) port.Presenter {
	return &orderHistoryJsonPresenter{writer}
}

// toOrderHistoryJsonResponse convert entity.OrderHistory to OrderHistoryJsonResponse
func toOrderHistoryJsonResponse(orderHistory *entity.OrderHistory) OrderHistoryJsonResponse {
	return OrderHistoryJsonResponse{
		ID:        orderHistory.ID,
		OrderID:   orderHistory.OrderID,
		Status:    orderHistory.Status.String(),
		CreatedAt: orderHistory.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// Present write the response to the client
func (p *orderHistoryJsonPresenter) Present(pp dto.PresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.OrderHistory:
		output := toOrderHistoryJsonResponse(v)
		p.writer.JSON(http.StatusOK, output)
	case []*entity.OrderHistory:
		orderHistoryOutputs := make([]OrderHistoryJsonResponse, len(v))
		for i, orderHistory := range v {
			orderHistoryOutputs[i] = toOrderHistoryJsonResponse(orderHistory)
		}

		output := &OrderHistoryJsonPaginatedResponse{
			JsonPagination: JsonPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			OrderHistories: orderHistoryOutputs,
		}
		p.writer.JSON(http.StatusOK, output)
	default:
		p.writer.JSON(
			http.StatusInternalServerError,
			domain.NewInternalError(errors.New(domain.ErrInternalError)),
		)
	}
}
