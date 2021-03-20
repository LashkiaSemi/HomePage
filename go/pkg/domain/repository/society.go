//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

import "homepage/pkg/domain/entity"

type SocietyRepository interface {
	FindAll() ([]*entity.Society, error)
	FindByID(id int) (*entity.Society, error)

	Create(*entity.Society) (int, error)
	UpdateByID(*entity.Society) error

	DeleteByID(id int) error
}
