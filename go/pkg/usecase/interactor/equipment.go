//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package interactor

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/service"
)

type equipmentInteractor struct {
	srv service.Equipment
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
func NewEquipmentInteractor(srv service.Equipment) EquipmentInteractor {
	return &equipmentInteractor{
		srv: srv,
	}
}

func (ei *equipmentInteractor) GetAll() ([]*entity.Equipment, error) {
	return ei.srv.GetAll()
}

func (ei *equipmentInteractor) GetByID(id int) (*entity.Equipment, error) {
	return ei.srv.GetByID(id)

}

func (ei *equipmentInteractor) Create(name, comment string, stock, tagID int) (int, error) {
	return ei.srv.Create(name, comment, stock, tagID)
}

func (ei *equipmentInteractor) UpdateByID(id int, name, comment string, stock, tagID int) error {
	return ei.srv.UpdateByID(id, name, comment, stock, tagID)
}

func (ei *equipmentInteractor) DeleteByID(id int) error {
	return ei.srv.DeleteByID(id)
}
