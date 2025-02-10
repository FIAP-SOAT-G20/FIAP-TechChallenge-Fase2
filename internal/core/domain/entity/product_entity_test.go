package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	tests := []struct {
		name        string
		productName string
		description string
		price       float64
		categoryID  uint64
		expectError bool
	}{
		{
			name:        "should create valid product",
			productName: "Valid Product",
			description: "Valid description",
			price:       99.99,
			categoryID:  1,
			expectError: false,
		},
		{
			name:        "should fail with empty name",
			productName: "",
			description: "Valid description",
			price:       99.99,
			categoryID:  1,
			expectError: true,
		},
		{
			name:        "should fail with short name",
			productName: "ab",
			description: "Valid description",
			price:       99.99,
			categoryID:  1,
			expectError: true,
		},
		{
			name:        "should fail with zero price",
			productName: "Valid Product",
			description: "Valid description",
			price:       0,
			categoryID:  1,
			expectError: true,
		},
		{
			name:        "should fail with negative price",
			productName: "Valid Product",
			description: "Valid description",
			price:       -10.00,
			categoryID:  1,
			expectError: true,
		},
		{
			name:        "should fail with zero category ID",
			productName: "Valid Product",
			description: "Valid description",
			price:       99.99,
			categoryID:  0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := NewProduct(tt.productName, tt.description, tt.price, tt.categoryID)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, product)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, product)
				assert.Equal(t, tt.productName, product.Name)
				assert.Equal(t, tt.description, product.Description)
				assert.Equal(t, tt.price, product.Price)
				assert.Equal(t, tt.categoryID, product.CategoryID)
				assert.NotZero(t, product.CreatedAt)
				assert.NotZero(t, product.UpdatedAt)
			}
		})
	}
}
