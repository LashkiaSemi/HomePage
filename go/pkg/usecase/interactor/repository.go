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
