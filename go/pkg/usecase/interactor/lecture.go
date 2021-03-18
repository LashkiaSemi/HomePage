package interactor

import (
	"homepage/pkg/domain/entity"

	"github.com/pkg/errors"
)

type lectureInteractor struct {
	LectureRepository
}

// LectureInteractor レクチャーのユースケースを実装
type LectureInteractor interface {
	GetAll() ([]*entity.Lecture, error)
	GetByID(id int) (*entity.Lecture, error)
	Create(studentID, title, file, comment string, activation int) (int, error)
	UpdateByID(id int, studentID, title, file, comment string, activation int) error
	DeleteByID(id int) error
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

func (li *lectureInteractor) Create(studentID, title, file, comment string, activation int) (int, error) {
	author, err := li.LectureRepository.FindAuthorByStudentID(studentID)
	if err != nil {
		err = errors.Wrap(err, "failed to get author")
		return 0, err
	}

	lecture := entity.Lecture{}
	lecture.Create(title, file, comment, activation, author)

	id, err := li.LectureRepository.Create(&lecture)
	if err != nil {
		err = errors.Wrap(err, "failed to insert db")
		return 0, err
	}
	return id, nil
}

func (li *lectureInteractor) UpdateByID(id int, studentID, title, file, comment string, activation int) error {
	author, err := li.LectureRepository.FindAuthorByStudentID(studentID)
	if err != nil {
		err = errors.Wrap(err, "failed to get autho")
		return err
	}
	lecture, err := li.LectureRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data")
		return err
	}

	newLecture := lecture.Update(title, file, comment, activation, author)

	// 永続化
	err = li.LectureRepository.UpdateByID(newLecture)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
		return err
	}
	return nil

}

func (li *lectureInteractor) DeleteByID(id int) error {
	return li.LectureRepository.DeleteByID(id)
}
