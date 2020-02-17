package interactor

import (
	"homepage/conf"
	"homepage/pkg/domain"
	"time"
)

// ResearchInteractor インタラクタ
type ResearchInteractor interface {
	FetchAll() (domain.Researches, error)
	FetchByID(resID int) (domain.Research, error)
	Add(title, author, file, comment string, isPublic int) (domain.Research, error)
	Update(resID int, title, author, file, comment string, isPublic int) (domain.Research, error)
	Delete(resID int) error
}

type researchInteractor struct {
	ResearchRepository
}

// NewResearchInteractor インタラクタの作成
func NewResearchInteractor(rr ResearchRepository) ResearchInteractor {
	return &researchInteractor{
		ResearchRepository: rr,
	}
}

func (ri *researchInteractor) FetchAll() (domain.Researches, error) {
	return ri.ResearchRepository.FindAll()
}

func (ri *researchInteractor) FetchByID(resID int) (domain.Research, error) {
	return ri.ResearchRepository.FindByID(resID)
}

func (ri *researchInteractor) Add(title, author, file, comment string, isPublic int) (res domain.Research, err error) {
	createdAt := time.Now()
	id, err := ri.ResearchRepository.Store(title, author, file, comment, createdAt, isPublic)
	if err != nil {
		return
	}
	res.ID = id
	res.Title = title
	res.Author = author
	res.File = file
	res.Comment = comment
	res.CreatedAt = createdAt.Format(conf.DateFormat)
	return
}

func (ri *researchInteractor) Update(resID int, title, author, file, comment string, isPublic int) (res domain.Research, err error) {
	updatedAt := time.Now()
	err = ri.ResearchRepository.Update(resID, title, author, file, comment, updatedAt, isPublic)
	if err != nil {
		return
	}
	return ri.ResearchRepository.FindByID(resID)
}

func (ri *researchInteractor) Delete(resID int) error {
	return ri.ResearchRepository.Delete(resID)
}
