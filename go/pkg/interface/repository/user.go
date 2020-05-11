package repository

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/usecase/interactor"
	"log"
)

type userRepository struct {
	SQLHandler
}

// NewUserRepository リポジトリの作成
func NewUserRepository(sh SQLHandler) interactor.UserRepository {
	return &userRepository{
		SQLHandler: sh,
	}
}

func (ur *userRepository) FindAll() ([]*model.User, error) {
	rows, err := ur.SQLHandler.Query(`
		SELECT users.id, name, grade
		FROM users
		INNER JOIN introductions ON users.id = user_id
	`)
	if err != nil {
		log.Println("userRepository: FindAll: ", err)
		return []*model.User{}, err
	}

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Grade); err != nil {
			log.Println("userRepository: FindAll: ", err)
			continue
		}
		users = append(users, &user)
	}
	return users, nil
}

func (ur *userRepository) FindByID(userID string) (*model.User, error) {
	row := ur.SQLHandler.QueryRow(`
		SELECT users.id, users.name, users.student_id, intr.department, intr.grade, intr.comments
		FROM users
		INNER JOIN introductions as intr
		ON intr.user_id = users.id
		WHERE users.id = ?
	`, userID)
	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.StudentID, &user.Department, &user.Grade, &user.Comment); err != nil {
		log.Println("userRepository: FindByID: ", err)
		return &model.User{}, err
	}
	return &user, nil
}

func (ur *userRepository) FindAuthInfoByStudentID(studentID string) (*model.User, error) {
	row := ur.SQLHandler.QueryRow(`
		SELECT id, student_id, password_digest, role 
		FROM users
		WHERE student_id = ?`,
		studentID,
	)
	var user model.User
	if err := row.Scan(&user.ID, &user.StudentID, &user.Password, &user.Role); err != nil {
		log.Println("userRepository: findAuthInfoByStudentID: ", err)
		return &user, err
	}
	return &user, nil
}
