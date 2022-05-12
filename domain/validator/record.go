package validator

import "github.com/SananGuliyev/gossignment/domain/io"

type RecordValidator interface {
	ValidateFilterInput(input *io.RecordFilterInput) error
}
