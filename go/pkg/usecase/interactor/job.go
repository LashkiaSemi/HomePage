package interactor

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/service"
)

type jobInteractor struct {
	srv service.Job
}

// JobInteractor 就職先のユースケースを実装
type JobInteractor interface {
	GetAll() ([]*entity.Job, error)
	GetByID(id int) (*entity.Job, error)
	Create(company, job string) (int, error)
	UpdateByID(id int, company, job string) error
	DeleteByID(id int) error
}

// NewJobInteractor インタラクタの作成
func NewJobInteractor(srv service.Job) JobInteractor {
	return &jobInteractor{
		srv: srv,
	}
}

func (ji *jobInteractor) GetAll() ([]*entity.Job, error) {
	return ji.srv.GetAll()
}

func (ji *jobInteractor) GetByID(id int) (*entity.Job, error) {
	return ji.srv.GetByID(id)
}

func (ji *jobInteractor) Create(company, job string) (int, error) {
	return ji.srv.Create(company, job)
}

func (ji *jobInteractor) UpdateByID(id int, company, job string) error {
	return ji.srv.UpdateByID(id, company, job)
}

func (ji *jobInteractor) DeleteByID(id int) error {
	return ji.srv.DeleteByID(id)
}
