package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type staffJsonPresenter struct{}

// StaffJsonResponse represents the response of a staff
func NewStaffJsonPresenter() port.StaffPresenter {
	return &staffJsonPresenter{}
}

// toStaffJsonResponse convert entity.Staff to StaffJsonResponse
func toStaffJsonResponse(staff *entity.Staff) StaffJsonResponse {
	return StaffJsonResponse{
		ID:        staff.ID,
		Name:      staff.Name,
		Role:      string(staff.Role),
		CreatedAt: staff.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: staff.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// Present write the response to the client
func (p *staffJsonPresenter) Present(pp dto.StaffPresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.Staff:
		output := toStaffJsonResponse(v)
		pp.Writer.JSON(http.StatusOK, output)
	case []*entity.Staff:
		staffOutputs := make([]StaffJsonResponse, len(v))
		for i, staff := range v {
			staffOutputs[i] = toStaffJsonResponse(staff)
		}

		output := &StaffJsonPaginatedResponse{
			JsonPagination: JsonPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			Staffs: staffOutputs,
		}
		pp.Writer.JSON(http.StatusOK, output)
	default:
		err := pp.Writer.Error(domain.NewInternalError(errors.New(domain.ErrInternalError)))
		if err != nil {
			pp.Writer.JSON(http.StatusInternalServerError, err)
		}
	}
}
