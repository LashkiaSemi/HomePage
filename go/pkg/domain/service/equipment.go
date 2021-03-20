package service

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/repository"

	"github.com/pkg/errors"
)

type Equipment interface {
	GetAll() ([]*entity.Equipment, error)
	GetByID(id int) (*entity.Equipment, error)

	Create(name, comment string, stock, tagID int) (int, error)
	UpdateByID(id int, name, comment string, stock, tagID int) error
	DeleteByID(id int) error
}

type equipment struct {
	repo repository.EquipmentRepository
}

func NewEquipment(repo repository.EquipmentRepository) Equipment {
	return &equipment{
		repo: repo,
	}
}

func (e *equipment) GetAll() ([]*entity.Equipment, error) {
	return e.repo.FindAll()
}

func (e *equipment) GetByID(id int) (*entity.Equipment, error) {
	return e.repo.FindByID(id)
}

func (e *equipment) Create(name, comment string, stock, tagID int) (int, error) {
	equipment := entity.NewEquipment(name, comment, stock, tagID)
	id, err := e.repo.Create(equipment)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert db")
	}
	return id, nil
}

func (e *equipment) UpdateByID(id int, name, comment string, stock, tagID int) error {
	equ, err := e.repo.FindByID(id)
	if err != nil {
		return errors.Wrap(err, "faield to get original data")
	}
	newEqu := equ.Update(name, comment, stock, tagID)
	if err = e.repo.UpdateByID(newEqu); err != nil {
		return errors.Wrap(err, "failed to update")
	}
	return nil
}

func (e *equipment) DeleteByID(id int) error {
	return e.repo.DeleteByID(id)
}
