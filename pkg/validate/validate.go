package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sebajax/go-clean-architecture/pkg/apperror"
)

// generate validator instance
var Validator = validator.New(validator.WithRequiredStructEnabled())

func Validate(body interface{}) (*apperror.AppError, error) {
	err := Validator.Struct(body)
	var schemaErr *apperror.AppError
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			schemaErr.BadRequest(fmt.Sprintf("Field %s Tag %s", err.Field(), err.Tag()))
			return schemaErr, err
		}
	}
	return nil, nil
}
