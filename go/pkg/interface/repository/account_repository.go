package repository

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

type accountRepository struct {
	SQLHandler
}

// NewAccountRepository アカウント管理用DB接続
func NewAccountRepository(sh SQLHandler) interactor.AccountRepository {
	return &accountRepository{
		SQLHandler: sh,
	}
}

func (ar *accountRepository) FindByID(userID int) (user domain.User, err error) {
	row := ar.SQLHandler.QueryRow(
		`SELECT users.id, name, password_digest, role, student_id, users.created_at, users.updated_at, department, grade, comments 
		FROM users
		INNER JOIN introductions
		ON users.id = introductions.user_id
		WHERE users.id = ?`,
		userID)
	if err = row.Scan(&user.ID, &user.Name, &user.Password, &user.Role, &user.StudentID, &user.CreatedAt, &user.UpdatedAt, &user.Department, &user.Grade, &user.Comment); err != nil {
		if err != ar.SQLHandler.ErrNoRows() {
			logger.Warn("accoutn findByID: content not found")
			return user, domain.NotFound(errors.New("content not found"))
		}
		logger.Error("account findByID: ", err)
		return user, domain.InternalServerError(err)
	}
	return user, nil
}

func (ar *accountRepository) FindByStudentID(studentID string) (user domain.User, err error) {
	row := ar.SQLHandler.QueryRow(
		`SELECT users.id, name, password_digest, role, student_id, users.created_at, users.updated_at, department, grade, comments 
		FROM users
		INNER JOIN introductions
		ON users.id = introductions.user_id
		WHERE users.student_id = ?`,
		studentID)
	if err = row.Scan(&user.ID, &user.Name, &user.Password, &user.Role, &user.StudentID, &user.CreatedAt, &user.UpdatedAt, &user.Department, &user.Grade, &user.Comment); err != nil {
		if err != ar.SQLHandler.ErrNoRows() {
			logger.Warn("accoutn findByStudentID: content not found")
			return user, domain.NotFound(errors.New("content not found"))
		}
		logger.Error("account findByStudentID: ", err)
		return user, domain.InternalServerError(err)
	}
	return user, nil
}

func (ar *accountRepository) Store(name, password, role, studentID, department, comment string, grade int, createdAt time.Time) error {
	return transact(ar.SQLHandler, func(tx Tx) error {
		ret, err := tx.Execute(
			"INSERT INTO users(name, password_digest, role, student_id, created_at, updated_at) VALUES (?,?,?,?,?,?)",
			name, password, role, studentID, createdAt, createdAt,
		)
		if err != nil {
			return err
		}

		id, err := ret.LastInsertId()
		if err != nil {
			return err
		}

		if _, err := tx.Execute(
			"INSERT INTO introductions(user_id, department, grade, comments, created_at, updated_at) VALUES (?,?,?,?,?,?)",
			id, department, grade, comment, createdAt, createdAt,
		); err != nil {
			return err
		}
		return nil
	})
}

func (ar *accountRepository) Update(userID int, name, password, role, studentID, department, comment string, grade int, updatedAt time.Time) error {
	return transact(ar.SQLHandler, func(tx Tx) error {
		// users table
		values := map[string]interface{}{
			"name":            name,
			"password_digest": password,
			"role":            role,
			"student_id":      studentID,
			"updated_at":      updatedAt,
		}
		conds := map[string]interface{}{
			"id": userID,
		}
		query, args, _ := makeUpdateQuery("users", values, conds)
		_, err := tx.Execute(query, args...)
		if err != nil {
			return err
		}

		// introduction
		values = map[string]interface{}{
			"department": department,
			"grade":      grade,
			"comments":   comment,
			"updated_at": updatedAt,
		}
		conds = map[string]interface{}{
			"user_id": userID,
		}
		query, args, _ = makeUpdateQuery("introductions", values, conds)
		if _, err := tx.Execute(query, args...); err != nil {
			return err
		}

		return nil
	})
}

func (ar *accountRepository) Delete(userID int) error {
	return transact(ar.SQLHandler, func(tx Tx) error {
		if _, err := tx.Execute("DELETE FROM users WHERE id=?", userID); err != nil {
			return err
		}

		if _, err := tx.Execute("DELETE FROM introductions WHERE user_id=?", userID); err != nil {
			return err
		}

		return nil
	})
}
