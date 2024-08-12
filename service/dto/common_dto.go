package dto

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func GetValidateError(err error) error {
	var validateErrors validation.Errors
	if errors.As(err, &validateErrors) {
		for _, msg := range validateErrors {
			return msg
		}
	}
	return err
}
