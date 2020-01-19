package datastore

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"homepage/conf"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/interface/repository"
)

type sqlHandler struct {
	DB *sql.DB
}

// NewSQLHandler データベースハンドラを作成
func NewSQLHandler() repository.SQLHandler {
	config := conf.LoadDatabaseConfig()

	conn, err := sql.Open(
		config["driver"],
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config["user"], config["password"], config["host"], config["port"], config["db"]),
	)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("connection database: ", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config["user"], config["password"], config["host"], config["port"], config["db"]))
	return &sqlHandler{DB: conn}
}

func (sh *sqlHandler) Execute(query string, args ...interface{}) (repository.Result, error) {
	res, err := sh.DB.Exec(query, args...)
	if err != nil {
		logger.Error(err)
		return &sqlResult{}, domain.InternalServerError(err)
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
		logger.Error(err)
		return &sqlRows{}, domain.InternalServerError(err)
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
		return &sqlTx{}, domain.InternalServerError(err)
	}
	return &sqlTx{Tx: tx}, nil
}

type sqlTx struct {
	Tx *sql.Tx
}

func (tx *sqlTx) Commit() error {
	return tx.Tx.Commit()
}

func (tx *sqlTx) Rollback() error {
	return tx.Tx.Rollback()
}

func (tx *sqlTx) Execute(query string, args ...interface{}) (repository.Result, error) {
	res, err := tx.Tx.Exec(query, args...)
	if err != nil {
		logger.Error(err)
		return &sqlResult{}, domain.InternalServerError(err)
	}
	return &sqlResult{Result: res}, nil
}

func (sh *sqlHandler) ErrNoRows() error {
	return sql.ErrNoRows
}
