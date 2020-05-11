package interactor

import (
	"homepage/pkg/entity"
)

type researchInteractor struct {
	ResearchRepository
}

// ResearchInteractor 卒業研究のユースケースを実装
type ResearchInteractor interface {
	GetAll() ([]*entity.Research, error)
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
