package throwable

import "net/http"

type InvalidBodyError struct {
	baseError
}

func NewInvalidBodyError(err string) *InvalidBodyError {
	return &InvalidBodyError{
		baseError{
			err:        err,
			code:       2,
			statusCode: http.StatusBadRequest,
		},
	}
}
