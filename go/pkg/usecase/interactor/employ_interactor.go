package interactor

import (
	"homepage/pkg/domain"
	"time"
)

// EmployInteractor インタラクタ
type EmployInteractor interface {
	FetchAll() (domain.Companies, error)
	FetchByID(compID int) (domain.Company, error)
	Add(company string) (domain.Company, error)
	Update(compID int, company string) (domain.Company, error)
	Delete(compID int) error
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

func (ei *employInteractor) FetchAll() (domain.Companies, error) {
	return ei.EmployRepository.FindAll()
}

func (ei *employInteractor) FetchByID(compID int) (domain.Company, error) {
	return ei.EmployRepository.FindByID(compID)
}

func (ei *employInteractor) Add(company string) (comp domain.Company, err error) {
	createdAt := time.Now()
	id, err := ei.EmployRepository.Store(company, createdAt)
	if err != nil {
		return
	}
	comp.ID = id
	comp.Company = company
	return
}

func (ei *employInteractor) Update(compID int, company string) (comp domain.Company, err error) {
	updatedAt := time.Now()
	err = ei.EmployRepository.Update(compID, company, updatedAt)
	if err != nil {
		return
	}
	return ei.EmployRepository.FindByID(compID)
}

func (ei *employInteractor) Delete(compID int) error {
	return ei.EmployRepository.Delete(compID)
}
