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

	Create(company, job string) (int, error)
	UpdateByID(id int, company, job string) error
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

func (ji *jobInteractor) Create(company, job string) (int, error) {
	// create obj
	data := entity.Job{}
	data.Create(company, job)

	// insert db
	id, err := ji.JobRepository.Create(&data)
	if err != nil {
		err = errors.Wrap(err, "interactor: failed to insert db")
		return 0, err
	}
	return id, nil
}

func (ji *jobInteractor) UpdateByID(id int, company, job string) error {
	data, err := ji.JobRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "can't find target data")
		return err
	}
	newData := data.Update(company, job)

	// update db
	err = ji.JobRepository.UpdateByID(newData)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
		return err
	}
	return nil
}
