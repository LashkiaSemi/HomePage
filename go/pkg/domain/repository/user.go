package repository

import "homepage/pkg/domain/entity"

type UserRepository interface {
	FindAll() ([]*entity.User, error)
	FindByID(userID int) (*entity.User, error)
	FindByStudentID(studentID string) (*entity.User, error)
	FindAuthInfoByStudentID(studentID string) (*entity.User, error)

	UpdateByID(*entity.User) error
	UpdatePasswordByStudentID(studentID, password string) error

	AdminCreate(*entity.User) (int, error)
	AdminUpdateByID(*entity.User) error

	DeleteByID(userID int) error
}
