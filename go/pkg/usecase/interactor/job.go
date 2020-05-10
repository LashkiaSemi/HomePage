package interactor

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/domain/service"
)

type jobInteractor struct {
	service.JobService
	JobRespotiroy
}

type JobInteractor interface {
	GetAll() ([]*model.Job, error)
}

func NewJobInteractor(js service.JobService, jr JobRespotiroy) JobInteractor {
	return &jobInteractor{
		JobService:    js,
		JobRespotiroy: jr,
	}
}

func (ji *jobInteractor) GetAll() ([]*model.Job, error) {
	return ji.JobRespotiroy.FindAll()
}
