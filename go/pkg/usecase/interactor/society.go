package interactor

import (
	"homepage/pkg/entity"

	"github.com/pkg/errors"
)

type societyInteractor struct {
	SocietyRepository
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
func NewSocietyInteractor(sr SocietyRepository) SocietyInteractor {
	return &societyInteractor{
		SocietyRepository: sr,
	}
}

func (si *societyInteractor) GetAll() ([]*entity.Society, error) {
	return si.SocietyRepository.FindAll()
}

func (si *societyInteractor) GetByID(id int) (*entity.Society, error) {
	return si.SocietyRepository.FindByID(id)
}

func (si *societyInteractor) Create(title, author, society, award, date string) (int, error) {
	// create obj
	data := entity.Society{}
	data.Create(title, author, society, award, date)

	// insert db
	id, err := si.SocietyRepository.Create(&data)
	if err != nil {
		err = errors.Wrap(err, "failed to insert db")
		return 0, err
	}
	return id, nil
}

func (si *societyInteractor) UpdateByID(id int, title, author, society, award, date string) error {
	data, err := si.SocietyRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to get original data")
		return err
	}
	newData := data.Update(title, author, society, award, date)

	// update db
	err = si.SocietyRepository.UpdateByID(newData)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
		return err
	}
	return nil
}

func (si *societyInteractor) DeleteByID(id int) error {
	return si.SocietyRepository.DeleteByID(id)
}
