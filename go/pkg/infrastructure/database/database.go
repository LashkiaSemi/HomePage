package database

import (
	"database/sql"
	"fmt"
	"homepage/pkg/configs"
	"homepage/pkg/interface/repository"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type sqlHandler struct {
	DB *sql.DB
}

// NewSQLHandler dbのコネクションを作成
func NewSQLHandler() repository.SQLHandler {
	conn, err := sql.Open(
		configs.DBDriver,
		fmt.Sprintf("%s:%s@%s(%s)/%s", configs.DBUser, configs.DBPassword, configs.DBProtocol, configs.DBTarget, configs.DBName),
	)
	if err != nil {
		log.Fatalf("failed to open sql: %v", err)
	}
	log.Printf("[info] DB connection success: driver='%s', target='@%s:(%s)/%s', user='%s'",
		configs.DBDriver,
		configs.DBProtocol, configs.DBTarget, configs.DBName,
		configs.DBUser,
	)

	return &sqlHandler{DB: conn}
}

func (sh *sqlHandler) ErrNoRows() error {
	return sql.ErrNoRows
}

func (sh *sqlHandler) Execute(query string, args ...interface{}) (repository.Result, error) {
	res, err := sh.DB.Exec(query, args...)
	if err != nil {
		return &sqlResult{}, err
	}
	return &sqlResult{Result: res}, nil
}

type sqlResult struct {
	Result sql.Result
}

func (r *sqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r *sqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (sh *sqlHandler) Query(query string, args ...interface{}) (repository.Rows, error) {
	rows, err := sh.DB.Query(query, args...)
	if err != nil {
		return &sqlRows{}, err
	}
	return &sqlRows{Rows: rows}, nil
}

type sqlRows struct {
	Rows *sql.Rows
}

func (r *sqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r *sqlRows) Next() bool {
	return r.Rows.Next()
}

func (r *sqlRows) Close() error {
	return r.Rows.Close()
}

func (sh *sqlHandler) QueryRow(query string, args ...interface{}) repository.Row {
	row := sh.DB.QueryRow(query, args...)
	return &sqlRow{Row: row}
}

type sqlRow struct {
	Row *sql.Row
}

func (r *sqlRow) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

// transaction
func (sh *sqlHandler) Begin() (repository.Tx, error) {
	tx, err := sh.DB.Begin()
	if err != nil {
		return &sqlTx{}, err
	}
	return &sqlTx{Tx: tx}, nil
}

type sqlTx struct {
	*sql.Tx
}

func (s *sqlTx) Commit() error {
	return s.Tx.Commit()
}

func (s *sqlTx) Rollback() error {
	return s.Tx.Rollback()
}

func (s *sqlTx) Execute(query string, args ...interface{}) (repository.Result, error) {
	result, err := s.Tx.Exec(query, args...)
	if err != nil {
		return &sqlResult{}, err
	}
	return &sqlResult{Result: result}, nil
}
