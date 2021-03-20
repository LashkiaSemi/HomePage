package repository

import "homepage/pkg/domain/entity"

type ResearchRepository interface {
	FindAll() ([]*entity.Research, error)
	FindByID(id int) (*entity.Research, error)

	Create(*entity.Research) (int, error)
	UpdateByID(*entity.Research) error

	DeleteByID(id int) error
}
