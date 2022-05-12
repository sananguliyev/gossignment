package interactor

import (
	"github.com/SananGuliyev/gossignment/domain/io"
	"github.com/SananGuliyev/gossignment/domain/repository"
	"github.com/SananGuliyev/gossignment/domain/throwable"
	"github.com/SananGuliyev/gossignment/domain/validator"
)

type RecordInteractor interface {
	FilterRecords(input *io.RecordFilterInput) (*io.RecordsOutput, error)
}

type recordInteractor struct {
	recordRepository repository.RecordRepository
	recordValidator  validator.RecordValidator
}

func NewRecordInteractor(
	recordRepository repository.RecordRepository,
	recordValidator validator.RecordValidator,
) RecordInteractor {
	return &recordInteractor{
		recordRepository: recordRepository,
		recordValidator:  recordValidator,
	}
}

func (s *recordInteractor) FilterRecords(input *io.RecordFilterInput) (*io.RecordsOutput, error) {
	var err error

	err = s.recordValidator.ValidateFilterInput(input)
	if err != nil {
		return nil, throwable.NewInvalidBodyError(err.Error())
	}

	records, err := s.recordRepository.FilterByTimeAndAmount(
		input.StartDate,
		input.EndDate,
		input.MinCount,
		input.MaxCount,
	)
	if err != nil {
		return nil, err
	}

	result := &io.RecordsOutput{Records: make([]*io.RecordOutput, len(records))}

	for i, record := range records {
		result.Records[i] = &io.RecordOutput{
			Key:        record.Key,
			CreatedAt:  record.CreatedAt,
			TotalCount: *record.TotalCount,
		}
	}

	return result, nil
}
