package interactor

import (
	"homepage/pkg/entity"
)

type equipmentInteractor struct {
	EquipmentRepository
}

// EquipmentInteractor 備品のユースケースを実現
type EquipmentInteractor interface {
	GetAll() ([]*entity.Equipment, error)
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
