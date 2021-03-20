//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package interactor

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/service"
)

type societyInteractor struct {
	srv service.Society
}

// SocietyInteractor 学会発表のユースケースを実装
type SocietyInteractor interface {
	GetAll() ([]*entity.Society, error)
	GetByID(id int) (*entity.Society, error)
	Create(title, author, society, award, date string) (int, error)
	UpdateByID(id int, title, author, society, award, date string) error
	DeleteByID(id int) error
}

// NewSocietyInteractor インタラクタの作成
func NewSocietyInteractor(srv service.Society) SocietyInteractor {
	return &societyInteractor{
		srv: srv,
	}
}

func (si *societyInteractor) GetAll() ([]*entity.Society, error) {
	return si.srv.GetAll()
}

func (si *societyInteractor) GetByID(id int) (*entity.Society, error) {
	return si.srv.GetByID(id)
}

func (si *societyInteractor) Create(title, author, society, award, date string) (int, error) {
	return si.srv.Create(title, author, society, award, date)
}

func (si *societyInteractor) UpdateByID(id int, title, author, society, award, date string) error {
	return si.srv.UpdateByID(id, title, author, society, award, date)
}

func (si *societyInteractor) DeleteByID(id int) error {
	return si.srv.DeleteByID(id)
}
