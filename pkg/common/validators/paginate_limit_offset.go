package validators

import (
	"github.com/go-playground/validator/v10"
)

func LimitAndOffset(fl validator.FieldLevel) bool {
	isValid := true

	switch fl.FieldName() {
	case "Limit":
		if fl.Field().Int() <= 0 {
			isValid = false
		}
	case "Offset":
		if fl.Field().Int() < 0 || fl.Field().String() == "" {
			isValid = false
		}
	}
	return isValid
}
