package interactor

import (
	"github.com/SananGuliyev/gossignment/domain/entity"
	"github.com/SananGuliyev/gossignment/domain/io"
	"github.com/SananGuliyev/gossignment/domain/repository"
	"github.com/SananGuliyev/gossignment/domain/throwable"
	"github.com/SananGuliyev/gossignment/domain/validator"
)

type MemInteractor interface {
	Create(input *io.MemInput) (*io.MemOutput, error)
	Read(key string) (*io.MemOutput, error)
}

type memInteractor struct {
	memRepository repository.MemRepository
	memValidator  validator.MemValidator
}

func NewMemInteractor(
	memRepository repository.MemRepository,
	memValidator validator.MemValidator,
) MemInteractor {
	return &memInteractor{
		memRepository: memRepository,
		memValidator:  memValidator,
	}
}

func (s *memInteractor) Create(input *io.MemInput) (*io.MemOutput, error) {
	var err error

	err = s.memValidator.ValidateCreateInput(input)
	if err != nil {
		return nil, throwable.NewInvalidBodyError(err.Error())
	}

	mem := &entity.Mem{
		Key:   input.Key,
		Value: input.Value,
	}

	err = s.memRepository.Create(mem)
	if err != nil {
		return nil, err
	}

	return &io.MemOutput{
		Key:   mem.Key,
		Value: mem.Value,
	}, nil
}

func (s *memInteractor) Read(key string) (*io.MemOutput, error) {
	if key == "" {
		return nil, throwable.NewMissingParamError("Key is required query parameter")
	}

	mem, err := s.memRepository.GetByKey(key)
	if err != nil {
		return nil, err
	}

	return &io.MemOutput{
		Key:   mem.Key,
		Value: mem.Value,
	}, nil
}
