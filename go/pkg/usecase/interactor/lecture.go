package interactor

import (
	"homepage/pkg/entity"
)

type lectureInteractor struct {
	LectureRepository
}

// LectureInteractor レクチャーのユースケースを実装
type LectureInteractor interface {
	GetAll() ([]*entity.Lecture, error)
}

// NewLectureInteractor インタラクタの作成
func NewLectureInteractor(lr LectureRepository) LectureInteractor {
	return &lectureInteractor{
		LectureRepository: lr,
	}
}

func (li *lectureInteractor) GetAll() ([]*entity.Lecture, error) {
	return li.LectureRepository.FindAll()
}
