package interactor

import (
	"homepage/pkg/entity"
)

type jobInteractor struct {
	JobRepository
}

// JobInteractor 就職先のユースケースを実装
type JobInteractor interface {
	GetAll() ([]*entity.Job, error)
}

// NewJobInteractor インタラクタの作成
func NewJobInteractor(jr JobRepository) JobInteractor {
	return &jobInteractor{
		JobRepository: jr,
	}
}

func (ji *jobInteractor) GetAll() ([]*entity.Job, error) {
	return ji.JobRepository.FindAll()
}
