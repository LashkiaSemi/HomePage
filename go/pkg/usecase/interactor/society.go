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
}

// NewSocietyInteractor インタラクタの作成
func NewSocietyInteractor(sr SocietyRepository) SocietyInteractor {
	return &societyInteractor{
		SocietyRepository: sr,
	}
}

func (si *societyInteractor) GetAll() ([]*entity.Society, error) {
	datas, err := si.SocietyRepository.FindAll()
	return datas, err
}

func (si *societyInteractor) GetByID(id int) (*entity.Society, error) {
	data, err := si.SocietyRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "GetByID")
	}
	return data, err
}
