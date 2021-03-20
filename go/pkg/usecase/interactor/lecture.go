package interactor

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/service"

	"github.com/pkg/errors"
)

type lectureInteractor struct {
	srv  service.Lecture
	user service.User
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
func NewLectureInteractor(srv service.Lecture, user service.User) LectureInteractor {
	return &lectureInteractor{
		srv:  srv,
		user: user,
	}
}

func (li *lectureInteractor) GetAll() ([]*entity.Lecture, error) {
	return li.srv.GetAll()
}

func (li *lectureInteractor) GetByID(id int) (*entity.Lecture, error) {
	return li.GetByID(id)
}

func (li *lectureInteractor) Create(studentID, title, file, comment string, activation int) (int, error) {
	author, err := li.user.GetByStudentID(studentID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get author")
	}

	id, err := li.srv.Create(title, file, comment, activation, author)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert db")
	}
	return id, nil
}

func (li *lectureInteractor) UpdateByID(id int, studentID, title, file, comment string, activation int) error {
	author, err := li.user.GetByStudentID(studentID)
	if err != nil {
		err = errors.Wrap(err, "failed to get autho")
		return err
	}
	return li.srv.UpdateByID(id, title, file, comment, activation, author)
}

func (li *lectureInteractor) DeleteByID(id int) error {
	return li.srv.DeleteByID(id)
}
