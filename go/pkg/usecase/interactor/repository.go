package interactor

import "homepage/pkg/domain/model"

type SocietyRepository interface {
	FindAll() ([]*model.Society, error)
}
