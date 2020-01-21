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

// ActivityRepository dbにつなぐ。実装は interface > *_repository
type ActivityRepository interface {
	FindActivities() (domain.Activities, error)
	FindActivityByID(actID int) (domain.Activity, error)
	StoreActivity(date time.Time, act string, createdAt time.Time) (int, error)
	UpdateActivity(actID int, date time.Time, act string, updatedAt time.Time) error
	DeleteActivity(actID int) error
}

// SocietyRepository dbにつなぐ。実装は interface > *_repository
type SocietyRepository interface {
	FindSocieties() (domain.Societies, error)
	FindSocietyByID(socID int) (domain.Society, error)
	StoreSociety(title, author, society, award string, date, createdAt time.Time) (int, error)
	UpdateSociety(socID int, title, author, society, award string, date, updatedAt time.Time) error
	DeleteSociety(socID int) error
}

// ResearchRepository dbにつなぐ。実装は interface > *_repository
type ResearchRepository interface {
	FindAll() (domain.Researches, error)
	FindByID(resID int) (domain.Research, error)
	Store(title, author, file, comment string, createdAt time.Time) (int, error)
	Update(resID int, title, author, file, comment string, updatedAt time.Time) error
	Delete(resID int) error
}

// EmployRepository dbにつなぐ。実装は interface > *_repository
type EmployRepository interface {
	FindAll() (domain.Companies, error)
	FindByID(compID int) (domain.Company, error)
	Store(company string, createdAt time.Time) (int, error)
	Update(compID int, company string, updatedAt time.Time) error
	Delete(compID int) error
}
