//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

import "homepage/pkg/domain/entity"

type TagRepository interface {
	FindAll() ([]*entity.Tag, error)
	FindByID(id int) (*entity.Tag, error)
	Create(data *entity.Tag) (int, error)
	UpdateByID(data *entity.Tag) error
	DeleteByID(id int) error
}
