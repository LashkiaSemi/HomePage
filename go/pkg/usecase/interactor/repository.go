package interactor

import (
	"homepage/pkg/domain"
	"time"
)

// AccountRepository dbにつなぐ。実装はinterface > *_repositoryで
type AccountRepository interface {
	FindAccountByUserID(userID int) (domain.User, error)
	FindAccountByStudentID(studentID string) (user domain.User, err error)
	StoreAccount(name, password, role, studentID, department, comment string, grade int, createdAt time.Time) error
	UpdateAccount(userID int, name, password, role, studentID, department, comment string, grade int, updatedAt time.Time) error
	DeleteAccount(userID int) error
}

// UserRepository dbにつなぐ。実装はinterface > *_repositoryで
type UserRepository interface {
	FindUsers() (domain.Users, error)
	FindUserByUserID(userID int) (domain.User, error)
	StoreUser(name, password, role, studentID, department, comment string, grade int, createdAt time.Time) error
	UpdateUser(userID int, name, password, role, studentID, department, comment string, grade int, updatedAt time.Time) error
	DeleteUser(userID int) error
}
