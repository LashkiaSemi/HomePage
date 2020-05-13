package interactor

import (
	"homepage/pkg/entity"

	"github.com/pkg/errors"
)

type lectureInteractor struct {
	LectureRepository
}

// LectureInteractor レクチャーのユースケースを実装
type LectureInteractor interface {
	GetAll() ([]*entity.Lecture, error)
	GetByID(id int) (*entity.Lecture, error)
	Create(studentID, title, file, comment string, activation int) (*entity.Lecture, error)
	UpdateByID(id int, studentID, title, file, comment string, activation int) (*entity.Lecture, error)
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

func (li *lectureInteractor) Create(studentID, title, file, comment string, activation int) (*entity.Lecture, error) {
	author, err := li.LectureRepository.FindAuthorByStudentID(studentID)
	if err != nil {
		err = errors.Wrap(err, "failed to get author")
		return &entity.Lecture{}, err
	}

	lecture := entity.Lecture{}
	lecture.Create(title, file, comment, activation, author)

	id, err := li.LectureRepository.Create(&lecture)
	if err != nil {
		err = errors.Wrap(err, "failed to insert db")
		return &entity.Lecture{}, err
	}
	lecture.ID = id
	return &lecture, nil
}

func (li *lectureInteractor) UpdateByID(id int, studentID, title, file, comment string, activation int) (*entity.Lecture, error) {
	author, err := li.LectureRepository.FindAuthorByStudentID(studentID)
	if err != nil {
		err = errors.Wrap(err, "failed to get autho")
		return &entity.Lecture{}, err
	}
	lecture, err := li.LectureRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data")
		return &entity.Lecture{}, err
	}

	newLecture := lecture.Update(title, file, comment, activation, author)

	// 永続化
	err = li.LectureRepository.UpdateByID(newLecture)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
		return &entity.Lecture{}, err
	}
	return newLecture, nil

}

func (li *lectureInteractor) DeleteByID(id int) error {
	return li.LectureRepository.DeleteByID(id)
}
