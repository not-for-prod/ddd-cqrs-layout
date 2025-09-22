package validator

import (
	"strconv"

	validator "github.com/go-playground/validator/v10"
)

const (
	pinLen = 6
)

func validatePin(fl validator.FieldLevel) bool {
	pin := fl.Field().String()
	if len(pin) != pinLen {
		return false
	}

	_, err := strconv.ParseInt(pin, 10, 64)
	return err == nil
}
