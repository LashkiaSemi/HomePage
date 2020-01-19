package interactor

import (
	"homepage/pkg/domain"
	"time"
)

type AccountRepository interface {
	FindAccountByUserID(userID int) (domain.User, error)
	FindAccountByStudentID(studentID string) (user domain.User, err error)
	StoreAccount(name, password, role, studentID, department, comment string, grade int, createdAt time.Time) error
}
