package interactor

import (
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"time"
)

type UserInteractor interface {
	FetchUsers() (domain.Users, error)
	FetchUserByUserID(userID int) (domain.User, error)
	AddUser(name, password, role, studentID, department, comment string, grade int) (domain.User, error)
	UpdateUser(userID int, name, password, role, studentID, department, comment string, grade int) (domain.User, error)
	DeleteUser(userID int) error
}

type userInteractor struct {
	UserRepository
	AuthHandler
}

func NewUserInteractor(ur UserRepository, ah AuthHandler) UserInteractor {
	return &userInteractor{
		UserRepository: ur,
		AuthHandler:    ah,
	}
}

func (ui *userInteractor) FetchUsers() (domain.Users, error) {
	return ui.UserRepository.FindUsers()
}

func (ui *userInteractor) FetchUserByUserID(userID int) (domain.User, error) {
	return ui.UserRepository.FindUserByUserID(userID)
}

func (ui *userInteractor) AddUser(name, password, role, studentID, department, comment string, grade int) (user domain.User, err error) {
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
	err = ui.UserRepository.StoreUser(name, hash, role, studentID, department, comment, grade, createdAt)
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

func (ui *userInteractor) UpdateUser(userID int, name, password, role, studentID, department, comment string, grade int) (user domain.User, err error) {
	var hash string
	if password == "" {
		hash, err = ui.AuthHandler.PasswordHash(password)
		if err != nil {
			logger.Error(err)
			return user, domain.InternalServerError(err)
		}
	}

	updatedAt := time.Now()

	err = ui.UserRepository.UpdateUser(userID, name, hash, role, studentID, department, comment, grade, updatedAt)
	if err != nil {
		return user, err
	}
	return ui.UserRepository.FindUserByUserID(userID)
}

func (ui *userInteractor) DeleteUser(userID int) error {
	return ui.UserRepository.DeleteUser(userID)
}
