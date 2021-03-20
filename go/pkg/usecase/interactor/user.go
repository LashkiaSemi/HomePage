//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package interactor

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/service"
)

type userInteractor struct {
	srv service.User
}

// UserInteractor ユーザのユースケースを実装
type UserInteractor interface {
	GetAll() ([]*entity.User, error)
	GetByID(userID int) (*entity.User, error)
	GetByStudentID(studentID string) (*entity.User, error)

	UpdateByID(userID int, name, studentID, department, comment string, grade int) (*entity.User, error)
	UpdatePasswordByStudentID(studentID, oldPassword, newPassword string) error

	// AuthenticationByStudentID 学籍番号からログイン機能を使う
	AuthenticationByStudentID(studentID, password string) error

	// AuthorizationByStudentID 学籍番号からadminの認可
	AdminAuthorizationByStudentID(studentID, password string) error

	// Admin系 roleのことがあるので、普通のupdateと分けてます
	AdminCreate(name, studentID, password, role, department, comment string, grade int) (int, error)
	AdminUpdate(userID int, name, studentID, role, department, comment string, grade int) error

	DeleteByID(id int) error
}

// NewUserInteractor インタラクタの作成
func NewUserInteractor(srv service.User) UserInteractor {
	return &userInteractor{
		srv: srv,
	}
}

func (ui *userInteractor) GetAll() ([]*entity.User, error) {
	return ui.srv.GetAll()
}

func (ui *userInteractor) GetByID(userID int) (*entity.User, error) {
	return ui.srv.GetByID(userID)
}

func (ui *userInteractor) GetByStudentID(studentID string) (*entity.User, error) {
	return ui.srv.GetByStudentID(studentID)
}

func (ui *userInteractor) UpdateByID(userID int, name, studentID, department, comment string, grade int) (*entity.User, error) {
	return ui.srv.UpdateByID(userID, name, studentID, department, comment, grade)
}

func (ui *userInteractor) UpdatePasswordByStudentID(studentID, oldPassword, newPassword string) error {
	return ui.srv.UpdatePasswordByStudentID(studentID, oldPassword, newPassword)
}

// admin
func (ui *userInteractor) AdminCreate(name, studentID, password, role, department, comment string, grade int) (int, error) {
	return ui.srv.AdminCreate(name, studentID, password, role, department, comment, grade)
}

func (ui *userInteractor) AdminUpdate(userID int, name, studentID, role, department, comment string, grade int) error {
	return ui.srv.AdminUpdate(userID, name, studentID, role, department, comment, grade)
}

func (ui *userInteractor) DeleteByID(id int) error {
	return ui.srv.DeleteByID(id)
}

// auth
func (ui *userInteractor) AuthenticationByStudentID(studentID, password string) error {
	return ui.srv.AuthenticationByStudentID(studentID, password)
}

func (ui *userInteractor) AdminAuthorizationByStudentID(studentID, password string) error {
	return ui.srv.AdminAuthorizationByStudentID(studentID, password)
}
