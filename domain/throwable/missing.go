package throwable

import "net/http"

type MissingParamError struct {
	baseError
}

func NewMissingParamError(err string) *NotFoundError {
	return &NotFoundError{
		baseError{
			err:        err,
			code:       4,
			statusCode: http.StatusBadRequest,
		},
	}
}
