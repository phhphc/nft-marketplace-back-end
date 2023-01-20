package models

type HTTPError struct {
	Code int
	Err  error
}

func NewHTTPError(code int, err error) *HTTPError {
	return &HTTPError{
		Code: code,
		Err:  err,
	}
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}
