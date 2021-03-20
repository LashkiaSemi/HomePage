//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package service

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/repository"

	"github.com/pkg/errors"
)

type Activity interface {
	GetAll() ([]*entity.Activity, error)
	GetByID(id int) (*entity.Activity, error)
	GetUpcoming() ([]*entity.Activity, error)
	GetForNotification() ([]*entity.Activity, error)

	Create(activity, showDate, date, annotation string, isImportant, isNotify int) (int, error)
	UpdateByID(id int, activity, showDate, date, annotation string, isImportant, isNotify int) error

	DeleteByID(id int) error
}

type activity struct {
	repo repository.ActivityRepository
}

func NewActivity(repo repository.ActivityRepository) Activity {
	return &activity{
		repo: repo,
	}
}

func (as *activity) GetAll() ([]*entity.Activity, error) {
	return as.repo.FindAll()
}

func (as *activity) GetByID(id int) (*entity.Activity, error) {
	return as.repo.FindByID(id)
}

func (as *activity) GetUpcoming() ([]*entity.Activity, error) {
	return as.repo.FindUpcoming()
}

func (as *activity) GetForNotification() ([]*entity.Activity, error) {
	return as.repo.FindByNotify()

}

func (as *activity) Create(activity, showDate, date, annotation string, isImportant, isNotify int) (int, error) {
	act := entity.NewActivity(activity, showDate, date, annotation, isImportant, isNotify)
	id, err := as.repo.Create(act)
	if err != nil {
		err = errors.Wrap(err, "failed to insert db")
		return 0, err
	}
	return id, nil

}

func (as *activity) UpdateByID(id int, activity, showDate, date, annotation string, isImportant, isNotify int) error {
	data, err := as.repo.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to get original data")
		return err
	}
	newData := data.Update(activity, showDate, date, annotation, isImportant, isNotify)

	// update db
	err = as.repo.UpdateByID(newData)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
	}
	return err
}

func (as *activity) DeleteByID(id int) error {
	return as.repo.DeleteByID(id)
}
