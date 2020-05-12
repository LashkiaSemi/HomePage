package interactor

import (
	"homepage/pkg/entity"

	"github.com/pkg/errors"
)

type researchInteractor struct {
	ResearchRepository
}

// ResearchInteractor 卒業研究のユースケースを実装
type ResearchInteractor interface {
	GetAll() ([]*entity.Research, error)
	GetByID(id int) (*entity.Research, error)
}

// NewResearchInteractor インタラクタを作成
func NewResearchInteractor(rr ResearchRepository) ResearchInteractor {
	return &researchInteractor{
		ResearchRepository: rr,
	}
}

func (ri *researchInteractor) GetAll() ([]*entity.Research, error) {
	return ri.ResearchRepository.FindAll()
}

func (ri *researchInteractor) GetByID(id int) (*entity.Research, error) {
	data, err := ri.ResearchRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "researchInteractor: GetByID")
	}
	return data, err
}
