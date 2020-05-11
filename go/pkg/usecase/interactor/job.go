package interactor

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/domain/service"
)

type jobInteractor struct {
	service.JobService
	JobRepository
}

// JobInteractor 就職先のユースケースを実装
type JobInteractor interface {
	GetAll() ([]*model.Job, error)
}

// NewJobInteractor インタラクタの作成
func NewJobInteractor(js service.JobService, jr JobRepository) JobInteractor {
	return &jobInteractor{
		JobService:    js,
		JobRepository: jr,
	}
}

func (ji *jobInteractor) GetAll() ([]*model.Job, error) {
	return ji.JobRepository.FindAll()
}
