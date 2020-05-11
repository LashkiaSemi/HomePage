package interactor

import (
	"homepage/pkg/domain/service"
	"homepage/pkg/entity"
)

type jobInteractor struct {
	service.JobService
	JobRepository
}

// JobInteractor 就職先のユースケースを実装
type JobInteractor interface {
	GetAll() ([]*entity.Job, error)
}

// NewJobInteractor インタラクタの作成
func NewJobInteractor(js service.JobService, jr JobRepository) JobInteractor {
	return &jobInteractor{
		JobService:    js,
		JobRepository: jr,
	}
}

func (ji *jobInteractor) GetAll() ([]*entity.Job, error) {
	return ji.JobRepository.FindAll()
}
