package validator

import (
	"github.com/SananGuliyev/gossignment/domain/io"
	valid "github.com/SananGuliyev/gossignment/domain/validator"
	"github.com/go-playground/validator/v10"
)

type memValidator struct {
	validator *validator.Validate
}

func NewMemValidator(validator *validator.Validate) valid.MemValidator {
	return &memValidator{
		validator: validator,
	}
}

func (v *memValidator) ValidateCreateInput(input *io.MemInput) error {
	var err error

	err = v.validator.Struct(input)
	if err != nil {
		return err
	}

	return nil
}
