package interactor

import (
	"homepage/conf"
	"homepage/pkg/domain"
	"time"
)

// SocietyInteractor インタラクタ
type SocietyInteractor interface {
	FetchSocieties() (domain.Societies, error)
	FetchSocietyByID(socID int) (domain.Society, error)
	AddSociety(title, author, society, award string, date time.Time) (domain.Society, error)
	UpdateSociety(socID int, title, author, society, award string, date time.Time) (domain.Society, error)
	DeleteSociety(socID int) error
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

func (si *societyInteractor) FetchSocieties() (domain.Societies, error) {
	return si.SocietyRepository.FindSocieties()
}

func (si *societyInteractor) FetchSocietyByID(socID int) (domain.Society, error) {
	return si.SocietyRepository.FindSocietyByID(socID)
}

func (si *societyInteractor) AddSociety(title, author, society, award string, date time.Time) (soc domain.Society, err error) {
	createdAt := time.Now()
	id, err := si.SocietyRepository.StoreSociety(title, author, society, award, date, createdAt)
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

func (si *societyInteractor) UpdateSociety(socID int, title, author, society, award string, date time.Time) (soc domain.Society, err error) {
	updatedAt := time.Now()
	err = si.SocietyRepository.UpdateSociety(socID, title, author, society, award, date, updatedAt)
	if err != nil {
		return soc, err
	}
	return si.SocietyRepository.FindSocietyByID(socID)
}

func (si *societyInteractor) DeleteSociety(socID int) error {
	return si.SocietyRepository.DeleteSociety(socID)
}
