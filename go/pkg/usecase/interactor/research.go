package interactor

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/service"
)

type researchInteractor struct {
	srv service.Research
}

// ResearchInteractor 卒業研究のユースケースを実装
type ResearchInteractor interface {
	GetAll() ([]*entity.Research, error)
	GetByID(id int) (*entity.Research, error)

	Create(title, author, file, comment string, activation int) (int, error)
	UpdateByID(id int, title, author, file, comment string, activation int) error
	DeleteByID(id int) error
}

// NewResearchInteractor インタラクタを作成
func NewResearchInteractor(srv service.Research) ResearchInteractor {
	return &researchInteractor{
		srv: srv,
	}
}

func (ri *researchInteractor) GetAll() ([]*entity.Research, error) {
	return ri.srv.GetAll()
}

func (ri *researchInteractor) GetByID(id int) (*entity.Research, error) {
	return ri.srv.GetByID(id)

}

func (ri *researchInteractor) Create(title, author, file, comment string, activation int) (int, error) {
	return ri.srv.Create(title, author, file, comment, activation)
}

func (ri *researchInteractor) UpdateByID(id int, title, author, file, comment string, activation int) error {
	return ri.srv.UpdateByID(id, title, author, file, comment, activation)
}

func (ri *researchInteractor) DeleteByID(id int) error {
	return ri.srv.DeleteByID(id)
}
