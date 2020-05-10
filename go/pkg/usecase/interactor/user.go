package interactor

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/domain/service"
)

type userInteractor struct {
	service.UserService
	UserRepository
	VerifyHandler
}

// UserInteractor ユーザのユースケースを実装
type UserInteractor interface {
	GetAll() ([]*model.User, error)
	// AuthenticationByStudentID 学籍番号からログイン機能を使う
	AuthenticationByStudentID(studentID, password string) error

	// AuthorizationByStudentID 学籍番号から認可
	// AuthorizationByStudentID()

}

// NewUserInteractor インタラクタの作成
func NewUserInteractor(us service.UserService, ur UserRepository, vh VerifyHandler) UserInteractor {
	return &userInteractor{
		UserService:    us,
		UserRepository: ur,
		VerifyHandler:  vh,
	}
}

func (ui *userInteractor) GetAll() ([]*model.User, error) {
	return ui.UserRepository.FindAll()
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
