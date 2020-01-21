package domain

// Error エラーの情報取得系ばっか
type Error interface {
	// Error 発生したerrorをstringに変換
	Error() string

	// GetStatusCode ステータスコードの取得
	StatusCode() int
}

// BadRequest クライアントエラー(400)の処理
func BadRequest(err error) Error {
	return newError(StatusBadRequest, err)
}

// Unauthorized 認証エラー(401)の処理
func Unauthorized(err error) Error {
	return newError(StatusUnauthorized, err)
}

// NotFound コンテンツが見つからない(404)の処理
func NotFound(err error) Error {
	return newError(StatusNotFound, err)
}

// MethodNotAllowed メソッドエラー(405)の処理
func MethodNotAllowed(err error) Error {
	return newError(StatusMethodNotAllowed, err)
}

// InternalServerError サーバのエラー(500)の処理
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

func (e *httpError) StatusCode() int {
	return e.Code
}

const (
	StatusDefault             = 500
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusInternalServerError = 500
)
