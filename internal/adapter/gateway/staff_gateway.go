package gateway

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type staffGateway struct {
	dataSource port.StaffDataSource
}

func NewStaffGateway(dataSource port.StaffDataSource) port.StaffGateway {
	return &staffGateway{
		dataSource: dataSource,
	}
}

func (pg *staffGateway) FindByID(ctx context.Context, id uint64) (*entity.Staff, error) {
	return pg.dataSource.FindByID(ctx, id)
}

func (pg *staffGateway) FindAll(ctx context.Context, name string, role valueobject.StaffRole, page, limit int) ([]*entity.Staff, int64, error) {
	filters := make(map[string]interface{})

	if name != "" {
		filters["name"] = name
	}
	if role != "" {
		filters["role"] = role.String()
	}

	return pg.dataSource.FindAll(ctx, filters, page, limit)
}

func (pg *staffGateway) Create(ctx context.Context, product *entity.Staff) error {
	return pg.dataSource.Create(ctx, product)
}

func (pg *staffGateway) Update(ctx context.Context, product *entity.Staff) error {
	return pg.dataSource.Update(ctx, product)
}

func (pg *staffGateway) Delete(ctx context.Context, id uint64) error {
	return pg.dataSource.Delete(ctx, id)
}
