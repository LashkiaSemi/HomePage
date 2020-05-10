package interactor

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/domain/service"
)

type jobInteractor struct {
	service.JobService
	JobRepository
}

type JobInteractor interface {
	GetAll() ([]*model.Job, error)
}

func NewJobInteractor(js service.JobService, jr JobRepository) JobInteractor {
	return &jobInteractor{
		JobService:    js,
		JobRepository: jr,
	}
}

func (ji *jobInteractor) GetAll() ([]*model.Job, error) {
	return ji.JobRepository.FindAll()
}
