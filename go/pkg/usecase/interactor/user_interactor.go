package interactor

import (
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"time"
)

// UserInteractor インたらクタ
type UserInteractor interface {
	FetchAll() (domain.Users, error)
	FetchByID(userID int) (domain.User, error)
	Add(name, password, role, studentID, department, comment string, grade int) (domain.User, error)
	Update(userID int, name, password, role, studentID, department, comment string, grade int) (domain.User, error)
	Delete(userID int) error
}

type userInteractor struct {
	UserRepository
	AuthHandler
}

// NewUserInteractor インたらクタの作成
func NewUserInteractor(ur UserRepository, ah AuthHandler) UserInteractor {
	return &userInteractor{
		UserRepository: ur,
		AuthHandler:    ah,
	}
}

func (ui *userInteractor) FetchAll() (domain.Users, error) {
	return ui.UserRepository.FindAll()
}

func (ui *userInteractor) FetchByID(userID int) (domain.User, error) {
	return ui.UserRepository.FindByID(userID)
}

func (ui *userInteractor) Add(name, password, role, studentID, department, comment string, grade int) (user domain.User, err error) {
	// password hashing
	hash, err := ui.AuthHandler.PasswordHash(password)
	if err != nil {
		return user, domain.InternalServerError(err)
	}

	// make role
	if role == "" {
		role = "member"
	}

	// get time now
	createdAt := time.Now()

	// do repository
	err = ui.UserRepository.Store(name, hash, role, studentID, department, comment, grade, createdAt)
	if err != nil {
		return user, err
	}

	// return
	// TODO: idだけ取得できてない
	user.Name = name
	user.Role = role
	user.StudentID = studentID
	user.Department = department
	user.Grade = grade
	user.Comment = comment
	user.CreatedAt = createdAt.Format("2006/01/02 15:04:05")
	user.UpdatedAt = user.CreatedAt
	return
}

func (ui *userInteractor) Update(userID int, name, password, role, studentID, department, comment string, grade int) (user domain.User, err error) {
	var hash string
	if password == "" {
		hash, err = ui.AuthHandler.PasswordHash(password)
		if err != nil {
			logger.Error(err)
			return user, domain.InternalServerError(err)
		}
	}

	updatedAt := time.Now()

	err = ui.UserRepository.Update(userID, name, hash, role, studentID, department, comment, grade, updatedAt)
	if err != nil {
		return user, err
	}
	return ui.UserRepository.FindByID(userID)
}

func (ui *userInteractor) Delete(userID int) error {
	return ui.UserRepository.Delete(userID)
}
