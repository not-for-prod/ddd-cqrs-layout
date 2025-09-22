package validator

import (
	"fmt"
	"sync"

	validator "github.com/go-playground/validator/v10"
)

const (
	pin = "pin"
)

var (
	once     sync.Once           //nolint:gochecknoglobals // no need
	validate *validator.Validate //nolint:gochecknoglobals // no need
)

func GetInstance() *validator.Validate {
	if validate == nil {
		once.Do(
			func() {
				// register custom example
				err := validator.New().RegisterValidation(pin, validatePin)
				if err != nil {
					panic(fmt.Sprintf("failed to validate pin: %s", err.Error()))
				}

				validate = validator.New()
			},
		)
	}

	return validate
}
