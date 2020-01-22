package interactor

import (
	"homepage/pkg/domain"
	"time"
)

// EmployInteractor インタラクタ
type EmployInteractor interface {
	FetchAll() (domain.Jobs, error)
	FetchByID(jobID int) (domain.Job, error)
	Add(company, job string) (domain.Job, error)
	Update(jobID int, company, job string) (domain.Job, error)
	Delete(jobID int) error
}

type employInteractor struct {
	EmployRepository
}

// NewEmployInteractor インタラクタの作成
func NewEmployInteractor(er EmployRepository) EmployInteractor {
	return &employInteractor{
		EmployRepository: er,
	}
}

func (ei *employInteractor) FetchAll() (domain.Jobs, error) {
	return ei.EmployRepository.FindAll()
}

func (ei *employInteractor) FetchByID(jobID int) (domain.Job, error) {
	return ei.EmployRepository.FindByID(jobID)
}

func (ei *employInteractor) Add(company, jobName string) (job domain.Job, err error) {
	createdAt := time.Now()
	id, err := ei.EmployRepository.Store(company, jobName, createdAt)
	if err != nil {
		return
	}
	job.ID = id
	job.Company = company
	job.Job = jobName
	return
}

func (ei *employInteractor) Update(jobID int, company, jobName string) (job domain.Job, err error) {
	updatedAt := time.Now()
	err = ei.EmployRepository.Update(jobID, company, jobName, updatedAt)
	if err != nil {
		return
	}
	return ei.EmployRepository.FindByID(jobID)
}

func (ei *employInteractor) Delete(jobID int) error {
	return ei.EmployRepository.Delete(jobID)
}
