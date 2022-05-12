package validator

import "github.com/SananGuliyev/gossignment/domain/io"

type MemValidator interface {
	ValidateCreateInput(input *io.MemInput) error
}
