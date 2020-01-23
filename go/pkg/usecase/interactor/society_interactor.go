package interactor

import (
	"homepage/conf"
	"homepage/pkg/domain"
	"time"
)

// SocietyInteractor インタラクタ
type SocietyInteractor interface {
	FetchAll() (domain.Societies, error)
	FetchByID(socID int) (domain.Society, error)
	Add(title, author, society, award string, date time.Time) (domain.Society, error)
	Update(socID int, title, author, society, award string, date time.Time) (domain.Society, error)
	Delete(socID int) error
}

type societyInteractor struct {
	SocietyRepository SocietyRepository
}

// NewSocietyInteractor インタラクタの作成
func NewSocietyInteractor(sr SocietyRepository) SocietyInteractor {
	return &societyInteractor{
		SocietyRepository: sr,
	}
}

func (si *societyInteractor) FetchAll() (domain.Societies, error) {
	return si.SocietyRepository.FindAll()
}

func (si *societyInteractor) FetchByID(socID int) (domain.Society, error) {
	return si.SocietyRepository.FindByID(socID)
}

func (si *societyInteractor) Add(title, author, society, award string, date time.Time) (soc domain.Society, err error) {
	createdAt := time.Now()
	id, err := si.SocietyRepository.Store(title, author, society, award, date, createdAt)
	if err != nil {
		return
	}
	soc.ID = id
	soc.Title = title
	soc.Author = author
	soc.Society = society
	soc.Award = award
	soc.Date = date.Format(conf.DateFormat)
	return
}

func (si *societyInteractor) Update(socID int, title, author, society, award string, date time.Time) (soc domain.Society, err error) {
	updatedAt := time.Now()
	err = si.SocietyRepository.Update(socID, title, author, society, award, date, updatedAt)
	if err != nil {
		return soc, err
	}
	return si.SocietyRepository.FindByID(socID)
}

func (si *societyInteractor) Delete(socID int) error {
	return si.SocietyRepository.Delete(socID)
}
