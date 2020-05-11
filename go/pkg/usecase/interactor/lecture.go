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
	GetByID(id int) (*entity.Lecture, error)
	UpdateByID(id int, title, comment string, activation int) (*entity.Lecture, error)
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

func (li *lectureInteractor) GetByID(id int) (*entity.Lecture, error) {
	return li.LectureRepository.FindByID(id)
}

func (li *lectureInteractor) UpdateByID(id int, title, comment string, activation int) (*entity.Lecture, error) {
	lecture, err := li.LectureRepository.FindByID(id)
	if err != nil {
		return &entity.Lecture{}, err
	}

	newLecture := lecture.Update(title, comment, activation)

	// 永続化
	err = li.LectureRepository.UpdateByID(newLecture)
	if err != nil {
		return &entity.Lecture{}, err
	}
	return newLecture, nil

}
