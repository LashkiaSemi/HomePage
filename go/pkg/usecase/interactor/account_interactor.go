package interactor

import (
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"time"

	"github.com/google/uuid"
)

// AccountInteractor
type AccountInteractor interface {
	FetchAccountByUserID(userID int) (domain.User, error)
	FetchAccountByStudentID(studentID string) (domain.User, error)
	AddAccount(name, password, role, studentID, department, comment string, grade int) (domain.User, error)
	UpdateAccount(userID int, name, password, role, studentID, department, comment string, grade int) (domain.User, error)
	DeleteAccount(userID int) error

	// Login ログイン機能
	Login(studentID, password string) (domain.Session, error)
}

type accountInteractor struct {
	AccountRepository
	AuthHandler
}

func NewAccountInteractor(ar AccountRepository, ah AuthHandler) AccountInteractor {
	return &accountInteractor{
		AccountRepository: ar,
		AuthHandler:       ah,
	}
}

func (ai *accountInteractor) FetchAccountByUserID(userID int) (domain.User, error) {
	return ai.AccountRepository.FindAccountByUserID(userID)
}

func (ai *accountInteractor) FetchAccountByStudentID(studentID string) (domain.User, error) {
	return ai.AccountRepository.FindAccountByStudentID(studentID)
}

func (ai *accountInteractor) AddAccount(name, password, role, studentID, department, comment string, grade int) (user domain.User, err error) {
	// password hashing
	hash, err := ai.AuthHandler.PasswordHash(password)
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
	err = ai.AccountRepository.StoreAccount(name, hash, role, studentID, department, comment, grade, createdAt)
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

func (ai *accountInteractor) UpdateAccount(userID int, name, password, role, studentID, department, comment string, grade int) (user domain.User, err error) {
	// passwordあるならハッシュ
	var hash string
	if password != "" {
		hash, err = ai.PasswordHash(password)
		if err != nil {
			logger.Error(err)
			return user, domain.InternalServerError(err)
		}
	}

	// time
	updatedAt := time.Now()

	err = ai.AccountRepository.UpdateAccount(userID, name, hash, role, studentID, department, comment, grade, updatedAt)
	user, err = ai.AccountRepository.FindAccountByUserID(userID)
	return
}

func (ai *accountInteractor) DeleteAccount(userID int) error {
	// TODO: 実装して
	err := ai.AccountRepository.DeleteAccount(userID)
	return err
}

func (ai *accountInteractor) Login(studentID, password string) (sess domain.Session, err error) {
	// データの取得
	user, err := ai.AccountRepository.FindAccountByStudentID(studentID)
	if err != nil {
		return sess, err
	}

	// パスワード認証
	err = ai.AuthHandler.PasswordVerify(user.Password, password)
	if err != nil {
		return sess, domain.BadRequest(err)
	}

	// sessionの作成
	sessionID, err := uuid.NewRandom()
	if err != nil {
		return sess, domain.InternalServerError(err)
	}

	sess.StudentID = user.StudentID
	sess.SessionID = sessionID.String()
	sess.UserID = user.ID
	sess.Role = user.Role

	return sess, nil
}
