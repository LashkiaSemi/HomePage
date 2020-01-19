package repository

type SQLHandler interface {
	ErrNoRows() error

	Execute(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row

	// trainsaction
	Begin() (Tx, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

type Row interface {
	Scan(v ...interface{}) error
}

type Tx interface {
	Commit() error
	Rollback() error
	Execute(query string, args ...interface{}) (Result, error)
}
