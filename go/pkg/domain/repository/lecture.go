//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

import "homepage/pkg/domain/entity"

type LectureRepository interface {
	FindAll() ([]*entity.Lecture, error)
	FindByID(id int) (*entity.Lecture, error)
	FindAuthorByStudentID(studentID string) (*entity.User, error)

	Create(*entity.Lecture) (int, error)
	UpdateByID(*entity.Lecture) error
	DeleteByID(id int) error
}
