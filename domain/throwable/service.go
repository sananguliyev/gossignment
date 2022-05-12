package throwable

type BaseError interface {
	Error() string
	Code() int
	StatusCode() int
}

type baseError struct {
	err        string
	code       int
	statusCode int
}

func (e *baseError) Error() string {
	return e.err
}

func (e *baseError) Code() int {
	return e.code
}

func (e *baseError) StatusCode() int {
	return e.statusCode
}
