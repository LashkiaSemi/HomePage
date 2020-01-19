package domain

type Error interface {
	Error() string
	GetStatusCode() int
}

// BadRequest 400エラーの処理
func BadRequest(err error) Error {
	return newError(StatusBadRequest, err)
}

func Unauthorized(err error) Error {
	return newError(StatusUnauthorized, err)
}

func MethodNotAllowed(err error) Error {
	return newError(StatusMethodNotAllowed, err)
}

func InternalServerError(err error) Error {
	return newError(StatusInternalServerError, err)
}

func newError(code int, err error) Error {
	if err != nil {
		return &httpError{
			Code:  code,
			error: err,
		}
	}
	return nil
}

type httpError struct {
	Code int
	error
}

func (e *httpError) Error() string {
	return e.error.Error()
}

func (e *httpError) GetStatusCode() int {
	return e.Code
}

const (
	StatusDefault             = 500
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusMethodNotAllowed    = 405
	StatusInternalServerError = 500
)
