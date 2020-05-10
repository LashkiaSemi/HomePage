package interactor

import "homepage/pkg/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindAuthInfoByStudentID(studentID string) (*model.User, error)
}

type SocietyRepository interface {
	FindAll() ([]*model.Society, error)
}

type JobRespotiroy interface {
	FindAll() ([]*model.Job, error)
}
