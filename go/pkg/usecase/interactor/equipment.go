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
	data, err := ei.EquipmentRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "equipmentInteractor: getByID")
	}
	return data, err
}
