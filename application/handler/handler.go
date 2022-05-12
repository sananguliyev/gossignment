package handler

import (
	"bytes"
	"encoding/json"
	"github.com/SananGuliyev/gossignment/domain/throwable"
	"github.com/elgopher/mapify"
	"net/http"
	"unicode"
	"unicode/utf8"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type Handler struct {
	mapper *mapify.Mapper
}

func (h *Handler) Respond(writer http.ResponseWriter, httpCode int, src interface{}) {
	var body []byte
	var err error

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if body, err = json.Marshal(src); err != nil {
		// h.Error could be used here, but it could cause infinity loop
		errorBody := `{"status": 500, "message": "Something happened wrong during generating response"}`
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(errorBody))
		return
	}

	body = bytes.ReplaceAll(body, []byte("\\u0026"), []byte("&"))

	writer.WriteHeader(httpCode)
	_, _ = writer.Write(body)
}

func (h *Handler) Error(writer http.ResponseWriter, err error) {
	var statusCode int
	var message string
	var code int

	switch e := err.(type) {
	case *json.UnsupportedTypeError, *json.UnmarshalTypeError, *json.SyntaxError:
		statusCode = http.StatusBadRequest
		code = 1
		message = "Request body is invalid"
	case throwable.BaseError:
		statusCode = e.StatusCode()
		code = e.Code()
		message = e.Error()
	default:
		statusCode = http.StatusInternalServerError
		message = e.Error()
		code = -1
	}

	errorResponse := &Response{
		Code:    code,
		Message: message,
	}

	h.Respond(writer, statusCode, errorResponse)
}

func (h *Handler) Success(writer http.ResponseWriter, httpCode int, src interface{}) {
	h.Respond(writer, httpCode, src)
}

func (h *Handler) SuccessWithMetadata(writer http.ResponseWriter, httpCode int, src interface{}) {
	mapper := h.getMapper()
	r, err := mapper.MapAny(src)
	if err != nil {
		h.Error(writer, err)
	}
	resultWithMetadata := r.(map[string]interface{})
	resultWithMetadata["code"] = 0
	resultWithMetadata["msg"] = "Success"

	h.Respond(writer, httpCode, resultWithMetadata)
}

func (h *Handler) getMapper() *mapify.Mapper {
	if h.mapper == nil {
		mapper := mapify.Mapper{}
		mapper.Rename = func(path string, e mapify.Element) (string, error) {
			r, n := utf8.DecodeRuneInString(e.Name())
			return string(unicode.ToLower(r)) + e.Name()[n:], nil
		}

		h.mapper = &mapper
	}

	return h.mapper
}
