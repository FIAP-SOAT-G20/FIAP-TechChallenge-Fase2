package gateway

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type customerGateway struct {
	dataSource port.CustomerDataSource
}

func NewCustomerGateway(dataSource port.CustomerDataSource) port.CustomerGateway {
	return &customerGateway{
		dataSource: dataSource,
	}
}

func (cg *customerGateway) FindByID(ctx context.Context, id uint64) (*entity.Customer, error) {
	return cg.dataSource.FindByID(ctx, id)
}

func (cg *customerGateway) FindByCPF(ctx context.Context, cpf string) (*entity.Customer, error) {
	return cg.dataSource.FindByCPF(ctx, cpf)
}

func (cg *customerGateway) FindAll(ctx context.Context, name string, page, limit int) ([]*entity.Customer, int64, error) {
	filters := make(map[string]interface{})

	if name != "" {
		filters["name"] = name
	}

	return cg.dataSource.FindAll(ctx, filters, page, limit)
}

func (cg *customerGateway) Create(ctx context.Context, customer *entity.Customer) error {
	return cg.dataSource.Create(ctx, customer)
}

func (cg *customerGateway) Update(ctx context.Context, customer *entity.Customer) error {
	return cg.dataSource.Update(ctx, customer)
}

func (cg *customerGateway) Delete(ctx context.Context, id uint64) error {
	return cg.dataSource.Delete(ctx, id)
}
