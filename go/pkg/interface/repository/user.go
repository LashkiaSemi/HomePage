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

func (ur *userRepository) FindAuthInfoByStudentID(studentID string) (*model.User, error) {
	row := ur.SQLHandler.QueryRow(`
		SELECT student_id, password_digest, role 
		FROM users
		WHERE student_id = ?`,
		studentID,
	)
	var user model.User
	if err := row.Scan(&user.StudentID, &user.Password, &user.Role); err != nil {
		log.Println("userRepository: findAuthInfoByStudentID: ", err)
		return &user, err
	}
	return &user, nil
}
