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
