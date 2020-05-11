package interactor

import (
	"homepage/pkg/domain/service"
	"homepage/pkg/entity"
	"log"
)

type userInteractor struct {
	service.UserService
	UserRepository
	VerifyHandler
}

// UserInteractor ユーザのユースケースを実装
type UserInteractor interface {
	GetAll() ([]*entity.User, error)
	GetByID(userID string) (*entity.User, error)
	UpdateByID(userID, name, studentID, department, comment string, grade int) (*entity.User, error)

	// AuthenticationByStudentID 学籍番号からログイン機能を使う
	AuthenticationByStudentID(studentID, password string) error

	// AuthorizationByStudentID 学籍番号からadminの認可
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

func (ui *userInteractor) GetAll() ([]*entity.User, error) {
	return ui.UserRepository.FindAll()
}

func (ui *userInteractor) GetByID(userID string) (*entity.User, error) {
	return ui.UserRepository.FindByID(userID)
}

func (ui *userInteractor) UpdateByID(userID, name, studentID, department, comment string, grade int) (*entity.User, error) {
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
