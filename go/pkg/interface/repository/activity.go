package repository

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"log"

	"github.com/pkg/errors"
)

type activityRepository struct {
	SQLHandler
}

// NewActivityRepository リポジトリの作成
func NewActivityRepository(sh SQLHandler) interactor.ActivityRepository {
	return &activityRepository{
		SQLHandler: sh,
	}
}

func (ar *activityRepository) FindAll() ([]*entity.Activity, error) {
	rows, err := ar.SQLHandler.Query(`
		SELECT id, activity, date
		FROM activities
		ORDER BY date DESC
	`)
	if err != nil {
		err = errors.Wrap(err, "can't get activities from db")
		return []*entity.Activity{}, err
	}
	var acts []*entity.Activity
	for rows.Next() {
		var act entity.Activity
		if err = rows.Scan(&act.ID, &act.Activity, &act.Date); err != nil {
			log.Println("activityRepository: findAll: skip scan: ", err)
			continue
		}
		acts = append(acts, &act)
	}
	return acts, nil
}

func (ar *activityRepository) FindByID(id int) (*entity.Activity, error) {
	row := ar.SQLHandler.QueryRow(`
		SELECT id, activity, date
		FROM activities
		WHERE id=?
	`, id)
	var data entity.Activity
	if err := row.Scan(&data.ID, &data.Activity, &data.Date); err != nil {
		errors.Wrap(err, "FindByID")
		return &data, err
	}
	return &data, nil
}
