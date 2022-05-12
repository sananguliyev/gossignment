package handler

import (
	"encoding/json"
	"github.com/SananGuliyev/gossignment/domain/interactor"
	"github.com/SananGuliyev/gossignment/domain/io"
	"net/http"
)

type RecordHandler struct {
	Handler
	recordInteractor interactor.RecordInteractor
}

func NewRecordHandler(recordInteractor interactor.RecordInteractor) *RecordHandler {
	return &RecordHandler{
		recordInteractor: recordInteractor,
	}
}

func (h *RecordHandler) Filter(writer http.ResponseWriter, request *http.Request) {
	var err error

	recordFilterInput := &io.RecordFilterInput{}
	err = json.NewDecoder(request.Body).Decode(recordFilterInput)
	if err != nil {
		h.Error(writer, err)
		return
	}

	records, err := h.recordInteractor.FilterRecords(recordFilterInput)
	if err != nil {
		h.Error(writer, err)
		return
	}

	h.SuccessWithMetadata(writer, http.StatusOK, records)
}
