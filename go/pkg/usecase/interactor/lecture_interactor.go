package interactor

import (
	"homepage/pkg/domain"
	"time"
)

// LectureInteractor インタラクタ
type LectureInteractor interface {
	FetchAll() (domain.Lectures, error)
	FetchByID(lecID int) (domain.Lecture, error)
	Add(title, file, comment string, userID int) (domain.Lecture, error)
	Update(lecID int, title, file, comment string, userID int) (domain.Lecture, error)
	Delete(lecID int) error
}

type lectureInteractor struct {
	LecutureRepository LectureRepository
}

// NewLectureInteractor インタラクタの作成
func NewLectureInteractor(lr LectureRepository) LectureInteractor {
	return &lectureInteractor{
		LecutureRepository: lr,
	}
}

func (li *lectureInteractor) FetchAll() (domain.Lectures, error) {
	return li.LecutureRepository.FindAll()
}

func (li *lectureInteractor) FetchByID(lecID int) (domain.Lecture, error) {
	return li.LecutureRepository.FindByID(lecID)
}

func (li *lectureInteractor) Add(title, file, comment string, userID int) (domain.Lecture, error) {
	createdAt := time.Now()
	id, err := li.LecutureRepository.Store(title, file, comment, userID, createdAt)
	if err != nil {
		return domain.Lecture{}, err
	}
	return li.LecutureRepository.FindByID(id)
}

func (li *lectureInteractor) Update(lecID int, title, file, comment string, userID int) (domain.Lecture, error) {
	updatedAt := time.Now()
	err := li.LecutureRepository.Update(lecID, title, file, comment, userID, updatedAt)
	if err != nil {
		return domain.Lecture{}, err
	}
	return li.LecutureRepository.FindByID(lecID)
}

func (li *lectureInteractor) Delete(lecID int) error {
	return li.LecutureRepository.Delete(lecID)
}
