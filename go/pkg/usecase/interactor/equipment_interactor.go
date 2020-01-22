package interactor

import (
	"homepage/pkg/domain"
	"time"
)

// EquipmentInteractor インタラクタ
type EquipmentInteractor interface {
	FetchAll() (domain.Equipments, error)
	FetchByID(equID int) (domain.Equipment, error)
	Add(name, note string, stock, tagID int) (domain.Equipment, error)
	Update(equID int, name, note string, stock, tagID int) (domain.Equipment, error)
	Delete(equID int) error
}

type equipmentInteractor struct {
	EquipmentRepository
}

// NewEquipmentInteractor インタラクタの作成
func NewEquipmentInteractor(er EquipmentRepository) EquipmentInteractor {
	return &equipmentInteractor{
		EquipmentRepository: er,
	}
}

func (ei *equipmentInteractor) FetchAll() (domain.Equipments, error) {
	return ei.EquipmentRepository.FindAll()
}

func (ei *equipmentInteractor) FetchByID(equID int) (domain.Equipment, error) {
	return ei.EquipmentRepository.FindByID(equID)
}

func (ei *equipmentInteractor) Add(name, note string, stock, tagID int) (equ domain.Equipment, err error) {
	createdAt := time.Now()
	id, err := ei.EquipmentRepository.Store(name, note, stock, tagID, createdAt)
	if err != nil {
		return
	}
	// タグの関係でもう一回db叩く
	return ei.EquipmentRepository.FindByID(id)
}

func (ei *equipmentInteractor) Update(equID int, name, note string, stock, tagID int) (equ domain.Equipment, err error) {
	updatedAt := time.Now()
	err = ei.EquipmentRepository.Update(equID, name, note, stock, tagID, updatedAt)
	if err != nil {
		return
	}
	return ei.EquipmentRepository.FindByID(equID)
}

func (ei *equipmentInteractor) Delete(equID int) error {
	return ei.EquipmentRepository.Delete(equID)
}
