package interactor

import "homepage/pkg/domain/model"

type SocietyRepository interface {
	FindAll() ([]*model.Society, error)
}

type JobRespotiroy interface {
	FindAll() ([]*model.Job, error)
}
