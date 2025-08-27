package internalerrors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj any) error {
	validate := validator.New()

	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)

	var errRequired, errMin, errMax, errEmail error

	for _, v := range validationErrors {
		field := strings.ToLower(v.StructField())

		switch v.Tag() {

		case "required":
			errRequired = fmt.Errorf("%s is required", field)
		case "min":
			errMin = fmt.Errorf("min length for %s is %v", field, v.Param())
		case "max":
			errMax = fmt.Errorf("max length for %s is %v", field, v.Param())
		case "email":
			errEmail = fmt.Errorf("%s is invalid", field)
		}
	}

	if err = errors.Join(errRequired, errMin, errMax, errEmail); err == nil {
		return errors.New("validation error not mapped")
	}

	// name is required\nemail is invalid

	return err
}
