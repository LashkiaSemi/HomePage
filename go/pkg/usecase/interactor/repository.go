package interactor

import (
	"homepage/pkg/entity"
)

// UserRepository ユーザのDBアクセスを実現
type UserRepository interface {
	FindAll() ([]*entity.User, error)
	FindByID(userID int) (*entity.User, error)
	FindByStudentID(studentID string) (*entity.User, error)
	FindAuthInfoByStudentID(studentID string) (*entity.User, error)

	UpdateByID(*entity.User) error
	UpdatePasswordByStudentID(studentID, password string) error
}

// SocietyRepository 学会発表のDBアクセスを実現
type SocietyRepository interface {
	FindAll() ([]*entity.Society, error)
}

// JobRepository 就職先のDBアクセスを実現
type JobRepository interface {
	FindAll() ([]*entity.Job, error)
}

// LectureRepository レクチャーのDBアクセスを実現
type LectureRepository interface {
	FindAll() ([]*entity.Lecture, error)
	FindByID(id int) (*entity.Lecture, error)
	FindAuthorByStudentID(studentID string) (*entity.User, error)

	Create(*entity.Lecture) (int, error)
	UpdateByID(*entity.Lecture) error
	DeleteByID(id int) error
}

// ResearchRepository 卒業研究関連のDBアクセスを実現
type ResearchRepository interface {
	FindAll() ([]*entity.Research, error)
}

// EquipmentRepository 備品関連のDBアクセスを実現
type EquipmentRepository interface {
	FindAll() ([]*entity.Equipment, error)
}
