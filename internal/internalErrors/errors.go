package internalerrors

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInternalServerError error = errors.New("internal server error")

func ProcessError(err error) error {
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrInternalServerError
		}
	}

	return err
}
