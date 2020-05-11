package interactor

import (
	"homepage/pkg/domain/service"
	"homepage/pkg/entity"
)

type lectureInteractor struct {
	service.LectureService
	LectureRepository
}

// LectureInteractor レクチャーのユースケースを実装
type LectureInteractor interface {
	GetAll() ([]*entity.Lecture, error)
}

// NewLectureInteractor インタラクタの作成
func NewLectureInteractor(ls service.LectureService, lr LectureRepository) LectureInteractor {
	return &lectureInteractor{
		LectureService:    ls,
		LectureRepository: lr,
	}
}

func (li *lectureInteractor) GetAll() ([]*entity.Lecture, error) {
	return li.LectureRepository.FindAll()
}
