package interactor

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/domain/service"
	"log"
)

type societyInteractor struct {
	service.SocietyService
	SocietyRepository
}

// SocietyInteractor 学会発表のユースケースを実装
type SocietyInteractor interface {
	GetAll() ([]*model.Society, error)
}

// NewSocietyInteractor インタラクタの作成
func NewSocietyInteractor(ss service.SocietyService, sr SocietyRepository) SocietyInteractor {
	return &societyInteractor{
		SocietyService:    ss,
		SocietyRepository: sr,
	}
}

func (si *societyInteractor) GetAll() ([]*model.Society, error) {
	log.Println(si.SocietyService.Create())
	datas, err := si.SocietyRepository.FindAll()
	return datas, err

}
