package structvalidator

import (
	coreerror "github.com/IbnAnjung/dealls/pkg/error"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(obj interface{}) error
}

type structValidator struct {
	validator *validator.Validate
}

func NewStructValidator() Validator {
	return &structValidator{
		validator: validator.New(),
	}
}

func (v *structValidator) Validate(obj interface{}) error {
	err := v.validator.Struct(obj)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "fail validate data")
			return e
		}

		e := coreerror.NewValidationError()
		e.Errors = map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			e.Errors[err.Field()] = err.Error()
		}

		err = e

		return err
	}

	return nil
}
