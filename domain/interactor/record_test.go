package interactor

import (
	"errors"
	"github.com/SananGuliyev/gossignment/domain/entity"
	"github.com/SananGuliyev/gossignment/domain/io"
	"github.com/SananGuliyev/gossignment/domain/throwable"
	"github.com/SananGuliyev/gossignment/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRecipeInteractor(t *testing.T) {
	expectedRecordRepository := &mocks.RecordRepository{}
	expectedRecordValidator := &mocks.RecordValidator{}

	expected := &recordInteractor{expectedRecordRepository, expectedRecordValidator}

	actual := NewRecordInteractor(expectedRecordRepository, expectedRecordValidator)

	assert.Equal(t, expected, actual)
}

func TestRecordInteractor_FilterRecords(t *testing.T) {
	someKey := "HemenGetiyorum"
	someCreatedAt := time.Now()
	someTotalCount := 1991
	someStartDate := io.Date{Time: time.Now()}
	someEndDate := io.Date{Time: time.Now().Add(1991)}
	var someMinCount int64 = 1903
	var someMaxCount int64 = 2018

	var expectedError error
	var expectedRecords = []*entity.Record{
		{
			Key:        someKey,
			CreatedAt:  someCreatedAt,
			TotalCount: &someTotalCount,
		},
	}
	var expectedOutput = &io.RecordsOutput{
		Records: []*io.RecordOutput{
			{
				Key:        someKey,
				CreatedAt:  someCreatedAt,
				TotalCount: someTotalCount,
			},
		},
	}

	someInput := io.RecordFilterInput{
		StartDate: someStartDate,
		EndDate:   someEndDate,
		MinCount:  someMinCount,
		MaxCount:  someMaxCount,
	}

	mockRecordRepository := &mocks.RecordRepository{}
	mockRecordRepository.On(
		"FilterByTimeAndAmount",
		someStartDate,
		someEndDate,
		someMinCount,
		someMaxCount,
	).
		Return(expectedRecords, expectedError).
		Once()

	mockRecordValidator := &mocks.RecordValidator{}
	mockRecordValidator.On("ValidateFilterInput", &someInput).
		Return(expectedError).
		Once()

	underTest := &recordInteractor{mockRecordRepository, mockRecordValidator}

	actualOutput, actualError := underTest.FilterRecords(&someInput)

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedError, actualError)
	mockRecordRepository.AssertExpectations(t)
	mockRecordValidator.AssertExpectations(t)
}

func TestRecordInteractor_FilterRecords_RepositoryError(t *testing.T) {
	someStartDate := io.Date{Time: time.Now()}
	someEndDate := io.Date{Time: time.Now().Add(1991)}
	var someMinCount int64 = 1903
	var someMaxCount int64 = 2018

	var expectedValidationError error
	var expectedError = errors.New("getiremedi")
	var expectedRecords []*entity.Record
	var expectedOutput *io.RecordsOutput

	someInput := io.RecordFilterInput{
		StartDate: someStartDate,
		EndDate:   someEndDate,
		MinCount:  someMinCount,
		MaxCount:  someMaxCount,
	}

	mockRecordRepository := &mocks.RecordRepository{}
	mockRecordRepository.On(
		"FilterByTimeAndAmount",
		someStartDate,
		someEndDate,
		someMinCount,
		someMaxCount,
	).
		Return(expectedRecords, expectedError).
		Once()

	mockRecordValidator := &mocks.RecordValidator{}
	mockRecordValidator.On("ValidateFilterInput", &someInput).
		Return(expectedValidationError).
		Once()

	underTest := &recordInteractor{mockRecordRepository, mockRecordValidator}

	actualOutput, actualError := underTest.FilterRecords(&someInput)

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedError, actualError)
	mockRecordRepository.AssertExpectations(t)
	mockRecordValidator.AssertExpectations(t)
}

func TestRecordInteractor_FilterRecords_ValidationError(t *testing.T) {
	var expectedError = throwable.NewInvalidBodyError("getiremedi")
	var expectedOutput *io.RecordsOutput

	someInput := io.RecordFilterInput{}

	mockRecordRepository := &mocks.RecordRepository{}

	mockRecordValidator := &mocks.RecordValidator{}
	mockRecordValidator.On("ValidateFilterInput", &someInput).
		Return(expectedError).
		Once()

	underTest := &recordInteractor{mockRecordRepository, mockRecordValidator}

	actualOutput, actualError := underTest.FilterRecords(&someInput)

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedError, actualError)
	mockRecordRepository.AssertExpectations(t)
	mockRecordValidator.AssertExpectations(t)
}
