package interactor

import (
	"homepage/pkg/entity"
	"log"
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
	// AuthorizationByStudentID()

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
