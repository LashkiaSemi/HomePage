package repository

import "homepage/pkg/domain/entity"

type EquipmentRepository interface {
	FindAll() ([]*entity.Equipment, error)
	FindByID(id int) (*entity.Equipment, error)

	Create(*entity.Equipment) (int, error)
	UpdateByID(*entity.Equipment) error

	DeleteByID(id int) error
}
