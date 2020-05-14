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
	var acts []*entity.Activity
	if err != nil {
		if err == ar.SQLHandler.ErrNoRows() {
			log.Printf("[warn] not data hit: %v", err)
			return acts, nil
		}
		err = errors.Wrap(err, "failed to execute query")
		return acts, err
	}
	for rows.Next() {
		var act entity.Activity
		if err = rows.Scan(&act.ID, &act.Activity, &act.Date); err != nil {
			log.Printf("[warn] rows.Scan skip: %v", err)
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
		err = errors.Wrap(err, "failed to bind data")
		return &data, err
	}
	return &data, nil
}

func (ar *activityRepository) Create(data *entity.Activity) (int, error) {
	result, err := ar.SQLHandler.Execute(`
		INSERT INTO activities(date, activity, created_at, updated_at)
		VALUES (?,?,?,?)
	`, data.Date, data.Activity, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrap(err, "failed to id new data")
		return 0, err
	}
	return int(id), nil
}

func (ar *activityRepository) UpdateByID(data *entity.Activity) error {
	_, err := ar.SQLHandler.Execute(`
		UPDATE activities
		SET date=?, activity=?, updated_at=?
		WHERE id=?
	`, data.Date, data.Activity, data.UpdatedAt, data.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return err
	}
	return nil
}

func (ar *activityRepository) DeleteByID(id int) error {
	_, err := ar.SQLHandler.Execute(`
		DELETE FROM activities
		WHERE id=?
	`, id)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}
	return err
}
