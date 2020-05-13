package repository

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"log"

	"github.com/pkg/errors"
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
		SELECT users.id, users.name, users.student_id, users.role, intr.department, intr.grade, intr.comments
		FROM users
		INNER JOIN introductions as intr
		ON intr.user_id = users.id
		WHERE users.id = ?
	`, userID)
	var user entity.User
	if err := row.Scan(&user.ID, &user.Name, &user.StudentID, &user.Role, &user.Department, &user.Grade, &user.Comment); err != nil {
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

func (ur *userRepository) UpdatePasswordByStudentID(studentID, password string) error {
	_, err := ur.SQLHandler.Execute(`
		UPDATE users 
		SET password_digest = ? 
		WHERE student_id = ?
	`, password, studentID)
	if err != nil {
		log.Println("userRepository: UpdatePassword: ", err)
		return err
	}
	return nil
}

func (ur *userRepository) AdminCreate(user *entity.User) (int, error) {
	// TODO: tx
	result, err := ur.SQLHandler.Execute(`
		INSERT INTO users(name, password_digest, role, created_at, updated_at, student_id)
		VALUES (?,?,?,?,?,?)
	`, user.Name, user.Password, user.Role, user.CreatedAt, user.UpdatedAt, user.StudentID)
	if err != nil {
		err = errors.Wrap(err, "can't create users")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrap(err, "can't get id")
		return 0, err
	}
	_, err = ur.SQLHandler.Execute(`
		INSERT INTO introductions(user_id, department, grade, comments, created_at, updated_at)
		VALUES (?,?,?,?,?,?)
	`, id, user.Department, user.Grade, user.Comment, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		err = errors.Wrap(err, "can't create introductions")
		return 0, err
	}
	return int(id), nil
}

func (ur *userRepository) AdminUpdateByID(user *entity.User) error {
	// TODO: tx
	_, err := ur.SQLHandler.Execute(`
		UPDATE users
		SET name=?, role=?, updated_at=?, student_id=?
		WHERE id=?
	`, user.Name, user.Role, user.UpdatedAt, user.StudentID, user.ID)
	if err != nil {
		err = errors.Wrap(err, "can't create users")
		return err
	}
	_, err = ur.SQLHandler.Execute(`
		UPDATE introductions
		SET department=?, grade=?, comments=?, updated_at=?
		WHERE user_id=?
	`, user.Department, user.Grade, user.Comment, user.UpdatedAt, user.ID)
	if err != nil {
		err = errors.Wrap(err, "can't create introductions")
		return err
	}
	return nil
}
