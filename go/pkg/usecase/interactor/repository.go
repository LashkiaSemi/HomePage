package interactor

import "homepage/pkg/domain/model"

// UserRepository ユーザのDBアクセスを実現
type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindAuthInfoByStudentID(studentID string) (*model.User, error)
}

// SOcietyRepository 学会発表のDBアクセスを実現
type SocietyRepository interface {
	FindAll() ([]*model.Society, error)
}

// JobRepository 就職先のDBアクセスを実現
type JobRepositroy interface {
	FindAll() ([]*model.Job, error)
}
