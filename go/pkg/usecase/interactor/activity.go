package interactor

import (
	"homepage/pkg/entity"

	"github.com/pkg/errors"
)

type activityInteractor struct {
	ActivityRepository
}

// ActivityInteractor 活動内容のユースケースを実現
type ActivityInteractor interface {
	GetAll() ([]*entity.Activity, error)
	GetByID(id int) (*entity.Activity, error)

	Create(activity, date string) (int, error)
	UpdateByID(id int, activity, date string) error

	DeleteByID(id int) error
}

// NewActivityInteractor インタラクタの作成
func NewActivityInteractor(ar ActivityRepository) ActivityInteractor {
	return &activityInteractor{
		ActivityRepository: ar,
	}
}

func (ai *activityInteractor) GetAll() ([]*entity.Activity, error) {
	acts, err := ai.ActivityRepository.FindAll()
	if err != nil {
		// TODO: いらんくね?
		err = errors.Wrap(err, "GetAll")
	}
	return acts, err
}

func (ai *activityInteractor) GetByID(id int) (*entity.Activity, error) {
	data, err := ai.ActivityRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "GetByID")
	}
	return data, nil
}

func (ai *activityInteractor) Create(activity, date string) (int, error) {
	// create obj
	act := entity.Activity{}
	act.Create(activity, date)

	// insert db
	id, err := ai.ActivityRepository.Create(&act)
	if err != nil {
		err = errors.Wrap(err, "interactor")
		return 0, err
	}
	return id, nil
}

func (ai *activityInteractor) UpdateByID(id int, activity, date string) error {
	data, err := ai.ActivityRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "can't find target data")
		return err
	}
	newData := data.Update(activity, date)

	// update db
	err = ai.ActivityRepository.UpdateByID(newData)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
		return err
	}
	return nil
}

func (ai *activityInteractor) DeleteByID(id int) error {
	return ai.ActivityRepository.DeleteByID(id)
}
