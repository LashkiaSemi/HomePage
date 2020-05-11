package interactor

import (
	"homepage/pkg/entity"
)

type societyInteractor struct {
	SocietyRepository
}

// SocietyInteractor 学会発表のユースケースを実装
type SocietyInteractor interface {
	GetAll() ([]*entity.Society, error)
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
