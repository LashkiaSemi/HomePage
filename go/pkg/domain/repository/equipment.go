//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

import "homepage/pkg/domain/entity"

type EquipmentRepository interface {
	FindAll() ([]*entity.Equipment, error)
	FindByID(id int) (*entity.Equipment, error)

	Create(*entity.Equipment) (int, error)
	UpdateByID(*entity.Equipment) error

	DeleteByID(id int) error
}
