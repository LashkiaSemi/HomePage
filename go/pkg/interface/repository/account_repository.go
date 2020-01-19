package repository

import (
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

func (ar *accountRepository) FindAccountByUserID(userID int) (user domain.User, err error) {
	row := ar.SQLHandler.QueryRow(
		`SELECT users.id, name, password_digest, role, student_id, users.created_at, users.updated_at, department, grade, comments 
		FROM users
		INNER JOIN introductions
		ON users.id = introductions.user_id
		WHERE users.id = ?`,
		userID)
	if err = row.Scan(&user.ID, &user.Name, &user.Password, &user.Role, &user.StudentID, &user.CreatedAt, &user.UpdatedAt, &user.Department, &user.Grade, &user.Comment); err != nil {
		if err != ar.SQLHandler.ErrNoRows() {
			logger.Error(err)
			domain.InternalServerError(err)
			return
		}
	}
	return user, nil
}

func (ar *accountRepository) FindAccountByStudentID(studentID string) (user domain.User, err error) {
	row := ar.SQLHandler.QueryRow(
		`SELECT users.id, name, password_digest, role, student_id, users.created_at, users.updated_at, department, grade, comments 
		FROM users
		INNER JOIN introductions
		ON users.id = introductions.user_id
		WHERE users.student_id = ?`,
		studentID)
	if err = row.Scan(&user.ID, &user.Name, &user.Password, &user.Role, &user.StudentID, &user.CreatedAt, &user.UpdatedAt, &user.Department, &user.Grade, &user.Comment); err != nil {
		if err != ar.SQLHandler.ErrNoRows() {
			logger.Error(err)
			domain.InternalServerError(err)
			return
		}
	}
	return user, nil
}

func (ar *accountRepository) StoreAccount(name, password, role, studentID, department, comment string, grade int, createdAt time.Time) error {
	return Transact(ar.SQLHandler, func(tx Tx) error {
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
