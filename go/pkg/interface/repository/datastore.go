package repository

// SQLHandler sqlの操作子
type SQLHandler interface {
	ErrNoRows() error

	Execute(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row
}

// Result sql.Result
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Rows sql.Rows
type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

// Row sql.Row
type Row interface {
	Scan(v ...interface{}) error
}
