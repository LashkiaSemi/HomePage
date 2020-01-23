package interactor

import (
	"homepage/conf"
	"homepage/pkg/domain"
	"time"
)

// ActivityInteractor インタラクタ
type ActivityInteractor interface {
	FetchAll() (domain.Activities, error)
	FetchByID(actID int) (domain.Activity, error)
	Add(date time.Time, activity string) (domain.Activity, error)
	Update(actID int, date time.Time, activity string) (domain.Activity, error)
	Delete(actID int) error
}

type activityInteractor struct {
	ActivityRepository
}

// NewActivityInteractor インたらクタの作成
func NewActivityInteractor(ar ActivityRepository) ActivityInteractor {
	return &activityInteractor{
		ActivityRepository: ar,
	}
}

func (ai *activityInteractor) FetchAll() (domain.Activities, error) {
	return ai.ActivityRepository.FindAll()
}

func (ai *activityInteractor) FetchByID(actID int) (domain.Activity, error) {
	return ai.ActivityRepository.FindByID(actID)
}

func (ai *activityInteractor) Add(date time.Time, activity string) (act domain.Activity, err error) {
	createdAt := time.Now()

	actID, err := ai.ActivityRepository.Store(date, activity, createdAt)
	if err != nil {
		return act, err
	}

	act.ID = actID
	act.Date = date.Format(conf.DateFormat)
	act.Activity = activity
	return act, nil
}

func (ai *activityInteractor) Update(actID int, date time.Time, activity string) (act domain.Activity, err error) {
	updatedAt := time.Now()

	err = ai.ActivityRepository.Update(actID, date, activity, updatedAt)
	if err != nil {
		return act, err
	}

	return ai.ActivityRepository.FindByID(actID)
}

func (ai *activityInteractor) Delete(actID int) error {
	return ai.ActivityRepository.Delete(actID)
}
