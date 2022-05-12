package handler

import (
	"encoding/json"
	"github.com/SananGuliyev/gossignment/domain/interactor"
	"github.com/SananGuliyev/gossignment/domain/io"
	"net/http"
)

type InMemoryHandler struct {
	Handler
	memInteractor interactor.MemInteractor
}

func NewInMemoryHandler(memService interactor.MemInteractor) *InMemoryHandler {
	return &InMemoryHandler{
		memInteractor: memService,
	}
}

func (h *InMemoryHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var err error

	memCreateInput := &io.MemInput{}
	err = json.NewDecoder(request.Body).Decode(memCreateInput)
	if err != nil {
		h.Error(writer, err)
		return
	}

	mem, err := h.memInteractor.Create(memCreateInput)
	if err != nil {
		h.Error(writer, err)
		return
	}

	h.Success(writer, http.StatusCreated, mem)
}

func (h *InMemoryHandler) Read(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get("key")

	mem, err := h.memInteractor.Read(key)
	if err != nil {
		h.Error(writer, err)
		return
	}

	h.Success(writer, http.StatusOK, mem)
}
