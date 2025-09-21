package validator

import (
	"strconv"

	validator "github.com/go-playground/validator/v10"
)

func validatePin(fl validator.FieldLevel) bool {
	pin := fl.Field().String()
	if len(pin) != 6 {
		return false
	}

	_, err := strconv.ParseInt(pin, 10, 64)
	if err != nil {
		return false
	}

	return true
}
