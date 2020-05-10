package interactor

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/domain/service"
)

type userInteractor struct {
	service.UserService
	UserRepository
}

type UserInteractor interface {
	GetAll() ([]*model.User, error)
}

func NewUserInteractor(us service.UserService, ur UserRepository) UserInteractor {
	return &userInteractor{
		UserService:    us,
		UserRepository: ur,
	}
}

func (ui *userInteractor) GetAll() ([]*model.User, error) {
	return ui.UserRepository.FindAll()
}
