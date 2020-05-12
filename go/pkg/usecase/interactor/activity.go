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
		err = errors.Wrap(err, "interactor")
	}
	return acts, err
}
