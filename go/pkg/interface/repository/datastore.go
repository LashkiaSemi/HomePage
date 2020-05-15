package repository

// SQLHandler sqlの操作子
type SQLHandler interface {
	ErrNoRows() error

	Execute(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row

	// Begin トランザクションの開始
	Begin() (Tx, error)
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

// Tx トランザクション用。sql.Tx
type Tx interface {
	Commit() error
	Rollback() error
	Execute(query string, args ...interface{}) (Result, error)
}
