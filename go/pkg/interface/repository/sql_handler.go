package repository

// SQLHandler sql用のハンドラ。db叩きたいときはこの子を噛ませる
type SQLHandler interface {
	ErrNoRows() error

	Execute(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row

	// trainsaction
	Begin() (Tx, error)
}

// Result sqlexecuteした時の戻り
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Rows sqlQueryやった時の戻り
type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

// Row sqlQueryRowやった時の戻り
type Row interface {
	Scan(v ...interface{}) error
}

// Tx トランザクションする時に使います
type Tx interface {
	Commit() error
	Rollback() error
	Execute(query string, args ...interface{}) (Result, error)
}
