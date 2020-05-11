package interactor

import (
	"homepage/pkg/domain/service"
	"homepage/pkg/entity"
)

type researchInteractor struct {
	service.ResearchService
	ResearchRepository
}

// ResearchInteractor 卒業研究のユースケースを実装
type ResearchInteractor interface {
	GetAll() ([]*entity.Research, error)
}

// NewResearchInteractor インタラクタを作成
func NewResearchInteractor(rs service.ResearchService, rr ResearchRepository) ResearchInteractor {
	return &researchInteractor{
		ResearchService:    rs,
		ResearchRepository: rr,
	}
}

func (ri *researchInteractor) GetAll() ([]*entity.Research, error) {
	return ri.ResearchRepository.FindAll()
}
