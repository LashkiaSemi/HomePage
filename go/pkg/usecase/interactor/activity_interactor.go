package interactor

import (
	"homepage/conf"
	"homepage/pkg/domain"
	"time"
)

type ActivityInteractor interface {
	FetchActivities() (domain.Activities, error)
	FetchActivityByID(actID int) (domain.Activity, error)
	AddActiviry(date time.Time, activity string) (domain.Activity, error)
	UpdateActiviry(actID int, date time.Time, activity string) (domain.Activity, error)
	DeleteActiviry(actID int) error
}

type activityInteractor struct {
	ActivityRepository
}

func NewActivityInteractor(ar ActivityRepository) ActivityInteractor {
	return &activityInteractor{
		ActivityRepository: ar,
	}
}

func (ai *activityInteractor) FetchActivities() (domain.Activities, error) {
	return ai.ActivityRepository.FindActivities()
}

func (ai *activityInteractor) FetchActivityByID(actID int) (domain.Activity, error) {
	return ai.ActivityRepository.FindActivityByID(actID)
}

func (ai *activityInteractor) AddActiviry(date time.Time, activity string) (act domain.Activity, err error) {
	createdAt := time.Now()

	actID, err := ai.ActivityRepository.StoreActivity(date, activity, createdAt)
	if err != nil {
		return act, err
	}

	act.ID = actID
	act.Date = date.Format(conf.DateFormat)
	act.Activity = activity
	return act, nil
}

func (ai *activityInteractor) UpdateActiviry(actID int, date time.Time, activity string) (act domain.Activity, err error) {
	updatedAt := time.Now()

	err = ai.ActivityRepository.UpdateActivity(actID, date, activity, updatedAt)
	if err != nil {
		return act, err
	}

	return ai.ActivityRepository.FindActivityByID(actID)
}

func (ai *activityInteractor) DeleteActiviry(actID int) error {
	return ai.ActivityRepository.DeleteActivity(actID)
}
