package shared

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(validationStruct interface{}) error {
	validate := validator.New()
	return validate.Struct(validationStruct)
}

func TranslateError(err error) string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var errMessages []string
		for _, e := range ve {
			switch e.Tag() {
			case "required":
				errMessages = append(errMessages, e.Field()+" is required")
			default:
				errMessages = append(errMessages, e.Field()+" is invalid")
			}
		}

		return strings.Join(errMessages, ", ")
	} else {
		return err.Error()
	}
}
