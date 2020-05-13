package interactor

import (
	"homepage/pkg/entity"
	"log"

	"github.com/pkg/errors"
)

type userInteractor struct {
	UserRepository
	VerifyHandler
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

	// roleのことがあるので...
	AdminCreate(name, studentID, password, role, department, comment string, grade int) (int, error)
	AdminUpdate(userID int, name, studentID, role, department, comment string, grade int) error

	DeleteByID(id int) error
}

// NewUserInteractor インタラクタの作成
func NewUserInteractor(ur UserRepository, vh VerifyHandler) UserInteractor {
	return &userInteractor{
		UserRepository: ur,
		VerifyHandler:  vh,
	}
}

func (ui *userInteractor) GetAll() ([]*entity.User, error) {
	return ui.UserRepository.FindAll()
}

func (ui *userInteractor) GetByID(userID int) (*entity.User, error) {
	return ui.UserRepository.FindByID(userID)
}

func (ui *userInteractor) GetByStudentID(studentID string) (*entity.User, error) {
	return ui.UserRepository.FindByStudentID(studentID)
}

func (ui *userInteractor) UpdateByID(userID int, name, studentID, department, comment string, grade int) (*entity.User, error) {
	user, err := ui.UserRepository.FindByID(userID)
	if err != nil {
		log.Println("userRepository: UpdateByID: ", err)
		return &entity.User{}, err
	}
	newUser := user.Update(name, studentID, department, comment, grade)
	err = ui.UserRepository.UpdateByID(newUser)
	if err != nil {
		log.Println("userRepository: UpdateByID: ", err)
		return &entity.User{}, err
	}
	return newUser, nil
}

func (ui *userInteractor) UpdatePasswordByStudentID(studentID, oldPassword, newPassword string) error {
	// 現在のパスワードがあってるか検証
	user, err := ui.UserRepository.FindAuthInfoByStudentID(studentID)
	if err != nil {
		log.Println("userInteractor: UpdatePassword: ", err)
		return err
	}
	err = ui.VerifyHandler.PasswordVerify(user.Password, oldPassword)
	if err != nil {
		log.Println("現在のパスワードが違う")
		return err
	}

	// 新規パスワードのハッシュ
	hash, err := ui.VerifyHandler.PasswordHash(newPassword)
	if err != nil {
		log.Println("パスワードのハッシュエラー")
		return err
	}

	// 新規パスワードを登録
	err = ui.UserRepository.UpdatePasswordByStudentID(studentID, hash)
	if err != nil {
		log.Println("登録でエラー")
		return err
	}
	return nil
}

// admin
func (ui *userInteractor) AdminCreate(name, studentID, password, role, department, comment string, grade int) (int, error) {
	// パスワードのハッシュ
	hash, err := ui.VerifyHandler.PasswordHash(password)
	if err != nil {
		err = errors.Wrap(err, "failed to hash password")
		return 0, err
	}

	// objectcreate
	user := entity.User{}
	user.Create(name, studentID, hash, role, department, comment, grade)

	id, err := ui.UserRepository.AdminCreate(&user)
	if err != nil {
		err = errors.Wrap(err, "adminCreate")
		return 0, err
	}

	return id, nil
}

func (ui *userInteractor) AdminUpdate(userID int, name, studentID, role, department, comment string, grade int) error {
	user, err := ui.UserRepository.FindByID(userID)
	if err != nil {
		log.Println("userRepository: UpdateByID: ", err)
		return err
	}

	newUser := user.AdminUpdate(name, studentID, department, comment, role, grade)
	err = ui.UserRepository.AdminUpdateByID(newUser)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
		return err
	}
	return nil
}

func (ui *userInteractor) DeleteByID(id int) error {
	return ui.UserRepository.DeleteByID(id)
}

// auth
func (ui *userInteractor) AuthenticationByStudentID(studentID, password string) error {
	user, err := ui.UserRepository.FindAuthInfoByStudentID(studentID)
	if err != nil {
		return err
	}

	err = ui.VerifyHandler.PasswordVerify(user.Password, password)
	if err != nil {
		return err
	}
	return nil
}

func (ui *userInteractor) AdminAuthorizationByStudentID(studentID, password string) error {
	user, err := ui.UserRepository.FindAuthInfoByStudentID(studentID)
	if err != nil {
		err = errors.Wrap(err, "failed to get auth info")
		return err
	}
	if user.Role != "admin" && user.Role != "owner" {
		err = errors.New("failed to authorized permission")
		return err
	}
	err = ui.VerifyHandler.PasswordVerify(user.Password, password)
	if err != nil {
		err = errors.Wrap(err, "failed to verify password")
		return err
	}
	return nil
}
