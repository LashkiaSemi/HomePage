package interactor

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/domain/service"
)

type equipmentInteractor struct {
	service.EquipmentService
	EquipmentRepository
}

// EquipmentInteractor 備品のユースケースを実現
type EquipmentInteractor interface {
	GetAll() ([]*model.Equipment, error)
}

// NewEquipmentInteractor インタラクタの作成
func NewEquipmentInteractor(es service.EquipmentService, er EquipmentRepository) EquipmentInteractor {
	return &equipmentInteractor{
		EquipmentService:    es,
		EquipmentRepository: er,
	}
}

func (ei *equipmentInteractor) GetAll() ([]*model.Equipment, error) {
	return ei.EquipmentRepository.FindAll()
}
