package interactor

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/domain/service"
)

type jobInteractor struct {
	service.JobService
	JobRepositroy
}

type JobInteractor interface {
	GetAll() ([]*model.Job, error)
}

func NewJobInteractor(js service.JobService, jr JobRepositroy) JobInteractor {
	return &jobInteractor{
		JobService:    js,
		JobRepositroy: jr,
	}
}

func (ji *jobInteractor) GetAll() ([]*model.Job, error) {
	return ji.JobRepositroy.FindAll()
}
