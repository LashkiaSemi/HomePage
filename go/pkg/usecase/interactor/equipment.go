package interactor

import (
	"homepage/pkg/domain/service"
	"homepage/pkg/entity"
)

type equipmentInteractor struct {
	service.EquipmentService
	EquipmentRepository
}

// EquipmentInteractor 備品のユースケースを実現
type EquipmentInteractor interface {
	GetAll() ([]*entity.Equipment, error)
}

// NewEquipmentInteractor インタラクタの作成
func NewEquipmentInteractor(es service.EquipmentService, er EquipmentRepository) EquipmentInteractor {
	return &equipmentInteractor{
		EquipmentService:    es,
		EquipmentRepository: er,
	}
}

func (ei *equipmentInteractor) GetAll() ([]*entity.Equipment, error) {
	return ei.EquipmentRepository.FindAll()
}
