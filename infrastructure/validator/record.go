package validator

import (
	"github.com/SananGuliyev/gossignment/domain/io"
	valid "github.com/SananGuliyev/gossignment/domain/validator"
	"github.com/go-playground/validator/v10"
)

type recordValidator struct {
	validator *validator.Validate
}

func NewRecordValidator(validator *validator.Validate) valid.RecordValidator {
	return &recordValidator{
		validator: validator,
	}
}

func (v *recordValidator) ValidateFilterInput(input *io.RecordFilterInput) error {
	var err error

	err = v.validator.Struct(input)
	if err != nil {
		return err
	}

	return nil
}
