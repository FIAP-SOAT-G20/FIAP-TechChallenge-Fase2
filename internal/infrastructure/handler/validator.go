package handler

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

		// Here you can register custom validation functions
		// instance.RegisterValidation("custom", customValidation)
	})

	return instance
}
