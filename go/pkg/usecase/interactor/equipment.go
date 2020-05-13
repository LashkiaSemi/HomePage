package interactor

import (
	"homepage/pkg/entity"

	"github.com/pkg/errors"
)

type equipmentInteractor struct {
	EquipmentRepository
}

// EquipmentInteractor 備品のユースケースを実現
type EquipmentInteractor interface {
	GetAll() ([]*entity.Equipment, error)
	GetByID(id int) (*entity.Equipment, error)

	Create(name, comment string, stock, tagID int) (int, error)
	UpdateByID(id int, name, comment string, stock, tagID int) error
	DeleteByID(id int) error
}

// NewEquipmentInteractor インタラクタの作成
func NewEquipmentInteractor(er EquipmentRepository) EquipmentInteractor {
	return &equipmentInteractor{
		EquipmentRepository: er,
	}
}

func (ei *equipmentInteractor) GetAll() ([]*entity.Equipment, error) {
	return ei.EquipmentRepository.FindAll()
}

func (ei *equipmentInteractor) GetByID(id int) (*entity.Equipment, error) {
	return ei.EquipmentRepository.FindByID(id)

}

func (ei *equipmentInteractor) Create(name, comment string, stock, tagID int) (int, error) {
	// create obj
	equ := entity.Equipment{}
	equ.Create(name, comment, stock, tagID)

	// insert db
	id, err := ei.EquipmentRepository.Create(&equ)
	if err != nil {
		err = errors.Wrap(err, "failed to insert db")
		return 0, err
	}
	return id, nil
}

func (ei *equipmentInteractor) UpdateByID(id int, name, comment string, stock, tagID int) error {
	data, err := ei.EquipmentRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to get original data")
		return err
	}
	newData := data.Update(name, comment, stock, tagID)

	// update db
	err = ei.EquipmentRepository.UpdateByID(newData)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
		return err
	}
	return nil
}

func (ei *equipmentInteractor) DeleteByID(id int) error {
	return ei.EquipmentRepository.DeleteByID(id)
}
