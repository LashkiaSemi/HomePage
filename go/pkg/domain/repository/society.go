package repository

import "homepage/pkg/domain/entity"

type SocietyRepository interface {
	FindAll() ([]*entity.Society, error)
	FindByID(id int) (*entity.Society, error)

	Create(*entity.Society) (int, error)
	UpdateByID(*entity.Society) error

	DeleteByID(id int) error
}
