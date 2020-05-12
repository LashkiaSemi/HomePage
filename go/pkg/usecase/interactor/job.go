package interactor

import (
	"homepage/pkg/entity"

	"github.com/pkg/errors"
)

type jobInteractor struct {
	JobRepository
}

// JobInteractor 就職先のユースケースを実装
type JobInteractor interface {
	GetAll() ([]*entity.Job, error)
	GetByID(id int) (*entity.Job, error)
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

func (ji *jobInteractor) GetByID(id int) (*entity.Job, error) {
	data, err := ji.JobRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "jobInteractor: GetByID")
	}
	return data, err
}
