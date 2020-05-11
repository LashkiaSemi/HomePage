package interactor

import "homepage/pkg/domain/model"

// UserRepository ユーザのDBアクセスを実現
type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindAuthInfoByStudentID(studentID string) (*model.User, error)
}

// SocietyRepository 学会発表のDBアクセスを実現
type SocietyRepository interface {
	FindAll() ([]*model.Society, error)
}

// JobRepository 就職先のDBアクセスを実現
type JobRepository interface {
	FindAll() ([]*model.Job, error)
}

// LectureRepository レクチャーのDBアクセスを実現
type LectureRepository interface {
	FindAll() ([]*model.Lecture, error)
}

// ResearchRepository 卒業研究関連のDBアクセスを実現
type ResearchRepository interface {
	FindAll() ([]*model.Research, error)
}

// EquipmentRepository 備品関連のDBアクセスを実現
type EquipmentRepository interface {
	FindAll() ([]*model.Equipment, error)
}
