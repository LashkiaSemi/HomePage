//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package service

import (
	"errors"
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/repository"
)

type User interface {
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

// TODO: verifyをここに置くのがナンセンスな気がする...?
type user struct {
	repo repository.UserRepository
	vh   VerifyHandler
}

func NewUser(repo repository.UserRepository, vh VerifyHandler) User {
	return &user{
		repo: repo,
		vh:   vh,
	}
}

func (u *user) GetAll() ([]*entity.User, error) {
	return u.repo.FindAll()
}

func (u *user) GetByID(userID int) (*entity.User, error) {
	return u.repo.FindByID(userID)
}

func (u *user) GetByStudentID(studentID string) (*entity.User, error) {
	return u.repo.FindByStudentID(studentID)
}

func (u *user) UpdateByID(userID int, name, studentID, department, comment string, grade int) (*entity.User, error) {
	us, err := u.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	newData := us.Update(name, studentID, department, comment, grade)
	if err = u.repo.UpdateByID(newData); err != nil {
		return nil, err
	}
	return newData, nil
}

func (u *user) UpdatePasswordByStudentID(studentID, oldPassword, newPassword string) error {
	us, err := u.repo.FindAuthInfoByStudentID(studentID)
	if err != nil {
		return err
	}
	if err = u.vh.PasswordVerify(us.Password, oldPassword); err != nil {
		return err
	}

	hash, err := u.vh.PasswordHash(newPassword)
	if err != nil {
		return err
	}
	if err = u.repo.UpdatePasswordByStudentID(studentID, hash); err != nil {
		return err
	}
	return nil
}

func (u *user) DeleteByID(id int) error {
	return u.repo.DeleteByID(id)
}

// AuthenticationByStudentID 学籍番号からログイン機能を使う
func (u *user) AuthenticationByStudentID(studentID, password string) error {
	us, err := u.repo.FindAuthInfoByStudentID(studentID)
	if err != nil {
		return err
	}
	if err = u.vh.PasswordVerify(us.Password, password); err != nil {
		return err
	}
	return nil
}

// AdminAuthorizationByStudentID 学籍番号からadminの認可
func (u *user) AdminAuthorizationByStudentID(studentID, password string) error {
	us, err := u.repo.FindAuthInfoByStudentID(studentID)
	if err != nil {
		return err
	}
	if !us.IsAdmin() {
		return errors.New("request user is not admin")
	}
	if err = u.vh.PasswordVerify(us.Password, password); err != nil {
		return err
	}
	return nil
}

// Admin系 roleのことがあるので、普通のupdateと分けてます
func (u *user) AdminCreate(name, studentID, password, role, department, comment string, grade int) (int, error) {
	hash, err := u.vh.PasswordHash(password)
	if err != nil {
		return 0, err
	}
	us := entity.NewUser(name, studentID, hash, role, department, comment, grade)
	id, err := u.repo.AdminCreate(us)
	if err != nil {
		return 0, err
	}
	return id, nil

}

func (u *user) AdminUpdate(userID int, name, studentID, role, department, comment string, grade int) error {
	us, err := u.repo.FindByID(userID)
	if err != nil {
		return err
	}
	newUser := us.AdminUpdate(name, studentID, department, comment, role, grade)
	if err = u.repo.AdminUpdateByID(newUser); err != nil {
		return err
	}
	return nil
}
