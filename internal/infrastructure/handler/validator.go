package handler

import (
	"sync"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/go-playground/validator/v10"
)

var (
	once     sync.Once
	instance *validator.Validate
)

func GetValidator() *validator.Validate {
	once.Do(func() {
		instance = validator.New()

		// Here you can register custom validation functions
		// instance.RegisterValidation("custom", customValidation)
		err := instance.RegisterValidation("order_status", OrderStatusValidator)
		if err != nil {
			panic(err)
		}
	})

	return instance
}

func OrderStatusValidator(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	return entity.IsValidOrderStatus(status)
}
