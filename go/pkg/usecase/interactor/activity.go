//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package interactor

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/service"

	"github.com/pkg/errors"
)

type activityInteractor struct {
	srv service.Activity
}

// ActivityInteractor 活動内容のユースケースを実現
type ActivityInteractor interface {
	GetAll() ([]*entity.Activity, error)
	GetByID(id int) (*entity.Activity, error)
	GetUpcoming() ([]*entity.Activity, error)
	GetForNotification() ([]*entity.Activity, error)

	Create(activity, showDate, date, annotation string, isImportant, isNotify int) (int, error)
	UpdateByID(id int, activity, showDate, date, annotation string, isImportant, isNotify int) error

	DeleteByID(id int) error
}

// NewActivityInteractor インタラクタの作成
func NewActivityInteractor(as service.Activity) ActivityInteractor {
	return &activityInteractor{
		srv: as,
	}
}

func (ai *activityInteractor) GetAll() ([]*entity.Activity, error) {
	return ai.srv.GetAll()
}

func (ai *activityInteractor) GetByID(id int) (*entity.Activity, error) {
	return ai.srv.GetByID(id)
}

func (ai *activityInteractor) GetUpcoming() ([]*entity.Activity, error) {
	return ai.srv.GetUpcoming()
}

func (ai *activityInteractor) GetForNotification() ([]*entity.Activity, error) {
	return ai.srv.GetForNotification()
}

func (ai *activityInteractor) Create(activity, showDate, date, annotation string, isImportant, isNotify int) (int, error) {
	id, err := ai.srv.Create(activity, showDate, date, annotation, isImportant, isNotify)
	if err != nil {
		err = errors.Wrap(err, "failed to create activity")
		return 0, err
	}
	return id, nil
}

func (ai *activityInteractor) UpdateByID(id int, activity, showDate, date, annotation string, isImportant, isNotify int) error {
	err := ai.srv.UpdateByID(id, activity, showDate, date, annotation, isImportant, isNotify)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
	}
	return err
}

func (ai *activityInteractor) DeleteByID(id int) error {
	return ai.srv.DeleteByID(id)
}
