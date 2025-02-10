package presenter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

func TestProductPresenter_ToOutput(t *testing.T) {
	presenter := NewProductPresenter()
	now := time.Now()

	product := &entity.Product{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	output := presenter.ToOutput(product)

	assert.NotNil(t, output)
	assert.Equal(t, product.ID, output.ID)
	assert.Equal(t, product.Name, output.Name)
	assert.Equal(t, product.Description, output.Description)
	assert.Equal(t, product.Price, output.Price)
	assert.Equal(t, product.CategoryID, output.CategoryID)
	assert.Equal(t, now.Format("2006-01-02T15:04:05Z07:00"), output.CreatedAt)
	assert.Equal(t, now.Format("2006-01-02T15:04:05Z07:00"), output.UpdatedAt)
}

func TestProductPresenter_ToPaginatedOutput(t *testing.T) {
	presenter := NewProductPresenter()
	now := time.Now()

	products := []*entity.Product{
		{
			ID:          1,
			Name:        "Product 1",
			Description: "Description 1",
			Price:       99.99,
			CategoryID:  1,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          2,
			Name:        "Product 2",
			Description: "Description 2",
			Price:       199.99,
			CategoryID:  2,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	output := presenter.ToPaginatedOutput(products, 10, 1, 2)

	assert.NotNil(t, output)
	assert.Equal(t, int64(10), output.Total)
	assert.Equal(t, 1, output.Page)
	assert.Equal(t, 2, output.Limit)
	assert.Len(t, output.Products, 2)

	// Verifica o primeiro produto
	assert.Equal(t, products[0].ID, output.Products[0].ID)
	assert.Equal(t, products[0].Name, output.Products[0].Name)
	assert.Equal(t, products[0].Description, output.Products[0].Description)
	assert.Equal(t, products[0].Price, output.Products[0].Price)
	assert.Equal(t, products[0].CategoryID, output.Products[0].CategoryID)
	assert.Equal(t, now.Format("2006-01-02T15:04:05Z07:00"), output.Products[0].CreatedAt)
	assert.Equal(t, now.Format("2006-01-02T15:04:05Z07:00"), output.Products[0].UpdatedAt)

	// Verifica o segundo produto
	assert.Equal(t, products[1].ID, output.Products[1].ID)
	assert.Equal(t, products[1].Name, output.Products[1].Name)
	assert.Equal(t, products[1].Description, output.Products[1].Description)
	assert.Equal(t, products[1].Price, output.Products[1].Price)
	assert.Equal(t, products[1].CategoryID, output.Products[1].CategoryID)
	assert.Equal(t, now.Format("2006-01-02T15:04:05Z07:00"), output.Products[1].CreatedAt)
	assert.Equal(t, now.Format("2006-01-02T15:04:05Z07:00"), output.Products[1].UpdatedAt)
}
