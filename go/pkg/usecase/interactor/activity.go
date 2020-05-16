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
	GetUpcoming() ([]*entity.Activity, error)

	Create(activity, showDate, lastDate, annotation string, isImportant int) (int, error)
	UpdateByID(id int, activity, showDate, lastDate, annotation string, isImportant int) error

	DeleteByID(id int) error
}

// NewActivityInteractor インタラクタの作成
func NewActivityInteractor(ar ActivityRepository) ActivityInteractor {
	return &activityInteractor{
		ActivityRepository: ar,
	}
}

func (ai *activityInteractor) GetAll() ([]*entity.Activity, error) {
	return ai.ActivityRepository.FindAll()
}

func (ai *activityInteractor) GetByID(id int) (*entity.Activity, error) {
	return ai.ActivityRepository.FindByID(id)
}

func (ai *activityInteractor) GetUpcoming() ([]*entity.Activity, error) {
	return ai.ActivityRepository.FindUpcoming()
}

func (ai *activityInteractor) Create(activity, showDate, lastDate, annotation string, isImportant int) (int, error) {
	// create obj
	act := entity.Activity{}
	act.Create(activity, showDate, lastDate, annotation, isImportant)

	// insert db
	id, err := ai.ActivityRepository.Create(&act)
	if err != nil {
		err = errors.Wrap(err, "failed to insert db")
		return 0, err
	}
	return id, nil
}

func (ai *activityInteractor) UpdateByID(id int, activity, showDate, lastDate, annotation string, isImportant int) error {
	data, err := ai.ActivityRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to get original data")
		return err
	}
	newData := data.Update(activity, showDate, lastDate, annotation, isImportant)

	// update db
	err = ai.ActivityRepository.UpdateByID(newData)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
	}
	return err
}

func (ai *activityInteractor) DeleteByID(id int) error {
	return ai.ActivityRepository.DeleteByID(id)
}
