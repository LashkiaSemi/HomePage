//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package service

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/repository"

	"github.com/pkg/errors"
)

type Lecture interface {
	GetAll() ([]*entity.Lecture, error)
	GetByID(id int) (*entity.Lecture, error)
	Create(title, file, comment string, activation int, author *entity.User) (int, error)
	UpdateByID(id int, title, file, comment string, activation int, author *entity.User) error
	DeleteByID(id int) error
}

type lecture struct {
	repo repository.LectureRepository
}

func NewLecture(repo repository.LectureRepository) Lecture {
	return &lecture{
		repo: repo,
	}
}

func (l *lecture) GetAll() ([]*entity.Lecture, error) {
	return l.repo.FindAll()
}

func (l *lecture) GetByID(id int) (*entity.Lecture, error) {
	return l.repo.FindByID(id)
}

func (l *lecture) Create(title, file, comment string, activation int, author *entity.User) (int, error) {
	lecture := entity.NewLecture(title, file, comment, activation, author)
	id, err := l.repo.Create(lecture)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert data")
	}
	return id, nil
}

func (l *lecture) UpdateByID(id int, title, file, comment string, activation int, author *entity.User) error {
	lec, err := l.repo.FindByID(id)
	if err != nil {
		return errors.Wrap(err, "failed to get origin data")
	}
	newLec := lec.Update(title, file, comment, activation, author)
	if err := l.repo.UpdateByID(newLec); err != nil {
		return errors.Wrap(err, "failed to update")
	}
	return nil
}

func (l *lecture) DeleteByID(id int) error {
	return l.repo.DeleteByID(id)
}
