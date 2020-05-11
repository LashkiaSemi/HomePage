package interactor

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/domain/service"
)

type researchInteractor struct {
	service.ResearchService
	ResearchRepository
}

// ResearchInteractor 卒業研究のユースケースを実装
type ResearchInteractor interface {
	GetAll() ([]*model.Research, error)
}

// NewResearchInteractor インタラクタを作成
func NewResearchInteractor(rs service.ResearchService, rr ResearchRepository) ResearchInteractor {
	return &researchInteractor{
		ResearchService:    rs,
		ResearchRepository: rr,
	}
}

func (ri *researchInteractor) GetAll() ([]*model.Research, error) {
	return ri.ResearchRepository.FindAll()
}
