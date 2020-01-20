package repository

import (
	"fmt"
	"homepage/pkg/domain/logger"
	"strings"
	"time"
)

// transact トランザクションを作成する
func transact(sh SQLHandler, txFunc func(Tx) error) (err error) {
	tx, err := sh.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}

// makeUpdateQuery updateのsqlクエリの作成。tableと、map[field]value
func makeUpdateQuery(table string, values map[string]interface{}, conds map[string]interface{}) (string, []interface{}, error) {
	var query = fmt.Sprintf("UPDATE %s SET", table)
	var args []interface{}
	// ここでvalue
	for key, value := range values {
		switch value.(type) {
		case int:
			if value != 0 {
				query += fmt.Sprintf(" %v=?,", key)
				args = append(args, value)
			}
		case int64:
			if value != 0 {
				query += fmt.Sprintf(" %v=?,", key)
				args = append(args, value)
			}
		case string:
			if value != "" {
				query += fmt.Sprintf(" %v=?,", key)
				args = append(args, value)
			}
		case time.Time:
			// TODO: ここnil？
			if value != nil {
				query += fmt.Sprintf(" %v=?,", key)
				args = append(args, value)
			}
		default:
			logger.Debug(fmt.Sprintf("%T", value))
		}
	}

	query = strings.TrimSuffix(query, ",")
	query += " WHERE"

	// ここでwhere
	for key, value := range conds {
		switch value.(type) {
		case int:
			if value != 0 {
				query += fmt.Sprintf(" %v=?", key)
				args = append(args, value)
			}
		case int64:
			if value != 0 {
				query += fmt.Sprintf(" %v=?", key)
				args = append(args, value)
			}
		case string:
			if value != "" {
				query += fmt.Sprintf(" %v=?", key)
				args = append(args, value)
			}
		case time.Time:
			// TODO: ここnil？
			if value != nil {
				query += fmt.Sprintf(" %v=?", key)
				args = append(args, value)
			}
		default:
			logger.Debug(fmt.Sprintf("%T", value))
		}
	}
	return query, args, nil
}
