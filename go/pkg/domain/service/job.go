//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package service

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/repository"

	"github.com/pkg/errors"
)

type Job interface {
	GetAll() ([]*entity.Job, error)
	GetByID(id int) (*entity.Job, error)
	Create(company, job string) (int, error)
	UpdateByID(id int, company, job string) error
	DeleteByID(id int) error
}

type job struct {
	repo repository.JobRepository
}

func NewJob(repo repository.JobRepository) Job {
	return &job{
		repo: repo,
	}
}

func (j *job) GetAll() ([]*entity.Job, error) {
	return j.repo.FindAll()
}

func (j *job) GetByID(id int) (*entity.Job, error) {
	return j.repo.FindByID(id)
}

func (j *job) Create(company, job string) (int, error) {
	obj := entity.NewJob(company, job)
	id, err := j.repo.Create(obj)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert data")
	}
	return id, nil
}

func (j *job) UpdateByID(id int, company, job string) error {
	data, err := j.repo.FindByID(id)
	if err != nil {
		return errors.Wrap(err, "failed to get origin data")
	}
	newData := data.Update(company, job)
	if err = j.repo.UpdateByID(newData); err != nil {
		return errors.Wrap(err, "failed to update")
	}
	return nil
}

func (j *job) DeleteByID(id int) error {
	return j.repo.DeleteByID(id)
}
