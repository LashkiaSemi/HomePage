package repository

import "homepage/pkg/domain/entity"

type JobRepository interface {
	FindAll() ([]*entity.Job, error)
	FindByID(id int) (*entity.Job, error)

	Create(*entity.Job) (int, error)
	UpdateByID(*entity.Job) error

	DeleteByID(id int) error
}
