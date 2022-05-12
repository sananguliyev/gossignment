package throwable

import "net/http"

type NotFoundError struct {
	baseError
}

func NewNotFoundError(err string) *NotFoundError {
	return &NotFoundError{
		baseError{
			err:        err,
			code:       3,
			statusCode: http.StatusNotFound,
		},
	}
}
