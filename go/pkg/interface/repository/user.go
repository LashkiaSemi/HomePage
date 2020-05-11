package repository

import (
	"homepage/pkg/entity"
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

func (ur *userRepository) FindAll() ([]*entity.User, error) {
	rows, err := ur.SQLHandler.Query(`
		SELECT users.id, name, grade
		FROM users
		INNER JOIN introductions ON users.id = user_id
	`)
	if err != nil {
		log.Println("userRepository: FindAll: ", err)
		return []*entity.User{}, err
	}

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Grade); err != nil {
			log.Println("userRepository: FindAll: ", err)
			continue
		}
		users = append(users, &user)
	}
	return users, nil
}

func (ur *userRepository) FindByID(userID int) (*entity.User, error) {
	row := ur.SQLHandler.QueryRow(`
		SELECT users.id, users.name, users.student_id, intr.department, intr.grade, intr.comments
		FROM users
		INNER JOIN introductions as intr
		ON intr.user_id = users.id
		WHERE users.id = ?
	`, userID)
	var user entity.User
	if err := row.Scan(&user.ID, &user.Name, &user.StudentID, &user.Department, &user.Grade, &user.Comment); err != nil {
		log.Println("userRepository: FindByID: ", err)
		return &entity.User{}, err
	}
	return &user, nil
}

func (ur *userRepository) FindByStudentID(studentID string) (*entity.User, error) {
	row := ur.SQLHandler.QueryRow(`
		SELECT users.id, users.name, users.student_id, intr.department, intr.grade, intr.comments
		FROM users
		INNER JOIN introductions as intr
		ON intr.user_id = users.id
		WHERE users.student_id = ?
	`, studentID)
	var user entity.User
	if err := row.Scan(&user.ID, &user.Name, &user.StudentID, &user.Department, &user.Grade, &user.Comment); err != nil {
		log.Println("userRepository: FindByID: ", err)
		return &entity.User{}, err
	}
	return &user, nil
}

func (ur *userRepository) FindAuthInfoByStudentID(studentID string) (*entity.User, error) {
	row := ur.SQLHandler.QueryRow(`
		SELECT id, student_id, password_digest, role 
		FROM users
		WHERE student_id = ?`,
		studentID,
	)
	var user entity.User
	if err := row.Scan(&user.ID, &user.StudentID, &user.Password, &user.Role); err != nil {
		log.Println("userRepository: findAuthInfoByStudentID: ", err)
		return &user, err
	}
	return &user, nil
}

func (ur *userRepository) UpdateByID(user *entity.User) error {
	// TODO: tx
	_, err := ur.SQLHandler.Execute(`
		UPDATE users
		SET name = ?, student_id=?, updated_at=?
		WHERE id=?
	`, user.Name, user.StudentID, user.UpdatedAt, user.ID)
	if err != nil {
		log.Println("userRepository: UpdateByID: ", err)
		return err
	}

	_, err = ur.SQLHandler.Execute(`
		UPDATE introductions
		SET department=?, comments=?, grade=?, updated_at=?
		WHERE user_id=?
	`, user.Department, user.Comment, user.Grade, user.UpdatedAt, user.ID)
	if err != nil {
		log.Println("userRepository: UpdateByID: ", err)
		return err
	}
	return nil
}
