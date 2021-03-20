//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package interactor

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/service"
)

type tagInteractor struct {
	srv service.Tag
}

// TagInteractor タグ関連のユースケースを実現
type TagInteractor interface {
	GetAll() ([]*entity.Tag, error)
	GetByID(id int) (*entity.Tag, error)

	Create(name string) (int, error)
	UpdateByID(id int, name string) error

	DeleteByID(id int) error
}

// NewTagInteractor インタラクタの作成
func NewTagInteractor(srv service.Tag) TagInteractor {
	return &tagInteractor{
		srv: srv,
	}
}

func (ti *tagInteractor) GetAll() ([]*entity.Tag, error) {
	return ti.srv.GetAll()
}

func (ti *tagInteractor) GetByID(id int) (*entity.Tag, error) {
	return ti.srv.GetByID(id)
}

func (ti *tagInteractor) Create(name string) (int, error) {
	return ti.srv.Create(name)
}

func (ti *tagInteractor) UpdateByID(id int, name string) error {
	return ti.srv.UpdateByID(id, name)
}

func (ti *tagInteractor) DeleteByID(id int) error {
	return ti.srv.DeleteByID(id)
}
