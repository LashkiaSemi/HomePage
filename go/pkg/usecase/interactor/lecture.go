package interactor

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/domain/service"
)

type lectureInteractor struct {
	service.LectureService
	LectureRepository
}

// LectureInteractor レクチャーのユースケースを実装
type LectureInteractor interface {
	GetAll() ([]*model.Lecture, error)
}

// NewLectureInteractor インタラクタの作成
func NewLectureInteractor(ls service.LectureService, lr LectureRepository) LectureInteractor {
	return &lectureInteractor{
		LectureService:    ls,
		LectureRepository: lr,
	}
}

func (li *lectureInteractor) GetAll() ([]*model.Lecture, error) {
	return li.LectureRepository.FindAll()
}
