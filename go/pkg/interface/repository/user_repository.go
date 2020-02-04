package repository

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

type userRepository struct {
	SQLHandler
}

// NewUserRepository ユーザ関連でデータベースを叩く
func NewUserRepository(sh SQLHandler) interactor.UserRepository {
	return &userRepository{
		SQLHandler: sh,
	}
}

func (ur *userRepository) FindAll() (users domain.Users, err error) {
	rows, err := ur.SQLHandler.Query(
		`SELECT users.id, name, password_digest, role, student_id, users.created_at, users.updated_at, department, grade, comments 
		FROM users
		INNER JOIN introductions
		ON users.id = introductions.user_id`,
	)
	if err != nil {
		return users, domain.InternalServerError(err)
	}

	for rows.Next() {
		var user domain.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Password, &user.Role, &user.StudentID, &user.CreatedAt, &user.UpdatedAt, &user.Department, &user.Grade, &user.Comment); err != nil {
			logger.Warn("find users: skip user data")
			continue
		}
		users = append(users, user)
	}
	return
}

func (ur *userRepository) FindByID(userID int) (user domain.User, err error) {
	row := ur.SQLHandler.QueryRow(
		`SELECT users.id, name, password_digest, role, student_id, users.created_at, users.updated_at, department, grade, comments 
		FROM users
		INNER JOIN introductions
		ON users.id = introductions.user_id
		WHERE users.id = ?`,
		userID,
	)

	if err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Role, &user.StudentID, &user.CreatedAt, &user.UpdatedAt, &user.Department, &user.Grade, &user.Comment); err != nil {
		if err == ur.SQLHandler.ErrNoRows() {
			logger.Warn("user findByID: content not found")
			return user, domain.NotFound(errors.New("content not found"))
		}
		logger.Error("user findByID: ", err)
		return user, domain.InternalServerError(err)
	}
	return
}

func (ur *userRepository) Store(name, password, role, studentID, department, comment string, grade int, createdAt time.Time) error {
	return transact(ur.SQLHandler, func(tx Tx) error {
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

func (ur *userRepository) Update(userID int, name, password, role, studentID, department, comment string, grade int, updatedAt time.Time) error {
	return transact(ur.SQLHandler, func(tx Tx) error {
		// users の更新
		query, args, _ := makeUpdateQuery(
			"users",
			map[string]interface{}{
				"name":            name,
				"password_digest": password,
				"role":            role,
				"student_id":      studentID,
				"updated_at":      updatedAt,
			},
			map[string]interface{}{
				"id": userID,
			},
		)
		_, err := tx.Execute(query, args...)
		if err != nil {
			return err
		}

		// introductionsの更新
		query, args, _ = makeUpdateQuery(
			"introductions",
			map[string]interface{}{
				"department": department,
				"comments":   comment,
				"grade":      grade,
				"updated_at": updatedAt,
			},
			map[string]interface{}{
				"user_id": userID,
			},
		)
		_, err = tx.Execute(query, args...)
		if err != nil {
			return err
		}
		return nil
	})
}

func (ur *userRepository) Delete(userID int) error {
	return transact(ur.SQLHandler, func(tx Tx) error {
		if _, err := tx.Execute("DELETE FROM users WHERE id=?", userID); err != nil {
			return err
		}

		if _, err := tx.Execute("DELETE FROM introductions WHERE user_id=?", userID); err != nil {
			return err
		}

		return nil
	})
}
