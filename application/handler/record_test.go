package handler

import (
	"bytes"
	"encoding/json"
	"github.com/SananGuliyev/gossignment/domain/io"
	"github.com/SananGuliyev/gossignment/domain/throwable"
	"github.com/SananGuliyev/gossignment/mocks"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecipeHandler_Get(t *testing.T) {
	someErrorMessage := "getiremedi"
	someRequestBody := []byte(`{"startDate":"2017-01-26","endDate":"2018-02-02","minCount":100,"maxCount":300}`)
	someRecords := make([]*io.RecordOutput, 0)

	tests := []struct {
		name               string
		expectedRecords    *io.RecordsOutput
		expectedResponse   interface{}
		expectedErr        error
		expectedStatusCode int
	}{
		{
			name:               "success case",
			expectedRecords:    &io.RecordsOutput{Records: someRecords},
			expectedResponse:   map[string]interface{}{"code": 0, "msg": "Success", "records": someRecords},
			expectedErr:        nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "error case",
			expectedRecords:    nil,
			expectedResponse:   map[string]interface{}{"code": 2, "msg": someErrorMessage},
			expectedErr:        throwable.NewInvalidBodyError(someErrorMessage),
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected, _ := json.Marshal(test.expectedResponse)
			expectedInput := &io.RecordFilterInput{}
			_ = json.Unmarshal(someRequestBody, expectedInput)

			request := httptest.NewRequest(
				http.MethodPost, "https://api.getir.com/records/filter", bytes.NewReader(someRequestBody))
			writer := httptest.NewRecorder()

			mockRecordInteractor := &mocks.RecordInteractor{}
			mockRecordInteractor.On("FilterRecords", expectedInput).
				Return(test.expectedRecords, test.expectedErr).
				Once()

			underTest := RecordHandler{recordInteractor: mockRecordInteractor}
			underTest.Filter(writer, request)

			response := writer.Result()
			actual, _ := ioutil.ReadAll(response.Body)

			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
			assert.Equal(t, expected, actual)
			mockRecordInteractor.AssertExpectations(t)
		})
	}
}
