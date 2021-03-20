package repository

import "homepage/pkg/domain/entity"

type TagRepository interface {
	FindAll() ([]*entity.Tag, error)
	FindByID(id int) (*entity.Tag, error)
	Create(data *entity.Tag) (int, error)
	UpdateByID(data *entity.Tag) error
	DeleteByID(id int) error
}
