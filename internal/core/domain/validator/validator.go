package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	once     sync.Once
	instance *validator.Validate
)

func GetValidator() *validator.Validate {
	once.Do(func() {
		instance = validator.New()

		// Aqui você pode registrar validações customizadas
		// instance.RegisterValidation("custom", customValidation)
	})

	return instance
}
