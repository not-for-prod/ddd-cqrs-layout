package validator

import (
	"log"
	"sync"

	validator "github.com/go-playground/validator/v10"
)

const (
	pin = "pin"
)

var (
	once     sync.Once
	validate *validator.Validate
)

func GetInstance() *validator.Validate {
	if validate == nil {
		once.Do(
			func() {
				// register custom example
				err := validator.New().RegisterValidation(pin, validatePin)
				if err != nil {
					log.Fatalf("error registering custom validation: %s", err)
				}

				validate = validator.New()
			},
		)
	}

	return validate
}
