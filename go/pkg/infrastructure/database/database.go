package database

import (
	"database/sql"
	"homepage/pkg/interface/repository"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type sqlHandler struct {
	DB *sql.DB
}

func NewSQLHandler() repository.SQLHandler {
	conn, err := sql.Open("mysql", "root:password@tcp(localhost:3307)/homepage")
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	return &sqlHandler{DB: conn}
}

func (sh *sqlHandler) ErrNoRows() error {
	return sql.ErrNoRows
}

func (sh *sqlHandler) Execute(query string, args ...interface{}) (repository.Result, error) {
	res, err := sh.DB.Exec(query, args...)
	if err != nil {
		log.Println("")
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
		log.Println("")
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
